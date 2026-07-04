function PIssueListShow(group) {
  const issues = PIssueRead(group);

  if (issues.length === 0) {
    return `<div class="PIssueRow PIssueOk"><span class="PIssueName"><span class="PIssueDot"></span>${PLanguageTextRead("noBlockingIssues")}</span><span class="PIssueValue">${PLanguageTextRead("ok")}</span></div>`;
  }

  return issues.slice(0, 4).map(issue => `
    <div class="PIssueRow ${PIssueClassRead(issue.tone)}">
      <span class="PIssueName"><span class="PIssueDot"></span>${LHtmlEscape(issue.name)}</span>
      <span class="PIssueValue">${LHtmlEscape(issue.value)}</span>
    </div>
  `).join("");
}

function PIssueRead(group) {
  const issues = [];

  (group.LReportSection || []).forEach(section => {
    const items = section.LReportItem || [];
    if (items.length === 0) return;

    const title = section.LReportTitle || section.LReportBadge || "Issue";
    const lowerTitle = title.toLowerCase();
    const tone = PIssueToneRead(section);

    if (lowerTitle.includes("numbering")) {
      issues.push({ name: PLanguageTextRead("missingPartNumbers"), value: PIssueNumberRead(items, group), count: PIssueMissingRead(group), tone });
      return;
    }

    if (lowerTitle.includes("compatibility")) {
      issues.push({ name: PIssueCompatibilityRead(title), value: PLanguageCountRead(items.length, "issue", "issues"), count: items.length, tone });
      return;
    }

    issues.push({ name: title, value: items.length === 1 ? items[0] : PLanguageCountRead(items.length, "issue", "issues"), count: items.length, tone });
  });

  return issues;
}

function PIssueCompatibilityRead(title) {
  return title.toLowerCase().includes("caution") ? PLanguageTextRead("compatibilityCaution") : PLanguageTextRead("resolutionMismatch");
}

function PIssueNumberRead(items, group) {
  const segments = PSegmentRead(group).filter(segment => segment.missing);

  if (segments.length > 0) {
    return segments.map(segment => segment.start === segment.end ? String(segment.start) : `${segment.start}–${segment.end}`).join(", ");
  }

  return String(items[0] || "-").replace(/^.*?between\s+/i, "");
}

function PIssueMissingRead(group) {
  const segments = PSegmentRead(group).filter(segment => segment.missing);
  return segments.length || 1;
}

function PIssueToneRead(section) {
  const tag = String(section.LReportTag || "").toLowerCase();

  if (tag === "error") return "Error";
  if (tag === "caution") return "Warning";
  if (tag === "notice") return "Notice";
  return "Error";
}

function PIssueClassRead(tone) {
  if (tone === "Warning") return "PIssueWarning";
  if (tone === "Notice") return "PIssueNotice";
  if (tone === "Ok") return "PIssueOk";
  return "PIssueError";
}
