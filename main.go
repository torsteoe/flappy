package main

import ( 
    "fmt"
    "os"
    "time"
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/ttf"
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
    time.Sleep(time.Second*5)
    running := true
    for running {
        for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
            switch event.(type) {
            case *sdl.QuitEvent:
                println("Quit")
                running = false
                break
            }
        }
    }

    return nil
}

func drawTitle(r * sdl.Renderer) error {

    r.Clear()
    f, err := ttf.OpenFont("fonts/test.ttf", 20)
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
