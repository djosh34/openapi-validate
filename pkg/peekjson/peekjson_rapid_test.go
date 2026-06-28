package peekjson

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"

	"pgregory.net/rapid"
)

type rapidDecoderOp uint8

const (
	rapidOpDecode rapidDecoderOp = iota
	rapidOpToken
	rapidOpMore
	rapidOpInputOffset
)

func TestRapidDecoderMatchesEncodingJSON(t *testing.T) {
	t.Parallel()

	rapid.Check(t, func(rt *rapid.T) {
		input := drawChaoticJSONStream(rt)
		useNumber := rapid.Bool().Draw(rt, "use number")
		opCount := rapid.IntRange(1, 80).Draw(rt, "op count")

		want := json.NewDecoder(strings.NewReader(input))
		got := NewDecoder(strings.NewReader(input))
		if useNumber {
			want.UseNumber()
			got.UseNumber()
		}

		history := make([]rapidDecoderOp, 0, opCount)
		for i := range opCount {
			assertRapidPeekBurst(rt, input, useNumber, history, got, fmt.Sprintf("before op %d", i))

			op := drawRapidDecoderOp(rt, fmt.Sprintf("op %d", i))
			record, stop := assertRapidPublicOpMatches(rt, op, want, got)
			if record {
				history = append(history, op)
			}
			if stop {
				return
			}
		}

		assertRapidPeekBurst(rt, input, useNumber, history, got, "after ops")
	})
}

func drawRapidDecoderOp(t *rapid.T, label string) rapidDecoderOp {
	switch rapid.IntRange(0, 9).Draw(t, label+" kind") {
	case 0, 1, 2, 3:
		return rapidOpDecode
	case 4, 5, 6, 7:
		return rapidOpToken
	case 8:
		return rapidOpMore
	default:
		return rapidOpInputOffset
	}
}

func assertRapidPeekBurst(
	t *rapid.T,
	input string,
	useNumber bool,
	history []rapidDecoderOp,
	got *Decoder,
	label string,
) {
	count := rapid.IntRange(0, 8).Draw(t, label+" peek count")
	for i := range count {
		wantTok, wantErr := rapidPeekOracle(input, useNumber, history)
		gotTok, gotErr := got.Peek()

		assertRapidErrorsEqual(t, wantErr, gotErr, label, "peek")
		if wantErr == nil {
			if gotTok == nil {
				t.Fatalf("%s peek %d: Peek returned nil token pointer without error", label, i)
			}
			assertRapidValuesEqual(t, wantTok, *gotTok, label, "peek token")
		}
	}
}

func rapidPeekOracle(input string, useNumber bool, history []rapidDecoderOp) (json.Token, error) {
	dec := json.NewDecoder(strings.NewReader(input))
	if useNumber {
		dec.UseNumber()
	}

	for _, op := range history {
		switch op {
		case rapidOpDecode:
			var v any
			if err := dec.Decode(&v); err != nil {
				return nil, fmt.Errorf("replay decode: %w", err)
			}
		case rapidOpToken:
			if _, err := dec.Token(); err != nil {
				return nil, fmt.Errorf("replay token: %w", err)
			}
		case rapidOpMore, rapidOpInputOffset:
		default:
			return nil, fmt.Errorf("unknown replay op %d", op)
		}
	}

	return dec.Token()
}

func assertRapidPublicOpMatches(
	t *rapid.T,
	op rapidDecoderOp,
	want *json.Decoder,
	got *Decoder,
) (record bool, stop bool) {
	switch op {
	case rapidOpDecode:
		var wantValue any
		var gotValue any
		wantErr := want.Decode(&wantValue)
		gotErr := got.Decode(&gotValue)

		assertRapidErrorsEqual(t, wantErr, gotErr, "public op", "Decode")
		assertRapidValuesEqual(t, wantValue, gotValue, "public op", "Decode value")

		return wantErr == nil, wantErr != nil && !errors.Is(wantErr, io.EOF)
	case rapidOpToken:
		wantTok, wantErr := want.Token()
		gotTok, gotErr := got.Token()

		assertRapidErrorsEqual(t, wantErr, gotErr, "public op", "Token")
		if wantErr == nil {
			assertRapidValuesEqual(t, wantTok, gotTok, "public op", "Token value")
		}

		return wantErr == nil, wantErr != nil && !errors.Is(wantErr, io.EOF)
	case rapidOpMore:
		wantMore := want.More()
		gotMore := got.More()
		assertRapidValuesEqual(t, wantMore, gotMore, "public op", "More")

		return false, false
	case rapidOpInputOffset:
		wantOffset := want.InputOffset()
		gotOffset := got.InputOffset()
		assertRapidValuesEqual(t, wantOffset, gotOffset, "public op", "InputOffset")

		return false, false
	default:
		t.Fatalf("unknown op %d", op)

		return false, true
	}
}

