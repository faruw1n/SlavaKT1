

//  пакет 1
// package animals

// type Swimmer interface {
//  Swim() int
// }

// type Animal interface {
//  Sound() string
//  Move() string
// }

// // Реализация структур животных
// type Dolphin struct {
//  speed string
// }

// func (d Dolphin) Sound() string {
//  return "Дельфины издают звуки, похожие на свист и щелчки."
// }

// func (d Dolphin) Move() string {
//  return "Дельфины могут плавать со скоростью до 60 км/ч."
// }

// func (d Dolphin) Swim() int {
//  return 10 // Хорошая оценка плавания дельфина
// }

// type Squirrel struct {
//  climbSkill string
// }

// func (s Squirrel) Sound() string {
//  return "Белки издают щебетание и шорох."
// }

// func (s Squirrel) Move() string {
//  return "Белки умеют быстро бегать и лазить по деревьям."
// }

// type Giraffe struct {
//  NoSwimming string
// }

// func (g Giraffe) Sound() string {
//  return "Жирафы издают низкие звуки, похожие на фырканье."
// }

// func (g Giraffe) Move() string {
//  return "Жирафы могут бегать со скоростью до 60 км/ч."
// }

// type Parrot struct {
//  AbilityTalk string
// }

// func (p Parrot) Move() string {
//  return "Попугаи могут быстро летать, но не очень быстро на земле."
// }

// func (p Parrot) Sound() string {
//  return "Попугаи могут имитировать звуки и говорить."
// }

// type Penguin struct {
//  CantFly string
// }

// func (p Penguin) Move() string {
//  return "Средняя скорость, которую пингвины развивают в воде, — 5–10 км/ч."
// }

// func (p Penguin) Sound() string {
//  return "На суше они общаются посредством громких криков."
// }

//  пакет 2
// package interaction

// import "fmt"

// func PrintAnimal(a animals.Animal) {
//  fmt.Println("Звуки:", a.Sound())
//  fmt.Println("Движения:", a.Move())
// }

// func PrintSwim(s animals.Swimmer) {
//  fmt.Println("Плавает: ", s.Swim())
// }

//  пакет 3git add README.md
// package utils

// import "fmt"

// // Функция для вывода сообщения
// func PrintMessage(message string) {
//  fmt.Println(message)
// }
