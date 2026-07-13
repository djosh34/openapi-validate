package generate

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"testing"

	"github.com/arran4/golang-diff/pkg/diff"
	"github.com/stretchr/testify/require"
)

// PrettyDiff prints a colored line diff for a failed test.
func PrettyDiff(t *testing.T, expected, actual string) {
	t.Helper()

	_, err := fmt.Print(coloredStandardDiff(expected, actual))
	require.NoError(t, err)
}

// coloredStandardDiff returns a colored unified diff.
func coloredStandardDiff(expected, actual string) string {
	expectedLines := diffInputLines(expected)
	actualLines := diffInputLines(actual)
	alignedLines := diff.AlignLines(expectedLines, actualLines, diff.NewOptions())

	output := []string{
		coloredDiffLine(ansiRed, "--- expected"),
		coloredDiffLine(ansiGreen, "+++ actual"),
		coloredDiffLine(
			ansiCyan,
			fmt.Sprintf("@@ -1,%d +1,%d @@", len(expectedLines), len(actualLines)),
		),
	}

	for _, alignedLine := range alignedLines {
		if alignedLine.Type == diff.DiffEqual {
			output = append(output, " ", alignedLine.Left, "\n")

			continue
		}

		if alignedLine.Left != "" {
			output = append(output, coloredDiffLine(ansiRed, "-"+alignedLine.Left))
		}

		if alignedLine.Right != "" {
			output = append(output, coloredDiffLine(ansiGreen, "+"+alignedLine.Right))
		}

		if alignedLine.Left == "" && alignedLine.Right == "" {
			output = append(
				output,
				coloredDiffLine(ansiRed, "-"),
				coloredDiffLine(ansiGreen, "+"),
			)
		}
	}

	return strings.Join(output, "")
}

// diffInputLines splits input without a final empty line.
func diffInputLines(input string) []string {
	if input == "" {
		return nil
	}

	lines := strings.Split(input, "\n")
	if strings.HasSuffix(input, "\n") {
		return lines[:len(lines)-1]
	}

	return lines
}

// coloredDiffLine decorates one diff line with an ANSI color.
func coloredDiffLine(color string, line string) string {
	return color + line + ansiReset + "\n"
}

// comparableFiles lists regular fixture files outside exceptions.
func comparableFiles(t *testing.T, root string, exceptions map[string]struct{}) map[string]struct{} {
	t.Helper()

	files := map[string]struct{}{}

	err := filepath.WalkDir(root, func(path string, dirEntry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		if rel == "." {
			return nil
		}

		rel = filepath.ToSlash(rel)
		if exceptedPath(rel, exceptions) {
			if dirEntry.IsDir() {
				return filepath.SkipDir
			}

			return nil
		}

		if dirEntry.IsDir() {
			return nil
		}

		files[rel] = struct{}{}

		return nil
	})
	require.NoError(t, err)

	return files
}

// exceptedPath reports whether rel is excluded from fixture comparison.
func exceptedPath(rel string, exceptions map[string]struct{}) bool {
	if _, ok := exceptions[rel]; ok {
		return true
	}

	for exception := range exceptions {
		if strings.HasPrefix(rel, exception+"/") {
			return true
		}
	}

	return false
}
