package heart

import (
	"github.com/JCGrant/heart/renderer"
)

func Beat() {
	var frames []*renderer.Frame
	for _, heart := range hearts {
		frame := renderer.NewFrame(heart)
		frames = append(frames, frame)
	}
	frameRate := 1
	renderer.Loop(frames, frameRate)
}
