package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DocumentKeyPrefix is the prefix to retrieve all Document
	DocumentKeyPrefix = "Document/value/"
)

// DocumentKey returns the store key to retrieve a Document from the index fields
func DocumentKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
