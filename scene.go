package main

import (
    "context"
    "time"
    "github.com/veandco/go-sdl2/sdl"
    "fmt"
    "github.com/veandco/go-sdl2/img"
)
type scene struct {
    time int
    bg *sdl.Texture
    bird *bird
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

    return &scene{bg: bg, bird: b}, nil
}

func (s *scene) run(ctx context.Context, r *sdl.Renderer) <-chan error {
    errc := make(chan error)
    go func() {
        defer close(errc)
        for range time.Tick(time.Millisecond *10) {
            select {
            case <-ctx.Done():
                return
            default:
                if err := s.paint(r); err != nil {
                    errc <- err
                }
            }
        }
    }()
    return errc
}
func (s *scene) paint(r *sdl.Renderer) error {
    r.Clear()

    if err := r.Copy(s.bg, nil, nil); err != nil {
        return fmt.Errorf("Could not copy background: %v", err)
    }
    if err := s.bird.paint(r); err != nil {
        return err
    }
    r.Present()
    return nil
}

func (s *scene) destroy() {
    s.bg.Destroy()
    s.bird.destroy()
}
