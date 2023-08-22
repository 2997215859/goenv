package env

import "os"

const ENV_DEV = "dev"
const ENV_TEST = "test"
const ENV_PROD = "prod"

func ENV() string {
	runtimeEnv := os.Getenv("RUNTIME_ENV")
	if runtimeEnv == "" {
		runtimeEnv = ENV_DEV
	}
	return runtimeEnv
}

