package environment

import (
	"flag"
	"os"
	"strings"
)

// AppEnvironment enum
type AppEnvironment string

// const declaration of environment
const (
	LOCAL   AppEnvironment = "local"
	DEV     AppEnvironment = "development"
	PROD    AppEnvironment = "production"
	STAGING AppEnvironment = "staging"
)

// declare of error getting environment
var (
	strEnv = ""
)

// IsLocal method of AppEnvironment
func (env AppEnvironment) IsLocal() bool {
	return env == LOCAL
}

// IsDev method of AppEnvironment
func (env AppEnvironment) IsDev() bool {
	return env == DEV
}

// IsStaging method of AppEnvironment
func (env AppEnvironment) IsStaging() bool {
	return env == STAGING
}

// IsProd method of AppEnvironment
func (env AppEnvironment) IsProd() bool {
	return env == PROD
}

// GetEnvFlag func to get command line flag "env"
func getEnvFlag(name, value, usage string) string {
	var flagEnv string
	flag.StringVar(&flagEnv, name, value, usage)
	flag.Parse()
	return flagEnv
}

// FromOsEnv func
func FromOsEnv() (AppEnvironment, error) {
	if strEnv == "" {
		strEnv = strings.Trim(strings.ToLower(os.Getenv("APP_ENV")), " ")
		if flag.Lookup("env") == nil {
			strEnv = getEnvFlag("env", "", "env Mode project local, development, staging, production")
		}
	}

	switch strEnv {
	case "development":
		return DEV, nil
	case "production":
		return PROD, nil
	case "staging":
		return STAGING, nil
	}

	return LOCAL, nil
}
