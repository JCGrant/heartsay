package renderer

import (
	"fmt"
	"strings"
)

type Frame struct {
	rows []string
}

func NewFrame(frameStr string) *Frame {
	return &Frame{
		strings.Split(frameStr, "\n"),
	}
}

func center(s string, w int) string {
	return fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(s))/2, s))
}

func generateTextRows(text string, textBounds Rect) ([]string, error) {
	dx := textBounds.x2 - textBounds.x1
	dy := textBounds.y2 - textBounds.y1
	words := strings.Split(text, " ")
	textRows := make([]string, dy+1)
	rowIndex := 0
	var word string
	for len(words) > 0 {
		word, words = words[0], words[1:]
		if len(word)+1+len(textRows[rowIndex]) > dx {
			rowIndex++
			if rowIndex > dy {
				return nil, fmt.Errorf("too many words")
			}
		}
		textRows[rowIndex] += word + " "
	}
	for i, _ := range textRows {
		textRows[i] = strings.Trim(textRows[i], " ")
		textRows[i] = center(textRows[i], dx)
	}
	return textRows, nil
}

func (f *Frame) AddText(text string, textBounds Rect) error {
	textRows, err := generateTextRows(text, textBounds)
	if err != nil {
		return fmt.Errorf("generating text failed: %s", err)
	}
	for i := range f.rows {
		if textBounds.y1 <= i && i <= textBounds.y2 {
			f.rows[i] = f.rows[i][:textBounds.x1] + textRows[i-textBounds.y1] + f.rows[i][textBounds.x2:]
		}
	}
	return nil
}

func (f *Frame) String() string {
	return strings.Join([]string(f.rows), "\n")
}
