package backend

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

type LPreference struct {
	LPreferenceInput      []string `json:"LPreferenceInput"`
	LPreferenceOutput     string   `json:"LPreferenceOutput"`
	LPreferenceMirror     bool     `json:"LPreferenceMirror"`
	LPreferenceTree       bool     `json:"LPreferenceTree"`
	LPreferenceSuffix     string   `json:"LPreferenceSuffix"`
	LPreferenceCaution    bool     `json:"LPreferenceCaution"`
	LPreferenceWarning    bool     `json:"LPreferenceWarning"`
	LPreferenceMarker     string   `json:"LPreferenceMarker"`
	LPreferencePattern    string   `json:"LPreferencePattern"`
	LPreferenceCustom     bool     `json:"LPreferenceCustom"`
	LPreferenceUnnumbered bool     `json:"LPreferenceUnnumbered"`
}

func LPreferenceLoad() (LPreference, error) {
	path, err := LPreferencePathResolve()
	if err != nil {
		return LPreference{}, err
	}

	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return LPreference{}, nil
	}
	if err != nil {
		return LPreference{}, err
	}

	var preference LPreference
	if err := json.Unmarshal(data, &preference); err != nil {
		return LPreference{}, err
	}

	return preference, nil
}

func LPreferenceSave(preference LPreference) error {
	path, err := LPreferencePathResolve()
	if err != nil {
		return err
	}

	preference.LPreferenceSuffix = strings.TrimSpace(preference.LPreferenceSuffix)
	preference.LPreferencePattern = strings.TrimSpace(preference.LPreferencePattern)

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(preference, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func LPreferencePathResolve() (string, error) {
	configRoot, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(configRoot, "Video Merger", "preference.json"), nil
}
