package main

import (
    "os"
    "fmt"
    "strconv"
    "flag"
    "bufio"
    "strings"
)


func addScore(s score) error {
    /*
    if, err := os.Create("highscores.txt")

    if err != nil {
        return err
    }
    */
    f, err := os.OpenFile("highscores.txt", os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    _, err = fmt.Fprintln(f, s.name + " " + strconv.Itoa(s.pipes))
    if err != nil {
        return err
    }

    err = f.Close()
    if err != nil {
        return err
    }
    return nil
}

func readHighscores() (highscores, error) {
    var hs highscores
    var highscores []int
    var highscoresNames []string 
    fptr := flag.String("fpath", "highscores.txt", "file path to read from")
    flag.Parse()

    f, err := os.Open(*fptr)
    if err != nil {
        return hs, err
    }


    s := bufio.NewScanner(f)
    for s.Scan() {

        scorer := strings.SplitN(strings.TrimSpace(s.Text()), ",", 2)
        name := scorer[0]
        val, err := strconv.Atoi(scorer[1])
        if err != nil {
            return hs, err
        }
        highscores = append(highscores, val)
        highscoresNames = append(highscoresNames, name)

    }

    err = s.Err()
    if err != nil {
        return hs, err
    }
    if err = f.Close(); err != nil {
        return hs, err
    }
    hs.scores = highscores
    hs.names = highscoresNames
    return hs, nil
}
