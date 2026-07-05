package backend

import "sort"

func LBatchClipCreate(files []LClip) []LBatch {
	groupMap := map[string][]LClip{}

	for _, file := range files {
		key := file.LBatchDirectory + "\x00" + file.LBatchName + "\x00" + file.LClipExtension
		groupMap[key] = append(groupMap[key], file)
	}

	var groups []LBatch

	for _, files := range groupMap {
		sort.SliceStable(files, func(i, j int) bool {
			if files[i].LClipNumber != files[j].LClipNumber {
				return files[i].LClipNumber < files[j].LClipNumber
			}

			return files[i].LClipPath < files[j].LClipPath
		})

		group := LBatch{
			LBatchName:      files[0].LBatchName,
			LClipExtension:  files[0].LClipExtension,
			LBatchDirectory: files[0].LBatchDirectory,
			LBatchClip:      files,
			LBatchNotice:    LSequenceFind(files),
		}

		groups = append(groups, group)
	}

	sort.Slice(groups, func(i, j int) bool {
		if groups[i].LBatchDirectory != groups[j].LBatchDirectory {
			return groups[i].LBatchDirectory < groups[j].LBatchDirectory
		}

		return groups[i].LBatchName < groups[j].LBatchName
	})

	return groups
}
