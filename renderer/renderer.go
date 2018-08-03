package renderer

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Renderer struct {
	frames    []*Frame
	frameRate int
}

func New(frames []*Frame, frameRate int) *Renderer {
	return &Renderer{
		frames:    frames,
		frameRate: frameRate,
	}
}

func (r *Renderer) clear() {
	out, _ := exec.Command("clear").Output()
	os.Stdout.Write(out)
}

func (r *Renderer) Render(currentFrame int) {
	fmt.Print(r.frames[currentFrame])
}

func (r *Renderer) Loop() {
	currentFrame := 0
	for _ = range time.Tick(time.Second / time.Duration(r.frameRate)) {
		r.clear()
		r.Render(currentFrame)
		currentFrame = (currentFrame + 1) % len(r.frames)
	}
}

func Loop(frames []*Frame, frameRate int) {
	r := New(frames, frameRate)
	r.Loop()
}
