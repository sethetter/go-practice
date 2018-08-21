package main

import "fmt"

type Calculator struct {
	acc float64
}

func Add(v float64) func(float64) float64 {
	return func(x float64) float64 {
		return v + x
	}
}

func (c *Calculator) Do(ops ...func(float64) float64) float64 {
	for _, op := range ops {
		c.acc = op(c.acc)
	}
	return c.acc
}

func main() {
	c := &Calculator{acc: 0.0}
	result := c.Do(Add(5), Add(2))
	fmt.Println("Result: ", result)
}
