package main

import(
	"github.com/Heisler0/MonteCarloSimulation-2D-Ising/go/configs"
	"github.com/Heisler0/MonteCarloSimulation-2D-Ising/go/mcs"
	//"fmt"
)

func main(){

	nrows, ncols := 20 + 2, 20 + 2

	_ = configs.Get(configs.CHECKERBOARD, nrows, ncols)
	inter := configs.Get(configs.INTERFACE, nrows, ncols)
	_ = configs.Get(configs.UNEQUAL, nrows, ncols)

	mcs.Simulate(inter, 22, 22, 1)
}
