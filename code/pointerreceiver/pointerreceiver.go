package main

import "fmt"

//START OMIT
type Object struct {
	Value int
}

func (o Object) SetValue(v int) {
	o.Value = v
}

func (o *Object) SetPointerValue(v int) {
	o.Value = v
}

//END OMIT

func main() {
	o := Object{Value: 42}

	fmt.Printf("Original value: %d\n", o.Value)

	o.SetValue(5)
	fmt.Printf("After setting on the object method: %d\n", o.Value)

	o.SetPointerValue(100)
	fmt.Printf("After setting on the pointer method: %d\n", o.Value)
}
