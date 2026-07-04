const LWindowStorageKey = "LWindowStorageKeyV44";
const LWindowMinimumWidth = 1200;
const LWindowMinimumHeight = 820;

function LWindowStateStart() {
  LWindowControlStart();
  LWindowStateLoad();

  let saveTimer = null;

  window.addEventListener("resize", () => {
    clearTimeout(saveTimer);
    saveTimer = setTimeout(LWindowStateSave, 300);
  });

  window.addEventListener("beforeunload", () => {
    LWindowStateSave();
  });
}

function LWindowControlStart() {
  document.getElementById("PFrameDash")?.addEventListener("click", () => window.runtime?.WindowMinimise?.());
  document.getElementById("PFrameBox")?.addEventListener("click", () => window.runtime?.WindowToggleMaximise?.());
  document.getElementById("PFrameExit")?.addEventListener("click", () => window.runtime?.Quit?.());
}

async function LWindowStateLoad() {
  const rawState = localStorage.getItem(LWindowStorageKey);

  if (!rawState || !window.runtime) {
    return;
  }

  try {
    const state = JSON.parse(rawState);

    if (
      Number.isFinite(state.width) &&
      Number.isFinite(state.height) &&
      state.width >= LWindowMinimumWidth &&
      state.height >= LWindowMinimumHeight
    ) {
      await window.runtime.WindowSetSize(state.width, state.height);
    }

    if (Number.isFinite(state.x) && Number.isFinite(state.y)) {
      await window.runtime.WindowSetPosition(state.x, state.y);
    }
  } catch {
    localStorage.removeItem(LWindowStorageKey);
  }
}

async function LWindowStateSave() {
  if (!window.runtime) {
    return;
  }

  try {
    const size = await window.runtime.WindowGetSize();
    const position = await window.runtime.WindowGetPosition();

    const state = {
      width: Math.max(size.w ?? size.width, LWindowMinimumWidth),
      height: Math.max(size.h ?? size.height, LWindowMinimumHeight),
      x: position.x,
      y: position.y,
    };

    if (
      Number.isFinite(state.width) &&
      Number.isFinite(state.height) &&
      Number.isFinite(state.x) &&
      Number.isFinite(state.y)
    ) {
      localStorage.setItem(LWindowStorageKey, JSON.stringify(state));
    }
  } catch {
    // Ignore window-state save failures.
  }
}
