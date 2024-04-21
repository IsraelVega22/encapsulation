package course

import "fmt"

type Course struct {
	Name    string
	Price   float32
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

func (c *Course) ChangePrice(price float32) {
	c.Price = price
}

//trabajar solo con los valores sin actualizar nada received normal
func (c Course) PrintClass() {
	text := "Las Clases son: "
	for _, class := range c.Classes {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])

}
