package bus

import (
    "time"
)

type Message struct {
    Data       map[string]interface{}
    Purpose    string
    Timestamp  time.Time
    Delay      time.Time
}

func NewMessage(data map[string]interface{}, purpose string, delay int) Message {
    return Message{
        Data: data,
        Purpose: purpose,
        Timestamp: time.Now(),
        Delay: timer(delay),
    }
}

//make a time.Time type n milliseconds in the future
func timer(ms int) (time.Time) {
    return time.Now().Add(time.Millisecond * time.Duration(ms))
}
