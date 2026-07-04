function PMeterReset() {
  PMeterFill.style.width = "0%";
  PMeterCount.textContent = "0 / 0";
  PMeterPercent.textContent = "0%";
  PStatusMessage.innerHTML = `<span class="PStatusDot">✓</span>${PLanguageTextRead("ready")}`;
  PStatusGroup.textContent = PLanguageCountRead(0, "group", "groups");
  PStatusClip.textContent = PLanguageCountRead(0, "clip", "clips");
  PStatusSize.textContent = "-";
  PStatusDuration.textContent = "-";
  PStatusWarning.textContent = PLanguageCountRead(0, "warning", "warnings");
}

function PResultReportSet(report) {
  if (!report) {
    return;
  }

  const percent = Number(report.LProgressPercent || 0);
  const processedFiles = Number(report.LProgressProcessed || 0);
  const totalFiles = Number(report.LProgressTotal || 0);
  const cleanPercent = Math.max(0, Math.min(100, percent));

  PMeterFill.style.width = `${cleanPercent}%`;
  PMeterCount.textContent = `${processedFiles} / ${totalFiles}`;
  PMeterPercent.textContent = `${cleanPercent}%`;
  PMeterText.textContent = PLanguageReportTextRead(report.LTaskMessage) || "";

  PResultShow(report);
}
