package appm

import (
    "strings"
    "strconv"
    "log"
    "fmt"
)

func HMSToSeconds (s string) (int, error) {
    // HMSToSeconds parses a duration string of the form H..HH:MM:SS.
    // and returns its value in seconds.
    if strings.Count(s, ":") != 2 {
        err := fmt.Errorf("Expected duration string of the form H..HH:MM:SS; received: %q", s)
        return 0, err
    }

    hms := strings.Split(s, ":")
    if len(hms) != 3 {
        err := fmt.Errorf("Expected duration string of the form H..HH:MM:SS; received: %q", s)
        return 0, err
    }

    hours, err := strconv.Atoi(hms[0])
    if err != nil {
        log.Fatal(err)
    }
    minutes, err := strconv.Atoi(hms[1])
    if err != nil {
        log.Fatal(err)
    }
    seconds, err := strconv.Atoi(hms[2])
    if err != nil {
        log.Fatal(err)
    }

    return hours*3600 + minutes*60 + seconds, nil
}

func SecondsToHMS(sec int) string {
    // 
    hours := sec/3600
    minutes := (sec - 3600*hours)/60
    seconds := sec - 3600*hours - 60*minutes
    return fmt.Sprintf("%02v:%02v:%02v", hours, minutes, seconds)
}
