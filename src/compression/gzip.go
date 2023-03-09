package compression

import (
	"bytes"
	"compress/gzip"
)

// Gzip implements the Compression interface
type Gzip struct {
}

func (compression *Gzip) Compress(b []byte) ([]byte, error) {
	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)
	_, e := writer.Write(b)
	if e != nil {
		return nil, e
	}
	writer.Close()
	return buf.Bytes(), nil
}
