package mypackage

import "fmt"

type Export struct {
}

func (c Export) Hello() {
	fmt.Println("Hello")
}

func (c Export) DoMagic() {
	fmt.Println("Magic function was called")
}
