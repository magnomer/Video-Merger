let PResultStateReport = null;
let PResultStateIndex = 0;
let PResultOutputTimer = null;
let PResultMergeUpdate = false;

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

function PResultRenderSet() {
  const groups = PResultStateReport?.LReportGroup || [];
  const selectedGroup = groups[PResultStateIndex] || groups[0] || null;
  const groupList = groups.length > 0
    ? groups.map((group, index) => PGroupShow(group, index)).join("")
    : `<div class="PPlaceholder">${PLanguageTextRead("noGroups")}</div>`;

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
  PPreviewStart(selectedGroup);
  PResultGroupStart();

  if (typeof PWorkSet === "function") {
    PWorkSet();
  }
}

function PResultGroupStart() {
  document.querySelectorAll(".PGroupRow").forEach(row => {
    row.addEventListener("click", () => {
      PResultStateIndex = Number(row.dataset.index || 0);
      PResultRenderSet();
    });
  });
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
