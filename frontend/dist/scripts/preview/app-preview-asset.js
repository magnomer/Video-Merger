function PPreviewUrlRead(path) {
  return `/LAssetVideoRead.mp4?path=${encodeURIComponent(String(path || ""))}`;
}

function PPreviewIconShow(name) {
  return `<img class="PPreviewIcon PPreviewIcon${name}" src="assets/${name}.svg" alt="" aria-hidden="true" />`;
}
