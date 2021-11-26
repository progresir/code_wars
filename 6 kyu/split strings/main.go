package main

import "fmt"

func Solution(str string) []string {
	res :=make([]string,0)
	for i:=0;len(str)/2>i;i++{
			res=append(res,str[(i*2):(i*2+2)])
	}
	if len(str)%2!=0{
		res=append(res,str[len(str)-1:]+"_")
	}
	return res
}

func main(){
	fmt.Println(Solution("abc"))
}