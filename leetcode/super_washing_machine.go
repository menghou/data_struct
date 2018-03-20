package leetcode

import "fmt"

/*
https://leetcode.com/problems/super-washing-machines/description/

You have n super washing machines on a line. Initially, each washing machine has some dresses or is empty.

For each move, you could choose any m (1 ≤ m ≤ n) washing machines,
and pass one dress of each washing machine to one of its adjacent washing machines at the same time .

Given an integer array representing the number of dresses in each washing machine from left to right
on the line, you should find the minimum number of moves to make all the washing machines
have the same number of dresses. If it is not possible to do it, return -1.

Example1

Input: [1,0,5]

Output: 3

Explanation:
1st move:    1     0 <-- 5    =>    1     1     4
2nd move:    1 <-- 1 <-- 4    =>    2     1     3
3rd move:    2     1 <-- 3    =>    2     2     2
Example2

Input: [0,3,0]

Output: 2

Explanation:
1st move:    0 <-- 3     0    =>    1     2     0
2nd move:    1     2 --> 0    =>    1     1     1
Example3

Input: [0,2,0]

Output: -1

Explanation:
It's impossible to make all the three washing machines have the same number of dresses.


Note:
The range of n is [1, 10000].
The range of dresses number in a super washing machine is [0, 1e5].
 */

func findMinMoves(machines []int) int {
	sum := 0
	for _, v := range machines {
		sum += v
	}
	if sum % len(machines) != 0 {
		return -1
	} else {
		minMoves := 0
		everyValueShouldBe := sum / len(machines)
		everyMachineMissZero := make([]int, len(machines))
		for i , value := range machines {
			everyMachineMissZero[i] = value - everyValueShouldBe
		}
		for i := range everyMachineMissZero {
			for everyMachineMissZero[i] < 0 {
				nearestIndex := findNearestOverIndex(i, everyMachineMissZero)
				fmt.Printf("%d - %d|%d\n", i, everyMachineMissZero[i], nearestIndex)
				if nearestIndex < 0 {
					panic(fmt.Sprintf("Err: %d %v", i, everyMachineMissZero))
				}
				everyMachineMissZero[i] += 1
				everyMachineMissZero[nearestIndex] += 1
				minMoves += 1
			}
		}
		return minMoves
	}
}

func findNearestOverIndex(index int, everyMachineMissZero []int) int {
	if everyMachineMissZero[index] > 0 {
		return -1
	}
	overIndex := make([]int, 0)
	for i := 0; i < len(everyMachineMissZero); i ++ {
		if everyMachineMissZero[i] > 0 {
			overIndex = append(overIndex, i)
		}
	}
	fmt.Printf("overIndex: %v\n", overIndex)
	if len(overIndex) == 0 {
		return -1
	}
	var minStep int = len(everyMachineMissZero)
	var minIndex int = -1
	for _, i := range overIndex {
		step := i - index
		if step < 0 {
			step = -step
		}
		if step < minStep {
			minStep	= step
			minIndex = i
		}
	}
	return minIndex
}
