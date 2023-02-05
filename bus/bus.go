package bus

import (
    "fmt"
    "time"
    "strconv"
)

type Bus struct {
    Purpose    string 
    Queue      chan Message
}

func New(purpose string) Bus {
    return Bus{
        Purpose: purpose,
        Queue: make(chan Message),
    }
}

func (b *Bus) Publish(m Message) {
    b.queue <- m
}

func (b *Bus) Subscribe(callback func(Message)) {
    go func() {
        for m := range b.queue {
            time.Sleep(time.Until(m.delay))
            callback(m)
        }
    }()
}

