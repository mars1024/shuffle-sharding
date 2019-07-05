package main

import (
	"fmt"
	"math/rand"
)

func shuffleSharding(v, queue uint64, handSize int) []int {
	if handSize < 1 {
		return nil
	}

	ret := make([]int, handSize)

	for i := 0; i < handSize; i++ {
		ret[i] = int(v % (queue - uint64(i)))
	}

	fmt.Printf("A set: %+v \n", ret)

	for i := 0; i < handSize; i++ {
		for j := 0; j < i; j++ {
			if ret[j] <= ret[i] {
				ret[i]++
			}
		}
	}

	fmt.Printf("I set: %+v \n", ret)

	return ret
}

func main() {
	numQueues := uint64(128)
	handSize := 6

	queue := make([]int, numQueues)
	retries := 1000000

	for i := 0; i < retries; i++ {
		hashValue := rand.Uint64()
		cards := shuffleSharding(hashValue, numQueues, handSize)

		for _, card := range cards {
			queue[card]++
		}
	}

	fmt.Printf("Queue statistics: %+v \n", queue)
}
