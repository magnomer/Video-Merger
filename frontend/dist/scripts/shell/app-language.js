const PLanguageStorageKey = "PLanguageStorageKeyV1";
const PLanguageFallback = "en";
const PLanguageMap = {
  en: {
    languageName: "English", languageShort: "EN",
    languageLabel: "Language", languageEnglish: "English",
    languageKorean: "Korean", title: "Video Merger",
    actions: "Application actions", info: "Info",
    settings: "Settings", windowControls: "Window controls",
    minimize: "Minimize", maximize: "Maximize",
    close: "Close",
    mergeSetup: "Merge setup",
    source: "Source",
    sourcePlaceholder: "Select files/folders or type paths manually, one per line",
    sourceFolder: "Select input folder",
    sourceOpen: "Open source folder",
    output: "Output",
    outputPlaceholder: "Select output folder or type manually",
    outputFolder: "Select output folder",
    outputOpen: "Open output folder",
    browse: "Browse",
    analyzing: "Analyzing...",
    startingMerge: "Starting merge...",
    stopping: "Stopping...",
    sourceOpenError: "Source folder open error.",
    sourcePickerError: "Folder picker error.",
    outputPickerError: "Output folder picker error.",
    outputOpenError: "Output folder open error.",
    analyzeError: "Analyze error.",
    mergeError: "Merge error.",
    stopError: "Stop error.",
    previewOpenError: "This clip cannot be opened by the preview player.",
    previewCompatibilityPreparing: "Direct preview failed. Creating a temporary preview copy...",
    compatible: "Compatible", incompatible: "Incompatible",
    caution: "Caution", notice: "Notice",
    processing: "Processing", canceled: "Canceled",
    notMerged: "Not merged", copied: "Copied",
    merged: "Merged", failed: "Failed",
    video: "Video", audio: "Audio",
    audioLoudness: "Audio loudness (LUFS)",
    codec: "Codec",
    streams: "Streams",
    totalDurationEstimated: "Total duration (estimated)",
    totalSizeEstimated: "Total size (estimated)",
    actualSize: "Actual size",
    actualDuration: "Actual duration",
    actualLufs: "Actual LUFS",
    estimatedSize: "Estimated size",
    estimatedDuration: "Estimated duration",
    suffix: "Suffix",
    suffixPlaceholder: "optional suffix, e.g. _merged",
    suffixHelp: "Suffix help",
    suffixTip: "Append a suffix before the output file extension.",
    includeSubfolders: "Include subfolders",
    sameAsInput: "Same as input",
    analyze: "Analyze", merge: "Merge",
    stop: "Stop", progress: "Progress",
    ignoreCautions: "Ignore cautions",
    ignoreCautionsHelp: "Ignore cautions help",
    ignoreCautionsTip: "Not strictly identical, but probably compatible.",
    forceMerge: "Force merge",
    forceMergeHelp: "Force merge despite warnings help",
    forceMergeTip: "Attempt merging even when compatibility checks report problems.",
    statusSummary: "Status summary",
    ready: "Ready", group: "group",
    groups: "groups", clip: "clip",
    clips: "clips", warning: "warning",
    warnings: "warnings",
    analysisCompleted: "Analysis completed",
    inspector: "Inspector",
    summary: "Summary",
    plannedOutput: "Planned output",
    copyPlannedOutput: "Copy planned output",
    copiedPlannedOutput: "Copied planned output",
    canMerge: "Can merge safely",
    cannotMerge: "Cannot merge safely",
    safeText: "This group passed blocking compatibility checks.",
    unsafeText: "This group has blocking issues.",
    noBlockingIssues: "No blocking issues",
    ok: "OK", issue: "issue",
    issues: "issues",
    missingPartNumbers: "Missing part numbers",
    compatibilityCaution: "Compatibility caution",
    resolutionMismatch: "Resolution mismatch",
    error: "Error", noGroups: "No matching groups found.",
    groupsTitle: "Groups",
    outputRefreshError: "Output refresh error",
    preview: "Preview",
    playback: "Playback",
    playPreview: "Play virtual preview",
    nextClip: "Next clip",
    previewPosition: "Virtual preview position",
    previewVolume: "Preview volume",
    fullscreenPreview: "Fullscreen preview",
    timeline: "Timeline",
    availableClip: "Available clip",
    missingParts: "Missing parts",
    compatibilityIssue: "Compatibility issue",
    resizeHeight: "Resize Inspector and Preview height",
    resizeWidth: "Resize Inspector and Preview width",
    aboutClose: "Close About dialog",
    aboutLead: "Merge numbered video parts with FFmpeg.",
    aboutVersion: "Version",
    aboutAuthor: "Author",
    aboutEngine: "Engine",
    aboutChecks: "Checks",
    aboutText: "Video Merger is designed for batch merging numbered media files while preserving video and audio streams.",
    aboutButton: "Close",
    settingClose: "Close Settings dialog",
    settingHeading: "Part detection",
    settingLead: "Choose how numbered parts of one group are recognized from filenames.",
    settingStyle: "Number style",
    settingCustom: "Custom (Regular expression)",
    settingPatternPlaceholder: "^(.+?)\\s*\\((\\d+)\\)$  (group 1 = name, group 2 = number)",
    settingUnnumbered: "Include files without a number",
    settingButton: "Close",
    settingSample: "Example:",
    settingSampleNoNumber: "no number",
    settingSampleNone: "not detected",
    settingSampleEmpty: "Enter a custom pattern.",
    settingSampleInvalid: "Invalid regular expression.",
    settingEngineHeading: "Engine",
    settingEngineLead: "Choose FFmpeg and where Video Merger stores temporary preview files.",
    settingFFmpeg: "FFmpeg path",
    settingFFmpegPlaceholder: "Use Windows PATH",
    settingTemporary: "Temporary files folder",
    settingTemporaryPlaceholder: "Use Windows temp folder",
    settingClean: "Clean up",
    settingCleanWorking: "Cleaning up...",
    settingCleanDone: "Removed {count} temporary files from {path}",
    settingCleanError: "Clean-up error.",
    settingFFmpegError: "FFmpeg picker error.",
    settingTemporaryError: "Temporary folder picker error.",
    unknown: "Unknown",
  },
  ko: {
    languageName: "한국어", languageShort: "KO",
    languageLabel: "언어", languageEnglish: "영어",
    languageKorean: "한국어", title: "비디오 병합기",
    actions: "프로그램 작업", info: "정보",
    settings: "설정", windowControls: "창 제어",
    minimize: "최소화", maximize: "최대화",
    close: "닫기",
    mergeSetup: "병합 설정",
    source: "원본",
    sourcePlaceholder: "파일/폴더를 선택하거나 경로를 한 줄에 하나씩 입력",
    sourceFolder: "입력 폴더 선택",
    sourceOpen: "소스 폴더 열기",
    output: "저장",
    outputPlaceholder: "출력 폴더를 선택하거나 직접 입력",
    outputFolder: "출력 폴더 선택",
    outputOpen: "출력 폴더 열기",
    browse: "선택",
    analyzing: "분석 중...",
    startingMerge: "병합 시작 중...",
    stopping: "중지 중...",
    sourceOpenError: "소스 폴더 열기 오류.",
    sourcePickerError: "폴더 선택 오류.",
    outputPickerError: "출력 폴더 선택 오류.",
    outputOpenError: "출력 폴더 열기 오류.",
    analyzeError: "분석 오류.",
    mergeError: "병합 오류.",
    stopError: "중지 오류.",
    previewOpenError: "이 클립은 미리보기 플레이어에서 열 수 없습니다.",
    previewCompatibilityPreparing: "직접 미리보기에 실패했습니다. 임시 미리보기 복사본을 만드는 중입니다...",
    compatible: "호환 가능", incompatible: "호환 불가",
    caution: "주의", notice: "알림",
    processing: "처리 중", canceled: "취소됨",
    notMerged: "미병합", copied: "복사됨",
    merged: "병합됨", failed: "실패",
    video: "비디오", audio: "오디오",
    audioLoudness: "오디오 음량 (LUFS)",
    codec: "코덱",
    streams: "스트림",
    totalDurationEstimated: "총 재생 시간 (예상)",
    totalSizeEstimated: "총 크기 (예상)",
    actualSize: "실제 크기",
    actualDuration: "실제 재생 시간",
    actualLufs: "실제 LUFS",
    estimatedSize: "예상 크기",
    estimatedDuration: "예상 재생 시간",
    suffix: "접미사",
    suffixPlaceholder: "선택 접미사, 예: _merged",
    suffixHelp: "접미사 도움말",
    suffixTip: "출력 파일 확장자 앞에 접미사를 붙입니다.",
    includeSubfolders: "하위 폴더 포함",
    sameAsInput: "원본과 같은 폴더",
    analyze: "분석", merge: "병합",
    stop: "중지", progress: "진행률",
    ignoreCautions: "주의 무시",
    ignoreCautionsHelp: "주의 무시 도움말",
    ignoreCautionsTip: "완전히 동일하지는 않지만 대체로 호환됩니다.",
    forceMerge: "강제 병합",
    forceMergeHelp: "경고가 있어도 강제 병합",
    forceMergeTip: "호환성 검사에서 문제가 보고되어도 병합을 시도합니다.",
    statusSummary: "상태 요약",
    ready: "준비됨", group: "그룹",
    groups: "그룹", clip: "클립",
    clips: "클립", warning: "경고",
    warnings: "경고",
    analysisCompleted: "분석 완료",
    inspector: "검사기",
    summary: "요약",
    plannedOutput: "예상 출력",
    copyPlannedOutput: "예상 출력 복사",
    copiedPlannedOutput: "예상 출력 복사됨",
    canMerge: "안전하게 병합 가능",
    cannotMerge: "안전하게 병합 불가",
    safeText: "이 그룹은 차단 호환성 검사를 통과했습니다.",
    unsafeText: "이 그룹에는 차단 문제가 있습니다.",
    noBlockingIssues: "차단 문제 없음",
    ok: "정상", issue: "문제",
    issues: "문제",
    missingPartNumbers: "누락된 파트 번호",
    compatibilityCaution: "호환성 주의",
    resolutionMismatch: "해상도 불일치",
    error: "오류", noGroups: "일치하는 그룹이 없습니다.",
    groupsTitle: "그룹",
    outputRefreshError: "출력 새로고침 오류",
    preview: "미리보기",
    playback: "재생",
    playPreview: "가상 미리보기 재생",
    nextClip: "다음 클립",
    previewPosition: "가상 미리보기 위치",
    previewVolume: "미리보기 볼륨",
    fullscreenPreview: "전체 화면 미리보기",
    timeline: "타임라인",
    availableClip: "사용 가능한 클립",
    missingParts: "누락된 파트",
    compatibilityIssue: "호환성 문제",
    resizeHeight: "검사기와 미리보기 높이 조정",
    resizeWidth: "검사기와 미리보기 너비 조정",
    aboutClose: "정보 창 닫기",
    aboutLead: "번호가 붙은 비디오 파트를 FFmpeg로 병합합니다.",
    aboutVersion: "버전",
    aboutAuthor: "제작자",
    aboutEngine: "엔진",
    aboutChecks: "검사",
    aboutText: "비디오 병합기는 비디오와 오디오 스트림을 유지하면서 번호가 붙은 미디어 파일을 일괄 병합하도록 설계되었습니다.",
    aboutButton: "닫기",
    settingClose: "설정 창 닫기",
    settingHeading: "파트 감지",
    settingLead: "파일 이름에서 한 그룹의 번호가 붙은 파트를 인식하는 방식을 선택합니다.",
    settingStyle: "번호 형식",
    settingCustom: "사용자 지정 (정규식)",
    settingPatternPlaceholder: "^(.+?)\\s*\\((\\d+)\\)$  (그룹 1 = 이름, 그룹 2 = 번호)",
    settingUnnumbered: "번호가 없는 파일 포함",
    settingButton: "닫기",
    settingSample: "예:",
    settingSampleNoNumber: "번호 없음",
    settingSampleNone: "감지되지 않음",
    settingSampleEmpty: "사용자 지정 패턴을 입력하세요.",
    settingSampleInvalid: "잘못된 정규식.",
    settingEngineHeading: "엔진",
    settingEngineLead: "FFmpeg 위치와 비디오 병합기가 임시 미리보기 파일을 저장할 위치를 선택합니다.",
    settingFFmpeg: "FFmpeg 경로",
    settingFFmpegPlaceholder: "윈도우 PATH 사용",
    settingTemporary: "임시 파일 폴더",
    settingTemporaryPlaceholder: "윈도우 임시 폴더 사용",
    settingClean: "정리",
    settingCleanWorking: "정리 중...",
    settingCleanDone: "{path}에서 임시 파일 {count}개를 제거했습니다",
    settingCleanError: "정리 오류.",
    settingFFmpegError: "FFmpeg 선택 오류.",
    settingTemporaryError: "임시 폴더 선택 오류.",
    unknown: "알 수 없음",
  },
};
function PLanguageRead() {
  const saved = localStorage.getItem(PLanguageStorageKey);
  return PLanguageMap[saved] ? saved : PLanguageFallback;
}
function PLanguageTextRead(key) {
  const language = PLanguageRead();
  return PLanguageMap[language]?.[key] || PLanguageMap[PLanguageFallback][key] || key;
}
function PLanguageReportTextRead(value) {
  const text = String(value || "");
  const map = {
    "Analysis completed.": "analysisCompleted",
    "Compatible": "compatible",
    "Incompatible": "incompatible",
    "Caution": "caution",
    "Notice": "notice",
    "Processing": "processing",
    "Canceled": "canceled",
    "Not merged": "notMerged",
    "Copied": "copied",
    "Merged": "merged",
    "Failed": "failed",
    "Actual size": "actualSize",
    "Actual duration": "actualDuration",
    "Actual LUFS": "actualLufs",
    "Estimated size": "estimatedSize",
    "Estimated duration": "estimatedDuration",
  };
  return map[text] ? PLanguageTextRead(map[text]) : text;
}
function PLanguageCountRead(count, singleKey, pluralKey) {
  return `${count} ${PLanguageTextRead(count === 1 ? singleKey : pluralKey)}`;
}
function PLanguageTextSet(selector, text) {
  const element = document.querySelector(selector);
  if (element) element.textContent = text;
}
function PLanguageAttributeSet(selector, attribute, text) {
  const element = document.querySelector(selector);
  if (element) element.setAttribute(attribute, text);
}
function PLanguageCheckSet(selector, text) {
  const label = document.querySelector(selector);
  if (!label) return;
  Array.from(label.childNodes).forEach(node => {
    if (node.nodeType === Node.TEXT_NODE) node.remove();
  });
  label.appendChild(document.createTextNode(text));
}
function PLanguageControlSet(selector, icon, text) {
  const button = document.querySelector(selector);
  if (!button) return;
  button.innerHTML = `<span class="PControlIcon">${icon}</span>${text}`;
}
function PLanguageStart() {
  const root = document.getElementById("PLanguage");
  const button = document.getElementById("PLanguageButton");
  const menu = document.getElementById("PLanguageMenu");
  if (!root || !button || !menu) return;
  button.addEventListener("click", event => {
    event.stopPropagation();
    PLanguageToggleSet(root, button, menu);
  });
  menu.querySelectorAll(".PLanguageOption").forEach(option => {
    option.addEventListener("click", () => {
      PLanguageSet(option.dataset.language || PLanguageFallback);
      PLanguageHide(root, button, menu);
    });
  });
  document.addEventListener("click", event => {
    if (!root.contains(event.target)) PLanguageHide(root, button, menu);
  });
  document.addEventListener("keydown", event => {
    if (event.key === "Escape") PLanguageHide(root, button, menu);
  });
  PLanguageApplySet();
}
function PLanguageToggleSet(root, button, menu) {
  if (menu.hidden) {
    PLanguageShow(root, button, menu);
    return;
  }
  PLanguageHide(root, button, menu);
}
function PLanguageShow(root, button, menu) {
  root.classList.add("PLanguageOpen");
  button.setAttribute("aria-expanded", "true");
  menu.hidden = false;
}
function PLanguageHide(root, button, menu) {
  root.classList.remove("PLanguageOpen");
  button.setAttribute("aria-expanded", "false");
  menu.hidden = true;
}
function PLanguageSet(language) {
  if (!PLanguageMap[language]) return;
  localStorage.setItem(PLanguageStorageKey, language);
  PLanguageApplySet();
}
function PLanguageApplySet() {
  const language = PLanguageRead();
  const text = PLanguageMap[language];
  document.documentElement.lang = language === "ko" ? "ko" : "en";
  document.title = text.title;
  PLanguageTextSet(".PFrameHeader h1", text.title);
  PLanguageTextSet("#PLanguageText", text.languageName);
  PLanguageTextSet("#PLanguageBadge", text.languageShort);
  PLanguageAttributeSet("#PLanguageButton", "aria-label", text.languageLabel);
  PLanguageAttributeSet("#PLanguageButton", "title", text.languageLabel);
  PLanguageAttributeSet(".PFrameAction", "aria-label", text.actions);
  PLanguageAttributeSet("#PFrameInfo", "aria-label", text.info);
  PLanguageAttributeSet("#PFrameInfo", "title", text.info);
  PLanguageAttributeSet(".PFrameAction > button:not(#PFrameInfo)", "aria-label", text.settings);
  PLanguageAttributeSet(".PFrameAction > button:not(#PFrameInfo)", "title", text.settings);
  PLanguageAttributeSet(".PFrameSystem", "aria-label", text.windowControls);
  PLanguageAttributeSet("#PFrameDash", "aria-label", text.minimize);
  PLanguageAttributeSet("#PFrameBox", "aria-label", text.maximize);
  PLanguageAttributeSet("#PFrameExit", "aria-label", text.close);
  PLanguageAttributeSet(".PEntry", "aria-label", text.mergeSetup);
  PLanguageTextSet("label[for='PSourceText']", text.source);
  PLanguageTextSet("#PSourceFolder span", text.browse);
  PLanguageAttributeSet("#PSourceText", "placeholder", text.sourcePlaceholder);
  PLanguageAttributeSet("#PSourceFolder", "aria-label", text.sourceFolder);
  PLanguageAttributeSet("#PSourceFolder", "title", text.sourceFolder);
  PLanguageAttributeSet("#PSourceExplorer", "aria-label", text.sourceOpen);
  PLanguageAttributeSet("#PSourceExplorer", "title", text.sourceOpen);
  PLanguageTextSet("label[for='POutputText']", text.output);
  PLanguageTextSet("#POutputFolder span", text.browse);
  PLanguageAttributeSet("#POutputText", "placeholder", text.outputPlaceholder);
  PLanguageAttributeSet("#POutputFolder", "aria-label", text.outputFolder);
  PLanguageAttributeSet("#POutputFolder", "title", text.outputFolder);
  PLanguageAttributeSet("#POutputExplorer", "aria-label", text.outputOpen);
  PLanguageAttributeSet("#POutputExplorer", "title", text.outputOpen);
  PLanguageTextSet("label[for='POptionField']", text.suffix);
  PLanguageAttributeSet("#POptionField", "placeholder", text.suffixPlaceholder);
  PLanguageAttributeSet(".POptionBox .PHelpButton", "aria-label", text.suffixHelp);
  PLanguageAttributeSet(".POptionBox .PHelpButton", "data-tooltip", text.suffixTip);
  PLanguageCheckSet("label.POptionCheck:has(#POptionSubfolder)", text.includeSubfolders);
  PLanguageCheckSet("label.POptionCheck:has(#POptionMirror)", text.sameAsInput);
  PLanguageControlSet("#PControlAnalysis", "▷", text.analyze);
  PLanguageControlSet("#PControlAssembly", "⌘", text.merge);
  PLanguageControlSet("#PControlTermination", "■", text.stop);
  PLanguageTextSet(".PMeterBlock span", text.progress);
  PLanguageCheckSet("label.POptionCheck:has(#POptionCaution)", text.ignoreCautions);
  PLanguageAttributeSet(".POptionHelp:has(#POptionCaution) .PHelpButton", "aria-label", text.ignoreCautionsHelp);
  PLanguageAttributeSet(".POptionHelp:has(#POptionCaution) .PHelpButton", "data-tooltip", text.ignoreCautionsTip);
  PLanguageCheckSet("label.POptionCheck:has(#POptionWarning)", text.forceMerge);
  PLanguageAttributeSet(".POptionHelp:has(#POptionWarning) .PHelpButton", "aria-label", text.forceMergeHelp);
  PLanguageAttributeSet(".POptionHelp:has(#POptionWarning) .PHelpButton", "data-tooltip", text.forceMergeTip);
  PLanguageAttributeSet(".PStatus", "aria-label", text.statusSummary);
  PLanguageTextSet(".PLanguageOption[data-language='en'] span", text.languageEnglish);
  PLanguageTextSet(".PLanguageOption[data-language='ko'] span", text.languageKorean);
  document.querySelectorAll(".PLanguageOption").forEach(option => {
    const active = option.dataset.language === language;
    option.classList.toggle("PLanguageActive", active);
    option.setAttribute("aria-checked", active ? "true" : "false");
  });
  PLanguageTextSet("#PSAboutTitle", text.title);
  PLanguageAttributeSet("#PSAboutIcon", "aria-label", text.aboutClose);
  PLanguageTextSet(".PSAboutBody > p:first-child", text.aboutLead);
  PLanguageTextSet(".PSAboutFact div:nth-child(1) dt", text.aboutVersion);
  PLanguageTextSet(".PSAboutFact div:nth-child(2) dt", text.aboutAuthor);
  PLanguageTextSet(".PSAboutFact div:nth-child(3) dt", text.aboutEngine);
  PLanguageTextSet(".PSAboutFact div:nth-child(4) dt", text.aboutChecks);
  PLanguageTextSet(".PSAboutBody > p:last-child", text.aboutText);
  PLanguageTextSet("#PSAboutButton", text.aboutButton);
  PLanguageTextSet("#PSSettingTitle", text.settings);
  PLanguageAttributeSet("#PSSettingIcon", "aria-label", text.settingClose);
  PLanguageTextSet(".PSSettingEngineHeading", text.settingEngineHeading);
  PLanguageTextSet(".PSSettingEngineLead", text.settingEngineLead);
  PLanguageTextSet(".PSSettingPartHeading", text.settingHeading);
  PLanguageTextSet(".PSSettingPartLead", text.settingLead);
  PLanguageTextSet("label[for='PSSettingStyle']", text.settingStyle);
  PLanguageCheckSet("label.PSSettingCheck:has(#PSSettingCustom)", text.settingCustom);
  PLanguageCheckSet("label.PSSettingCheck:has(#PSSettingUnnumbered)", text.settingUnnumbered);
  PLanguageAttributeSet("#PSSettingPattern", "placeholder", text.settingPatternPlaceholder);
  PLanguageTextSet("label[for='PSSettingFFmpeg']", text.settingFFmpeg);
  PLanguageAttributeSet("#PSSettingFFmpeg", "placeholder", text.settingFFmpegPlaceholder);
  PLanguageTextSet("label[for='PSSettingTemporary']", text.settingTemporary);
  PLanguageAttributeSet("#PSSettingTemporary", "placeholder", text.settingTemporaryPlaceholder);
  PLanguageTextSet("#PSSettingFFmpegButton", text.browse);
  PLanguageTextSet("#PSSettingTemporaryButton", text.browse);
  PLanguageTextSet("#PSSettingCleanButton", text.settingClean);
  PLanguageTextSet("#PSSettingButton", text.settingButton);
  if (typeof PSSettingSampleSet === "function") {
    PSSettingSampleSet();
  }
  if (typeof PResultRenderSet === "function") {
    PResultRenderSet();
  }
  if (typeof PResultStateReport !== "undefined" && PResultStateReport) {
    PStatusReportSet(PResultStateReport);
  } else if (typeof PMeterReset === "function") {
    PMeterReset();
  }
}
