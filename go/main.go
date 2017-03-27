package main

import(
	"github.com/Heisler0/MonteCarloSimulation-2D-Ising/go/configs"
	"github.com/Heisler0/MonteCarloSimulation-2D-Ising/go/mcs"
	"math"
	"fmt"
)

func main(){

	nrows, ncols := 20 + 2, 20 + 2
	config := configs.INTERFACE
	npass, nequil := 5000, 2500
	high_temp, low_temp, temp_interval, temp := 400.0, 200.0, 5.0, 0.0

	nscans := (high_temp - low_temp)/temp_interval + 1.0

	ch := make(chan float64, int(nscans)*8)

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
					mag = <-mch
					mag_avg += mag
					mag2_avg += math.Pow(mag, 2)
					eng = <-ech
					energy_avg += eng
					energy2_avg += math.Pow(eng, 2)
				}
				mcs.Simulate(A, nrows, ncols, beta)
			}
			ch<-temp
			ch<-math.Abs(mag_avg/counter)
			ch<-mag2_avg/counter
			ch<-beta*(mag2_avg/counter - math.Pow(mag_avg, 2))
			ch<-temp
			ch<-energy_avg/counter
			ch<-energy2_avg/counter
			ch<-beta*(energy2_avg/counter - math.Pow(energy_avg, 2))
		}(temp)
	}
	for i:=0.0; i<nscans; i++{
		fmt.Println(<-ch,<-ch,<-ch,<-ch)
		fmt.Println(<-ch,<-ch,<-ch,<-ch)
	}
}
