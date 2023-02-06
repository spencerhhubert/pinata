package main

import (
    "fmt"
    "time"
    "github.com/spencerhhubert/pinata/bus"
    "github.com/spencerhhubert/pinata/message"
    "github.com/spencerhhubert/pinata/subscriber"
)

func simpleCallback(mes message.Message) {
    fmt.Println(mes.Data["data"])
}

func main() {
    data1 := map[string]interface{}{ "data": "one" }
    data2 := map[string]interface{}{ "data": "two" }
    data3 := map[string]interface{}{ "data": "three" }

    mes1 := message.New(data1, "whocares", 100)
    mes2 := message.New(data2, "whocares", 100)
    mes3 := message.New(data3, "whocares", 100)

    sub1 := subscriber.New(simpleCallback)
    bus := bus.New("whocares")
    time.Sleep(time.Second*1)

    go bus.Sub(sub1)
    go bus.Publish(mes1)
    go bus.Run()
    go bus.Publish(mes2)
    go bus.Publish(mes3)
    time.Sleep(time.Second*1)

    bus.Unsub(sub1)
    bus.Kill()
}
