let PResultStateReport = null;
let PResultStateIndex = 0;
let PResultOutputTimer = null;
let PResultPreviewTimer = null;
let PResultMergeUpdate = false;
let PResultAnalysisKey = "";
const PResultPreviewDelay = 40;

function PResultErrorShow(title, error) {
  PMeterText.textContent = title;
  PResultStateReport = null;
  PResultRenderSet();

  const groupList = document.querySelector(".PGroupList");
  if (groupList) {
    groupList.innerHTML = `
      <div class="PPlaceholder">
        <span><strong>${PLanguageTextRead("error")}:</strong> ${LHtmlEscape(String(error))}</span>
      </div>
    `;
  }
}


function PResultAnalysisStart(options) {
  PResultAnalysisKey = PResultAnalysisKeyRead(options);
  PResultTaskStart("analysis", "Analysis started.", "analyzing");
}

function PResultMergeStart(options) {
  PResultAnalysisKey = "";
  PResultTaskStart("merge", "Starting merge.", "startingMerge");
}

function PResultTaskStart(kind, message, languageKey) {
  PResultShow({
    LReportKind: kind,
    LTaskMessage: message,
    LTaskCancel: false,
    LProgressTotal: 0,
    LProgressProcessed: 0,
    LProgressPercent: 0,
    LReportGroup: [],
  });

  PMeterFill.style.width = "0%";
  PMeterCount.textContent = "0 / 0";
  PMeterPercent.textContent = "0%";
  PMeterText.textContent = PLanguageTextRead(languageKey);
}

function PResultShow(report) {
  if (PResultMergeUpdate && report?.LReportKind === "merge") {
    PResultMergeSet(report);
    return;
  }

  PResultStateReport = report;
  const groups = report?.LReportGroup || [];

  if (PResultStateIndex >= groups.length) {
    PResultStateIndex = 0;
  }

  PResultRenderSet();
  PStatusReportSet(report);
}


function PResultPlaceholderRead(groups) {
  const isAnalysisRunning = PResultTaskRunningCheck("analysis", "Analysis completed.");
  const isMergeRunning = PResultTaskRunningCheck("merge", "Processing completed.");

  if (groups.length > 0) {
    return "";
  }

  if (isAnalysisRunning) {
    return PLanguageTextRead("analyzing");
  }

  return isMergeRunning ? PLanguageTextRead("startingMerge") : PLanguageTextRead("noGroups");
}

function PResultTaskRunningCheck(kind, finalMessage) {
  return PResultStateReport?.LReportKind === kind
    && !PResultStateReport?.LTaskCancel
    && PResultStateReport?.LTaskMessage !== finalMessage;
}

function PResultRenderSet() {
  PResultPreviewCancel();

  if (typeof PPreviewStop === "function") {
    PPreviewStop();
  }

  const groups = PResultStateReport?.LReportGroup || [];
  const selectedGroup = groups[PResultStateIndex] || groups[0] || null;
  const groupList = groups.length > 0
    ? groups.map((group, index) => PGroupShow(group, index)).join("")
    : `<div class="PPlaceholder">${PResultPlaceholderRead(groups)}</div>`;

  PResult.innerHTML = `
    <div class="PResult">
      <div class="PResultTitle">${PLanguageTextRead("groupsTitle")} (${groups.length})</div>
      <div class="PGroupList">
        ${groupList}
      </div>
      <div class="PWork">
        <aside class="PInspector">
          ${PInspectorShow(selectedGroup)}
        </aside>
        <section class="PPreview">
          ${PPreviewShow(selectedGroup)}
        </section>
      </div>
    </div>
  `;

  PInspectorCopyStart();
  PResultGroupStart();

  if (typeof PPreviewVideoMount === "function") {
    PPreviewVideoMount();
  }

  if (typeof PWorkSet === "function") {
    PWorkSet();
  }

  PPreviewScheduleStart(selectedGroup);
}

function PResultGroupStart() {
  document.querySelectorAll(".PGroupCard").forEach(card => {
    card.addEventListener("pointerdown", event => {
      if (event.button !== 0) {
        return;
      }

      PResultSelectionSet(Number(card.dataset.index || 0));
    });

    card.addEventListener("click", () => {
      PResultSelectionSet(Number(card.dataset.index || 0));
    });
  });
}

function PResultSelectionSet(index) {
  const groups = PResultStateReport?.LReportGroup || [];
  if (groups.length === 0) {
    return;
  }

  const nextIndex = Math.max(0, Math.min(index, groups.length - 1));
  if (nextIndex === PResultStateIndex) {
    return;
  }

  // Kill any in-flight preview work for the previous group at once so nothing
  // from the old selection keeps running once the user moves on.
  if (typeof PPreviewStop === "function") {
    PPreviewStop();
  }

  PResultStateIndex = nextIndex;
  const selectedGroup = groups[PResultStateIndex];

  // Reflect the selection immediately, regardless of whether the player is
  // ready. The card highlight and inspector are cheap and must never wait on
  // the media pipeline.
  PResultGroupSelectionSet();
  PResultInspectorSet(selectedGroup);

  // Defer the heavy preview rebuild (video element + timeline) so rapid group
  // switching only materialises the latest choice once, instead of thrashing
  // the media pipeline with a torn-down/recreated <video> per click.
  PResultPreviewSchedule();
}

