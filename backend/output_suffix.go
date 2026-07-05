package backend

import (
	"errors"
	"strings"
)

func LSuffixCheck(suffix string) error {
	if suffix == "" {
		return nil
	}

	invalidCharacters := []string{`<`, `>`, `:`, `"`, `/`, `\`, `|`, `?`, `*`}

	for _, invalid := range invalidCharacters {
		if strings.Contains(suffix, invalid) {
			return errors.New("suffix contains a character that cannot be used in a Windows filename")
		}
	}

	return nil
}
