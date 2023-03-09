package compression

type Compression interface {
	Compress(b []byte) ([]byte, error)
}
