package heartsay

import (
	"fmt"

	"github.com/JCGrant/heartsay/renderer"
)

func Say(text string) error {
	var frames []*renderer.Frame
	for _, heart := range hearts {
		frame := renderer.NewFrame(heart)
		err := frame.AddText(text, textBounds)
		if err != nil {
			return fmt.Errorf("adding text failed: %s", err)
		}
		frames = append(frames, frame)
	}
	frameRate := 1
	renderer.Loop(frames, frameRate)
	return nil
}
