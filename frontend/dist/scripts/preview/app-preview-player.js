let PPreviewSessionStopSet = null;
let PPreviewSessionValue = 0;
let PPreviewRequestValue = 0;
let PPreviewRequestGroup = null;
let PPreviewVideoElement = null;
let PPreviewActive = null;

// The WebView2 media pipeline keeps a <video> loading/decoding until it is
// explicitly unloaded or garbage collected. Rebuilding the preview markup on
// every group switch would orphan a still-loading element each time, so the
// browser ends up juggling several concurrent media loads (CPU spikes, frozen
// UI). To avoid that we keep a single element alive and re-mount it.
function PPreviewVideoRead() {
  if (!PPreviewVideoElement) {
    const video = document.createElement("video");
    video.id = "PPreviewVideo";
    video.className = "PPreviewVideo";
    video.setAttribute("preload", "metadata");
    video.setAttribute("playsinline", "");
    PPreviewVideoElement = video;
    PPreviewVideoBind(video);
  }

  return PPreviewVideoElement;
}

function PPreviewVideoBind(video) {
  video.addEventListener("play", () => {
    if (PPreviewActive) PPreviewPlaySet(PPreviewActive.state, PPreviewActive.view);
  });
  video.addEventListener("pause", () => {
    if (PPreviewActive) PPreviewPauseSet(PPreviewActive.state, PPreviewActive.view);
  });
  video.addEventListener("timeupdate", () => {
    const active = PPreviewActive;
    if (!active || !PPreviewSessionCheck(active.state)) {
      return;
    }

    PPreviewTimeSet(active.state, active.offsets, active.view.video, active.view.slider, active.view.now);
  });
  video.addEventListener("ended", () => {
    if (PPreviewActive) PPreviewEndSet(PPreviewActive.state, PPreviewActive.view, PPreviewActive.load, PPreviewActive.offsets);
  });
  video.addEventListener("error", () => {
    if (PPreviewActive) PPreviewErrorSet(PPreviewActive.state, PPreviewActive.offsets, PPreviewActive.view, PPreviewActive.load);
  });
}

function PPreviewVideoReset() {
  const video = PPreviewVideoElement;
  if (!video) {
    return;
  }

  try {
    video.pause();
  } catch (_) {}

  video.onloadedmetadata = null;
  video.removeAttribute("src");
  video.removeAttribute("data-preview-asset");
  video.removeAttribute("data-preview-compatibility");

  try {
    video.load();
  } catch (_) {}
}

function PPreviewVideoMount() {
  const stage = document.querySelector(".PPreviewStage");
  if (!stage) {
    return;
  }

  const template = stage.querySelector("#PPreviewVideo");
  const video = PPreviewVideoRead();
  if (template === video) {
    return;
  }

  // Abort whatever the shared element was doing before it is reused for the
  // newly selected group.
  PPreviewVideoReset();

  if (template) {
    template.replaceWith(video);
  } else {
    stage.insertBefore(video, stage.firstChild);
  }
}

function PPreviewStop() {
  const stoppedSession = String(PPreviewSessionValue);
  PPreviewLocalStopSet();
  PPreviewBackendStopSet(stoppedSession);
}

function PPreviewQuitStop() {
  PPreviewLocalStopSet();
  PPreviewBackendStopSet("");
}

function PPreviewLocalStopSet() {
  PPreviewSessionValue += 1;

  if (typeof PPreviewSessionStopSet === "function") {
    PPreviewSessionStopSet();
    PPreviewSessionStopSet = null;
  }
}

function PPreviewBackendStopSet(session) {
  try {
    const stop = window.go?.bridge?.LProgram?.LAssetPreviewStop;
    if (typeof stop === "function") {
      Promise.resolve(stop(String(session ?? ""))).catch(() => null);
    }
  } catch (_) {}
}

function PPreviewFrameWait() {
  return new Promise(resolve => {
    requestAnimationFrame(() => requestAnimationFrame(resolve));
  });
}

