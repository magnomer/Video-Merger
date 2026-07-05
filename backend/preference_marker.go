package backend

import (
	"errors"
	"regexp"
	"strings"
)

// LMarker is the resolved, runtime detection rule that decides how a filename
// is split into a shared group name and a part number. It is derived from
// LPreference and is never persisted directly.
type LMarker struct {
	LMarkerPattern    *regexp.Regexp
	LMarkerUnnumbered bool
}

// LMarkerPreset maps a preset id to a stem pattern. Each pattern matches the
// filename with its extension already removed, capturing the shared name in
// group 1 and the part number in group 2.
var LMarkerPreset = map[string]string{
	"paren":      `^(.+?)\s*\((\d+)\)$`,
	"bracket":    `^(.+?)\s*\[(\d+)\]$`,
	"underscore": `^(.+?)_(\d+)$`,
	"dash":       `^(.+?)\s*-\s*(\d+)$`,
	"space":      `^(.+?)\s+(\d+)$`,
	"dot":        `^(.+?)\.(\d+)$`,
}

const LMarkerDefault = "paren"

func LMarkerResolve(options LPreference) (LMarker, error) {
	pattern, err := LMarkerPatternResolve(options)
	if err != nil {
		return LMarker{}, err
	}

	return LMarker{
		LMarkerPattern:    pattern,
		LMarkerUnnumbered: options.LPreferenceUnnumbered,
	}, nil
}

func LMarkerPatternResolve(options LPreference) (*regexp.Regexp, error) {
	if options.LPreferenceCustom {
		expression := strings.TrimSpace(options.LPreferencePattern)
		if expression == "" {
			return nil, errors.New("custom detection pattern is empty")
		}

		pattern, err := regexp.Compile(expression)
		if err != nil {
			return nil, errors.New("custom detection pattern is not a valid regular expression")
		}

		if pattern.NumSubexp() < 2 {
			return nil, errors.New("custom detection pattern needs two capture groups: shared name and part number")
		}

		return pattern, nil
	}

	expression, ok := LMarkerPreset[options.LPreferenceMarker]
	if !ok {
		expression = LMarkerPreset[LMarkerDefault]
	}

	return regexp.MustCompile(expression), nil
}

func LMarkerDefaultCreate() LMarker {
	return LMarker{
		LMarkerPattern:    regexp.MustCompile(LMarkerPreset[LMarkerDefault]),
		LMarkerUnnumbered: false,
	}
}
