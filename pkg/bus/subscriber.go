package bus

type Subscriber struct {
    Callback    func(Message)
    Queue       chan Message
    Id          int
}

func NewSubscriber(callback func(Message)) Subscriber {
    return Subscriber{ Callback: callback, Queue: make(chan Message), Id: 0 }
}

