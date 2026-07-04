function PPreviewShow(group) {
  if (!group) {
    return `
      <div class="PPreviewTop">
        <h3>${PLanguageTextRead("preview")}</h3>
        <span id="PPreviewTag" class="PPreviewTag">0 / 0 ${PLanguageTextRead("playback")}</span>
      </div>
      <div class="PPreviewStage PPreviewStageEmpty">
        <video id="PPreviewVideo" class="PPreviewVideo" preload="metadata" playsinline></video>
        <div class="PPreviewControl">
          <button id="PPreviewPlay" class="PPreviewPlay" type="button" disabled aria-label="${PLanguageTextRead("playPreview")}">${PPreviewIconShow("Play")}</button>
          <button id="PPreviewNext" class="PPreviewSkip" type="button" disabled aria-label="${PLanguageTextRead("nextClip")}">${PPreviewIconShow("Next")}</button>
          <span id="PPreviewNow" class="PPreviewTime">00:00:00</span>
          <input id="PPreviewSlider" class="PPreviewSlider" type="range" min="0" max="0" step="0.01" value="0" disabled aria-label="${PLanguageTextRead("previewPosition")}" />
          <span id="PPreviewTotal" class="PPreviewTime">00:00:00</span>
          <span class="PPreviewVolume">${PPreviewIconShow("Volume")}<input id="PPreviewVolume" class="PPreviewVolumeSlider" type="range" min="0" max="100" step="1" value="72" disabled aria-label="${PLanguageTextRead("previewVolume")}" /></span>
          <button id="PPreviewFull" class="PPreviewFull" type="button" disabled aria-label="${PLanguageTextRead("fullscreenPreview")}">${PPreviewIconShow("Full")}</button>
        </div>
      </div>
      <div class="PTimeline">
        <div class="PTimelineHead"><span>${PLanguageTextRead("timeline")}</span></div>
        <div class="PTimelineScale"><span>00:00:00</span><span>00:00:00</span></div>
        <div class="PTimelineTrack"><div class="PPlaceholder">${PLanguageTextRead("noGroups")}</div></div>
      </div>
    `;
  }

  return `
    <div class="PPreviewTop">
      <h3>${PLanguageTextRead("preview")}: ${LHtmlEscape(PPreviewNameRead(group))}</h3>
      <span id="PPreviewTag" class="PPreviewTag">${PLanguageTextRead("playback")} ⓘ</span>
    </div>
    <div class="PPreviewStage">
      <video id="PPreviewVideo" class="PPreviewVideo" preload="metadata" playsinline></video>
      <div class="PPreviewControl">
        <button id="PPreviewPlay" class="PPreviewPlay" type="button" aria-label="${PLanguageTextRead("playPreview")}">${PPreviewIconShow("Play")}</button>
        <button id="PPreviewNext" class="PPreviewSkip" type="button" aria-label="${PLanguageTextRead("nextClip")}">${PPreviewIconShow("Next")}</button>
        <span id="PPreviewNow" class="PPreviewTime">00:00:00</span>
        <input id="PPreviewSlider" class="PPreviewSlider" type="range" min="0" max="0" step="0.01" value="0" aria-label="${PLanguageTextRead("previewPosition")}" />
        <span id="PPreviewTotal" class="PPreviewTime">${LHtmlEscape(group.LReportDuration || "00:00")}</span>
        <span class="PPreviewVolume">${PPreviewIconShow("Volume")}<input id="PPreviewVolume" class="PPreviewVolumeSlider" type="range" min="0" max="100" step="1" value="72" aria-label="${PLanguageTextRead("previewVolume")}" /></span>
        <button id="PPreviewFull" class="PPreviewFull" type="button" aria-label="${PLanguageTextRead("fullscreenPreview")}">${PPreviewIconShow("Full")}</button>
      </div>
    </div>
    <div class="PTimeline">
      <div class="PTimelineHead"><span>${PLanguageTextRead("timeline")}</span></div>
      <div class="PTimelineScale"><span>00:00:00</span><span>${LHtmlEscape(group.LReportDuration || "-")}</span></div>
      <div class="PTimelineTrack">${PTimelineShow(group)}</div>
      <div class="PLegend">
        <span><i class="PLegendBox"></i>${PLanguageTextRead("availableClip")}</span>
        <span><i class="PLegendBox PLegendMissing"></i>${PLanguageTextRead("missingParts")}</span>
        <span><i class="PLegendIssue">▲</i>${PLanguageTextRead("compatibilityIssue")}</span>
      </div>
    </div>
  `;
}
