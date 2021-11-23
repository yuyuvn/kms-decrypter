// Package message represents the message bus.
// Message Payloads pass through a Bus channel.
package message

// Payload represents a encrypted file
type Payload struct {
	FilePath string
}

// Bus is a channel where message Payloads pass.
type Bus chan Payload