function PPreviewScheduleStart(group) {
  PPreviewRequestValue += 1;
  PPreviewRequestGroup = group || null;
  const request = PPreviewRequestValue;
  const session = PPreviewSessionValue;

  PPreviewFrameWait().then(() => {
    if (request !== PPreviewRequestValue || session !== PPreviewSessionValue) {
      return;
    }

    if (PPreviewRequestGroup) {
      PPreviewStart(PPreviewRequestGroup, session);
    }
  });
}

function PPreviewStart(group, session) {
  const view = PPreviewViewRead(group);
  if (!view.video || view.files.length === 0 || session !== PPreviewSessionValue) {
    return;
  }

  const state = { files: view.files, index: 0, seeking: false, playing: false, resume: false, pendingSecond: null, readyKey: 0, active: true, loaded: false, session };
  const offsets = PPreviewOffsetRead(view.files);
  const totalSeconds = offsets[offsets.length - 1] || 0;

  view.slider.max = String(totalSeconds);
  view.total.textContent = PPreviewTimeRead(totalSeconds) || group.LReportDuration || "00:00";
  view.video.volume = Math.max(0, Math.min(1, Number(view.volume?.value || 72) / 100));

  const load = (index, shouldPlay, targetSecond = 0, compatibility = false) => PPreviewLoad(index, shouldPlay, targetSecond, state, offsets, view, compatibility);

  // The shared <video> keeps its listeners across groups, so route its events
  // to whichever session is current instead of re-binding them every time.
  PPreviewActive = { state, offsets, view, load };
  PPreviewSessionStopSet = () => PPreviewSessionStop(state, view);
  PPreviewEventStart(state, offsets, view, load);

  // Do NOT touch the media pipeline on selection. Selecting a group must never
  // make the webview load a video (no thumbnail, no metadata fetch). The clip
  // is only loaded once the user actually asks to play/seek it (state.loaded).
  view.tag.textContent = `1 / ${state.files.length} ${PLanguageTextRead("playback")}`;
}

function PPreviewSessionStop(state, view) {
  state.active = false;
  state.seeking = false;
  state.playing = false;
  state.resume = false;
  state.loaded = false;
  state.readyKey += 1;

  if (view?.notice) {
    view.notice.hidden = true;
  }

  if (!view?.video) {
    return;
  }

  try {
    view.video.pause();
  } catch (_) {}

  view.video.onloadedmetadata = null;
  view.video.removeAttribute("src");
  view.video.removeAttribute("data-preview-asset");
  view.video.removeAttribute("data-preview-compatibility");

  try {
    view.video.load();
  } catch (_) {}
}

function PPreviewViewRead(group) {
  return {
    video: PPreviewVideoRead(),
    play: document.getElementById("PPreviewPlay"),
    next: document.getElementById("PPreviewNext"),
    slider: document.getElementById("PPreviewSlider"),
    now: document.getElementById("PPreviewNow"),
    total: document.getElementById("PPreviewTotal"),
    tag: document.getElementById("PPreviewTag"),
    full: document.getElementById("PPreviewFull"),
    volume: document.getElementById("PPreviewVolume"),
    notice: document.getElementById("PPreviewNotice"),
    timeline: document.querySelector(".PTimelineTrack"),
    files: (group?.LReportFile || []).slice().sort((a, b) => a.LReportNumber - b.LReportNumber).filter(file => file.LReportAsset),
  };
}

function PPreviewSessionCheck(state) {
  return state.active && state.session === PPreviewSessionValue;
}

