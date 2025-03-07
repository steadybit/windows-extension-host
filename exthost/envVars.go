// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2025 Steadybit GmbH

package exthost

import (
	"github.com/rs/zerolog/log"
	"golang.org/x/exp/slices"
	"os"
	"regexp"
	"strings"
)

const SteadybitLabelPrefix = "steadybit_label_"

func readWhitelistedEnvVars() []string {
	var whitelistedEnvVars []string
	discoveryEnvList, ok := os.LookupEnv("STEADYBIT_DISCOVERY_ENV_LIST")
	if ok {
		re := regexp.MustCompile(`[\\s,;]+`)
		whitelistedEnvVars = re.Split(strings.Replace(strings.ToLower(discoveryEnvList), " ", "", -1), -1)
	} else {
		log.Trace().Msg("STEADYBIT_DISCOVERY_ENV_LIST not set, using default")
		whitelistedEnvVars = []string{}
	}
	return whitelistedEnvVars
}

func getEnvironmentVariables() map[string]string {
	envVars := make(map[string]string)
	env := os.Environ()
	whitelistedEnvVars := readWhitelistedEnvVars()
	for _, e := range env {
		pair := strings.Split(e, "=")
		key := strings.ToLower(pair[0])
		if slices.Contains(whitelistedEnvVars, key) {
			envVars[key] = pair[1]
		}
	}
	return envVars
}
func getLabels() map[string]string {
	labels := map[string]string{}
	env := os.Environ()
	for _, e := range env {
		pair := strings.Split(e, "=")
		key := strings.ToLower(pair[0])
		if strings.HasPrefix(key, SteadybitLabelPrefix) {
			labels[key[len(SteadybitLabelPrefix):]] = pair[1]
		}
	}
	return labels
}
