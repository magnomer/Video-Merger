const PSSettingPresetMap = {
  paren: /^(.+?)\s*\((\d+)\)$/,
  bracket: /^(.+?)\s*\[(\d+)\]$/,
  underscore: /^(.+?)_(\d+)$/,
  dash: /^(.+?)\s*-\s*(\d+)$/,
  space: /^(.+?)\s+(\d+)$/,
  dot: /^(.+?)\.(\d+)$/,
};

const PSSettingSampleMap = {
  paren: "Trip (1).mp4",
  bracket: "Trip [1].mp4",
  underscore: "Trip_1.mp4",
  dash: "Trip - 1.mp4",
  space: "Trip 1.mp4",
  dot: "Trip.1.mp4",
};

function PSSettingStart() {
  PFrameSetting.addEventListener("click", PSSettingOpen);
  PSSettingButton.addEventListener("click", PSSettingClose);
  PSSettingIcon.addEventListener("click", PSSettingClose);

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

  [PSSettingStyle, PSSettingCustom, PSSettingPattern, PSSettingUnnumbered].forEach(element => {
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
  let pattern;
  let sampleName;

  if (PSSettingCustom.checked) {
    sampleName = PSSettingSampleMap.paren;
    const expression = PSSettingPattern.value.trim();

    if (expression === "") {
      PSSettingSample.textContent = PLanguageTextRead("settingSampleEmpty");
      return;
    }

    try {
      pattern = new RegExp(expression);
    } catch {
      PSSettingSample.textContent = PLanguageTextRead("settingSampleInvalid");
      return;
    }
  } else {
    const style = PSSettingStyle.value;
    pattern = PSSettingPresetMap[style] || PSSettingPresetMap.paren;
    sampleName = PSSettingSampleMap[style] || PSSettingSampleMap.paren;
  }

  const stem = sampleName.replace(/\.[^.]+$/, "");
  const matches = stem.match(pattern);
  const prefix = `${PLanguageTextRead("settingSample")} ${sampleName} → `;

  if (matches && matches[2] !== undefined) {
    PSSettingSample.textContent = `${prefix}${(matches[1] || "").trim()} · ${matches[2]}`;
    return;
  }

  if (PSSettingUnnumbered.checked) {
    PSSettingSample.textContent = `${prefix}${stem} (${PLanguageTextRead("settingSampleNoNumber")})`;
    return;
  }

  PSSettingSample.textContent = `${prefix}${PLanguageTextRead("settingSampleNone")}`;
}
