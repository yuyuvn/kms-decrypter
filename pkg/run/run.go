package run

import (
	"context"
	"log"

	"github.com/yuyuvn/kms-decrypter/pkg/aws"
	"github.com/yuyuvn/kms-decrypter/pkg/config"
	"github.com/yuyuvn/kms-decrypter/pkg/file"
	"github.com/yuyuvn/kms-decrypter/pkg/message"
	"golang.org/x/sync/errgroup"
)

func Run(cfg config.Config) {
	// Create shared message bus
	ch := make(message.Bus, 10000)
	g, gctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		file.List(gctx, cfg.Source, ch)
		close(ch)
		return nil
	})

	for i := 0; i < cfg.Concurrency; i++ {
		g.Go(func() error {
			return aws.Decrypt(gctx, ch, cfg.Source, cfg.Target, cfg.Quiet)
		})
	}

	// Block and wait for goroutines
	err := g.Wait()
	if err != nil && err != context.Canceled {
		log.Fatalln(err)
	}
}
