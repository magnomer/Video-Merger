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

type LManifestWails struct {
	LProgramName   string `json:"name"`
	LProgramAuthor struct {
		LProgramName  string `json:"name"`
		LProgramEmail string `json:"email"`
	} `json:"author"`
	LProgramInfo struct {
		LProgramVersion string `json:"productVersion"`
	} `json:"info"`
}

func (a *LProgram) LProgramProfileRead() LProgramProfile {
	info := LProgramProfile{
		LProgramName:       "Video Merger",
		LProgramVersion:    "0.0",
		LProgramAuthorName: "Unknown",
	}

	var config LManifestWails
	if err := json.Unmarshal(a.LManifestWailsData, &config); err != nil {
		return info
	}

	if config.LProgramInfo.LProgramVersion != "" {
		info.LProgramVersion = config.LProgramInfo.LProgramVersion
	}

	if config.LProgramAuthor.LProgramName != "" {
		info.LProgramAuthorName = config.LProgramAuthor.LProgramName
	}

	info.LProgramAuthorEmail = config.LProgramAuthor.LProgramEmail

	return info
}
