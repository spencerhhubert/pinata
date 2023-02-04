package main

import (
    "fmt"
    "time"
    "strconv"
)

type Message struct {
    purpose    string
    timestamp  time.Time
    data       map[string]interface{}
    timer      time.Time
}

type Bus struct {
    purpose    string 
    name       string
    queue      []Message
}

//make a time.Time type n milliseconds in the future
func timer(ms int) (time.Time) {
    return time.Now().Add(time.Millisecond * time.Duration(ms))
}

func main() {
    msg := Message{
        purpose: "servo",
        timestamp: time.Now(), 
        data: map[string]interface{}{ "angle": 127, "speed": 3 },
        timer: timer(0),
    }
    content := msg.data["angle"].(int)
    fmt.Printf(strconv.Itoa(content))
}
