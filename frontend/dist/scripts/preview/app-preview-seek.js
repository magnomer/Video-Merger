function PPreviewSliderStart(event, state, video) {
  if (!PPreviewSessionCheck(state)) {
    return;
  }

  state.seeking = true;
  state.resume = !video.paused || state.playing;
  event.currentTarget.setPointerCapture?.(event.pointerId);
}

function PPreviewSliderStop(event, state, video) {
  if (!PPreviewSessionCheck(state)) {
    return;
  }

  state.seeking = false;
  event.currentTarget.releasePointerCapture?.(event.pointerId);

  if (state.resume) {
    video.play().catch(() => {});
  }
}

function PPreviewSeekSet(target, state, offsets, video, load) {
  if (!PPreviewSessionCheck(state)) {
    return;
  }

  let index = offsets.findIndex((offset, i) => i < offsets.length - 1 && target >= offset && target < offsets[i + 1]);
  if (index < 0) index = state.files.length - 1;
  const shouldPlay = state.seeking ? state.resume : (!video.paused && !video.ended);
  load(index, shouldPlay, target - offsets[index]);
}
