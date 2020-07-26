package atom

import (
	"bytes"
	"fmt"
	"io"
)

type HeaderType string

const (
	HeaderTypeFTYP HeaderType = "ftyp"
	HeaderTypeMOOV HeaderType = "moov"
	HeaderTypeUUID HeaderType = "uuid"
)

var HeaderTypeMap = map[string]HeaderType{
	string(HeaderTypeFTYP): HeaderTypeFTYP,
	string(HeaderTypeMOOV): HeaderTypeMOOV,
	string(HeaderTypeUUID): HeaderTypeUUID,
}

type Header struct {
	Size uint64
	Type HeaderType

	IsFullBoxHeader bool
	Version         byte
	Flags           [3]byte
}

func ReadHeader(r io.Reader) (*Header, error) {
	buf := new(bytes.Buffer)
	_, err := io.CopyN(buf, r, 8)
	if err != nil {
		return nil, fmt.Errorf("ReadHeader[read size & type failed], %w", err)
	}
	hdr := &Header{
		Size: uint64(ByteOrder.Uint32(buf.Bytes()[0:4])),
		// TODO: parse unknown type of box.
		Type: HeaderType(buf.Bytes()[4:]),
	}
	if hdr.Size == 1 {
		buf.Reset()
		_, err = io.CopyN(buf, r, 8)
		if err != nil {
			return nil, fmt.Errorf("ReadHeader[read largesize failed], %w", err)
		}
		hdr.Size = ByteOrder.Uint64(buf.Bytes())
	}
	if hdr.Type == HeaderTypeUUID {
		hdr.IsFullBoxHeader = true
		buf.Reset()
		_, err = io.CopyN(buf, r, 4)
		if err != nil {
			return nil, fmt.Errorf("ReadHeader[read full-box-header`s version & flags failed], %w", err)
		}
		hdr.Version = buf.Bytes()[0]
		copy(hdr.Flags[:], buf.Bytes()[1:])
	}
	return hdr, nil
}

func WriteHeader(w io.Writer, hdr *Header) error {
	panic("impl me")
}
