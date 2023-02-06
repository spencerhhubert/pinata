package main

import (
    "fmt"
    "time"
    "strconv"
    "github.com/spencerhhubert/pinata/pkg/bus"
)

func simpleCallback(mes bus.Message) {
    fmt.Println(mes.Data)
}

func main() {
    sub1 := bus.NewSubscriber(simpleCallback)
    bus1 := bus.NewBus("whocares")
    bus1.Wg.Add(1)
    go bus1.Sub(sub1)
    go bus1.Run()
    for i := 0; i <= 10; i++ {
        data := strconv.Itoa(i)
        msg := bus.NewMessage(data, "servo", 50)
        go bus1.Publish(msg)
    }
    go func() {
        time.Sleep(time.Second*2)
        msg := bus.NewMessage("hello", "servo", 50)
        bus1.Publish(msg)
    }()
    bus1.Wg.Wait()
    bus1.Unsub(sub1)
    bus1.Kill()
}
