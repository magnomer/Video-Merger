function PStatusReportSet(report) {
  if (typeof PStatusMessage === "undefined") {
    return;
  }

  const groups = report?.LReportGroup || [];
  const clips = groups.reduce((sum, group) => sum + PGroupClipRead(group), 0);
  const warnings = groups.reduce((sum, group) => sum + PGroupWarningRead(group), 0);
  const first = groups[0] || {};

  PStatusMessage.innerHTML = `<span class="PStatusDot">✓</span>${LHtmlEscape(PLanguageReportTextRead(report?.LTaskMessage) || PLanguageTextRead("analysisCompleted"))}`;
  PStatusGroup.textContent = PLanguageCountRead(groups.length, "group", "groups");
  PStatusClip.textContent = PLanguageCountRead(clips, "clip", "clips");
  PStatusSize.textContent = first.LReportSize || "-";
  PStatusDuration.textContent = first.LReportDuration || "-";
  PStatusWarning.textContent = PLanguageCountRead(warnings, "warning", "warnings");
}
