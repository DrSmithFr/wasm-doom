package render

import (
    "go-webgl/game"
    "image/color"
)

func RenderMinimap(r Renderer, p game.Player, m game.Map) {
    r.Clear()
    r.SetColor(color.RGBA{255, 255, 255, 255})

    screenW, screenH := r.Size()

    // render player position
    r.DrawCircle(
        float64(screenW/2),
        float64(screenH/2),
        5,
    )

    // draw a line in the direction the player is facing
    dx, dy := p.Position.GetDelta()
    r.DrawLine(
        float64(screenW/2),
        float64(screenH/2),
        dx*game.MovementMomentum+float64(screenW/2),
        dy*game.MovementMomentum+float64(screenH/2),
        1,
    )

    // render sectors
    for _, s := range m.Sectors {
        for _, wall := range s.Walls {
            x1, x2, y1, y2, _, _ := wall.RelativePosition(p)

            r.DrawLine(
                x1+float64(screenW/2),
                y1+float64(screenH/2),
                x2+float64(screenW/2),
                y2+float64(screenH/2),
                1,
            )
        }
    }

    r.Flush()
}