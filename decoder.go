package aliyun_opensearch_go_sdk

import (
	"encoding/xml"
	"io"
)

type OpenSearchDecoder interface {
	Decode(reader io.Reader, v interface{}) (err error)
	DecodeError(bodyBytes []byte, resource string) (decodedError error, err error)
	Test() bool
}

type xmlDecoder struct {
}

func NewOpenSearchDecoder() OpenSearchDecoder {
	return &xmlDecoder{}
}

func (p *xmlDecoder) Test() bool {
	return false
}

func (p *xmlDecoder) Decode(reader io.Reader, v interface{}) (err error) {
	decoder := xml.NewDecoder(reader)
	err = decoder.Decode(&v)

	return
}

func (p *xmlDecoder) DecodeError(bodyBytes []byte, resource string) (decodedError error, err error) {
	return
}

func ParseError(resource string) (err error) {
	return
}