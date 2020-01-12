package main

import (
    "sync"
    "fmt"
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/img"
)
type pipes struct {
    mu sync.RWMutex
    texture *sdl.Texture
    speed int32
    pipes []*pipe
}


func newPipes(r *sdl.Renderer) (*pipe, error) {
    texture, err := img.LoadTexture(r, "res/images/pipe.png")
    if err != nil {
        return nil, fmt.Errorf("Could not load pipe image")
    }
    return &pipes {
        texture: texture,
        speed: 2
    }, nil
}
type pipe struct {
    mu sync.RWMutex
    texture *sdl.Texture
    x int32
    h int32
    w int32
    inverted bool
}


func newPipe(r *sdl.Renderer) (*pipe, error) {
    texture, err := img.LoadTexture(r, "res/images/pipe.png")
    if err != nil {
        return nil, fmt.Errorf("Could not load pipe image")
    }
    return &pipe {
        texture: texture,
        x: 400,
        h: 100,
        w: 50,
        speed: 1,
        inverted: false,
    }, nil
}


func (p *pipe) update() {
    p.mu.Lock() 
    defer p.mu.Unlock()
    p.x -= p.speed
}

func (p *pipe) restart() {
    p.mu.Lock() 
    defer p.mu.Unlock()
    p.x = 400
}

func (p *pipe) paint(r *sdl.Renderer) error {
    p.mu.RLock()
    defer p.mu.RUnlock()
    rect := &sdl.Rect{X: p.x, Y: 600-p.h , W: p.w, H: p.h}
    flip := sdl.FLIP_NONE
    if p.inverted {
        rect.Y = 0
        flip = sdl.FLIP_VERTICAL
    }
    if err := r.CopyEx(p.texture, nil, rect, 0, nil, flip ); err != nil {
        return fmt.Errorf("Could not copy pipe: %v", err)
    }
    return nil
}

func (p *pipe) destroy() {
    p.mu.Lock() 
    defer p.mu.Unlock()
    p.texture.Destroy()
}



