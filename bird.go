package main

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/img"
    "fmt"
    "strconv"
    "sync"
)
type bird struct {
    mu sync.RWMutex
    time int
    textures []*sdl.Texture
    y, speed float64
    dead bool
}
const (
    gravity = 0.05
    jumpSpeed = 1
)
func newBird(r *sdl.Renderer) (*bird, error) {

    var textures []*sdl.Texture
    for i:= 1; i<= 4; i++ {
        iString := strconv.Itoa(i)
        bird, err := img.LoadTexture(r, "res/images/frame-" + iString + ".png")
        if err != nil {
            return nil, fmt.Errorf("Could not load bird image: %v", err)
        }
        textures = append(textures, bird)
    }
    return &bird{textures: textures, y:300, speed:0}, nil
}
func (b *bird) update() {
    b.mu.Lock()
    defer b.mu.Unlock()
    b.time++
    b.y -= b.speed
    if b.y < 0  {
        b.speed = -b.speed
        b.y = 0
        b.dead = true
    }
    b.speed += gravity
}

func (b *bird) paint(r *sdl.Renderer) error {
    b.mu.RLock()
    defer b.mu.RUnlock()
    i := b.time/10 % 4
    rect := &sdl.Rect{X: 10, Y: 600-int32(b.y), W:50, H:43}
    if err := r.Copy(b.textures[i], nil, rect); err != nil {
        return fmt.Errorf("Could not copy bird: %v", err)
    }
    return nil
}

func (b *bird) restart() {
    b.mu.Lock()
    defer b.mu.Unlock()
    b.y = 300
    b.speed = 0 
    b.dead = false
}

func (b *bird) destroy() {
    b.mu.Lock()
    defer b.mu.Unlock()
    for _, t := range b.textures { 
        t.Destroy()
    }
}
func (b *bird) jump() {
    b.mu.Lock()
    defer b.mu.Unlock()
    b.speed = -jumpSpeed
}
func (b *bird) isDead() bool {
    b.mu.RLock()
    b.mu.RUnlock()
    return b.dead
}
