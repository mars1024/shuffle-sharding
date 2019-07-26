package shufflesharding

import (
	"fmt"
	"math/big"
)

const maxHashBits = 64

func ValidateParameters(queueSize, handSize uint32) bool {
	if handSize == 0 || queueSize == 0 || handSize > queueSize {
		return false
	}

	p := big.NewInt(0)
	for i := uint32(handSize); i > 0; i-- {
		p.Add(p.Mul(p, big.NewInt(int64(queueSize-i+1))), big.NewInt(int64(queueSize-i)))
	}

	return p.BitLen() <= maxHashBits
}

func shuffleSharding(v, queue uint64, handSize int) []int {
	if handSize < 1 {
		return nil
	}

	as := make([]int, handSize)

	for i := 0; i < handSize; i++ {
		as[i] = int(v % (queue - uint64(i)))
	}

	fmt.Printf("A set: %+v \n", as)

	// you can choose two kind of algorithms to get real indices
	// ii := getIndicesByHandQueue(int(queue), handSize, as)
	ii := getIndices(handSize, as)

	fmt.Printf("I set: %+v \n", ii)

	return as
}

func getIndicesByHandQueue(queue, handSize int, as []int) []int {
	handQueue := make([]int, queue)

	for i := 0; i < queue; i++ {
		handQueue[i] = 0
	}

	for i := 0; i < handSize; i++ {
		ii, ai := 0, as[i]+1
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
		as[i] = ii
	}

	return as
}

func getIndices(handSize int, as []int) []int {
	hands := make([]int, handSize)

	for i := 0; i < handSize; i++ {
		ii := as[i]
		for j := i - 1; j >= 0; j-- {
			if ii >= as[j] {
				ii++
			}
		}
		hands[i] = ii
	}
	return hands
}
