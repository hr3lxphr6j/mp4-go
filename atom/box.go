package atom

import (
	"fmt"
	"io"
)

type Box interface {
	Header() *Header
	Boxes() []Box
}

type RawBox struct {
	Header *Header
	Body   io.Reader
}

func ReadRawBox(r io.Reader) (*RawBox, error) {
	hdr, err := ReadHeader(r)
	if err != nil {
		return nil, fmt.Errorf("ReadRawBox, %w", err)
	}
	return &RawBox{
		Header: hdr,
		Body:   io.LimitReader(r, int64(hdr.Size)),
	}, nil
}
