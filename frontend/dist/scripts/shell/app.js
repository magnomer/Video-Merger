const PSourceText = document.getElementById("PSourceText");
const POutputText = document.getElementById("POutputText");
const POptionSubfolder = document.getElementById("POptionSubfolder");
const POptionMirror = document.getElementById("POptionMirror");
const POptionField = document.getElementById("POptionField");
const POptionCaution = document.getElementById("POptionCaution");
const POptionWarning = document.getElementById("POptionWarning");

const PSourceExplorer = document.getElementById("PSourceExplorer");
const PSourceFolder = document.getElementById("PSourceFolder");
const PSourceIcon = document.getElementById("PSourceIcon");
const POutputFolder = document.getElementById("POutputFolder");
const POutputExplorer = document.getElementById("POutputExplorer");
const PFrameInfo = document.getElementById("PFrameInfo");
const PSAboutBackdrop = document.getElementById("PSAboutBackdrop");
const PSAboutButton = document.getElementById("PSAboutButton");
const PSAboutIcon = document.getElementById("PSAboutIcon");

const PFrameSetting = document.getElementById("PFrameSetting");
const PSSettingBackdrop = document.getElementById("PSSettingBackdrop");
const PSSettingButton = document.getElementById("PSSettingButton");
const PSSettingIcon = document.getElementById("PSSettingIcon");
const PSSettingStyle = document.getElementById("PSSettingStyle");
const PSSettingCustom = document.getElementById("PSSettingCustom");
const PSSettingPattern = document.getElementById("PSSettingPattern");
const PSSettingUnnumbered = document.getElementById("PSSettingUnnumbered");
const PSSettingSample = document.getElementById("PSSettingSample");

const PControlAnalysis = document.getElementById("PControlAnalysis");
const PControlAssembly = document.getElementById("PControlAssembly");
const PControlTermination = document.getElementById("PControlTermination");

const PMeterFill = document.getElementById("PMeterFill");
const PMeterCount = document.getElementById("PMeterCount");
const PMeterText = document.getElementById("PMeterText");
const PMeterPercent = document.getElementById("PMeterPercent");
const PResult = document.getElementById("PResult");
const PStatusMessage = document.getElementById("PStatusMessage");
const PStatusGroup = document.getElementById("PStatusGroup");
const PStatusClip = document.getElementById("PStatusClip");
const PStatusSize = document.getElementById("PStatusSize");
const PStatusDuration = document.getElementById("PStatusDuration");
const PStatusWarning = document.getElementById("PStatusWarning");
const PFrameVersion = document.getElementById("PFrameVersion");
const PSAboutVersion = document.getElementById("PSAboutVersion");
const PSAboutAuthor = document.getElementById("PSAboutAuthor");

function PFrameStart() {
  PResultEventStart();
  PControlStart();
  PSAboutStart();
  PSSettingStart();
  POptionStart();
  LWindowStateStart();
  PLanguageStart();
  PResultShow(null);
  PFrameProfileLoad();
}

function PResultEventStart() {
  if (!window.runtime?.EventsOn) {
    return;
  }

  window.runtime.EventsOn("LReportEvent", report => {
    PResultReportSet(report);
  });
}

PFrameStart();
