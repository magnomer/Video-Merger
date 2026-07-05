function PInspectorShow(group) {
  if (!group) {
    return `
      <div class="PInspectorTitle">${PLanguageTextRead("inspector")}</div>
      <div class="PInspectorBody">
        <div class="PPlaceholder">${PLanguageTextRead("noGroups")}</div>
      </div>
    `;
  }

  const alertState = PInspectorStateRead(group);
  const alertClass = alertState === "ok" ? " PInspectorOk" : alertState === "caution" ? " PInspectorCaution" : "";
  const alertTitle = alertState === "ok" ? PLanguageTextRead("canMerge") : alertState === "caution" ? PLanguageTextRead("canMergeWithCaution") : PLanguageTextRead("cannotMerge");
  const alertText = alertState === "ok" ? PLanguageTextRead("safeText") : alertState === "caution" ? PLanguageTextRead("cautionText") : PLanguageTextRead("unsafeText");

  return `
    <div class="PInspectorTitle">${PLanguageTextRead("inspector")}</div>
    <div class="PInspectorBody">
      <div class="PInspectorAlert${alertClass}">
        ${PInspectorIconShow(alertState)}
        <div><strong>${alertTitle}</strong><p>${alertText}</p></div>
      </div>
      <div class="PIssueList">${PIssueListShow(group)}</div>
      <div class="PSectionTitle">${PLanguageTextRead("summary")}</div>
      ${PMeasureShow(PMeasureRead(group))}
      <div class="PSectionTitle">${PInspectorOutputTitleRead(group)}</div>
      <div class="POutputLine">
        <div class="POutputRow">${PIconFolderShow()}<span>${LHtmlEscape(group.LReportOutputText || "-")}</span></div>
        <button class="POutputCopy" type="button" aria-label="${PLanguageTextRead("copyPlannedOutput")}">${PIconCopyShow()}</button>
      </div>
    </div>
  `;
}

function PInspectorOutputTitleRead(group) {
  if (!group || !group.LReportOutputTitle) {
    return PLanguageTextRead("plannedOutput");
  }

  if (group.LReportOutputTitle === "Planned output") {
    return PLanguageTextRead("plannedOutput");
  }

  if (group.LReportOutputTitle === "Output") {
    return PLanguageTextRead("output");
  }

  return LHtmlEscape(group.LReportOutputTitle);
}

function PInspectorStateRead(group) {
  const tag = String(group?.LReportCompatibilityTag || "").toLowerCase();

  if (tag === "ok" || tag === "notice") {
    return "ok";
  }

  if (tag === "caution") {
    return "caution";
  }

  return "error";
}
