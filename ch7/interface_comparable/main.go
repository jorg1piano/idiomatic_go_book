package main

import "fmt"

type Doubler interface {
	Double()
}

type DoubleInt int
type DoubleIntSlice []int

func (d *DoubleInt) Double() {
	*d = *d * 2
}

func (d DoubleIntSlice) Double() {
	for i := 0; i < len(d); i++ {
		d[i] = d[i] * 2
	}
}

func DoublerCompare(d1, d2 Doubler) {
	fmt.Println(d1 == d2)
}

func main() {
	var di DoubleInt = 10
	di.Double()

	fmt.Println(di)

	var dis = DoubleIntSlice{1, 2, 3}
	dis.Double()

	for _, v := range dis {
		fmt.Println(v)
	}

	DoublerCompare(&di, dis)
	DoublerCompare(dis, dis)

}
