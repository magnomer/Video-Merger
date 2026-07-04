function PTimelineShow(group) {
  const segments = PSegmentRead(group);
  const issueIndexes = PIssueIndexRead(group);

  return segments.map(segment => {
    const missingClass = segment.missing ? " PSegmentMissing" : "";
    const issueClass = issueIndexes.has(segment.start) ? " PSegmentIssue" : "";
    const label = segment.start === segment.end ? String(segment.start) : `${segment.start}–${segment.end}`;
    const fileLabel = segment.missing ? `Missing (${segment.count} parts)` : LHtmlEscape(segment.name || `${group.LReportName} (${segment.start}).mp4`);
    const duration = segment.missing ? "" : LHtmlEscape(PPreviewDurationRead(group, segment));

    return `
      <div class="PSegment${missingClass}${issueClass}">
        <div><strong>${label}</strong><span>${fileLabel}</span><span>${duration}</span></div>
      </div>
    `;
  }).join("");
}

function PSegmentRead(group) {
  const files = (group.LReportFile || []).slice().sort((a, b) => a.LReportNumber - b.LReportNumber);
  const segments = [];
  let nextNumber = files[0]?.LReportNumber || 1;

  files.forEach(file => {
    const number = Number(file.LReportNumber);
    if (number > nextNumber) {
      segments.push({ start: nextNumber, end: number - 1, count: number - nextNumber, missing: true });
    }
    segments.push({ start: number, end: number, count: 1, missing: false, name: file.LReportName, duration: file.LReportDurationSecond });
    nextNumber = number + 1;
  });

  return segments.length > 0 ? segments : [{ start: 1, end: 1, count: 1, missing: true }];
}

function PIssueIndexRead(group) {
  const indexes = new Set();

  PIssueRead(group)
    .filter(issue => issue.name === "Missing part numbers")
    .forEach(issue => {
      const matches = String(issue.value).match(/\d+/g) || [];
      matches.forEach(value => indexes.add(Number(value)));
    });

  return indexes;
}

function PGroupClipRead(group) {
  return (group.LReportFile || []).length;
}

function PGroupWarningRead(group) {
  return PIssueRead(group).reduce((sum, issue) => sum + (issue.count || 1), 0);
}

function PPreviewNameRead(group) {
  const first = (group.LReportFile || [])[0];
  return first?.LReportName || `${group.LReportName}.mp4`;
}

function PPreviewDurationRead(group, segment) {
  if (segment.count <= 1) {
    return PPreviewTimeRead(Number(segment.duration || 0));
  }
  return group.LReportDuration || "";
}
