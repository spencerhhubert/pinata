package main 

import (
    "time"
    "fmt"
    "reflect"
    "golang.org/x/exp/constraints"
)

type valid interface {
   constraints.Ordered | []byte 
}

type Message[X comparable, Y valid] struct {
    Data map[X]Y 
    Timestamp time.Time
    Delay time.Time
}

func NewMessage[X comparable, Y valid](data map[X]Y, delay int) *Message[X,Y] {
    msg := new(Message[X,Y])
    msg.Data = data
    return msg
}

