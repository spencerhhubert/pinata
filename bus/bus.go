package bus

import (
    "time"
    "sync"
    "fmt"
    "github.com/spencerhhubert/pinata/message"
    "github.com/spencerhhubert/pinata/subscriber"
)

type Bus struct {
    Purpose        string 
    Queue          chan message.Message
    subscribers    map[int]subscriber.Subscriber
    Wg             sync.WaitGroup
    rate           int
}

func New(purpose string) Bus {
    return Bus{
        Purpose:         purpose,
        Queue:           make(chan message.Message, 1),
        subscribers:     make(map[int]subscriber.Subscriber),
        rate:            200,
    }
}

func (b *Bus) Publish(m message.Message) {
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

func (b *Bus) Sub(s subscriber.Subscriber) {
    b.Wg.Add(1)
    s.Id = len(b.subscribers)
    b.subscribers[s.Id] = s
    for m := range s.Queue {
        fmt.Println("Tried to callback")
        time.Sleep(time.Until(m.Delay))
        s.Callback(m)
    }
}

func (b *Bus) Unsub(s subscriber.Subscriber) {
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
