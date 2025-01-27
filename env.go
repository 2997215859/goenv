package env

import "os"

const ENV_DEV = "dev"
const ENV_TEST = "test"
const ENV_PROD = "prod"

func ENV() string {
	runtimeEnv := os.Getenv("runtime_env")
	if runtimeEnv == "" {
		runtimeEnv = ENV_DEV
	}
	return runtimeEnv
}
