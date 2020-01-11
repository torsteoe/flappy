package main

import ( 
    "fmt"
    "os"
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/ttf"
    "github.com/veandco/go-sdl2/img"
    "time"

)

func main() {
    if err:=run(); err != nil {
        fmt.Fprintf(os.Stderr, "%v", err)
        os.Exit(2)
    }
}

func run() error {
    err:= sdl.Init(sdl.INIT_EVERYTHING)
    if err != nil {
        return fmt.Errorf("could not initialize SDL: $v", err)
    }
    defer sdl.Quit()
    err = ttf.Init()
    if err = ttf.Init(); err!=nil {
        return fmt.Errorf("could not initialize ttf: %v", err)
    }

    ttf.Quit()

    w, r, err := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
    if err != nil {
        return fmt.Errorf("Could not create window. %v", err)
    }
    defer w.Destroy()
    _ = r
    if err := drawTitle(r); err != nil {
        return  fmt.Errorf("Could not draw title: %v", err)
    }
    start := time.Now()
    running := true
    for running {
        for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
            switch event.(type) {
            case *sdl.QuitEvent:
                println("Quit")
                running = false
                break
            }
            if time.Since(start).Seconds() > 5 {
                if err := drawBackground(r); err != nil {
                    return fmt.Errorf("Could not draw background: %v", err)
                }
            }
        }
    }



    return nil
}
func drawBackground(r * sdl.Renderer) error {
    r.Clear()

    t, err := img.LoadTexture(r, "res/images/background.png")
    if err != nil {
        return fmt.Errorf("Could not fetch background image: %v", err)
    }
    if err := r.Copy(t, nil, nil); err != nil {
        return fmt.Errorf("Could not copy background: %v", err)
    }

    r.Present()
    return nil
}
func drawTitle(r * sdl.Renderer) error {

    r.Clear()
    f, err := ttf.OpenFont("res/fonts/test.ttf", 20)
    if err != nil {
        return fmt.Errorf("Could not load font: %v", err)
    }
    defer f.Close()

    c := sdl.Color{ R: 255, G: 100,  B: 0, A: 255 }
    s, err :=f.RenderUTF8Blended("Flappy Gopher",c)

    if err != nil {
        return fmt.Errorf("Could not render text: %v", err)
    }
    defer s.Free()

    t, err := r.CreateTextureFromSurface(s)
    if err != nil {
        return fmt.Errorf("could not render title: %v", err)
    }
    defer t.Destroy()

    err = r.Copy(t, nil, nil)
    if err !=nil {
        fmt.Errorf("Could not copy texture: %v", err)
    }
    r.Present()

    return nil
}
