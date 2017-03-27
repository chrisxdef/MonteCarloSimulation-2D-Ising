package mcs

import(
	"math/rand"
	"math"
)

func Simulate(A []int, nrows int, ncols int, beta float64){

	//Randomly choose a spin
	m, n := rand.Intn(nrows-2) + 1, rand.Intn(ncols-2) + 1
	trial_spin := A[m*nrows+n]

	//Find the change in energy
	deltaU := float64(-1.0 * trial_spin * (A[(m-1)*nrows+n]+A[(m+1)*nrows+n]+A[m*nrows+(n-1)]+A[m*nrows+(n+1)]) * 2)
	log_eta := rand.Float64() + math.Pow(1.0, -10.0)

	if math.Exp(-1.0 * beta * deltaU) > log_eta{
		A[m*nrows+n] = trial_spin
		if m == 1{ A[(nrows-1)*nrows+n] = trial_spin }
		if m == nrows-2{ A[nrows+n] = trial_spin }
		if n == 1{ A[m*nrows+(ncols-1)] = trial_spin }
		if n == ncols-2{ A[m*nrows+1] = trial_spin }
	}
}
