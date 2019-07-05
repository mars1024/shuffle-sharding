package main

import (
	"fmt"
	"math"
	"math/rand"
)

var numQueues uint64

func shuffleDealAndPick(v, nq uint64,
	mr func(int /*in [0, nq-1]*/) int, /*in [0, numQueues-1] and excluding previously determined members of I*/
	nRem, minLen, bestIdx int) int {
	if nRem < 1 {
		return bestIdx
	}
	vNext := v / nq
	ai := int(v - nq*vNext)
	ii := mr(ai)
	i := numQueues - nq // i is used only for debug printing
	mrNext := func(a int /*in [0, nq-2]*/) int /*in [0, numQueues-1] and excluding I[0], I[1], ... ii*/ {
		if a < ai {
			fmt.Printf("mr[%v](%v) going low\n", i, a)
			return mr(a)
		}
		fmt.Printf("mr[%v](%v) going high\n", i, a)
		return mr(a + 1)
	}
	lenI := lengthOfQueue(ii)
	fmt.Printf("Considering A[%v]=%v, I[%v]=%v, qlen[%v]=%v\n\n", i, ai, i, ii, i, lenI)
	if lenI < minLen {
		minLen = lenI
		bestIdx = ii
	}
	return shuffleDealAndPick(vNext, nq-1, mrNext, nRem-1, minLen, bestIdx)
}

func lengthOfQueue(i int) int {
	return i % 10 // hack for this PoC
}

func main() {
	numQueues = uint64(128)
	handSize := 12
	hashValue := rand.Uint64()
	queueIndex := shuffleDealAndPick(hashValue, numQueues, func(i int) int { return i }, handSize, math.MaxInt32, -1)
	fmt.Printf("For V=%v, numQueues=%v, handSize=%v, chosen queue is %v\n", hashValue, numQueues, handSize, queueIndex)
}