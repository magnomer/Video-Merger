function PPreviewSliderStart(event, state, video) {
  state.seeking = true;
  state.resume = !video.paused || state.playing;
  event.currentTarget.setPointerCapture?.(event.pointerId);
}

function PPreviewTimelineSet(event, state, offsets, video, load) {
  const track = event.currentTarget;
  const bounds = track.getBoundingClientRect();
  const totalSeconds = offsets[offsets.length - 1] || 0;

  if (bounds.width <= 0 || totalSeconds <= 0) {
    return;
  }

  const ratio = Math.max(0, Math.min(1, (event.clientX - bounds.left) / bounds.width));
  state.seeking = true;
  state.resume = !video.paused || state.playing;
  PPreviewSeekSet(ratio * totalSeconds, state, offsets, video, load);
  state.seeking = false;

  if (state.resume) {
    video.play().catch(() => {});
  }
}

function PPreviewSliderStop(event, state, video) {
  state.seeking = false;
  event.currentTarget.releasePointerCapture?.(event.pointerId);

  if (state.resume) {
    video.play().catch(() => {});
  }
}

function PPreviewSeekSet(target, state, offsets, video, load) {
  let index = offsets.findIndex((offset, i) => i < offsets.length - 1 && target >= offset && target < offsets[i + 1]);
  if (index < 0) index = state.files.length - 1;
  const shouldPlay = state.seeking ? state.resume : (!video.paused && !video.ended);
  load(index, shouldPlay, target - offsets[index]);
}
