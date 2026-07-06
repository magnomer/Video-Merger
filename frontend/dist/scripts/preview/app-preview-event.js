function PPreviewEventStart(state, offsets, view, load) {
  view.play.addEventListener("click", () => {
    if (!PPreviewSessionCheck(state)) {
      return;
    }

    // First press loads the clip; nothing is fetched until the user asks.
    if (!state.loaded) {
      view.play.innerHTML = PPreviewIconShow("Pause");
      load(state.index, true, 0);
      return;
    }

    if (view.video.paused) {
      view.play.innerHTML = PPreviewIconShow("Pause");
      view.video.play().catch(() => {});
    } else {
      view.video.pause();
    }
  });
  view.next.addEventListener("click", () => {
    if (!PPreviewSessionCheck(state)) {
      return;
    }

    const shouldPlay = state.loaded ? !view.video.paused : true;
    load((state.index + 1) % state.files.length, shouldPlay);
  });
  view.slider.addEventListener("pointerdown", event => PPreviewSliderStart(event, state, view.video));
  view.slider.addEventListener("input", () => PPreviewSeekSet(Number(view.slider.value), state, offsets, view.video, load));
  view.slider.addEventListener("pointerup", event => PPreviewSliderStop(event, state, view.video));
  view.slider.addEventListener("pointercancel", event => PPreviewSliderStop(event, state, view.video));
  PPreviewSegmentEventStart(view.timeline, state, view.video, load);
  // Media-element events (play/pause/timeupdate/ended/error) are bound once on
  // the shared <video> in PPreviewVideoBind; only the freshly rebuilt controls
  // are wired here to avoid stacking duplicate listeners on every group switch.
  view.volume?.addEventListener("input", () => {
    if (!PPreviewSessionCheck(state)) {
      return;
    }

    view.video.volume = Math.max(0, Math.min(1, Number(view.volume.value || 0) / 100));
  });
  view.full.addEventListener("click", () => {
    if (!PPreviewSessionCheck(state)) {
      return;
    }

    view.video.requestFullscreen?.();
  });
}


function PPreviewSegmentEventStart(timeline, state, video, load) {
  timeline?.querySelectorAll(".PSegment[data-preview-index]").forEach(segment => {
    segment.addEventListener("click", event => {
      event.stopPropagation();
      if (!PPreviewSessionCheck(state)) {
        return;
      }
      const index = Number(segment.dataset.previewIndex);
      if (!Number.isInteger(index) || index < 0 || index >= state.files.length) {
        return;
      }

      const shouldPlay = state.loaded ? (!video.paused && !video.ended) : true;
      load(index, shouldPlay, 0);
    });
  });
}

function PPreviewPlaySet(state, view) {
  if (!PPreviewSessionCheck(state)) {
    return;
  }

  state.playing = true;
  view.play.innerHTML = PPreviewIconShow("Pause");
  if (view.notice) view.notice.hidden = true;
}

function PPreviewPauseSet(state, view) {
  if (!PPreviewSessionCheck(state)) {
    return;
  }

  state.playing = false;
  view.play.innerHTML = PPreviewIconShow("Play");
}

function PPreviewEndSet(state, view, load, offsets) {
  if (!PPreviewSessionCheck(state)) {
    return;
  }

  if (state.index < state.files.length - 1) {
    load(state.index + 1, true);
    return;
  }

  view.play.innerHTML = PPreviewIconShow("Play");
  PPreviewSeekSet(0, state, offsets, view.video, load);
}

function PPreviewErrorSet(state, offsets, view, load) {
  if (!PPreviewSessionCheck(state)) {
    return;
  }

  if (view.video.dataset.previewCompatibility !== "true") {
    const asset = view.video.dataset.previewAsset || state.files[state.index]?.LReportAsset;
    if (asset) {
      const targetSecond = Number.isFinite(state.pendingSecond)
        ? Math.max(0, state.pendingSecond - (offsets[state.index] || 0))
        : Math.max(0, view.video.currentTime || 0);
      if (view.notice) {
        view.notice.hidden = false;
        view.notice.textContent = PLanguageTextRead("previewCompatibilityPreparing");
      }
      load(state.index, state.resume || state.playing, targetSecond, true);
      return;
    }
  }

  if (view.notice) {
    view.notice.hidden = false;
    view.notice.textContent = PLanguageTextRead("previewOpenError");
  }
}
