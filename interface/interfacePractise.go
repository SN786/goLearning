package main

import "fmt"

type I interface{
    voice()
    vegornonveg()
    petorwild()
}


type Animal struct{
    aName string
    voicei string
    canBePet bool
    veg bool
}

func(animal Animal) voice(){

    fmt.Println("The voice of" + animal.aName + "is"+animal.voicei)
}
func(animal Animal) vegornonveg(){
    if animal.veg == true{
        fmt.Println(animal.aName + "is "+ "Vegetarian")
    }else{
        fmt.Println(animal.aName + "is "+ "Non-Vegetarian")
    }

    
}
func(animal Animal) petorwild(){
    if animal.canBePet == true{
        fmt.Println(animal.aName + "can be "+ "Pet")
    }else{
        fmt.Println(animal.aName + "can't be "+ "Pet")
    }
}
func main(){
    var dog I = Animal{"doggy","barks",true,false}
    dog.voice()
    dog.vegornonveg()
    dog.petorwild()


    var elephant I = Animal{"elephantii","trumpet",false,true}

    elephant.voice()
    elephant.vegornonveg()
    elephant.petorwild()

}