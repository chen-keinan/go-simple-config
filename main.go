package main

import "fmt"

func main() {
	c := New()
	err := c.Load("./config.default.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(c.GetValueString("POLICY.retention"))
}
