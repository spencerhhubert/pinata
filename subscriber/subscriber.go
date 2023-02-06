package subscriber

import (
    "github.com/spencerhhubert/pinata/message"
)

type Subscriber struct {
    Callback    func(message.Message)
    Queue       chan message.Message
    Id          int
}

func New(callback func(message.Message)) Subscriber {
    return Subscriber{ Callback: callback, Queue: make(chan message.Message, 1), Id: 0 }
}

