function PSAboutStart() {
  PFrameInfo.addEventListener("click", () => {
    PSAboutOpen();
  });

  PSAboutButton.addEventListener("click", PSAboutClose);
  PSAboutIcon.addEventListener("click", PSAboutClose);

  PSAboutBackdrop.addEventListener("click", event => {
    if (event.target === PSAboutBackdrop) {
      PSAboutClose();
    }
  });

  window.addEventListener("keydown", event => {
    if (event.key === "Escape" && !PSAboutBackdrop.hidden) {
      PSAboutClose();
    }
  });
}

function PSAboutOpen() {
  PSAboutBackdrop.hidden = false;
  PSAboutButton.focus();
}

function PSAboutClose() {
  PSAboutBackdrop.hidden = true;
  PFrameInfo.focus();
}

async function PFrameProfileLoad() {
  if (!window.go?.bridge?.LProgram?.LProgramProfileRead) {
    PFrameVersion.textContent = PLanguageTextRead("aboutVersion");
    PSAboutVersion.textContent = PLanguageTextRead("unknown");
    PSAboutAuthor.textContent = PLanguageTextRead("unknown");
    return;
  }

  try {
    const info = await window.go.bridge.LProgram.LProgramProfileRead();
    const version = info.LProgramVersion || PLanguageTextRead("unknown");
    PFrameVersion.textContent = `${PLanguageTextRead("aboutVersion")} ${version}`;
    PSAboutVersion.textContent = version;
    PSAboutAuthor.textContent = info.LProgramAuthorName || PLanguageTextRead("unknown");
  } catch {
    PFrameVersion.textContent = PLanguageTextRead("aboutVersion");
    PSAboutVersion.textContent = PLanguageTextRead("unknown");
    PSAboutAuthor.textContent = PLanguageTextRead("unknown");
  }
}
