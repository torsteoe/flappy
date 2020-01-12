package main

import (
    "github.com/veandco/go-sdl2/sdl"
    "fmt"
    "github.com/veandco/go-sdl2/ttf"
    "strconv"
)

type score struct {
    pipes int
    seconds int
}

func (score *score) drawScore(r * sdl.Renderer) error {

    f, err := ttf.OpenFont("res/fonts/test.ttf", 20)
    if err != nil {
        return fmt.Errorf("Could not load font: %v", err)
    }
    defer f.Close()

    c := sdl.Color{ R: 255, G: 100,  B: 0, A: 255 }
    s, err :=f.RenderUTF8Blended("Score: " + strconv.Itoa(score.pipes),c)

    if err != nil {
        return fmt.Errorf("Could not render text: %v", err)
    }
    defer s.Free()

    t, err := r.CreateTextureFromSurface(s)
    if err != nil {
        return fmt.Errorf("could not render title: %v", err)
    }
    defer t.Destroy()

    rect := &sdl.Rect{X: 575, Y: 400 , W: 100, H: 100}
    err = r.Copy(t, nil, rect)
    if err !=nil {
        fmt.Errorf("Could not copy texture: %v", err)
    }
    r.Present()
    return nil
}
