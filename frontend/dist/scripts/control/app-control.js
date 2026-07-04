function POutputStateSet() {
  const sameAsInput = POptionMirror.checked;

  POutputText.disabled = sameAsInput;
  POutputFolder.disabled = sameAsInput;
  POutputExplorer.disabled = false;
}

function PControlStateSet(isBusy, canStop) {
  PControlAnalysis.disabled = isBusy;
  PControlAssembly.disabled = isBusy;
  PControlTermination.disabled = !canStop;

  PSourceExplorer.disabled = isBusy;
  PSourceFolder.disabled = isBusy;
  POutputFolder.disabled = isBusy || POptionMirror.checked;
  POutputExplorer.disabled = isBusy;
  PFrameInfo.disabled = isBusy;

  PSourceText.disabled = isBusy;
  POutputText.disabled = isBusy || POptionMirror.checked;
  POptionSubfolder.disabled = isBusy;
  POptionMirror.disabled = isBusy;
  POptionField.disabled = isBusy;
  POptionCaution.disabled = isBusy;
  POptionWarning.disabled = isBusy;
}
