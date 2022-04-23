package sops

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os/exec"
	"path"

	"github.com/yuyuvn/kms-decrypter/pkg/file"
	"github.com/yuyuvn/kms-decrypter/pkg/message"
)

type SopsDecrypter struct {
}

func (sd *SopsDecrypter) Init() {

}

// Decrypt get input from bus then decrypt file content and write to output
// Inputs:
// 		ctx is the context of the method call
//		bus is the channel from foler walker
//		source is encrypted files folder
//		target is output destination
//		quiet if enabled, no output if no error
func (sd *SopsDecrypter) Decrypt(ctx context.Context, bus message.Bus, source string, target string, quiet bool) error {
	for bus != nil {
		select {
		// Exit early if context done.
		case <-ctx.Done():
			return ctx.Err()
		// Get Messages from Bus
		case payload, ok := <-bus:
			if !ok {
				bus = nil
				break
			}

			var content bytes.Buffer
			filePath := payload.FilePath

			cmd := exec.Command("sops", "--decrypt", path.Join(source, filePath))
			cmd.Stdout = &content

			if err := cmd.Run(); err != nil {
				log.Fatalln("can't decode file:", path.Join(source, filePath), err)
			}

			if err := file.Write(ctx, path.Join(target, filePath), content.Bytes()); err != nil {
				log.Fatalln("can't write to file:", filePath, err)
			}

			if !quiet {
				fmt.Println("Decrypted: ", filePath)
			}
		}
	}

	return nil
}
