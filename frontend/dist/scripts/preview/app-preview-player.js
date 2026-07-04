function PPreviewStart(group) {
  const view = PPreviewViewRead(group);
  if (!view.video || view.files.length === 0) {
    return;
  }

  const state = { files: view.files, index: 0, seeking: false, playing: false, resume: false, pendingSecond: null, readyKey: 0 };
  const offsets = PPreviewOffsetRead(view.files);
  const totalSeconds = offsets[offsets.length - 1] || 0;

  view.slider.max = String(totalSeconds);
  view.total.textContent = PPreviewTimeRead(totalSeconds) || group.LReportDuration || "00:00";
  view.video.volume = Math.max(0, Math.min(1, Number(view.volume?.value || 72) / 100));

  const load = (index, shouldPlay, targetSecond = 0) => PPreviewLoad(index, shouldPlay, targetSecond, state, offsets, view);

  PPreviewEventStart(state, offsets, view, load);
  load(0, false);
}

function PPreviewViewRead(group) {
  return {
    video: document.getElementById("PPreviewVideo"),
    play: document.getElementById("PPreviewPlay"),
    next: document.getElementById("PPreviewNext"),
    slider: document.getElementById("PPreviewSlider"),
    now: document.getElementById("PPreviewNow"),
    total: document.getElementById("PPreviewTotal"),
    tag: document.getElementById("PPreviewTag"),
    full: document.getElementById("PPreviewFull"),
    volume: document.getElementById("PPreviewVolume"),
    timeline: document.querySelector(".PTimelineTrack"),
    files: (group?.LReportFile || []).slice().sort((a, b) => a.LReportNumber - b.LReportNumber).filter(file => file.LReportPath),
  };
}

function PPreviewLoad(index, shouldPlay, targetSecond, state, offsets, view) {
  state.index = Math.max(0, Math.min(index, state.files.length - 1));
  state.resume = shouldPlay;
  PPreviewTargetSet(state, view.slider, view.now, (offsets[state.index] || 0) + Math.max(0, targetSecond || 0));
  const readyKey = state.readyKey + 1;
  state.readyKey = readyKey;
  const ready = () => {
    if (readyKey === state.readyKey) {
      PPreviewTargetClear(state, offsets, view.video, view.slider, view.now);
    }
  };
  view.tag.textContent = `${state.index + 1} / ${state.files.length} ${PLanguageTextRead("playback")}`;

  const source = PPreviewUrlRead(state.files[state.index].LReportPath);
  if (view.video.src.endsWith(source)) {
    const second = Math.max(0, targetSecond);
    PPreviewFrameShow(view);
    PPreviewFrameReadySet(view, second, false, ready);
    view.video.currentTime = second;
    PPreviewTimeSet(state, offsets, view.video, view.slider, view.now);
    if (shouldPlay) view.video.play().catch(() => {});
    return;
  }

  PPreviewFrameShow(view);
  view.video.onloadedmetadata = () => {
    const second = Math.max(0, targetSecond);
    PPreviewFrameReadySet(view, second, true, ready);
    view.video.currentTime = second;
    PPreviewTimeSet(state, offsets, view.video, view.slider, view.now);
    if (state.resume) view.video.play().catch(() => {});
  };
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
