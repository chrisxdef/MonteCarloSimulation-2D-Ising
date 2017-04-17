package main

import(
	"github.com/Heisler0/MonteCarloSimulation-2D-Ising/go/configs"
	"github.com/Heisler0/MonteCarloSimulation-2D-Ising/go/mcs"
	"math"
	"fmt"
	"encoding/json"
	"os"
	"io/ioutil"
	"runtime"
)

type IsingInput struct{
	Nrows int
	Ncols int
	Config int
	Npass int
	Nequil int
	High_temp float64
	Low_temp float64
	Temp_interval float64
}


func main(){

	runtime.GOMAXPROCS(4)

	filename := os.Args[1]
	file, e := ioutil.ReadFile(filename)
	if e != nil {
		fmt.Println("Read error: ", e)
		os.Exit(1)
	}

	var input IsingInput
	e = json.Unmarshal(file, &input)
	if e != nil {
		fmt.Println("Json error: ", e)
	}

	nrows, ncols := input.Nrows + 2, input.Ncols + 2
	config := input.Config
	npass, nequil := input.Npass, input.Nequil
	high_temp, low_temp, temp_interval, temp := input.High_temp, input.Low_temp, input.Temp_interval, 0.0

	nscans := (high_temp - low_temp)/temp_interval + 1.0

	ch := make(chan float64, int(nscans)*7)

	for iscan := 0.0; iscan < nscans; iscan++{
		temp = high_temp - temp_interval*iscan
		go func(temp float64){
			beta := 1.0/temp
			counter := 0.0
			energy_avg, energy2_avg, mag_avg, mag2_avg := 0.0, 0.0, 0.0, 0.0
			mag, eng := 0.0, 0.0
			A := configs.Get(config, nrows, ncols)
			for ipass := 0; ipass < npass; ipass++{
				if ipass >= nequil{
					counter++
					mch := make(chan float64)
					go func(){
						sum := 0
						for i:=1; i<nrows-1; i++{
							for j:=1; j<ncols-1; j++{
								sum += A[i*nrows+j]
							}
						}
						mch<-float64(sum)
					}()
					ech := make(chan float64)
					go func(){
						sum := 0
						for i:=1; i<nrows-1; i++{
							for j:=1; j<ncols-1; j++{
								sum -= A[i*nrows+j]*(A[(i-1)*nrows+j] + A[(i+1)*nrows+j] + A[i*nrows+(j-1)] + A[i*nrows+(j+1)])
							}
						}
						ech<-float64(sum)

					}()
					mag = (<-mch)/float64(nrows*ncols)
					mag_avg += mag
					mag2_avg += math.Pow(mag, 2)
					eng = <-ech/float64(nrows*ncols*2)
					energy_avg += eng
					energy2_avg += math.Pow(eng, 2)
				}
				mcs.Simulate(A, nrows, ncols, beta)
			}

			ch<-temp

			ch<-math.Abs(mag_avg/counter)
			ch<-mag2_avg/counter
			ch<-beta*(mag2_avg/counter - math.Pow(mag_avg/counter, 2))
			ch<-energy_avg/counter
			ch<-energy2_avg/counter
			ch<-math.Pow(beta, 2)*(energy2_avg/counter - math.Pow(energy_avg/counter, 2))
		}(temp)
	}
	output, e := os.Create("../output/results-"+filename[9:len(filename)-5]+"-go.csv")
	if e != nil{
		fmt.Println("Cannot create file: ", e)
	}
	defer output.Close()

	fmt.Fprintf(output, "temperature,ave_magnetization,ave_magnetization^2,susceptibility,ave_energy,ave_energy^2,C_v\n")
	for i:=0.0; i<nscans; i++{
		fmt.Fprintf(output, "%f,%f,%f,%f,%f,%f,%f\n",<-ch,<-ch,<-ch,<-ch,<-ch,<-ch,<-ch)

	}
}
