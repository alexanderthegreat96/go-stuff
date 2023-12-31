package main

import (
	"fmt"

	"github.com/alexanderthegreat96/go-stuff/package-helpers/helpers"
)

var (
	p, q, r string
	n       int
	a, b, c *string
	d       *int
)

func init() {
	p = "Drop by drop, the bucket is filled."
	q = "Persistently, partiently, you are bound to succeed."
	r = "The meaning of life is..."
	n = 42
	a = &p // points ot p
	b = &q // points to q
	c = &r // points to r
	d = &n // points to n
}

func main() {
	fmt.Printf("%v \t is of type \t %T \n", p, p)
	fmt.Printf("%v \t is of type \t %T \n", q, q)
	fmt.Printf("%v \t is of type \t %T \n", r, r)
	fmt.Printf("%v \t is of type \t %T \n", n, n)
	fmt.Printf("%v \t is of type \t %T \n", a, a)
	fmt.Printf("%v \t is of type \t %T \n", b, b)
	fmt.Printf("%v \t is of type \t %T \n", c, c)
	fmt.Printf("%v \t is of type \t %T \n", d, d)

	helpers.Exit()
}
