package main

import (
	"github.com/yuyuvn/kms-decrypter/pkg/config"
	"github.com/yuyuvn/kms-decrypter/pkg/run"
)

func main() {
	// parse config flags, will exit in case of errors.
	cfg := config.Parse()

	run.Run(cfg)
}
