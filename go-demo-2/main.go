package main

import (
	"fmt"

	"example.com/go-demo-2/saturday"
)

type domatesSalcasi struct{}
type biberSalcasi struct{}

type kisi struct {
	ad  string
	yas int
}

func (j kisi) String() string {
	return fmt.Sprintf("%v (%v years)", j.ad, j.yas)
}

func do(k interface{}) {
	switch g := k.(type) {
	case int:
		{
			fmt.Printf("g is %d and g's type is %T\n", g, g)
		}
	case string:
		{
			fmt.Printf("g is %s and g's type is %T\n", g, g)
		}
	default:
		{
			fmt.Printf("I do not know about type %T\n", g)
		}

	}
}

func (s domatesSalcasi) Ye() { fmt.Println("Domates salcasi yenildi") }
func (s biberSalcasi) Ye()   { fmt.Println("Biber salcasi yenildi") }

type Salca interface {
	Ye()
}

func main() {
	fmt.Println(saturday.WhenSaturday())
	var salcam Salca = &biberSalcasi{}
	salcam.Ye()
	salcam = &domatesSalcasi{}
	salcam.Ye()

	i := salcam.(*domatesSalcasi)
	fmt.Println(i)
	i, ok := salcam.(*domatesSalcasi)
	fmt.Println(i, ok)

	do(21)
	do("merhaba")
	do(true)
	q := kisi{"Mustafa", 43}
	fmt.Println(q)
}
