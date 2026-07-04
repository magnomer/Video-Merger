function PPreviewEventStart(state, offsets, view, load) {
  view.play.addEventListener("click", () => view.video.paused ? view.video.play().catch(() => {}) : view.video.pause());
  view.next.addEventListener("click", () => load((state.index + 1) % state.files.length, !view.video.paused));
  view.slider.addEventListener("pointerdown", event => PPreviewSliderStart(event, state, view.video));
  view.slider.addEventListener("input", () => PPreviewSeekSet(Number(view.slider.value), state, offsets, view.video, load));
  view.slider.addEventListener("pointerup", event => PPreviewSliderStop(event, state, view.video));
  view.slider.addEventListener("pointercancel", event => PPreviewSliderStop(event, state, view.video));
  view.timeline?.addEventListener("click", event => PPreviewTimelineSet(event, state, offsets, view.video, load));
  view.video.addEventListener("play", () => PPreviewPlaySet(state, view));
  view.video.addEventListener("pause", () => PPreviewPauseSet(state, view));
  view.video.addEventListener("timeupdate", () => PPreviewTimeSet(state, offsets, view.video, view.slider, view.now));
  view.video.addEventListener("ended", () => PPreviewEndSet(state, view, load, offsets));
  view.video.addEventListener("error", () => PPreviewErrorSet(view));
  view.volume?.addEventListener("input", () => { view.video.volume = Math.max(0, Math.min(1, Number(view.volume.value || 0) / 100)); });
  view.full.addEventListener("click", () => view.video.requestFullscreen?.());
}

function PPreviewPlaySet(state, view) {
  state.playing = true;
  view.play.innerHTML = PPreviewIconShow("Pause");
  view.notice.hidden = true;
}

function PPreviewPauseSet(state, view) {
  state.playing = false;
  view.play.innerHTML = PPreviewIconShow("Play");
}

function PPreviewEndSet(state, view, load, offsets) {
  if (state.index < state.files.length - 1) {
    load(state.index + 1, true);
    return;
  }

  view.play.innerHTML = PPreviewIconShow("Play");
  PPreviewSeekSet(0, state, offsets, view.video, load);
}

function PPreviewErrorSet(view) {
  view.notice.hidden = false;
  view.notice.textContent = PLanguageTextRead("previewOpenError");
}
