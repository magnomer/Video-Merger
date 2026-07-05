package backend

func LBatchPathCreate(paths []string) []LBatch {
	var files []LClip

	marker := LMarkerDefaultCreate()

	for _, path := range paths {
		file, ok := LClipParse(path, marker)
		if !ok {
			continue
		}

		files = append(files, file)
	}

	return LBatchClipCreate(files)
}
