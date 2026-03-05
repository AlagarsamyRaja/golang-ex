package fibonacciSeries

func CheckFibonacci(out chan int, num int) {
	a := -1
	b := 1

	for i := 1; i <= num; i++ {
		c := a + b
		out <- c
		a, b = b, c
	}

	close(out)
}
