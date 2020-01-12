package main

import (
    "math/rand"
    "sync"
    "fmt"
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/img"
    "time"
)
type pipes struct {
    mu sync.RWMutex
    texture *sdl.Texture
    speed float64
    pipes []*pipe
    interval int64
}


func newPipes(r *sdl.Renderer) (*pipes, error) {
    texture, err := img.LoadTexture(r, "res/images/pipe.png")
    if err != nil {
        return nil, fmt.Errorf("Could not load pipe image")
    }
    ps := &pipes {
        texture: texture,
        speed: 3,
        interval: 100000,
    }
    go func() {
        for {
            ps.mu.Lock()
            ps.pipes = append(ps.pipes, newPipe())
            ps.mu.Unlock()
            for i:= 0; i<int(ps.interval); i++ {
                time.Sleep(time.Microsecond)
            }
            if ps.interval > 45000 {
                ps.interval -= 10000
                ps.speed += 0.01
            } else {
                ps.interval -= 0
                ps.speed += 0
            }
        }
    }()
    return ps, nil
}
type pipe struct {
    mu sync.RWMutex
    x int32
    h int32
    w int32
    inverted bool
}

func (ps *pipes) touch(b *bird) {
    ps.mu.RLock()
    defer ps.mu.RUnlock()
    for _, p := range ps.pipes {
        p.touch(b)
    }
}
func (p *pipe) touch(b *bird) {
    p.mu.RLock()
    defer p.mu.RUnlock()
    b.touch(p)
}
func (ps *pipes) update() {
    ps.mu.Lock()
    defer ps.mu.Unlock()
    var rem []*pipe
    for _, p := range ps.pipes {
        p.mu.Lock()
        p.x -= int32(ps.speed)
        p.mu.Unlock()
        if p.x+p.w >0 {
            rem = append(rem, p)
        } 
    }
    ps.pipes = rem
}

func (ps *pipes) restart() {
    ps.mu.Lock()
    defer ps.mu.Unlock()
    ps.pipes = nil
    ps.interval = 100000
}

func (ps *pipes) paint(r *sdl.Renderer) error {
    ps.mu.RLock()
    defer ps.mu.RUnlock()
    for _, p := range ps.pipes {
        if err := p.paint(r, ps.texture); err != nil {
            return err
        }
    }
    return nil
}

func (ps *pipes) destroy() {
    ps.mu.Lock()
    defer ps.mu.Unlock()
    ps.texture.Destroy()
}

func newPipe() (*pipe) {
    return &pipe {
        x: 700,
        h: 100 + int32(rand.Intn(300)),
        w: 50,
        inverted: rand.Float32()>0.5,
    }
}




func (p *pipe) paint(r *sdl.Renderer, texture *sdl.Texture) error {
    p.mu.RLock()
    defer p.mu.RUnlock()
    rect := &sdl.Rect{X: p.x, Y: 600-p.h , W: p.w, H: p.h}
    flip := sdl.FLIP_NONE
    if p.inverted {
        rect.Y = 0
        flip = sdl.FLIP_VERTICAL
    }
    if err := r.CopyEx(texture, nil, rect, 0, nil, flip ); err != nil {
        return fmt.Errorf("Could not copy pipe: %v", err)
    }
    return nil
}




