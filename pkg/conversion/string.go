package conversion

import "strings"

func StringToBoolPtr(s string) *bool {
	trueValues := map[string]bool{"true": true, "1": true, "yes": true, "on": true}
	falseValues := map[string]bool{"false": false, "0": false, "no": false, "off": false}

	normalized := strings.ToLower(strings.TrimSpace(s))

	if val, ok := trueValues[normalized]; ok {
		return &val
	} else if val, ok := falseValues[normalized]; ok {
		return &val
	}

	return nil
}
