package menus

import (
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/ge"
	"github.com/quasilyte/roboden-game/session"
)

func addDemoBackground(state *session.State, scene *ge.Scene) {
	if state.DemoFrame == nil {
		return
	}
	s := ge.NewSprite(scene.Context())
	s.Centered = false
	s.SetColorScale(ge.ColorScale{R: 0.30, G: 0.30, B: 0.30, A: 1})
	s.SetImage(resource.Image{Data: state.DemoFrame})
	scene.AddGraphics(s)
}