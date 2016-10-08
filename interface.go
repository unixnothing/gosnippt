package main

import "fmt"

type animal interface {
    run()
    shout()
}

type cat struct {
    name string
}

type dog struct {
    name string
}

func (c cat) run() {
    fmt.Println("slow run")
}

func (c cat) shout() {
    fmt.Println("miao miao")
}

func (d *dog) run() {
    fmt.Println("fast run")
}

func (d *dog) shout() {
    fmt.Println("wang wang")
}

func check(a animal) {
    a.run()
    a.shout()
}

func main() {
    cat := cat{"Tom"}
    catty := &cat
    check(cat)
    check(catty) //Go automatically convert pointer to value for method calls

    ////dog dosn't impl the animal, but *dog
    //dog := dog{"Mike"}
    //check(dog)

    doggy := &dog{"Mike"}
    check(doggy)
}

