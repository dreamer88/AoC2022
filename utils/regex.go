package utils

import "regexp"

func MatchNamedGroups(reg *regexp.Regexp, str string) map[string]string {
	match := reg.FindStringSubmatch(str)
	result := make(map[string]string)

	for i, name := range reg.SubexpNames()[1:] {
		result[name] = match[i]
	}

	return result
}
