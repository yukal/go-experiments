package gob

import (
	"fmt"
	"regexp"
)

type RegexpWrapper struct {
	*regexp.Regexp
}

func (r RegexpWrapper) GobEncode() ([]byte, error) {
	if r.Regexp == nil {
		return nil, nil
	}

	return []byte(r.Regexp.String()), nil
}

func (r *RegexpWrapper) GobDecode(data []byte) error {
	if len(data) == 0 {
		r.Regexp = nil
		return nil
	}

	pattern := string(data)

	if compiledRegexp, err := regexp.Compile(pattern); err != nil {
		return fmt.Errorf("unable compile regexp '%s': %w", pattern, err)
	} else {
		r.Regexp = compiledRegexp
	}

	return nil
}
