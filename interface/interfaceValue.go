package main

import "fmt"

type s struct{
	name string
	age int
}
func findTypeValue(i interface{}){

	fmt.Printf("type %T and value %v",i,i)

}

func main(){
	
	s:=s{"sharif",19}
	findTypeValue(s)
}