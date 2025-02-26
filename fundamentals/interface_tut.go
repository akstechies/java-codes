package main

import "fmt"

type Animal interface {
	AngrySound()
	HappySound()
}

type Cat string

func (cat Cat) Attack() {
	fmt.Println("Cat is fcking attacking")
}

func (cat Cat) Name() string {
	return string(cat)
}

func (cat Cat) AngrySound() {
	fmt.Println("Cat Hisssss")
}

func (cat Cat) HappySound() {
	fmt.Println("Cat Meoww")
}

func InterfaceTut() {

	var kitty Animal = Cat("Kitty")
	kitty.AngrySound()
	kitty.HappySound()

	fmt.Println("==----")

	var cat1 Cat = kitty.(Cat)
	fmt.Println(cat1.Name())
	cat1.Attack()
	cat1.AngrySound()
	cat1.HappySound()

	fmt.Println("==----")

	var cat2 Cat = "Huli Huli"
	fmt.Println(cat2.Name())
	cat2.Attack()
	cat2.AngrySound()
	cat2.HappySound()
}
