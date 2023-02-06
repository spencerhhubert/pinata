package main

import (
    "fmt"
    "time"
    "github.com/spencerhhubert/pinata/pkg/bus"
)

func simpleCallback(mes bus.Message) {
    fmt.Println(mes.Data["data"])
}

func main() {
    data1 := map[string]interface{}{ "data": "one" }
    data2 := map[string]interface{}{ "data": "two" }
    data3 := map[string]interface{}{ "data": "three" }

    mes1 := bus.NewMessage(data1, "whocares", 100)
    mes2 := bus.NewMessage(data2, "whocares", 100)
    mes3 := bus.NewMessage(data3, "whocares", 100)

    sub1 := bus.NewSubscriber(simpleCallback)
    bus1 := bus.NewBus("whocares")
    time.Sleep(time.Second*1)

    go bus1.Sub(sub1)
    go bus1.Publish(mes1)
    go bus1.Run()
    go bus1.Publish(mes2)
    go bus1.Publish(mes3)
    time.Sleep(time.Second*1)

    bus1.Unsub(sub1)
    bus1.Kill()
}
