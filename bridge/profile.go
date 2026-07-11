package bridge

import (
	"encoding/json"
)

type LProgramProfile struct {
	LProgramName        string `json:"LProgramName"`
	LProgramVersion     string `json:"LProgramVersion"`
	LProgramAuthorName  string `json:"LProgramAuthorName"`
	LProgramAuthorEmail string `json:"LProgramAuthorEmail"`
}

type LVersionManifest struct {
	LProgramVersion string `json:"version"`
	LProgramName    string `json:"name"`
	LProgramAuthor  struct {
		LProgramName  string `json:"name"`
		LProgramEmail string `json:"email"`
	} `json:"author"`
}

func (a *LProgram) LProgramProfileRead() LProgramProfile {
	info := LProgramProfile{
		LProgramName:       "Video Merger",
		LProgramVersion:    "0.0",
		LProgramAuthorName: "Unknown",
	}

	var config LVersionManifest
	if err := json.Unmarshal(a.LVersionData, &config); err != nil {
		return info
	}

	if config.LProgramVersion != "" {
		info.LProgramVersion = config.LProgramVersion
	}

	if config.LProgramName != "" {
		info.LProgramName = config.LProgramName
	}

	if config.LProgramAuthor.LProgramName != "" {
		info.LProgramAuthorName = config.LProgramAuthor.LProgramName
	}

	info.LProgramAuthorEmail = config.LProgramAuthor.LProgramEmail

	return info
}
