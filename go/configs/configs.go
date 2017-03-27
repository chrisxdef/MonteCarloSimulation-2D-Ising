package configs

import(
	//"fmt"
)

const CHECKERBOARD int= 1
const INTERFACE int = 2
const UNEQUAL int = 3

func Get(configType int, nrows int, ncols int) []int{

	var result []int

	switch configType{
		case CHECKERBOARD:
			result = checkerboard(nrows, ncols)
		case INTERFACE:
			result = interfaceboard(nrows, ncols)
		case UNEQUAL:
			result = unequalinterface(nrows, ncols)
	}

	return result
}

func checkerboard(rows int, cols int) []int{
	result := make([]int, rows*cols)
	result[0] = 1
	for i := 1; i<rows; i++{
		result[i] = -1 * result[i-1]
	}

	ch := make(chan int, cols)
	for j := 0; j<cols-1; j++{
		go func(col int){
			for k := 1; k<rows; k++{
				result[k*rows+col] = -1 * result[(k-1)*rows+col]
			}
			ch<-1
		}(j)
	}
	<-ch
	return result
}

func interfaceboard(rows int, cols int) []int{
	result := make([]int, rows*cols)
	ch := make(chan int, rows)
	for i := 0; i<rows; i++{
		go func(row int){
			for j := 0; j<cols/2; j++{
				result[row*rows+j] = 1
			}
			ch<-1
		}(i)
		go func(row int){
			for j := cols/2; j<cols; j++{
				result[row*rows+j] = -1
			}
			ch<-1
		}(i)
	}
	<-ch
	return result
}

func unequalinterface(rows int, cols int) []int{
	result := make([]int, rows*cols)
	ch := make(chan int, rows)
	for i := 0; i<rows; i++{
		go func(row int){
			for j := 0; j<cols/2; j++{
				result[row*rows+j] = 1
			}
			ch<-1
		}(i)
		go func(row int){
			for j := cols/2; j<cols; j++{
				result[row*rows+j] = -1
			}
			ch<-1
		}(i)
	}
	<-ch
	return result
}





