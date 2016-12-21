package main

import "fmt"

func NextPrime() <-chan int {

	out := make(chan int, 1)

	go func() {

		out <- 2
		composites := make(map[int][]int)
		num := 3

		for {

			if _, ok := composites[num]; !ok {

				out <- num
				composites[num * num] = []int{num}

			} else {

				for _, prime := range composites[num] {

					next := num + prime

					for next % 2 == 0 {
						next += prime
					}

					if _, ok := composites[next]; ok {

						composites[next] = append(composites[next], prime)

					} else {

						composites[next] = []int{prime}

					}

				}
				delete(composites, num)
			}

			num += 2
		}
	}()

	return out
}

func main() {

	primes := NextPrime()

	for i := 0; i < 100; i++ {
		<-primes
	}

	fmt.Println(<-primes)
}