func assertRapidErrorsEqual(t *rapid.T, want error, got error, scope string, op string) {
	if want == nil && got == nil {
		return
	}
	if want == nil || got == nil {
		t.Fatalf("%s %s: error mismatch: want %v, got %v", scope, op, want, got)
	}
	if reflect.TypeOf(want) != reflect.TypeOf(got) || want.Error() != got.Error() {
		t.Fatalf(
			"%s %s: error mismatch:\nwant (%T) %q\ngot  (%T) %q",
			scope,
			op,
			want,
			want.Error(),
			got,
			got.Error(),
		)
	}
}

func assertRapidValuesEqual(t *rapid.T, want any, got any, scope string, op string) {
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("%s %s mismatch:\nwant (%T) %#v\ngot  (%T) %#v", scope, op, want, want, got, got)
	}
}

func drawChaoticJSONStream(t *rapid.T) string {
	maxDepth := rapid.IntRange(0, 5).Draw(t, "max depth")
	valueCount := rapid.IntRange(1, 6).Draw(t, "top-level value count")

	var b strings.Builder
	b.WriteString(drawJSONWhitespace(t, "stream prefix whitespace"))
	for i := range valueCount {
		b.WriteString(drawJSONValue(t, maxDepth, fmt.Sprintf("root %d", i)))
		b.WriteString(drawJSONWhitespace(t, fmt.Sprintf("root %d suffix whitespace", i)))
	}

	input := b.String()
	switch rapid.IntRange(0, 9).Draw(t, "chaos mode") {
	case 0:
		cut := rapid.IntRange(0, len(input)).Draw(t, "truncate position")

		return input[:cut]
	case 1:
		pos := rapid.IntRange(0, len(input)).Draw(t, "insert junk position")

		return input[:pos] + drawJSONJunk(t, "insert junk") + input[pos:]
	case 2:
		return input + drawJSONJunk(t, "append junk")
	case 3:
		if len(input) == 0 {
			return input
		}
		start := rapid.IntRange(0, len(input)-1).Draw(t, "delete start")
		end := rapid.IntRange(start+1, len(input)).Draw(t, "delete end")

		return input[:start] + input[end:]
	default:
		return input
	}
}

func drawJSONValue(t *rapid.T, depth int, label string) string {
	maxKind := 3
	if depth > 0 {
		maxKind = 5
	}

	switch rapid.IntRange(0, maxKind).Draw(t, label+" kind") {
	case 0:
		return "null"
	case 1:
		if rapid.Bool().Draw(t, label+" bool") {
			return "true"
		}

		return "false"
	case 2:
		return drawJSONStringLiteral(t, label+" string")
	case 3:
		return drawJSONNumber(t, label+" number")
	case 4:
		return drawJSONArray(t, depth, label+" array")
	default:
		return drawJSONObject(t, depth, label+" object")
	}
}

func drawJSONArray(t *rapid.T, depth int, label string) string {
	length := rapid.IntRange(0, 8).Draw(t, label+" len")

	var b strings.Builder
	b.WriteByte('[')
	b.WriteString(drawJSONWhitespace(t, label+" open whitespace"))
	for i := range length {
		if i > 0 {
			b.WriteString(drawJSONWhitespace(t, fmt.Sprintf("%s comma %d left whitespace", label, i)))
			b.WriteByte(',')
			b.WriteString(drawJSONWhitespace(t, fmt.Sprintf("%s comma %d right whitespace", label, i)))
		}
		b.WriteString(drawJSONValue(t, depth-1, fmt.Sprintf("%s item %d", label, i)))
	}
	b.WriteString(drawJSONWhitespace(t, label+" close whitespace"))
	b.WriteByte(']')

	return b.String()
}

func drawJSONObject(t *rapid.T, depth int, label string) string {
	length := rapid.IntRange(0, 8).Draw(t, label+" len")

	var b strings.Builder
	b.WriteByte('{')
	b.WriteString(drawJSONWhitespace(t, label+" open whitespace"))
	for i := range length {
		if i > 0 {
			b.WriteString(drawJSONWhitespace(t, fmt.Sprintf("%s comma %d left whitespace", label, i)))
			b.WriteByte(',')
			b.WriteString(drawJSONWhitespace(t, fmt.Sprintf("%s comma %d right whitespace", label, i)))
		}
		b.WriteString(drawJSONStringLiteral(t, fmt.Sprintf("%s key %d", label, i)))
		b.WriteString(drawJSONWhitespace(t, fmt.Sprintf("%s colon %d left whitespace", label, i)))
		b.WriteByte(':')
		b.WriteString(drawJSONWhitespace(t, fmt.Sprintf("%s colon %d right whitespace", label, i)))
		b.WriteString(drawJSONValue(t, depth-1, fmt.Sprintf("%s value %d", label, i)))
	}
	b.WriteString(drawJSONWhitespace(t, label+" close whitespace"))
	b.WriteByte('}')

	return b.String()
}

