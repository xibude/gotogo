package main

import "fmt"



type  user struct{
	name string
	age int
}

type admin struct {
	user
	level int
}

type classmate struct{
	u user
	grade int32
}

func main () {
	u1 := user{"tom",13}
	a1 := admin{user{"lucy",18}, 1}
	a2 := admin{u1, 1}
	a1 = a2
	c1 := classmate{user{"jack",20},2}

	fmt.Println("u1's name is", u1.name)
	fmt.Println("adm's name is", a1.name, ". age is", a1.age, "level is", a1.level )
	fmt.Println("c1's name is",c1.u.name)
}


