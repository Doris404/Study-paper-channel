package main

import (
    "fmt"
    "reflect"
)

//类型转换与赋值
type T struct {
    A int `newT:"AA"`
    B string `newT:"BB"`
}

type NewT struct {
    AA int
    BB string
}

func main() {
    t := T{
        A: 111,
        B: "hello",
    }
    tt := reflect.TypeOf(t)
    tv := reflect.ValueOf(t)

    newT := new(NewT)
    newTv := reflect.ValueOf(newT)

    for i := 0; i < tt.NumField(); i++ {
        field := tt.Field(i)
        newTTag := field.Tag.Get("newT")

        tValue := tv.Field(i)
        newTv.Elem().FieldByName(newTTag).Set(tValue)
    }

    fmt.Println(*newT)
}
