function PSSettingStart() {
  PFrameSetting.addEventListener("click", PSSettingOpen);
  PSSettingButton.addEventListener("click", PSSettingClose);
  PSSettingIcon.addEventListener("click", PSSettingClose);
  PSSettingFFmpegButton.addEventListener("click", PSSettingFFmpegOpen);
  PSSettingTemporaryButton.addEventListener("click", PSSettingTemporaryOpen);
  PSSettingCleanButton.addEventListener("click", PSSettingCleanRun);

  PSSettingBackdrop.addEventListener("click", event => {
    if (event.target === PSSettingBackdrop) {
      PSSettingClose();
    }
  });

  window.addEventListener("keydown", event => {
    if (event.key === "Escape" && !PSSettingBackdrop.hidden) {
      PSSettingClose();
    }
  });

  [PSSettingStyle, PSSettingCustom, PSSettingPattern, PSSettingUnnumbered, PSSettingFFmpeg, PSSettingTemporary].forEach(element => {
    element.addEventListener("input", PSSettingChangeSet);
    element.addEventListener("change", PSSettingChangeSet);
  });

  PSSettingStateSet();
}

function PSSettingChangeSet() {
  PSSettingStateSet();

  if (typeof POptionSave === "function") {
    POptionSave();
  }
}

function PSSettingOpen() {
  PSSettingBackdrop.hidden = false;
  PSSettingStateSet();
  PSSettingButton.focus();
}

function PSSettingClose() {
  PSSettingBackdrop.hidden = true;
  PFrameSetting.focus();
}

function PSSettingStateSet() {
  const custom = PSSettingCustom.checked;
  PSSettingPattern.disabled = !custom;
  PSSettingStyle.disabled = custom;
  PSSettingSampleSet();
}

function PSSettingSampleSet() {
  PSSettingSample.textContent = `Trip.mp4      → group: Trip, number: 0
Trip (1).mp4 → group: Trip, number: 1
Trip (2).mp4 → group: Trip, number: 2`;
}



async function PSSettingFFmpegOpen() {
  try {
    const selectedPath = await window.go.bridge.LProgram.LPickerFFmpegOpen();

    if (!selectedPath) {
      return;
    }

    PSSettingFFmpeg.value = selectedPath;
    PSSettingChangeSet();
  } catch (error) {
    PSSettingCleanStatus.textContent = `${PLanguageTextRead("settingFFmpegError")} ${error}`;
  }
}

async function PSSettingTemporaryOpen() {
  try {
    const selectedFolder = await window.go.bridge.LProgram.LPickerFolderOpen();

    if (!selectedFolder) {
      return;
    }

    PSSettingTemporary.value = selectedFolder;
    PSSettingChangeSet();
  } catch (error) {
    PSSettingCleanStatus.textContent = `${PLanguageTextRead("settingTemporaryError")} ${error}`;
  }
}

async function PSSettingCleanRun() {
  if (!window.go?.bridge?.LProgram?.LTemporaryClean) {
    return;
  }

  try {
    PSSettingCleanButton.disabled = true;
    PSSettingCleanStatus.textContent = PLanguageTextRead("settingCleanWorking");

    if (typeof POptionSave === "function") {
      await POptionSave();
    }

    const result = await window.go.bridge.LProgram.LTemporaryClean(POptionRead());
    PSSettingCleanStatus.textContent = PLanguageTextRead("settingCleanDone")
      .replace("{count}", result.LTemporaryCount)
      .replace("{path}", result.LTemporaryPath);
  } catch (error) {
    PSSettingCleanStatus.textContent = `${PLanguageTextRead("settingCleanError")} ${error}`;
  } finally {
    PSSettingCleanButton.disabled = false;
  }
}
