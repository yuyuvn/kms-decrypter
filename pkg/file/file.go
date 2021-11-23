package file

import (
	"context"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/yuyuvn/kms-decrypter/pkg/message"
)

// List walk then send detected file to bus
func List(ctx context.Context, root string, bus message.Bus) error {
	handler := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			rel, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}

			select {
			case <-ctx.Done():
				return ctx.Err()
			case bus <- message.Payload{FilePath: rel}:
				// TODO: log something?
			}
		}
		return nil
	}

	return filepath.WalkDir(root, handler)
}

// Read get file content
func Read(ctx context.Context, path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln("Can't read file ", path, err)
	}

	return content, nil
}

// Write write file content and folder if needed
func Write(ctx context.Context, path string, content []byte) error {
	os.MkdirAll(filepath.Dir(path), os.ModePerm)

	return os.WriteFile(path, content, 0644)
}
