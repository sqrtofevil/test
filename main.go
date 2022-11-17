package main

import (
	"fmt"
)

func main() {
	var answer, result int
  var flag = true
  result=5

  for flag {}

  fmt.Println("Загадайте число от 1 до 10")
  fmt.Println("Число которое вы загадали это ",result," Да-1/Нет-0")
  fmt.Scan(&answer)
  if answer!=1 {
    fmt.Println("Ваше число меньше или больше? Если меньше, нажмите 1, если больше, нажмите 0")
    var temp, temp2 int
    fmt.Scan(&temp)
    for i:=0;i<3;i++ {
      if temp==0 {
        result++
        fmt.Println("Число которое вы загадали это ",result," Да-1/Нет-0")
        fmt.Scan(&temp2)
        if temp2==1 {fmt.Println("Ура! Я победил"); flag=true ;break}
      }
      if temp==1 {
        result--
        fmt.Println("Число которое вы загадали это ",result," Да-1/Нет-0")
        fmt.Scan(&temp2)
        if temp2==1 {fmt.Println("Ура! Я победил"); flag=true; continue}
      }
    }
  } else {fmt.Println("Ура! Я победил"); flag=true}
  if !flag {fmt.Println("Кожаный мешок победил")}
}
