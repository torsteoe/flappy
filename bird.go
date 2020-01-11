package main

import (
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/img"
    "fmt"
    "strconv"
)
type bird struct {
    time int
    textures []*sdl.Texture
    y, speed float64
}
const gravity = 0.0098
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

func (b *bird) paint(r *sdl.Renderer) error {
    b.time++
    b.y -= b.speed
    if b.y < 0  {
        b.speed = -b.speed
        b.y = 0
    }
    b.speed += gravity
    i := b.time/10 % 4
    rect := &sdl.Rect{X: 10, Y: 600-int32(b.y), W:50, H:43}
    if err := r.Copy(b.textures[i], nil, rect); err != nil {
        return fmt.Errorf("Could not copy bird: %v", err)
    }
    return nil
}

func (b *bird) destroy() {
    for _, t := range b.textures { 
        t.Destroy()
    }
}
