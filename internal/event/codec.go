// Package event provides encoding and decoding for events.
//
// YOUR TASK (Milestone 1):
// Implement JSON serialization for events.
package event

// TODO: Import "encoding/json"

// Codec handles encoding and decoding of events.
// TODO: Implement these methods:
//
// Encode(event *Event) ([]byte, error)
//   - Use json.Marshal to convert event to JSON bytes
//   - Return the bytes and any error
//
// Decode(data []byte) (*Event, error)
//   - Use json.Unmarshal to parse JSON bytes into Event
//   - Return pointer to event and any error
//
// Example usage:
//   codec := &JSONCodec{}
//   bytes, err := codec.Encode(event)
//   event, err := codec.Decode(bytes)
