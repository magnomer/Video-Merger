function PGroupShow(group, index) {
  const selected = index === PResultStateIndex ? " PGroupSelected" : "";
  const stateClass = PGroupClassRead(group);
  const compatibilityClass = PBadgeClassRead(group.LReportCompatibilityTag);
  const taskClass = PBadgeClassRead(group.LReportTaskTag);

  return `
    <button class="PGroupCard ${stateClass}${selected}" type="button" data-index="${index}" aria-pressed="${index === PResultStateIndex}">
      <div class="PGroupMain">
        <span class="PGroupIcon" aria-hidden="true"><img src="./assets/Video.svg" alt="" /></span>
        <div class="PGroupName">${LHtmlEscape(group.LReportName)}</div>
      </div>
      <div class="PGroupMeta">
        <span>${PIconClipShow()}${PLanguageCountRead(PGroupClipRead(group), "clip", "clips")}</span>
        <span>${PIconDiskShow()}${LHtmlEscape(group.LReportSize)}</span>
        <span>${PIconClockShow()}${LHtmlEscape(group.LReportDuration)}</span>
        <span>${PIconSpeakerShow()}${LHtmlEscape(group.LReportLoudness || "-")}</span>
      </div>
      <div class="PGroupAside">
        <span class="PBadge ${compatibilityClass}">${LHtmlEscape(PLanguageReportTextRead(group.LReportCompatibility))}</span>
        <span class="PBadge ${taskClass}">${LHtmlEscape(PLanguageReportTextRead(group.LReportTask))}</span>
        <span class="PGroupChevron" aria-hidden="true">${PIconChevronShow()}</span>
      </div>
    </button>
  `;
}


function PGroupClassRead(group) {
  const task = group?.LReportTask || "";

  if (task === "Processing") return "PGroupProcessing";
  if (task === "Merged" || task === "Copied") return "PGroupProcessed";
  if (task === "Failed" || task === "Canceled") return "PGroupFailed";

  return "PGroupPending";
}
