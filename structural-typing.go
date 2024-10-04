package main
import "fmt"

type Walker interface {
  Walk()
}

type Human struct {}

func (h Human) Walk() {
  fmt.Println("Human is walking...")
}

type Robot struct {}

func (r Robot) Walk(){
  fmt.Println("Robot is walking...")
}

func main(){
  var w Walker
  var h Human
  var r Robot

  w = h
  w.Walk()

  w = r
  w.Walk()
}




