package conf

import (
	"jrpc/modules"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()

	modules.InitAllModules()
}
