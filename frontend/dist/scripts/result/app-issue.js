function PIssueListShow(group) {
  const issues = PIssueRead(group);

  if (issues.length === 0) {
    return `<div class="PIssueRow PIssueOk"><span class="PIssueName"><span class="PIssueDot"></span>${PLanguageTextRead("noBlockingIssues")}</span><span class="PIssueValue">${PLanguageTextRead("ok")}</span></div>`;
  }

  return issues.slice(0, 4).map(issue => PIssueRowShow(issue)).join("");
}

function PIssueRowShow(issue) {
  const details = (issue.details || []).map(value => String(value || "").trim()).filter(Boolean);
  const className = PIssueClassRead(issue.tone);

  if (details.length === 0) {
    return `
      <div class="PIssueRow ${className}">
        <span class="PIssueName"><span class="PIssueDot"></span>${LHtmlEscape(issue.name)}</span>
        <span class="PIssueValue">${LHtmlEscape(issue.value)}</span>
      </div>
    `;
  }

  return `
    <details class="PIssueBox ${className}">
      <summary class="PIssueRow">
        <span class="PIssueName"><span class="PIssueDot"></span>${LHtmlEscape(issue.name)}</span>
        <span class="PIssueValue">${LHtmlEscape(issue.value)}</span>
      </summary>
      <div class="PIssueDetail">
        <ul>${details.map(detail => `<li>${LHtmlEscape(PIssueDetailTextRead(detail))}</li>`).join("")}</ul>
      </div>
    </details>
  `;
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
      issues.push({
        name: PLanguageTextRead("missingPartNumbers"),
        value: PIssueNumberRead(items, group),
        count: PIssueMissingRead(group),
        tone,
        details: PIssueNumberDetailRead(items, group),
      });
      return;
    }

    if (lowerTitle.includes("compatibility")) {
      const value = lowerTitle.includes("caution") ? PLanguageCountRead(items.length, "cautionCount", "cautionCounts") : PLanguageCountRead(items.length, "issue", "issues");
      issues.push({ name: PIssueCompatibilityRead(title), value, count: items.length, tone, details: items });
      return;
    }

    issues.push({ name: PIssueTitleRead(title), value: PIssueSummaryRead(title, items), count: items.length, tone, details: items });
  });

  return issues;
}

function PIssueTitleRead(title) {
  const lowerTitle = String(title || "").toLowerCase();

  if (lowerTitle.includes("notice")) {
    return PLanguageTextRead("notice");
  }

  return title;
}

function PIssueDetailTextRead(value) {
  const text = String(value || "").trim();
  const map = {
    "Frame-rate metadata differs slightly, but packet timing examination and a concat dry run found no merge problem.": "frameRateNoiseNotice",
    "Small frame-rate metadata differences were verified by packet timing and concat dry run.": "frameRateNoiseNotice",
    "Single-file group. It will be copied instead of merged.": "singleFileNotice",
  };

  return map[text] ? PLanguageTextRead(map[text]) : text;
}

function PIssueCompatibilityRead(title) {
  return title.toLowerCase().includes("caution") ? PLanguageTextRead("compatibilityCaution") : PLanguageTextRead("resolutionMismatch");
}

function PIssueNumberRead(items, group) {
  const segments = PIssueMissingSegmentRead(group);

  if (segments.length > 0) {
    return segments.map(segment => segment.start === segment.end ? String(segment.start) : `${segment.start}–${segment.end}`).join(", ");
  }

  return String(items[0] || "-").replace(/^.*?between\s+/i, "");
}

function PIssueSummaryRead(title, items) {
  const count = items.length;
  const lowerTitle = String(title || "").toLowerCase();

  if (lowerTitle.includes("notice")) {
    return PLanguageCountRead(count, "noticeCount", "noticeCounts");
  }

  return PLanguageCountRead(count, "issue", "issues");
}

function PIssueNumberDetailRead(items, group) {
  const segments = PIssueMissingSegmentRead(group);

  if (segments.length === 0) {
    return items;
  }

  return segments.map(segment => segment.start === segment.end ? String(segment.start) : `${segment.start}–${segment.end}`);
}

function PIssueMissingRead(group) {
  const segments = PIssueMissingSegmentRead(group);
  return segments.length || 1;
}

function PIssueMissingSegmentRead(group) {
  const files = (group.LReportFile || []).slice().sort((a, b) => Number(a.LReportNumber) - Number(b.LReportNumber));
  const segments = [];

  if (files.length === 0) {
    return segments;
  }

  const first = Number(files[0].LReportNumber);
  if (Number.isFinite(first) && first !== 0 && first !== 1) {
    segments.push({ start: 1, end: first - 1 });
  }

  for (let index = 1; index < files.length; index += 1) {
    const previous = Number(files[index - 1].LReportNumber);
    const current = Number(files[index].LReportNumber);

    if (Number.isFinite(previous) && Number.isFinite(current) && current > previous + 1) {
      segments.push({ start: previous + 1, end: current - 1 });
    }
  }

  return segments;
}

function PIssueToneRead(section) {
  const tag = String(section.LReportTag || "").toLowerCase();

  if (tag === "error") return "Error";
  if (tag === "caution") return "Caution";
  if (tag === "notice") return "Notice";
  return "Error";
}

function PIssueClassRead(tone) {
  if (tone === "Caution") return "PIssueCaution";
  if (tone === "Warning") return "PIssueWarning";
  if (tone === "Notice") return "PIssueNotice";
  if (tone === "Ok") return "PIssueOk";
  return "PIssueError";
}
