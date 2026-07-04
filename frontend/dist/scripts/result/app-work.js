const PWorkDefaultHeight = 620;
const PWorkMinimumHeight = 440;
const PWorkDefaultWidth = 394;
const PWorkMinimumWidth = 280;
const PWorkMinimumPreviewWidth = 360;
const PWorkHeightStorageKey = "PWorkTopHeightV1";
const PWorkWidthStorageKey = "PWorkWidthV1";
let PWorkObserver = null;

function PWorkLoad() {
  const result = document.getElementById("PResult");
  if (!result) return;

  PWorkSet();

  if (PWorkObserver) return;
  PWorkObserver = new MutationObserver(() => PWorkSet());
  PWorkObserver.observe(result, { childList: true, subtree: true });
}

function PWorkSet() {
  const work = document.querySelector(".PWork");
  if (!work) return;

  PWorkHandleShow(work);
  PWorkWidthShow(work);
  PWorkHeightSet(work, PWorkHeightRead());
  PWorkWidthSet(work, PWorkWidthRead());
}

function PWorkHandleShow(work) {
  const parent = work.parentNode;
  if (!parent) return;

  parent.querySelectorAll(':scope > .PWorkHandle[data-work-axis="height"]').forEach(handle => {
    if (handle.nextElementSibling !== work) handle.remove();
  });

  const previous = work.previousElementSibling;
  if (previous?.classList?.contains("PWorkHandle") && previous.dataset.workAxis === "height") return;

  const handle = document.createElement("div");
  handle.className = "PWorkHandle";
  handle.dataset.workAxis = "height";
  handle.setAttribute("role", "separator");
  handle.setAttribute("aria-orientation", "horizontal");
  handle.setAttribute("aria-label", PLanguageTextRead("resizeHeight"));
  handle.tabIndex = 0;
  handle.addEventListener("pointerdown", event => PWorkResizeStart(event, work));
  handle.addEventListener("keydown", event => PWorkKeySet(event, work, "height"));
  work.insertAdjacentElement("beforebegin", handle);
}

function PWorkWidthShow(work) {
  const preview = work.querySelector(":scope > .PPreview");
  if (!preview) return;

  work.querySelectorAll(':scope > .PWorkHandle[data-work-axis="width"]').forEach(handle => {
    if (handle.nextElementSibling !== preview) handle.remove();
  });

  const previous = preview.previousElementSibling;
  if (previous?.classList?.contains("PWorkHandle") && previous.dataset.workAxis === "width") return;

  const handle = document.createElement("div");
  handle.className = "PWorkHandle";
  handle.dataset.workAxis = "width";
  handle.setAttribute("role", "separator");
  handle.setAttribute("aria-orientation", "vertical");
  handle.setAttribute("aria-label", PLanguageTextRead("resizeWidth"));
  handle.tabIndex = 0;
  handle.addEventListener("pointerdown", event => PWorkWidthStart(event, work));
  handle.addEventListener("keydown", event => PWorkKeySet(event, work, "width"));
  preview.insertAdjacentElement("beforebegin", handle);
}

function PWorkResizeStart(event, work) {
  if (event.button !== 0) return;

  const startY = event.clientY;
  const startHeight = work.getBoundingClientRect().height;
  const handle = event.currentTarget;
  handle.setPointerCapture(event.pointerId);

  handle.addEventListener("pointermove", PWorkPointerSet);
  handle.addEventListener("pointerup", PWorkPointerStop, { once: true });
  handle.addEventListener("pointercancel", PWorkPointerStop, { once: true });

  function PWorkPointerSet(moveEvent) {
    const nextHeight = startHeight - (moveEvent.clientY - startY);
    PWorkHeightSet(work, nextHeight);
  }

  function PWorkPointerStop() {
    handle.removeEventListener("pointermove", PWorkPointerSet);
    PWorkHeightSave(work);
  }
}

function PWorkWidthStart(event, work) {
  if (event.button !== 0) return;

  const startX = event.clientX;
  const startWidth = PWorkCurrentRead(work);
  const handle = event.currentTarget;
  handle.setPointerCapture(event.pointerId);

  handle.addEventListener("pointermove", PWorkPointerSet);
  handle.addEventListener("pointerup", PWorkPointerStop, { once: true });
  handle.addEventListener("pointercancel", PWorkPointerStop, { once: true });

  function PWorkPointerSet(moveEvent) {
    const nextWidth = startWidth + moveEvent.clientX - startX;
    PWorkWidthSet(work, nextWidth);
  }

  function PWorkPointerStop() {
    handle.removeEventListener("pointermove", PWorkPointerSet);
    PWorkWidthSave(work);
  }
}

