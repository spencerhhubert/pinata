package main

import (
    "fmt"
    "time"
    "github.com/spencerhhubert/pinata/message"
)

func main() {
    data := map[string]interface{}{ "angle": 127, "speed": 3 }
    m := message.New(data, "servo", 127)
    fmt.Printf("%T\n", m)
    time.Sleep(time.Second)
    fmt.Println(m.Purpose)
}
