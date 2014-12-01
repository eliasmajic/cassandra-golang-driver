package gocql

import (
	//"./snappy"
)

type Compressor interface {
	Name() string
	Encode(data []byte) ([]byte, error)
	Decode(data []byte) ([]byte, error)
}

// SnappyCompressor implements the Compressor interface and can be used to
// compress incoming and outgoing frames. The snappy compression algorithm
// aims for very high speeds and reasonable compression.
type SnappyCompressor struct{}

func (s SnappyCompressor) Name() string {
	return "snappy"
}

func (s SnappyCompressor) Encode(data []byte) ([]byte, error) {
	return data, nil
	//return snappy.Encode(nil, data)
}

func (s SnappyCompressor) Decode(data []byte) ([]byte, error) {
	return data, nil
	//return snappy.Decode(nil, data)
}
