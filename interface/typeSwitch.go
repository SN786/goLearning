package main

import (  
    "fmt"
)
type st struct{
	name string
}
func findType(i interface{}) {  
    switch i.(type) {
	case st:
		v,_:=i.(st)
		fmt.Printf("It struct and value is %s\n",v.name)
    case string:
        fmt.Printf("It string and value is %s\n", i.(string))
    case int:
        fmt.Printf("It is int and  value is %d\n", i.(int))
    default:
        fmt.Printf("Unknown type\n")
    }
}
func main() {  
    findType("sohrat")
    findType(77)
    findType(89.98)
	var v=st{"sohart"}
	findType(v)
}