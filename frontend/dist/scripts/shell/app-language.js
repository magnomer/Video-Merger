const PLanguageStorageKey = "PLanguageStorageKeyV1";
const PLanguageFallback = "en";
/* Language dictionaries are defined in app-language-map.js. */
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
    "Analysis started.": "analyzing",
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
  PLanguageCheckSet("label.PPathwayCheck:has(#POptionSubfolder)", text.includeSubfolders);
  PLanguageCheckSet("label.PPathwayCheck:has(#POptionMirror)", text.sameAsInput);
  PLanguageControlSet("#PControlAnalysis", "▷", text.analyze);
  PLanguageControlSet("#PControlAssembly", "⌘", text.merge);
  PLanguageControlSet("#PControlTermination", "■", text.stop);
  PLanguageTextSet(".PMeterBlock span", text.progress);
  PLanguageCheckSet("label.PControlCheck:has(#POptionCaution)", text.ignoreCautions);
  PLanguageAttributeSet(".PControlHelp:has(#POptionCaution) .PHelpButton", "aria-label", text.ignoreCautionsHelp);
  PLanguageAttributeSet(".PControlHelp:has(#POptionCaution) .PHelpButton", "data-tooltip", text.ignoreCautionsTip);
  PLanguageCheckSet("label.PControlCheck:has(#POptionWarning)", text.forceMerge);
  PLanguageAttributeSet(".PControlHelp:has(#POptionWarning) .PHelpButton", "aria-label", text.forceMergeHelp);
  PLanguageAttributeSet(".PControlHelp:has(#POptionWarning) .PHelpButton", "data-tooltip", text.forceMergeTip);
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
