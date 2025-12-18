// Package event provides encoding and decoding for events.
package event

// Encoder defines the interface for encoding events to bytes.
type Encoder interface {
	// Encode serializes an event to bytes.
	Encode(event *ReviewEvent) ([]byte, error)
}

// Decoder defines the interface for decoding events from bytes.
type Decoder interface {
	// Decode deserializes bytes to an event.
	Decode(data []byte) (*ReviewEvent, error)
}

// Codec provides both encoding and decoding capabilities.
type Codec interface {
	Encoder
	Decoder
}

// JSONCodec implements Codec using JSON serialization.
type JSONCodec struct{}

// NewJSONCodec creates a new JSON codec.
func NewJSONCodec() *JSONCodec {
	return &JSONCodec{}
}

// Encode serializes a ReviewEvent to JSON bytes.
func (c *JSONCodec) Encode(event *ReviewEvent) ([]byte, error) {
	// TODO: Marshal event to JSON
	panic("TODO")
}

// Decode deserializes JSON bytes to a ReviewEvent.
func (c *JSONCodec) Decode(data []byte) (*ReviewEvent, error) {
	// TODO: Unmarshal JSON to event
	panic("TODO")
}

