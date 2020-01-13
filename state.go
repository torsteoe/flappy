package main

import (
    "log"
    "github.com/veandco/go-sdl2/sdl"
    "fmt"
    "time"
    "strconv"
)

type State struct{
    name string
}

func (s *scene) run(events <-chan sdl.Event, r *sdl.Renderer) <-chan error {
    errc := make(chan error)

    go func() {
        defer close(errc)
        tick := time.Tick(time.Millisecond * 10)
        s.bird.score.name = "Gopher" 
        s.state.name = "Start Screen"
        addStartScreen(r)
        for s.state.name != "quit"{
                select {
                case e:= <-events:
                    s.handleEvent(e)
                case <-tick:
                    switch (s.state.name) {
                    case "idle":
                        s.pipes.idle()
                        s.bird.idle()
                        if  err := s.paint(r); err != nil {
                            errc <- err
                        }

                   case "running":
                        s.update()
                        if s.bird.isDead() {
                            if err:=drawTitle(r, "Score: "+ strconv.Itoa(s.bird.score.pipes)); err != nil {
                                fmt.Printf("Could not draw title: %v", err)
                            }
                            addScore(s.bird.score)
                            s.state.name = "Game over"
                        } else if  err := s.paint(r); err != nil {
                            errc <- err
                        }
                   case "Game over":
                   case "Start screen":
                       addStartScreen(r)
                   }

            }
        }
    }()
    return errc
}
func (s *scene) handleEvent(event sdl.Event) {
    switch e :=event.(type) {
    case *sdl.QuitEvent:
        highScores, err := readHighscores()
        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Println(highScores.scores)
            fmt.Println(highScores.names)
        }
        s.state.name="quit"
    case *sdl.MouseButtonEvent:
        switch (s.state.name) {
        case "idle":
            s.state.name = "running"
        case "running":
            s.bird.jump()
        case "Game over":
            s.restart()
            s.state.name = "idle"
        case "Start Screen":
            s.restart()
            s.state.name = "idle"
        }
    case *sdl.WindowEvent, *sdl.MouseMotionEvent:
    case *sdl.KeyboardEvent:
        switch (s.state.name) {
        case "idle":
            s.state.name = "running"
        case "running":
            s.bird.jump()
        case "Game over":
            s.restart()
            s.state.name = "idle"
        case "Start Screen":
            s.restart()
            s.state.name = "idle"
        }
    default:
        log.Printf("Unknown event %T", e)
    }
}



