package env

import (
	"os"
)

const RUN_EVN_KEY = "WNCBB_RUN_ENV"

const (
	pro = "pro"
	dev = "dev"
)

var runEnv string
var isProduct bool

func Init() {
	runEnvStr := os.Getenv(RUN_EVN_KEY)
	if runEnvStr == pro {
		isProduct = true
		return
	}
	isProduct = false
	runEnv = dev
}

func IsProduct() bool {
	return isProduct
}
