function PPreviewUrlRead(asset, compatibility, session) {
  const query = new URLSearchParams({
    asset: String(asset || ""),
    session: String(session ?? ""),
  });
  if (compatibility) query.set("preview", "compatibility");
  return `/LAssetVideoRead.mp4?${query.toString()}`;
}

function PPreviewIconShow(name) {
  return `<img class="PPreviewIcon PPreviewIcon${name}" src="assets/${name}.svg" alt="" aria-hidden="true" />`;
}
