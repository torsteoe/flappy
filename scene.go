package main

import (
    "time"
    "github.com/veandco/go-sdl2/sdl"
    "fmt"
    "github.com/veandco/go-sdl2/img"
    "log"
    "strconv"
)
type scene struct {
    time int
    bg *sdl.Texture
    bird *bird
    pipes *pipes
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

    return &scene{bg: bg, bird: b, pipes: p}, nil
}
func (s *scene) update() {
    s.bird.update()
    s.pipes.update()
    s.pipes.touch(s.bird)
}
func (s *scene) restart() {
    s.bird.restart()
    s.pipes.restart()
}
func (s *scene) run(events <-chan sdl.Event, r *sdl.Renderer) <-chan error {
    errc := make(chan error)
    go func() {
        defer close(errc)
        tick := time.Tick(time.Millisecond * 10)
        running := true
        for running{
            select {
            case e:= <-events:
                running = s.handleEvent(e )
            case <-tick:
                s.update()
                if s.bird.isDead() {
                    if err:=drawTitle(r, "Score: "+strconv.Itoa(s.bird.score.pipes)); err != nil {
                        fmt.Printf("Could not draw title: %v", err)
                    }
                    time.Sleep(time.Second*5)
                    s.restart()
                }
                if err := s.paint(r); err != nil {
                    errc <- err
                }
            }
        }
    }()
    return errc
}

func (s *scene) handleEvent(event sdl.Event) bool {
    switch e :=event.(type) {
    case *sdl.QuitEvent:
        return false
    case *sdl.MouseButtonEvent:
        s.bird.jump()
    case *sdl.WindowEvent, *sdl.MouseMotionEvent:
    default:
        log.Printf("Unknown event %T", e)
    }
    return true
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
    r.Present()
    return nil
}

func (s *scene) destroy() {
    s.bg.Destroy()
    s.bird.destroy()
    s.pipes.destroy()
}
