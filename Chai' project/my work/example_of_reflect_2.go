package main

import (
    "fmt"
    "reflect"
)

//类型转换与赋值
type T struct {
    A int `newT:"AA"`
    B string `newT:"BB"`
    // C string `newT:"CC"`
}

type NewT struct {
    AA int
    BB string
}

// T和NewT之间存在一种一一映射关系，A -> AA, B -> BB

func main() {

    // t tv tt是一组
    t := T{
        A: 111,
        B: "hello",
    }
    tt := reflect.TypeOf(t)
    tv := reflect.ValueOf(t)

    // newT newTv 是一组 最开始它们都是空的 ，后续会将他们填充
    newT := new(NewT)
    newTv := reflect.ValueOf(newT)
    fmt.Println("newT : ",newT)
    fmt.Println("newTv : ",newTv)

    
    for i := 0; i < tt.NumField(); i++ {
        fmt.Println("i : ",i,"############")
        field := tt.Field(i)
        fmt.Println("field : ",field)
        newTTag := field.Tag.Get("newT")//计算出对应的在newT中的tag名字
        fmt.Println("newTTag : ",newTTag)
        tValue := tv.Field(i)
        fmt.Println("tValue : ",tValue)
        newTv.Elem().FieldByName(newTTag).Set(tValue)//填充的步骤
    }

    fmt.Println("newT : ",newT)
    fmt.Println("newT : ",*newT)
}
