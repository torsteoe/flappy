package main

import (
    "fmt"
    "os"
    "github.com/veandco/go-sdl2/sdl"
    "github.com/veandco/go-sdl2/ttf"
    "time"
    "runtime"
)

func main() {
    if err:=run(); err != nil {
        fmt.Fprintf(os.Stderr, "%v", err)
        os.Exit(2)
    }
    time.Sleep(time.Second)
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

    s, err := newScene(r)
    if err !=  nil{
        return fmt.Errorf("Could not draw background: %v", err)
    }
    defer s.destroy()
    //time.AfterFunc(20*time.Second, cancel)

    events := make(chan sdl.Event)
    errc := s.run(events, r)

    runtime.LockOSThread()

    for {
        select {
        case events <- sdl.WaitEvent():
        case err := <-errc:
            return err
        }
    }


    return nil
}

func drawTitle(r * sdl.Renderer, text string) error {

    r.Clear()
    f, err := ttf.OpenFont("res/fonts/test.ttf", 20)
    if err != nil {
        return fmt.Errorf("Could not load font: %v", err)
    }
    defer f.Close()

    c := sdl.Color{ R: 255, G: 100,  B: 0, A: 255 }
    s, err :=f.RenderUTF8Blended(text,c)

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
