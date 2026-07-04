function POptionStart() {
  POptionLoad();
  POutputStateSet();
  POptionAutoStart();

  if (typeof PResultOutputStart === "function") {
    PResultOutputStart();
  }
}

function POptionRead() {
  return {
    LPreferenceInput: PSourceInputRead(),
    LPreferenceOutput: POutputText.value.trim(),
    LPreferenceMirror: POptionMirror.checked,
    LPreferenceTree: POptionSubfolder.checked,
    LPreferenceSuffix: POptionField.value.trim(),
    LPreferenceCaution: POptionCaution.checked,
    LPreferenceWarning: POptionWarning.checked,
    LPreferenceMarker: PSSettingStyle.value,
    LPreferencePattern: PSSettingPattern.value.trim(),
    LPreferenceCustom: PSSettingCustom.checked,
    LPreferenceUnnumbered: PSSettingUnnumbered.checked,
  };
}

function PSourceInputRead() {
  return PSourceText.value
    .split(/\r?\n/)
    .map(path => path.trim())
    .filter(path => path.length > 0);
}

function POptionWrite(preference) {
  PSourceText.value = (preference.LPreferenceInput || []).join("\n");
  PSourceIconSet();
  POutputText.value = preference.LPreferenceOutput || "";
  POptionSubfolder.checked = Boolean(preference.LPreferenceTree);
  POptionMirror.checked = Boolean(preference.LPreferenceMirror);
  POptionField.value = preference.LPreferenceSuffix || "";
  POptionCaution.checked = Boolean(preference.LPreferenceCaution);
  POptionWarning.checked = Boolean(preference.LPreferenceWarning);
  PSSettingStyle.value = preference.LPreferenceMarker || "paren";
  PSSettingPattern.value = preference.LPreferencePattern || "";
  PSSettingCustom.checked = Boolean(preference.LPreferenceCustom);
  PSSettingUnnumbered.checked = Boolean(preference.LPreferenceUnnumbered);
  POutputStateSet();

  if (typeof PSSettingStateSet === "function") {
    PSSettingStateSet();
  }
}

function POptionAutoStart() {
  const elementsToSave = [
    PSourceText,
    POutputText,
    POptionSubfolder,
    POptionMirror,
    POptionField,
    POptionCaution,
    POptionWarning,
  ];

  elementsToSave.forEach(element => {
    element.addEventListener("input", POptionSave);
    element.addEventListener("change", POptionSave);
  });
}

async function POptionSave() {
  if (!window.go?.bridge?.LProgram?.LPreferenceSave) {
    return;
  }

  try {
    await window.go.bridge.LProgram.LPreferenceSave(POptionRead());
  } catch {
    // Preference save failures should not interrupt the UI flow.
  }
}

async function POptionLoad() {
  if (!window.go?.bridge?.LProgram?.LPreferenceLoad) {
    return;
  }

  try {
    const preference = await window.go.bridge.LProgram.LPreferenceLoad();
    POptionWrite(preference || {});
  } catch {
    // Start with empty controls when preferences cannot be loaded.
  }
}
