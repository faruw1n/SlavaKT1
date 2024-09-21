package main

import "fmt"

type Swimmer interface {
 Swim() int
}

type Animal interface {
 Sound() string
 Move() string
}

type Dolphin struct {
 speed string
}

func (d Dolphin) Sound() string {
 return "Дельфины издают звуки, похожие на свист и щелчки."
}

func (d Dolphin) Move() string {
 return "Дельфины могут плавать со скоростью до 60 км/ч."
}

func (d Dolphin) Swim() int {
 return 10 
}

type Squirrel struct {
 climbSkill string
}

func (s Squirrel) Sound() string {
 return "Белки издают щебетание и шорох."
}

func (s Squirrel) Move() string {
 return "Белки умеют быстро бегать и лазить по деревьям."
}

type Giraffe struct {
 NoSwimming string
}

func (g Giraffe) Sound() string {
 return "Жирафы издают низкие звуки, похожие на фырканье."
}

func (g Giraffe) Move() string {
 return "Жирафы могут бегать со скоростью до 60 км/ч."
}

type Parrot struct {
 AbilityTalk string
}

func (p Parrot) Move() string {
 return "Попугаи могут быстро летать, но не очень быстро на земле."
}

func (p Parrot) Sound() string {
 return "Попугаи могут имитировать звуки и говорить."
}

type Penguin struct {
 CantFly string
}

func (p Penguin) Move() string {
 return "Средняя скорость, которую пингвины развивают в воде, — 5–10 км/ч."
}

func (p Penguin) Sound() string {
 return "На суше они общаются посредством громких криков."
}

func printAnimal(a Animal) {
 fmt.Println("Звуки:", a.Sound())
 fmt.Println("Движение:", a.Move())
}

func printSwim(s Swimmer) {
 fmt.Println("Плавает: ", s.Swim())
}

func main() {
 dolphin := Dolphin{
  speed: "Достигает 60 км/ч",
 }

 squirrel := Squirrel{
  climbSkill: "Умеет прекрасно лазать по деревьям.",
 }

 giraffe := Giraffe{
  NoSwimming: "Не умеет плавать.",
 }

 parrot := Parrot{
  AbilityTalk: "Может имитировать человеческую речь.",
 }

 penguin := Penguin{
  CantFly: "Единственная птица, которая не умеет летать, но отлично плавает в воде.",
 }

 var input int
 fmt.Println("Введи цифру от 1 до 5:")
 fmt.Scanf("%d\n", &input)

 switch input {
 case 1:
  fmt.Println("Информация о дельфине:")
  printAnimal(dolphin)
  fmt.Println("Уникальная информация:", dolphin.speed)
 case 2:
  fmt.Println("Информация о белке:")
  printAnimal(squirrel)
  fmt.Println("Уникальная информация:", squirrel.climbSkill)
 case 3:
  fmt.Println("Информация о жирафе:")
  printAnimal(giraffe)
  fmt.Println("Уникальная информация:", giraffe.NoSwimming)
 case 4:
  fmt.Println("Информация о попугае:")
  printAnimal(parrot)
  fmt.Println("Уникальная информация:", parrot.AbilityTalk)
 case 5:
  fmt.Println("Информация о пингвине:")
  printAnimal(penguin)
  fmt.Println("Уникальная информация:", penguin.CantFly)
 default:
  fmt.Println("Вы превысили ограничение цифр")
 }
}


