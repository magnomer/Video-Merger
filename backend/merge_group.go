package backend

import "context"

func LMergerGroupCreate(groupResult LBatchResult) LBatch {
	return LBatch{
		LBatchName:      groupResult.LBatchName,
		LClipExtension:  LBatchExtensionRead(groupResult.LBatchClip),
		LBatchDirectory: groupResult.LBatchDirectory,
		LBatchClip:      groupResult.LBatchClip,
		LBatchNotice:    groupResult.LBatchNotice,
	}
}

func LMergerGroupRun(
	LRuntimeContext context.Context,
	options LPreference,
	groupResult LBatchResult,
	group LBatch,
) LMergerResult {
	return LMergerCancelRun(
		LRuntimeContext,
		group,
		groupResult.LBatchCompatibility,
		groupResult.LBatchPlan,
		options.LPreferenceCaution,
		options.LPreferenceWarning,
	)
}

func LMergerGroupStatusSet(groupResult LBatchResult, mergeResult LMergerResult) LBatchResult {
	if mergeResult.LTaskSuccess {
		groupResult.LProgressStatus = "Finished"
	} else {
		groupResult.LProgressStatus = "Failed"
	}

	return groupResult
}
