package chess

import (
	"bytes"
	"compress/gzip"
	"io"

	http "github.com/Carcraftz/fhttp"

	"github.com/dsnet/compress/brotli"
)

func decompressBrotli(data []byte) (resData []byte, err error) {
	br, err := brotli.NewReader(bytes.NewReader(data), &brotli.ReaderConfig{})
	if err != nil {
		return []byte{}, err
	}

	respBody, err := io.ReadAll(br)
	if err != nil {
		return []byte{}, err
	}

	return respBody, nil
}

// DecodeBody takes in the headers and the body reader from a request and correctly decodes it accordingly
func (c *Client) DecodeBody(headers http.Header, body io.Reader) ([]byte, error) {
	if len(headers["Content-Encoding"]) == 0 {
		b, err := io.ReadAll(body)
		if err != nil {
			return []byte{}, err
		}
		return b, err
	}
	if headers["Content-Encoding"][0] == "gzip" {
		reader, err := gzip.NewReader(body)
		if err != nil {
			return []byte{}, err
		}
		b, err := io.ReadAll(reader)
		if err != nil {
			return []byte{}, err
		}
		return b, err
	}
	bb, err := io.ReadAll(body)
	if err != nil {
		return []byte{}, err
	}
	b, err := decompressBrotli(bb)
	if err != nil {
		return []byte{}, err
	}
	return b, err
}
