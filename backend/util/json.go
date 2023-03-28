package util

import (
	"bytes"
	"io"
)

func ReadEsBody(res io.Reader) string {
	b := bytes.Buffer{}
	b.ReadFrom(res)
	return b.String()
}
