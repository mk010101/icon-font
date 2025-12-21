// Run this to normalize Google icons (SVG)
package main

import (
	"bytes"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	FixMaterialSVGs("../icon-font/svg")
}

var (
	viewBoxRe  = regexp.MustCompile(`viewBox="([^"]*?)\s-960\s([^"]*?)"`)
	svgOpenRe  = regexp.MustCompile(`(?s)<svg\b[^>]*>`)
	svgCloseRe = regexp.MustCompile(`</svg>`)
)

func FixMaterialSVGs(dir string) error {
	return filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || filepath.Ext(path) != ".svg" {
			return err
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// Only process Material-style SVGs
		if !bytes.Contains(data, []byte(`-960`)) {
			return nil
		}

		out := fixSVG(data)
		if !bytes.Equal(data, out) {
			return os.WriteFile(path, out, 0644)
		}

		return nil
	})
}

func fixSVG(input []byte) []byte {
	// Fix viewBox
	out := viewBoxRe.ReplaceAll(input, []byte(`viewBox="$1 0 $2"`))

	open := svgOpenRe.FindIndex(out)
	close := svgCloseRe.FindIndex(out)

	if open == nil || close == nil {
		return out
	}

	openEnd := open[1]
	closeStart := close[0]

	// Inject <g> wrapper
	var buf bytes.Buffer
	buf.Write(out[:openEnd])
	buf.WriteString(`<g transform="translate(0 960)">`)
	buf.Write(out[openEnd:closeStart])
	buf.WriteString(`</g>`)
	buf.Write(out[closeStart:])

	return buf.Bytes()
}
