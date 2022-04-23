package aws

import (
	"context"
	"fmt"
	"log"
	"path"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/yuyuvn/kms-decrypter/pkg/file"
	"github.com/yuyuvn/kms-decrypter/pkg/message"
)

type AwsDecrypter struct {
}

type KMSDecryptAPI interface {
	Decrypt(ctx context.Context,
		params *kms.DecryptInput,
		optFns ...func(*kms.Options)) (*kms.DecryptOutput, error)
}

func (ad *AwsDecrypter) Init() {

}

// DecodeData decrypts some text that was encrypted with an AWS Key Management Service (AWS KMS) customer master key (CMK).
// Inputs:
//     ctx is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If success, a DecryptOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to Decrypt.
func (ad *AwsDecrypter) decodeData(ctx context.Context, api KMSDecryptAPI, input *kms.DecryptInput) (*kms.DecryptOutput, error) {
	return api.Decrypt(ctx, input)
}

// decrypt decrypt text then return content
// Inputs:
//     ctx is the context of the method call, which includes the AWS Region.
//     data is the encoded string.
// Output:
//     If success, a Decrypt object containing the decrypted string and nil.
//     Otherwise, nil and an error from the call to Decrypt.
func (ad *AwsDecrypter) decrypt(ctx context.Context, blob []byte) ([]byte, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := kms.NewFromConfig(cfg)

	input := &kms.DecryptInput{
		CiphertextBlob: blob,
	}

	result, err := ad.decodeData(ctx, client, input)
	if err != nil {
		return nil, err
	}

	return result.Plaintext, nil
}

// Decrypt get input from bus then decrypt file content and write to output
// Inputs:
// 		ctx is the context of the method call
//		bus is the channel from foler walker
//		source is encrypted files folder
//		target is output destination
//		quiet if enabled, no output if no error
func (ad *AwsDecrypter) Decrypt(ctx context.Context, bus message.Bus, source string, target string, quiet bool) error {
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

			filePath := payload.FilePath
			blob, err := file.Read(ctx, path.Join(source, filePath))
			if err != nil {
				log.Fatalln("can't read file:", path.Join(source, filePath), err)
			}

			content, err := ad.decrypt(ctx, blob)
			if err != nil {
				log.Fatalln("can't decode file:", path.Join(source, filePath), err)
			}

			if err := file.Write(ctx, path.Join(target, filePath), content); err != nil {
				log.Fatalln("can't write to file:", filePath, err)
			}

			if !quiet {
				fmt.Println("Decrypted: ", filePath)
			}
		}
	}

	return nil
}
