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

	// you can choose two kind of algorithms to get real indices
	// getIndicesByHandQueue(int(queue), handSize, ret)
	getIndicesByConflictMap(handSize, ret)

	fmt.Printf("I set: %+v \n", ret)

	return ret
}

func getIndicesByHandQueue(queue, handSize int, hands []int) {
	handQueue := make([]int, queue)

	for i := 0; i < queue; i++ {
		handQueue[i] = 0
	}

	for i := 0; i < handSize; i++ {
		ii, ai := 0, hands[i]+1
		for {
			if handQueue[ii] == 0 {
				ai--
			}
			if ai == 0 {
				handQueue[ii] = 1
				break
			}
			ii++
		}
		hands[i] = ii
	}
}

func getIndicesByConflictMap(handSize int, hands []int){
	conflict := make(map[int]bool)

	for i := 0; i < handSize; i++ {
		for j := 0; j < i; j++ {
			if hands[j] <= hands[i] {
				hands[i]++
			}
		}
		for {
			if ! conflict[hands[i]] {
				break
			}
			hands[i]++
		}
		conflict[hands[i]] = true
	}
}

func main() {
	numQueues := uint64(128)
	handSize := 8

	queue := make([]int, numQueues)
	retries := 10000

	for i := 0; i < retries; i++ {
		hashValue := rand.Uint64()
		indices := shuffleSharding(hashValue, numQueues, handSize)
		for _, index := range indices {
			// fill queue
			queue[index]++
		}
	}

	fmt.Printf("Queue statistics: %+v \n", queue)
}
