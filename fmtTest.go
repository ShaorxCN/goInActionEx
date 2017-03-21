package main

import (
	"fmt"
)

type person struct {
	name string
}

func main() {
	s := "13"
	d := 13
	p := person{name: "gkond"}
	fmt.Printf("%v\n", p)  //{gkond}
	fmt.Printf("%+v\n", p) //{name:gkond}
	fmt.Printf("%#v\n", p) //main.person{name:"gkond"}
	fmt.Printf("%T\n", p)  //main.person
	fmt.Printf("%T\n", d)  //int
	fmt.Printf("%q\n", s)  //"13"
	fmt.Printf("%q\n", d)  //'\r'

	fmt.Printf("%%\n")         //"13"
	fmt.Printf("%t\n", 1 == 2) //false
	fmt.Printf("%b\n", d)      //1101
	fmt.Printf("%b\n", 1.13)   //5089067578928660p-52
	fmt.Printf("%c\n", 97)     //a
	fmt.Printf("%d\n", d)      //13
	fmt.Printf("%o\n", d)      //15
	fmt.Printf("%x\n", d)      //d
	fmt.Printf("%x\n", s)      //3133
	fmt.Printf("%X\n", d)      //D
	fmt.Printf("%X\n", s)      //3133

	fmt.Printf("%U\n", d)     //U+000D
	fmt.Printf("%e\n", 11.13) //1.113000e+01
	fmt.Printf("%E\n", 11.13) //1.113000E+01
	fmt.Printf("%f\n", 11.13) //11.130000
	fmt.Printf("%F\n", 11.13) //11.130000
	fmt.Printf("%g\n", 11.13) //11.13
	fmt.Printf("%G\n", 11.13) //11.13
	fmt.Printf("%s\n", s)     //13
	fmt.Printf("%p\n", &s)    //0xc04203e1d0

	fmt.Printf("start%9.5fend\n", 1.13)   //start  1.13000end
	fmt.Printf("start%.1fend\n", 1.13)    //start1.1end
	fmt.Printf("start%3.fend\n", 1111.13) //start1111end

	fmt.Printf("%+d\n", d) //+13
	fmt.Printf("%+q\n", s) //"13"

	fmt.Printf("start%-9.5fend\n", 1.13) //start1.13000  end

	fmt.Printf("%#p\n", &s)      //c04203e1d0
	fmt.Printf("%#q\n", s)       //`13`
	fmt.Printf("%#x\n", d)       //0xd
	fmt.Printf("%09.5f\n", 1.13) //001.13000

	fmt.Printf("start% d\n", d) //start 13
	fmt.Printf("% X\n", s)      //31 33

}
