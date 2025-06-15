package compress

import (
	"bytes"
	"compress/gzip"
	"io"
)

// Compress applies GZIP compression to the input data.
func Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)

	_, err := gz.Write(data)
	if err != nil {
		return nil, err
	}

	err = gz.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Decompress reverses GZIP compression on the input data.
func Decompress(data []byte) ([]byte, error) {
	buf := bytes.NewReader(data)
	gz, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}
	defer gz.Close()

	var out bytes.Buffer
	_, err = io.Copy(&out, gz)
	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}
