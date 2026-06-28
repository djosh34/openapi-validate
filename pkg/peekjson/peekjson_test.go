package peekjson

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPublicReaderReadFullyFromInternalBuffer(t *testing.T) {
	upstream := strings.NewReader("upstream")
	d := decoderWithBuffer(upstream, "buffer")
	p := make([]byte, len("buffer"))

	n, err := publicReader{d: d}.Read(p)

	require.NoError(t, err)
	require.Equal(t, len("buffer"), n)
	require.Equal(t, "buffer", string(p[:n]))
	require.Empty(t, d.lookAheadBuffer.String())
	require.Equal(t, "upstream", remainingString(t, upstream))
}

func TestPublicReaderReadFullyFromUpstream(t *testing.T) {
	upstream := strings.NewReader("upstream")
	d := decoderWithBuffer(upstream, "")
	p := make([]byte, len("up"))

	n, err := publicReader{d: d}.Read(p)

	require.NoError(t, err)
	require.Equal(t, len("up"), n)
	require.Equal(t, "up", string(p[:n]))
	require.Equal(t, "stream", remainingString(t, upstream))
}

func TestPublicReaderReadFillsFromInternalBufferThenPartialUpstream(t *testing.T) {
	upstream := strings.NewReader("stream")
	d := decoderWithBuffer(upstream, "buf")
	p := make([]byte, len("bufst"))

	n, err := publicReader{d: d}.Read(p)

	require.NoError(t, err)
	require.Equal(t, len("bufst"), n)
	require.Equal(t, "bufst", string(p[:n]))
	require.Empty(t, d.lookAheadBuffer.String())
	require.Equal(t, "ream", remainingString(t, upstream))
}

func TestPublicReaderReadDrainsInternalBufferAndUpstreamBeforeEOF(t *testing.T) {
	upstream := strings.NewReader("stream")
	d := decoderWithBuffer(upstream, "buf")
	p := make([]byte, len("bufstream")+10)

	n, err := publicReader{d: d}.Read(p)

	require.ErrorIs(t, err, io.EOF)
	require.Equal(t, len("bufstream"), n)
	require.Equal(t, "bufstream", string(p[:n]))
	require.Empty(t, d.lookAheadBuffer.String())
	require.Empty(t, remainingString(t, upstream))
}

func TestPublicReaderReadDrainsUpstreamBeforeEOF(t *testing.T) {
	upstream := strings.NewReader("upstream")
	d := decoderWithBuffer(upstream, "")
	p := make([]byte, len("upstream")+10)

	n, err := publicReader{d: d}.Read(p)

	require.ErrorIs(t, err, io.EOF)
	require.Equal(t, len("upstream"), n)
	require.Equal(t, "upstream", string(p[:n]))
	require.Empty(t, remainingString(t, upstream))
}

func TestPeekReaderBuffersReadBytes(t *testing.T) {
	upstream := strings.NewReader("peeked")
	d := &Decoder{sourceReader: upstream}
	p := make([]byte, len("peek"))

	n, err := peekReader{d: d}.Read(p)

	require.NoError(t, err)
	require.Equal(t, len("peek"), n)
	require.Equal(t, "peek", string(p[:n]))
	require.Equal(t, "peek", d.lookAheadBuffer.String())
	require.Equal(t, "ed", remainingString(t, upstream))
}

func TestPeekReaderBuffersBytesReturnedWithError(t *testing.T) {
	readErr := errors.New("read failed")
	d := &Decoder{sourceReader: errorAfterBytesReader{data: "peek", err: readErr}}
	p := make([]byte, len("peek"))

	n, err := peekReader{d: d}.Read(p)

	require.ErrorIs(t, err, readErr)
	require.Equal(t, len("peek"), n)
	require.Equal(t, "peek", string(p[:n]))
	require.Equal(t, "peek", d.lookAheadBuffer.String())
}

func decoderWithBuffer(upstream io.Reader, buffered string) *Decoder {
	d := &Decoder{sourceReader: upstream}
	_, _ = d.lookAheadBuffer.WriteString(buffered)
	return d
}

func remainingString(t *testing.T, r io.Reader) string {
	t.Helper()

	remaining, err := io.ReadAll(r)
	require.NoError(t, err)
	return string(remaining)
}

type errorAfterBytesReader struct {
	data string
	err  error
}

func (r errorAfterBytesReader) Read(p []byte) (int, error) {
	return copy(p, r.data), r.err
}
