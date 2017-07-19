package main

import (
	"log"
)

func main(){
	s := "123中国人!"

	//按照字符遍历
	for k,v := range s{
		log.Println(k,v)
	}

	//按照byte遍历
	for i := 0;i<len(s);i++{
		log.Println(i,s[i])
	}

	//s[1] = 2 不允许
	b := []byte(s)
	b[0] = 51
	b[4] = 223
	log.Println(b)
	log.Println(string(b))//这是针对字节修改的，如果是想将数字转化为字符串请使用strconv包

	r := []rune(s)//字符缓冲区
	log.Println(r)
	r[0] = '3'
	r[3] = '是'
	log.Println(r,string(r))


}