function PResultGroupSelectionSet() {
  document.querySelectorAll(".PGroupCard").forEach(card => {
    const selected = Number(card.dataset.index || 0) === PResultStateIndex;
    card.classList.toggle("PGroupSelected", selected);
    card.setAttribute("aria-pressed", selected ? "true" : "false");
  });
}

function PResultInspectorSet(selectedGroup) {
  const inspector = document.querySelector(".PInspector");
  if (!inspector) {
    return;
  }

  inspector.innerHTML = PInspectorShow(selectedGroup);
  PInspectorCopyStart();
}

function PResultPreviewSchedule() {
  PResultPreviewCancel();
  PResultPreviewTimer = setTimeout(() => {
    PResultPreviewTimer = null;

    const groups = PResultStateReport?.LReportGroup || [];
    const selectedGroup = groups[PResultStateIndex] || groups[0] || null;
    PResultPreviewSet(selectedGroup);
  }, PResultPreviewDelay);
}

function PResultPreviewCancel() {
  if (PResultPreviewTimer !== null) {
    clearTimeout(PResultPreviewTimer);
    PResultPreviewTimer = null;
  }
}

function PResultPreviewSet(selectedGroup) {
  const preview = document.querySelector(".PPreview");
  if (preview) {
    preview.innerHTML = PPreviewShow(selectedGroup);
  }

  if (typeof PPreviewVideoMount === "function") {
    PPreviewVideoMount();
  }

  if (typeof PWorkSet === "function") {
    PWorkSet();
  }

  PPreviewScheduleStart(selectedGroup);
}

function PResultMergeSet(report) {
  PResultStateReport = report;
  const groups = report?.LReportGroup || [];

  if (groups.length === 0 || !document.querySelector(".PResult")) {
    PResultRenderSet();
    PStatusReportSet(report);
    return;
  }

  if (PResultStateIndex >= groups.length) {
    PResultStateIndex = 0;
  }

  const selectedGroup = groups[PResultStateIndex] || groups[0];
  const groupList = document.querySelector(".PGroupList");
  const inspector = document.querySelector(".PInspector");

  if (groupList) {
    groupList.innerHTML = groups.map((group, index) => PGroupShow(group, index)).join("");
    PResultGroupStart();
  }

  if (inspector) {
    inspector.innerHTML = PInspectorShow(selectedGroup);
    PInspectorCopyStart();
  }

  PStatusReportSet(report);
}

function PResultMergeUpdateSet(isUpdate) {
  PResultMergeUpdate = Boolean(isUpdate);
}

function PResultAnalysisCheck() {
  return PResultStateReport?.LReportKind === "analysis" && (PResultStateReport.LReportGroup || []).length > 0;
}

function PResultAnalysisCompatibleCheck(options) {
  return PResultAnalysisCheck() && PResultAnalysisKey !== "" && PResultAnalysisKey === PResultAnalysisKeyRead(options);
}

function PResultAnalysisKeyRead(options) {
  return JSON.stringify({
    LPreferenceInput: (options?.LPreferenceInput || []).map(value => String(value).trim()).filter(Boolean),
    LPreferenceTree: Boolean(options?.LPreferenceTree),
    LPreferenceMarker: String(options?.LPreferenceMarker || "").trim(),
    LPreferencePattern: String(options?.LPreferencePattern || "").trim(),
    LPreferenceCustom: Boolean(options?.LPreferenceCustom),
    LPreferenceUnnumbered: Boolean(options?.LPreferenceUnnumbered),
    LPreferenceFFmpeg: String(options?.LPreferenceFFmpeg || "").trim(),
  });
}


function PResultOutputStart() {
  [POutputText, POptionMirror, POptionField].forEach(element => {
    element.addEventListener("input", PResultOutputSet);
    element.addEventListener("change", PResultOutputSet);
  });
}

function PResultOutputSet() {
  if (!PResultStateReport || PResultStateReport.LReportKind !== "analysis") {
    return;
  }

  clearTimeout(PResultOutputTimer);
  PResultOutputTimer = setTimeout(PResultOutputApplySet, 120);
}

async function PResultOutputApplySet() {
  if (!window.go?.bridge?.LProgram?.LReportOutputSet || !PResultStateReport || PResultStateReport.LReportKind !== "analysis") {
    return;
  }

  try {
    const report = await window.go.bridge.LProgram.LReportOutputSet(PResultStateReport, POptionRead());
    PResultStateReport = report;
    PResultRenderSet();
    PStatusReportSet(report);
  } catch (error) {
    PMeterText.textContent = `${PLanguageTextRead("outputRefreshError")}: ${String(error)}`;
  }
}
