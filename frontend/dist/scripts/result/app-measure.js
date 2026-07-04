function PMeasureRead(group) {
  const metrics = [];
  (group.LReportSection || []).forEach(section => {
    if (section.LReportMetric && section.LReportMetric.length > 0) {
      metrics.push(...section.LReportMetric);
    }
  });

  const size = PMeasureValueRead(metrics, ["size"]) || group.LReportSize || "-";
  const duration = PMeasureValueRead(metrics, ["duration"]) || group.LReportDuration || "-";

  return [
    { LReportLabel: PLanguageTextRead("audioLoudness"), LReportValue: group.LReportLoudness || "-" },
    { LReportLabel: PLanguageTextRead("codec"), LReportValue: "H.264 / AAC" },
    { LReportLabel: PLanguageTextRead("streams"), LReportValue: `${PLanguageTextRead("video")}: 1  •  ${PLanguageTextRead("audio")}: 1` },
    { LReportLabel: PLanguageTextRead("totalDurationEstimated"), LReportValue: duration },
    { LReportLabel: PLanguageTextRead("totalSizeEstimated"), LReportValue: size },
  ];
}

function PMeasureValueRead(metrics, tokens) {
  const metric = metrics.find(item => tokens.some(token => String(item.LReportLabel || "").toLowerCase().includes(token)));
  return metric?.LReportValue || "";
}

function PMeasureShow(metrics) {
  return `
    <dl class="PMeasure">
      ${metrics.map(metric => `
        <div><dt>${LHtmlEscape(PLanguageReportTextRead(metric.LReportLabel))}</dt><dd>${LHtmlEscape(metric.LReportValue)}</dd></div>
      `).join("")}
    </dl>
  `;
}
