function PInspectorCopyStart() {
  const copyButton = document.querySelector(".POutputCopy");
  const outputText = document.querySelector(".POutputRow span")?.textContent || "";

  if (!copyButton || outputText === "" || outputText === "-") {
    return;
  }

  copyButton.addEventListener("click", async () => {
    try {
      await navigator.clipboard.writeText(outputText);
      copyButton.setAttribute("aria-label", "Copied planned output");
    } catch {
      PInspectorTextCopy(outputText);
    }
  });
}

function PInspectorTextCopy(text) {
  const field = document.createElement("textarea");
  field.value = text;
  field.setAttribute("readonly", "");
  field.style.position = "fixed";
  field.style.opacity = "0";
  document.body.appendChild(field);
  field.select();
  document.execCommand("copy");
  field.remove();
}
