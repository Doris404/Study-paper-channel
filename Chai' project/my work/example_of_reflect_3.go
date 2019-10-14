package main

import (
    "fmt"
    "reflect"
)
//定义一个Enum类型
type Enum int
const Zero Enum = 0
type Student struct {
    Name string
    Age  int
}

func main() {

    //定义一个Student类型的变量
    var stu Student

    //获取结构体实例的反射类型对象
    typeOfStu := reflect.TypeOf(stu)

    //显示反射类型对象的名称和种类
    fmt.Println("typeOfStu's Name : ",typeOfStu.Name()," typeOfStu's Kind : ",typeOfStu.Kind())

    //用一个指针试一下
    var stuPointer = &Student{Name : "Kitty", Age : 12}

    //获取一个stuPointer的反射
    typeOfstuPointer := reflect.TypeOf(stuPointer)

    //显示typeOfPointer的Name和Kind
    fmt.Println("typeOfstuPointer's Name : ",typeOfstuPointer.Name()," typeOfstuPointer's Kind : ",typeOfstuPointer.Kind())

    //指针的新花样
    typeOfStuNew := typeOfstuPointer.Elem()

    //打印出来
    fmt.Println("typeOfStuNew's Name : ",typeOfStuNew.Name()," typeOfStuNew's Kind : ",typeOfStuNew.Kind())

    //获取Zero常量的反射类型对象
    typeOfZero := reflect.TypeOf(Zero)

    //显示反射类型对象的名称和种类
    fmt.Println("typeOfZero's Name : ",typeOfZero.Name()," typeOfZero's kind : " ,typeOfZero.Kind())

}