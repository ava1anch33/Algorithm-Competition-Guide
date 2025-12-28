package main

import (
	"fmt"
	"math/bits"
)

const (
	MaxN = 20
	INF  = 0x3f3f3f3f // 约 10^9，作为无穷大
)

// ---------------------------------------------------------
// 1. (Binary State Compression)
// ---------------------------------------------------------

// Define State as uint64，can handle up to 64 bit state compress
type State uint64

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

// NewStateFromArray
func NewStateFromArray(bools []bool) State {
	var s State = 0
	for i, v := range bools {
		if v {
			s.SetTrue(i)
		}
	}
	return s
}

// ---------------------------------------------------------
// shorten Hamilton distance
// ---------------------------------------------------------
func hamilton(n int, weight [][]int) int {
	// 1. 初始化 DP 表
	// f[state][j] 表示：经过的点集合为 state，且当前停在 j 点的最短距离
	// 状态总数：(1 << n)，每个状态有 n 个结尾点
	// 为了内存安全，这里使用一维切片模拟二维数组，或者直接开大数组
	// 这里演示标准二维逻辑
	limit := 1 << n
	f := make([][]int, limit)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = INF
		}
	}

	// 2. 起点初始化
	// 状态 1 (二进制 ...001) 代表只经过了 0 号点
	// 当前停在 0 号点，距离为 0
	f[1][0] = 0

	// 3. 状态转移
	// i 代表当前的状态 (mask)
	for i := 1; i < limit; i++ {
		// j 代表当前停在哪个点
		for j := 0; j < n; j++ {
			// 如果 i 的第 j 位是 1 (说明 j 在集合 i 中)
			if (i>>j)&1 == 1 {
				
				// preState: 除去 j 之前的状态 (i XOR 2^j)
				// 比如 i=111(7), j=1(中间位), preState=101(5)
				preState := i ^ (1 << j)
				
				// 如果 preState 为 0，说明 i 只有 j 这一位是 1 (即起点情况)
				// 我们在循环外已经初始化 f[1][0]=0，这里直接跳过空状态
				if preState == 0 {
					continue
				}

				// k 代表上一步停在哪
				for k := 0; k < n; k++ {
					// 只有当 preState 中包含 k 时，才能从 k 走到 j
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

	// 4. 返回结果
	// 状态：(1<<n)-1 代表所有位都是 1 (所有点都走过)
	// 终点：n-1
	return f[limit-1][n-1]
}

func main() {
	fmt.Println("=== Part 1: Fast Power Algorithm ===")
	// 计算 3^5 % 100
	res := Power(3, 5, 100)
	fmt.Printf("3^5 %% 100 = %d\n\n", res)

	fmt.Println("=== Part 2: State Compression ===")

	initialData := []bool{true, false, true}
	myState := NewStateFromArray(initialData)
	fmt.Printf("Initial State: %b (Decimal: %d)\n", myState, myState)

	myState.SetTrue(3)
	fmt.Printf("After setting bit 3: %b\n", myState)

	myState.Toggle(0)
	fmt.Printf("After toggling bit 0: %b\n", myState)

	bit2 := myState.GetKth(2)
	ones := myState.CountOnes()
	fmt.Printf("Value at bit 2: %d\n", bit2)
	fmt.Printf("Total bits set to 1: %d\n", ones)

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
	
	fmt.Println(hamilton(n, weights)) // should be 20
}