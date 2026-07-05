function PBadgeClassRead(tag) {
  if (tag === "ok") return "PBadgeOk";
  if (tag === "notice") return "PBadgeNotice";
  if (tag === "caution") return "PBadgeWarn";
  if (tag === "error") return "PBadgeError";
  return "PBadgeNeutral";
}

function PInspectorIconShow(state) {
  if (state === true || state === "ok") {
    return `<svg viewBox="0 0 24 24" aria-hidden="true"><path d="M20 6 9 17l-5-5" /></svg>`;
  }

  if (state === "caution") {
    return `<svg viewBox="0 0 24 24" aria-hidden="true"><path d="M12 3 3 20h18L12 3Z" /><path d="M12 9v5" /><path d="M12 17h.01" /></svg>`;
  }

  return `<svg viewBox="0 0 24 24" aria-hidden="true"><path d="M12 3 4 7v5c0 5 3.5 8 8 9c4.5-1 8-4 8-9V7l-8-4Z" /><path d="M12 8v5" /><path d="M12 16h.01" /></svg>`;
}
