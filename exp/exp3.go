package main

import "fmt"

func main() {
	mapDemo()
	sliceDemo()
}

func sliceDemo() {
	s := make([]int,0)
	s = append(s, 1,2,3)
	for _,v:=range s{
		if v==2{
			v=10
		}
	}
	fmt.Println(s)
}

func mapDemo() {
	m := make(map[string]interface{})
	m["a"]=1
	m["b"]=2
	m["c"]=3
	for _,v := range m{
		if v==2 {
			v=10
		}
	}

	fmt.Print(m)
}
