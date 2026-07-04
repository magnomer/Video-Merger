function PPreviewOffsetRead(files) {
  const offsets = [0];
  files.forEach(file => offsets.push(offsets[offsets.length - 1] + Math.max(0, Number(file.LReportDurationSecond || 0))));
  return offsets;
}

function PPreviewTimeSet(state, offsets, video, slider, now) {
  if (Number.isFinite(state.pendingSecond)) {
    PPreviewTargetSet(state, slider, now, state.pendingSecond);
    return;
  }

  const current = (offsets[state.index] || 0) + (video.currentTime || 0);
  slider.value = String(current);
  now.textContent = PPreviewTimeRead(current);
}

function PPreviewTargetSet(state, slider, now, second) {
  state.pendingSecond = Math.max(0, Number(second) || 0);
  slider.value = String(state.pendingSecond);
  now.textContent = PPreviewTimeRead(state.pendingSecond);
}

function PPreviewTargetClear(state, offsets, video, slider, now) {
  state.pendingSecond = null;
  PPreviewTimeSet(state, offsets, video, slider, now);
}

function PPreviewTimeRead(seconds) {
  if (!Number.isFinite(seconds) || seconds <= 0) return "00:00:00";
  const value = Math.floor(seconds);
  const hour = Math.floor(value / 3600);
  const minute = Math.floor((value % 3600) / 60);
  const second = value % 60;
  return `${String(hour).padStart(2, "0")}:${String(minute).padStart(2, "0")}:${String(second).padStart(2, "0")}`;
}
