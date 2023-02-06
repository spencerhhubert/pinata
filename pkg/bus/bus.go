package bus

import (
    "time"
    "sync"
)

type Bus struct {
    Purpose        string 
    Queue          chan Message
    subscribers     map[int]Subscriber
    Wg             sync.WaitGroup
    rate           int
}

func NewBus(purpose string) Bus {
    return Bus{
        Purpose:         purpose,
        Queue:           make(chan Message),
        subscribers:     make(map[int]Subscriber),
        rate:            10,
    }
}

func (b *Bus) Publish(m Message) {
    b.Queue <- m
}

func (b *Bus) Run() {
    for m := range b.Queue {
        for _,s := range b.subscribers {
            s.Queue <- m
        }
        time.Sleep(time.Millisecond * time.Duration(b.rate)) //refresh rate
    }
}

func (b *Bus) Sub(s Subscriber) {
    b.Wg.Add(1)
    s.Id = len(b.subscribers)
    b.subscribers[s.Id] = s
    for m := range s.Queue {
        time.Sleep(time.Until(m.Delay))
        s.Callback(m)
    }
}

func (b *Bus) Unsub(s Subscriber) {
    delete(b.subscribers, s.Id)
    b.Wg.Done()
    close(s.Queue)
}

func (b *Bus) Kill() {
    for _,s := range(b.subscribers) {
        b.Unsub(s)
    }
    b.Wg.Wait()
    close(b.Queue)
}
