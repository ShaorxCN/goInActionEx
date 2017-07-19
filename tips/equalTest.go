package main

import "log"


type test struct{
	a string
	b int
}

func main(){
	a := 1
	b := "s"
	c := "s"
	d := 1
	e := new(test)
	f := &test{}
	var g test
	h :=test{a:"name",b:1}
	i :=test{a:"name",b:1}

	log.Println(a==d)
	log.Println(b==c)
	log.Println(e ==f)
	log.Println(g)
	log.Println(h ==i)
	log.Println(e,f)
}
