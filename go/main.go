package main

import(
	"github.com/Heisler0/MonteCarlo/go/configs"
	"fmt"
)

func main(){
	_ = configs.Get(configs.CHECKERBOARD, 20, 20)
	inter := configs.Get(configs.INTERFACE, 20, 20)
	for i := 0; i<len(inter); i++{
		fmt.Print(inter[i])
	}
	_ = configs.Get(configs.UNEQUAL, 20, 20)
}
