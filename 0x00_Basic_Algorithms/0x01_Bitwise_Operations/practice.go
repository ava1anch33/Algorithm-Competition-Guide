package main

import (
	"fmt"
	"math/bits"
)

const (
	MaxN = 20
	INF  = 0x3f3f3f3f
)

// ---------------------------------------------------------
// 1. (Binary State Compression)
// ---------------------------------------------------------

// Define State as uint64，can handle up to 64 bit state compress
type State uint64

// NewStateFromArray
func (s State) NewStateFromArray(bools []bool) State {
	var ss State = 0
	for i, v := range bools {
		if v {
			s.SetTrue(i)
		}
	}
	return ss
}

func (s State) GetKth(k int) int {
	return int((s >> k) & 1)
}

func (s *State) SetTrue(k int) {
	*s |= (1 << k)
}

func (s *State) SetFalse(k int) {
	*s &= ^(1 << k)
}

func (s *State) Toggle(k int) {
	*s ^= (1 << k)
}

func (s State) CountOnes() int {
	return bits.OnesCount64(uint64(s))
}

func (s State) test() {
	fmt.Println("=== Part 2: State Compression ===")
	initialData := []bool{true, false, true}
	myState := s.NewStateFromArray(initialData)
	fmt.Printf("Initial State: %b (Decimal: %d)\n", myState, myState)
	myState.SetTrue(3)
	fmt.Printf("After setting bit 3: %b\n", myState)
	myState.Toggle(0)
	fmt.Printf("After toggling bit 0: %b\n", myState)

	bit2 := myState.GetKth(2)
	ones := myState.CountOnes()
	fmt.Printf("Value at bit 2: %d\n", bit2)
	fmt.Printf("Total bits set to 1: %d\n", ones)
}

// ---------------------------------------------------------
// 2. (a^b) % p
// ---------------------------------------------------------

func Power(a, b, p int64) int64 {
	ans := int64(1) % p
	a %= p
	for ; b > 0; b >>= 1 {
		if b&1 == 1 {
			ans = (ans * a) % p
		}
		a = (a * a) % p
	}
	return ans
}

// ---------------------------------------------------------
// shorten Hamilton distance
// ---------------------------------------------------------
type Hamilton struct{}

func (h Hamilton) hamilton(n int, weight [][]int) int {
	limit := 1 << n
	f := make([][]int, limit)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = INF
		}
	}

	f[1][0] = 0

	for i := 1; i < limit; i++ {
		for j := 0; j < n; j++ {
			if (i>>j)&1 == 1 {
				preState := i ^ (1 << j)
				if preState == 0 {
					continue
				}

				for k := 0; k < n; k++ {
					if (preState>>k)&1 == 1 {
						cost := f[preState][k] + weight[k][j]
						if cost < f[i][j] {
							f[i][j] = cost
						}
					}
				}
			}
		}
	}
	return f[limit-1][n-1]
}

func (h Hamilton) test() {
	// test data: simple triangle graph
	// 0 -> 1: 10
	// 1 -> 2: 10
	// 0 -> 2: 100
	// best：0 -> 1 -> 2 (20)
	n := 3
	weights := [][]int{
		{0, 10, 100},
		{10, 0, 10},
		{100, 10, 0},
	}
	fmt.Println("=== Part 3: Hamilton shortest distance ===")
	fmt.Println(h.hamilton(n, weights)) // should be 20
}

// ---------------------------------------------------------
// wake up difficult greedy
// ---------------------------------------------------------
type Door struct {
	Op string
	Val int
}

type SleepyDragonModule struct{}

// calculate specific bit through all bitwise
func (s SleepyDragonModule) check(bit int, now int, doors []Door) int {
	for _, door := range doors {
		x := (door.Val >> bit) & 1;
		switch door.Op {
			case "AND": now &= x; break;
			case "OR": now |= x; break;
			default: now ^= x;break;
		}
	}
	return now;
}

func (s SleepyDragonModule) getAttack(n int, m int, doors []Door) int {
	val := 0;
	ans := 0;
	for bit := 29; bit >= 0; bit-- {
		res0 := s.check(bit, 0, doors)
		res1 := s.check(bit, 1, doors)
		if val+(1<<bit) <= m && res1 > res0 {
			val += (1 << bit)
			ans |= (res1 << bit)
		} else {
			ans |= (res0 << bit)
		}
	}
	return ans
}

func (s SleepyDragonModule) Test() {
	fmt.Println("--- 🧪 Testing Sleepy Dragon (Go Greedy) ---")
	doors := []Door{
		{"AND", 5}, // 0101
		{"OR", 6},  // 0110
		{"XOR", 7}, // 0111
	}
	n := 3
	m := 10
	result := s.getAttack(n, m, doors)
	fmt.Printf("Max Attack: %d ✅\n\n", result)
}

func main() {
	fmt.Println("=== Part 1: Fast Power Algorithm ===")
	res := Power(3, 5, 100)
	fmt.Printf("3^5 %% 100 = %d\n\n", res)

	var state State
	state.test()

	hal := Hamilton{}
	hal.test()

	sdk := SleepyDragonModule{}
	sdk.Test()
}