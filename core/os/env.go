package os

import (
	"strings"
)

func EnvMapToSlices(envMap map[string]string) []string {
	var result = make([]string, 0, len(envMap))

	for k, v := range envMap {
		result = append(result, k+"="+v)
	}

	return result
}

func MergeEnvAsMap(env []string, additional map[string]string) map[string]string {
	var result = make(map[string]string, len(env))

	for _, v := range env {
		envSplit := strings.SplitN(v, "=", 2) //nolint:mnd // _
		result[envSplit[0]] = envSplit[1]
	}

	for k, v := range additional {
		result[k] = v
	}

	return result
}