func drawJSONStringLiteral(t *rapid.T, label string) string {
	length := rapid.IntRange(0, 32).Draw(t, label+" len")

	var b strings.Builder
	b.WriteByte('"')
	for i := range length {
		b.WriteString(drawJSONStringSegment(t, fmt.Sprintf("%s segment %d", label, i)))
	}
	b.WriteByte('"')

	return b.String()
}

func drawJSONStringSegment(t *rapid.T, label string) string {
	switch rapid.IntRange(0, 7).Draw(t, label+" kind") {
	case 0:
		return `\"`
	case 1:
		return `\\`
	case 2:
		return rapid.SampledFrom([]string{`\/`, `\b`, `\f`, `\n`, `\r`, `\t`}).
			Draw(t, label+" escaped control")
	case 3:
		return fmt.Sprintf(`\u%04x`, rapid.Uint16().Draw(t, label+" unicode escape"))
	case 4:
		hi := rapid.Uint16Range(0xd800, 0xdbff).Draw(t, label+" high surrogate")
		lo := rapid.Uint16Range(0xdc00, 0xdfff).Draw(t, label+" low surrogate")

		return fmt.Sprintf(`\u%04x\u%04x`, hi, lo)
	default:
		return string(drawJSONSafeStringByte(t, label+" raw byte"))
	}
}

func drawJSONSafeStringByte(t *rapid.T, label string) byte {
	for {
		b := rapid.ByteRange(0x20, 0x7e).Draw(t, label)
		if b != '"' && b != '\\' {
			return b
		}
	}
}

func drawJSONNumber(t *rapid.T, label string) string {
	var b strings.Builder
	if rapid.Bool().Draw(t, label+" negative") {
		b.WriteByte('-')
	}

	if rapid.Bool().Draw(t, label+" zero integer") {
		b.WriteByte('0')
	} else {
		b.WriteByte(rapid.ByteRange('1', '9').Draw(t, label+" first digit"))
		b.WriteString(drawJSONDigits(t, 0, 40, label+" integer tail"))
	}

	if rapid.Bool().Draw(t, label+" has fraction") {
		b.WriteByte('.')
		b.WriteString(drawJSONDigits(t, 1, 40, label+" fraction"))
	}

	if rapid.Bool().Draw(t, label+" has exponent") {
		b.WriteByte(rapid.SampledFrom([]byte{'e', 'E'}).Draw(t, label+" exponent marker"))
		if rapid.Bool().Draw(t, label+" exponent sign") {
			b.WriteByte(rapid.SampledFrom([]byte{'+', '-'}).Draw(t, label+" exponent sign byte"))
		}
		if rapid.IntRange(0, 4).Draw(t, label+" exponent chaos") == 0 {
			b.WriteString(rapid.SampledFrom([]string{"309", "400", "9999", "000001"}).
				Draw(t, label+" large exponent"))
		} else {
			b.WriteString(drawJSONDigits(t, 1, 8, label+" exponent"))
		}
	}

	return b.String()
}

func drawJSONDigits(t *rapid.T, minLen int, maxLen int, label string) string {
	length := rapid.IntRange(minLen, maxLen).Draw(t, label+" len")

	var b strings.Builder
	for i := range length {
		b.WriteByte(rapid.ByteRange('0', '9').Draw(t, fmt.Sprintf("%s digit %d", label, i)))
	}

	return b.String()
}

func drawJSONWhitespace(t *rapid.T, label string) string {
	length := rapid.IntRange(0, 8).Draw(t, label+" len")

	var b strings.Builder
	for i := range length {
		b.WriteByte(rapid.SampledFrom([]byte{' ', '\n', '\r', '\t'}).
			Draw(t, fmt.Sprintf("%s byte %d", label, i)))
	}

	return b.String()
}

func drawJSONJunk(t *rapid.T, label string) string {
	length := rapid.IntRange(1, 8).Draw(t, label+" len")

	var b strings.Builder
	for i := range length {
		b.WriteByte(rapid.Byte().Draw(t, fmt.Sprintf("%s byte %d", label, i)))
	}

	return b.String()
}
