package bus

import (
    "time"
    "reflect"
    "golang.org/x/exp/constraints"
)

type validData interface {
    constraints.Ordered | []byte
}

type Message[X comparable, Y validData] struct {
    Data map[X]Y 
    Timestamp time.Time
    Delay time.Time
}

func NewMessage(data map[X]Y, delay int) Message {
    return Message[X,Y]{
        Data: data,
        Timestamp: time.Now(),
        Delay: timer(delay),
    }
}

//make a time.Time type n milliseconds in the future
func timer(ms int) (time.Time) {
    return time.Now().Add(time.Millisecond * time.Duration(ms))
}
