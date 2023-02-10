package bus

import (
    "time"
    "sync"
    "golang.org/x/exp/constraints"
)

type validData interface {
   constraints.Ordered | []byte 
}


type Bus[X comparable, Y validData] struct {
    Queue chan Message[X,Y]
    subscribers    map[int]Subscriber
    Wg             sync.WaitGroup
    rate           int
}





type Bus[D Data] struct {
    Queue          chan Message[D]
    subscribers    map[int]Subscriber
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
