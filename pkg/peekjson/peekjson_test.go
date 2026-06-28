package peekjson

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReaderRead(t *testing.T) {
	readErr := errors.New("read failed")
	publicRead := func(d *Decoder) io.Reader { return publicReader{d: d} }
	peekRead := func(d *Decoder) io.Reader { return peekReader{d: d} }

	for name, tt := range map[string]struct {
		reader                  func(*Decoder) io.Reader
		bytesToRead             int
		internalBuffer          string
		sourceReader            io.Reader
		expectedOutput          string
		expectedLookAheadBuffer string
		expectedRemainingSource string
		expectedBytesRead       int
		expectedErr             error
	}{
		"public reader fully from internal buffer": {
			reader:                  publicRead,
			bytesToRead:             len("buffer"),
			internalBuffer:          "buffer",
			sourceReader:            strings.NewReader("upstream"),
			expectedOutput:          "buffer",
			expectedLookAheadBuffer: "",
			expectedBytesRead:       len("buffer"),
			expectedErr:             nil,
		},
		"public reader fully from upstream": {
			reader:                  publicRead,
			bytesToRead:             len("up"),
			internalBuffer:          "",
			sourceReader:            strings.NewReader("upstream"),
			expectedOutput:          "up",
			expectedLookAheadBuffer: "",
			expectedBytesRead:       len("up"),
			expectedErr:             nil,
		},
		"public reader fills from internal buffer then partial upstream": {
			reader:                  publicRead,
			bytesToRead:             len("bufst"),
			internalBuffer:          "buf",
			sourceReader:            strings.NewReader("stream"),
			expectedOutput:          "bufst",
			expectedLookAheadBuffer: "",
			expectedBytesRead:       len("bufst"),
			expectedErr:             nil,
		},
		"public reader drains internal buffer and upstream before eof": {
			reader:                  publicRead,
			bytesToRead:             len("bufstream") + 10,
			internalBuffer:          "buf",
			sourceReader:            strings.NewReader("stream"),
			expectedOutput:          "bufstream",
			expectedLookAheadBuffer: "",
			expectedBytesRead:       len("bufstream"),
			expectedErr:             nil,
		},
		"public reader drains upstream before eof": {
			reader:                  publicRead,
			bytesToRead:             len("upstream") + 10,
			internalBuffer:          "",
			sourceReader:            strings.NewReader("upstream"),
			expectedOutput:          "upstream",
			expectedLookAheadBuffer: "",
			expectedBytesRead:       len("upstream"),
			expectedErr:             nil,
		},
		"peek reader buffers read bytes": {
			reader:                  peekRead,
			bytesToRead:             len("peek"),
			sourceReader:            strings.NewReader("peeked"),
			expectedOutput:          "peek",
			expectedLookAheadBuffer: "peek",
			expectedRemainingSource: "ed",
			expectedBytesRead:       len("peek"),
			expectedErr:             nil,
		},
		"peek reader buffers bytes returned with error": {
			reader:                  peekRead,
			bytesToRead:             len("peek"),
			sourceReader:            errorAfterBytesReader{data: "peek", err: readErr},
			expectedOutput:          "peek",
			expectedLookAheadBuffer: "peek",
			expectedBytesRead:       len("peek"),
			expectedErr:             readErr,
		},
	} {
		t.Run(name, func(t *testing.T) {
			// Arrange
			d := &Decoder{sourceReader: tt.sourceReader}
			d.lookAheadBuffer = *bytes.NewBufferString(tt.internalBuffer)

			p := make([]byte, tt.bytesToRead)

			// Act
			n, err := tt.reader(d).Read(p)

			// Assert
			if tt.expectedErr == nil {
				require.NoError(t, err)
			} else {
				require.ErrorIs(t, err, tt.expectedErr)
			}

			require.Equal(t, tt.expectedBytesRead, n)
			require.Equal(t, tt.expectedOutput, string(p[:n]))
			require.Equal(t, tt.expectedLookAheadBuffer, d.lookAheadBuffer.String())

			if tt.expectedRemainingSource != "" {
				require.Equal(t, tt.expectedRemainingSource, remainingString(t, tt.sourceReader))
			}
		})
	}
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
