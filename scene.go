package main

import (
    "github.com/veandco/go-sdl2/sdl"
    "fmt"
    "github.com/veandco/go-sdl2/img"
    "github.com/veandco/go-sdl2/ttf"
)
type scene struct {
    time int
    bg *sdl.Texture
    bird *bird
    pipes *pipes
    state State
}

func newScene(r *sdl.Renderer) (*scene, error) {
    bg, err := img.LoadTexture(r, "res/images/background.png")
    if err != nil {
        return nil, fmt.Errorf("Could not load background image: %v", err)
    }
    b, err := newBird(r)
    if err != nil {
        return nil, fmt.Errorf("Could not fetch new bird: %v", err)
    }
    p, err := newPipes(r)
    if err != nil {
        return nil, fmt.Errorf("Could not fetch new pipe: %v", err)
    }
    state := State{
        name:"idle",
    }
    return &scene{bg: bg, bird: b, pipes: p, state:state}, nil
}
func (s *scene) update() {
    s.bird.update()
    s.pipes.update()
    s.pipes.touch(s.bird)
}
func (s *scene) restart() {
    s.bird.restart()
    s.pipes.restart()
    s.state.name = "idle"
}
func addStartScreen(r *sdl.Renderer) error {
    r.Clear()
    ss, err := img.LoadTexture(r, "res/images/startScreen.png")
    if err != nil {
        return fmt.Errorf("Could not load start screen image: %v", err)
    }

    if err := r.Copy(ss, nil, nil); err != nil {
        return fmt.Errorf("Could not copy start screen: %v", err)
    }

    whiteRect := &sdl.Rect{X: 200, Y: 50, W: 400, H: 200}
	r.SetDrawColor(255, 255, 255, 255)
	r.FillRect(whiteRect)
    f, err := ttf.OpenFont("res/fonts/test.ttf", 20)
    if err != nil {
        return fmt.Errorf("Could not load font: %v", err)
    }
    defer f.Close()

    c := sdl.Color{ R: 0, G: 0,  B: 0, A: 255 }
    s, err :=f.RenderUTF8Blended("Press any key to play",c)

    if err != nil {
        return fmt.Errorf("Could not render text: %v", err)
    }
    defer s.Free()

    t, err := r.CreateTextureFromSurface(s)
    if err != nil {
        return fmt.Errorf("could not render title: %v", err)
    }
    defer t.Destroy()

    rect := &sdl.Rect{X: 200, Y: 50, W: 400, H: 200}
    err = r.Copy(t, nil, rect)
    if err !=nil {
        fmt.Errorf("Could not copy texture: %v", err)
    }
    r.Present()
    return nil

}

func (s *scene) paint(r *sdl.Renderer) error {
    r.Clear()

    if err := r.Copy(s.bg, nil, nil); err != nil {
        return fmt.Errorf("Could not copy background: %v", err)
    }
    if err := s.bird.paint(r); err != nil {
        return err
    }
    if err := s.pipes.paint(r); err != nil {
        return err
    }
    if err := s.bird.score.drawScore(r); err != nil {
        return err
    }

    r.Present()
    return nil
}

func (s *scene) destroy() {
    s.bg.Destroy()
    s.bird.destroy()
    s.pipes.destroy()
}