function PPreviewLoad(index, shouldPlay, targetSecond, state, offsets, view, compatibility = false) {
  if (!PPreviewSessionCheck(state)) {
    return;
  }

  state.loaded = true;
  state.index = Math.max(0, Math.min(index, state.files.length - 1));
  state.resume = shouldPlay;
  PPreviewTargetSet(state, view.slider, view.now, (offsets[state.index] || 0) + Math.max(0, targetSecond || 0));
  const readyKey = state.readyKey + 1;
  state.readyKey = readyKey;
  const ready = () => {
    if (PPreviewSessionCheck(state) && readyKey === state.readyKey) {
      PPreviewTargetClear(state, offsets, view.video, view.slider, view.now);
    }
  };
  view.tag.textContent = `${state.index + 1} / ${state.files.length} ${PLanguageTextRead("playback")}`;

  // Tell the user the webview is fetching the clip. The compatibility retry
  // keeps its own notice, and PPreviewPlaySet clears this once playback begins.
  if (view.notice && shouldPlay && !compatibility) {
    view.notice.hidden = false;
    view.notice.textContent = PLanguageTextRead("previewLoading");
  }

  const source = PPreviewUrlRead(state.files[state.index].LReportAsset, compatibility, state.session);
  if (view.video.src.endsWith(source)) {
    const second = Math.max(0, targetSecond);
    PPreviewFrameShow(view);
    PPreviewFrameReadySet(view, second, false, ready);
    view.video.currentTime = second;
    PPreviewTimeSet(state, offsets, view.video, view.slider, view.now);
    if (PPreviewSessionCheck(state) && shouldPlay) view.video.play().catch(() => {});
    return;
  }

  PPreviewFrameShow(view);
  view.video.dataset.previewAsset = state.files[state.index].LReportAsset;
  view.video.dataset.previewCompatibility = compatibility ? "true" : "false";
  view.video.onloadedmetadata = () => {
    if (!PPreviewSessionCheck(state) || readyKey !== state.readyKey) {
      return;
    }

    const second = Math.max(0, targetSecond);
    PPreviewFrameReadySet(view, second, true, ready);
    view.video.currentTime = second;
    PPreviewTimeSet(state, offsets, view.video, view.slider, view.now);
    if (PPreviewSessionCheck(state) && state.resume) view.video.play().catch(() => {});
  };

  if (!PPreviewSessionCheck(state)) {
    return;
  }

  view.video.src = source;
}

function PPreviewFrameShow(view) {
  if (!view.video || view.video.readyState < 2 || view.video.videoWidth <= 0 || view.video.videoHeight <= 0) {
    return;
  }

  const canvas = PPreviewFrameRead(view);
  const width = Math.max(1, Math.round(view.video.clientWidth));
  const height = Math.max(1, Math.round(view.video.clientHeight));
  canvas.width = width;
  canvas.height = height;

  const context = canvas.getContext("2d");
  const scale = Math.min(width / view.video.videoWidth, height / view.video.videoHeight);
  const drawWidth = Math.round(view.video.videoWidth * scale);
  const drawHeight = Math.round(view.video.videoHeight * scale);
  const drawX = Math.round((width - drawWidth) / 2);
  const drawY = Math.round((height - drawHeight) / 2);

  try {
    context.clearRect(0, 0, width, height);
    context.drawImage(view.video, drawX, drawY, drawWidth, drawHeight);
    canvas.hidden = false;
  } catch (_) {
    canvas.hidden = true;
  }
}

function PPreviewFrameRead(view) {
  if (view.frame) {
    return view.frame;
  }

  view.frame = document.createElement("canvas");
  view.frame.className = "PPreviewFrame";
  view.frame.hidden = true;
  view.video.insertAdjacentElement("afterend", view.frame);
  return view.frame;
}

function PPreviewFrameReadySet(view, targetSecond, waitInitial, ready) {
  let isDone = false;
  const hide = () => {
    if (isDone) {
      return;
    }

    isDone = true;
    ready?.();
    requestAnimationFrame(() => requestAnimationFrame(() => PPreviewFrameHide(view)));
  };

  if (!waitInitial && Math.abs(view.video.currentTime - targetSecond) < 0.04) {
    hide();
    return;
  }

  view.video.addEventListener("seeked", hide, { once: true });
  if (targetSecond <= 0.04) {
    view.video.addEventListener("loadeddata", hide, { once: true });
    view.video.addEventListener("canplay", hide, { once: true });
  }
}

function PPreviewFrameHide(view) {
  if (view.frame) {
    view.frame.hidden = true;
  }
}