function PWorkKeySet(event, work, axis) {
  const step = event.shiftKey ? 40 : 10;

  if (axis === "height") {
    const currentHeight = work.getBoundingClientRect().height;

    if (event.key === "ArrowUp") {
      event.preventDefault();
      PWorkHeightSet(work, currentHeight + step);
      PWorkHeightSave(work);
    }

    if (event.key === "ArrowDown") {
      event.preventDefault();
      PWorkHeightSet(work, currentHeight - step);
      PWorkHeightSave(work);
    }

    if (event.key === "Home") {
      event.preventDefault();
      PWorkHeightSet(work, PWorkMinimumHeight);
      PWorkHeightSave(work);
    }

    if (event.key === "End") {
      event.preventDefault();
      PWorkHeightSet(work, PWorkDefaultHeight);
      PWorkHeightSave(work);
    }
  }

  if (axis === "width") {
    const currentWidth = PWorkCurrentRead(work);

    if (event.key === "ArrowLeft") {
      event.preventDefault();
      PWorkWidthSet(work, currentWidth - step);
      PWorkWidthSave(work);
    }

    if (event.key === "ArrowRight") {
      event.preventDefault();
      PWorkWidthSet(work, currentWidth + step);
      PWorkWidthSave(work);
    }

    if (event.key === "Home") {
      event.preventDefault();
      PWorkWidthSet(work, PWorkMinimumWidth);
      PWorkWidthSave(work);
    }

    if (event.key === "End") {
      event.preventDefault();
      PWorkWidthSet(work, PWorkDefaultWidth);
      PWorkWidthSave(work);
    }
  }
}

function PWorkHeightRead() {
  const saved = Number(localStorage.getItem(PWorkHeightStorageKey));
  return Number.isFinite(saved) && saved >= PWorkMinimumHeight ? saved : PWorkDefaultHeight;
}

function PWorkWidthRead() {
  const saved = Number(localStorage.getItem(PWorkWidthStorageKey));
  return Number.isFinite(saved) && saved >= PWorkMinimumWidth ? saved : PWorkDefaultWidth;
}

function PWorkHeightSet(work, height) {
  const nextHeight = Math.max(PWorkMinimumHeight, Math.round(Number(height) || PWorkDefaultHeight));
  work.style.setProperty("--PWorkHeight", `${nextHeight}px`);

  const handle = work.previousElementSibling;
  if (handle?.classList?.contains("PWorkHandle") && handle.dataset.workAxis === "height") {
    handle.setAttribute("aria-valuemin", String(PWorkMinimumHeight));
    handle.setAttribute("aria-valuenow", String(nextHeight));
  }
}

function PWorkWidthSet(work, width) {
  const nextWidth = Math.min(PWorkMaximumRead(work), Math.max(PWorkMinimumWidth, Math.round(Number(width) || PWorkDefaultWidth)));
  work.style.setProperty("--PWorkInspectorWidth", `${nextWidth}px`);

  const handle = work.querySelector(':scope > .PWorkHandle[data-work-axis="width"]');
  if (handle) {
    handle.setAttribute("aria-valuemin", String(PWorkMinimumWidth));
    handle.setAttribute("aria-valuemax", String(PWorkMaximumRead(work)));
    handle.setAttribute("aria-valuenow", String(nextWidth));
  }
}

function PWorkHeightSave(work) {
  const height = Math.round(work.getBoundingClientRect().height);
  localStorage.setItem(PWorkHeightStorageKey, String(height));
}

function PWorkWidthSave(work) {
  localStorage.setItem(PWorkWidthStorageKey, String(PWorkCurrentRead(work)));
}

function PWorkCurrentRead(work) {
  const inspector = work.querySelector(":scope > .PInspector");
  return Math.round(inspector?.getBoundingClientRect().width || PWorkWidthRead());
}

function PWorkMaximumRead(work) {
  const maximum = work.getBoundingClientRect().width - PWorkMinimumPreviewWidth - 18;
  return Math.max(PWorkMinimumWidth, Math.round(maximum));
}

document.addEventListener("DOMContentLoaded", PWorkLoad);
