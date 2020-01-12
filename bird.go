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
    x int32 
    y float64
    w, h int32
    speed float64
    dead bool
    score score
    receivingScore bool
    
}
const (
    gravity = 0.10
    jumpSpeed = 3 
)

func (b *bird) idle() {
}
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
    score := score{
        pipes:0,
        seconds:0,
    }
    return &bird{textures: textures, y:300, speed:0, x:10, w: 50, h: 43, receivingScore:false, score: score }, nil
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
    rect := &sdl.Rect{X: b.x, Y: 600-int32(b.y), W:b.w, H:b.w}
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
    b.score.pipes = 0
    b.receivingScore = false
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
func (b *bird) touch(p *pipe) {
    b.mu.Lock()
    defer b.mu.Unlock()

    if p.x > b.x +b.w { //pipe too far right
        return
    }
    if p.x +p.w < b.x { //pipe too far left
        b.receivingScore = false
        return
    }
    if !p.inverted && p.h < int32(b.y) + (b.h)/2 { //Pipe is too low
        if !b.receivingScore {
            b.score.pipes += 1
            b.receivingScore = true
        }
        return
    }
    if p.inverted && 600 - int32(b.y)+b.h/2>p.h{ //Pipe is too high
        if !b.receivingScore {
            b.score.pipes += 1
            b.receivingScore = true
        }
        return
    }
    fmt.Println("dead", 600-int32(b.y)+b.h/2, b.h, p.h)
    fmt.Println("Score", b.score.pipes)
    b.dead = true
}
