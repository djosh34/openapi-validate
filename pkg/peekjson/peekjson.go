package peekjson

import (
	"bytes"
	"encoding/json"
	"io"
)

type Decoder struct {
	json.Decoder

	sourceReader    io.Reader
	lookAheadBuffer bytes.Buffer

	peekDec     *json.Decoder
	peekedToken *json.Token
	peekedErr   error
}

func NewDecoder(r io.Reader) *Decoder {
	d := &Decoder{sourceReader: r}

	d.Decoder = *json.NewDecoder(publicReader{d})
	d.peekDec = json.NewDecoder(peekReader{d})
	d.peekDec.UseNumber()

	return d
}

func (d *Decoder) Peek() (*json.Token, error) {
	if d.peekedToken != nil {
		return d.peekedToken, d.peekedErr
	}

	tok, err := d.peekDec.Token()
	if err != nil {
		d.peekedErr = err
		return nil, err
	}

	return &tok, err
}

type peekReader struct {
	d *Decoder
}

func (r peekReader) Read(p []byte) (int, error) {
	n, err := r.d.sourceReader.Read(p)
	if err != nil {
		return n, err
	}

	return r.d.lookAheadBuffer.Write(p[:n])
}

type publicReader struct {
	d *Decoder
}

func (r publicReader) Read(p []byte) (int, error) {
	d := r.d

	d.peekedToken = nil
	d.peekedErr = nil

	// TODO
	// Read first from buffer, then wrong read reader
	return 0, nil
}
