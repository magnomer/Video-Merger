function PControlStart() {
  PSourceExplorer.addEventListener("click", PSourceExplorerOpen);
  PSourceFolder.addEventListener("click", PSourceFolderOpen);
  PSourceText.addEventListener("input", () => {
    PControlSourceChangeStop();
    PSourceIconSet();
  });
  POutputFolder.addEventListener("click", POutputFolderOpen);
  POutputExplorer.addEventListener("click", POutputExplorerOpen);
  POptionMirror.addEventListener("change", POutputStateSet);
  PControlAnalysis.addEventListener("click", PControlInspectionStart);
  PControlAssembly.addEventListener("click", PControlMergerStart);
  PControlTermination.addEventListener("click", PControlTaskStop);
}

async function PSourceExplorerOpen() {
  try {
    await window.go.bridge.LProgram.LDirectorySourceOpen(PSourceText.value);
  } catch (error) {
    PResultErrorShow(PLanguageTextRead("sourceOpenError"), error);
  }
}

async function PSourceFolderOpen() {
  try {
    const selectedFolder = await window.go.bridge.LProgram.LPickerFolderOpen();

    if (!selectedFolder) {
      return;
    }

    await PControlSourceChangeStop();
    PSourceText.value = selectedFolder;
    PSourceIconFolderSet();
    POptionSave();
  } catch (error) {
    PResultErrorShow(PLanguageTextRead("sourcePickerError"), error);
  }
}

function PSourceIconSet() {
  const sourceValue = PSourceText.value.trim();
  const sourceLines = sourceValue.split(/\r?\n/).filter(Boolean);
  const sourcePattern = /\.(mp4|mov|mkv|m4v)$/i;

  if (sourceLines.length > 1 || sourceLines.some(sourceLine => sourcePattern.test(sourceLine))) {
    PSourceIconFileSet();
    return;
  }

  PSourceIconFolderSet();
}

function PSourceIconFileSet() {
  PSourceIcon.src = "assets/File.svg";
}

function PSourceIconFolderSet() {
  PSourceIcon.src = "assets/Folder.svg";
}

async function POutputFolderOpen() {
  try {
    const selectedFolder = await window.go.bridge.LProgram.LPickerFolderOpen();

    if (!selectedFolder) {
      return;
    }

    POutputText.value = selectedFolder;
    POptionMirror.checked = false;
    POutputStateSet();
    POptionSave();
  } catch (error) {
    PResultErrorShow(PLanguageTextRead("outputPickerError"), error);
  }
}

async function POutputExplorerOpen() {
  try {
    await window.go.bridge.LProgram.LDirectoryOutputOpen(POutputText.value, PSourceText.value, POptionMirror.checked);
  } catch (error) {
    PResultErrorShow(PLanguageTextRead("outputOpenError"), error);
  }
}

async function PControlSourceChangeStop() {
  if (!PControlBusyState) {
    return;
  }

  try {
    PMeterText.textContent = PLanguageTextRead("stopping");
    await window.go.bridge.LProgram.LTaskStop();
  } catch (_) {}
}

async function PControlInspectionStart() {
  const options = POptionRead();

  POptionSave();
  PMeterReset();
  PControlStateSet(true, true);
  PResultMergeUpdateSet(false);
  PResultAnalysisStart(options);

  try {
    const report = await window.go.bridge.LProgram.LInspectionStart(options);
    PResultReportSet(report);
  } catch (error) {
    PResultErrorShow(PLanguageTextRead("analyzeError"), error);
  } finally {
    PControlStateSet(false, false);
  }
}

async function PControlMergerStart() {
  const options = POptionRead();

  const hasCompatibleAnalysis = PResultAnalysisCompatibleCheck(options);

  POptionSave();
  PMeterReset();
  PControlStateSet(true, true);
  PMeterText.textContent = PLanguageTextRead("startingMerge");
  PResultMergeUpdateSet(hasCompatibleAnalysis);

  if (!hasCompatibleAnalysis) {
    PResultMergeStart(options);
  }

  try {
    const report = await window.go.bridge.LProgram.LMergerRun(options);
    PResultReportSet(report);
  } catch (error) {
    PResultErrorShow(PLanguageTextRead("mergeError"), error);
  } finally {
    PResultMergeUpdateSet(false);
    PControlStateSet(false, false);
  }
}

async function PControlTaskStop() {
  try {
    PMeterText.textContent = PLanguageTextRead("stopping");
    await window.go.bridge.LProgram.LTaskStop();
  } catch (error) {
    PResultErrorShow(PLanguageTextRead("stopError"), error);
  }
}
