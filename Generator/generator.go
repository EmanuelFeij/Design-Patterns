package Generator

import "fmt"

func Count(n int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			out <- i
		}
	}()

	return out
}

func main() {
	for i := range Count(200) {
		fmt.Println(i)
	}
}
