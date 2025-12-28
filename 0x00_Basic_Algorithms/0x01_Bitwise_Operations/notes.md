# 0x01 Bitwise Operations

![Book](https://img.shields.io/badge/Guide-0x01-blueviolet.svg)
![Language](https://img.shields.io/badge/Language-c++-blue.svg)

Bitwise operations are the fundamental units of computation. Mastering these operations allows for high-performance optimizations and a deeper understanding of memory manipulation in competitive programming.

---

## 💡 Fundamental Bitwise Operations

| Operation | Symbol | Go Syntax | JS Syntax | Rule (Binary) | Use Case |
| :--- | :---: | :--- | :--- | :--- | :--- |
| **AND** | `&` | `a & b` | `a & b` | `1` iff both bits are `1` | Parity check (`n & 1`), Bit masking |
| **OR** | `\|` | `a \| b` | `a \| b` | `1` if any bit is `1` | Setting specific bits |
| **XOR** | `^` | `a ^ b` | `a ^ b` | `1` if bits are different | Toggle bits, Addition w/o carry |
| **NOT** | `~` / `^` | `^a` | `~a` | Invert all bits | Flipping all states |
| **L-Shift** | `<<` | `a << k` | `a << k` | Shift left by `k` bits | Multiply by $2^k$ |
| **R-Shift** | `>>` | `a >> k` | `a >> k` | Shift right by `k` bits | Divide by $2^k$ (floor) |

---

## 0x0101: Fast Exponentiation (a^b mod p)

### 📐 Mathematical Principle

Every positive integer $b$ can be uniquely decomposed into a sum of powers of 2 (Binary Representation). If $b$ has $k$ digits in binary, where $c_i \in \{0, 1\}$ is the digit at position $i$:

$$
b = \sum_{i=0}^{k-1} c_i 2^i
$$

Thus, the power $a^b$ can be rewritten as a product:
$$
a^b = a^{\sum c_i 2^i} = \prod_{i=0}^{k-1} a^{c_i 2^i} = a^{c_{k-1}2^{k-1}} \times \dots \times a^{c_02^0}
$$

> **Key Observation:** > 1. $k = \lfloor \log_2 b \rfloor + 1$
> 2. Each term $a^{2^i}$ can be computed by squaring the previous term: $a^{2^i} = (a^{2^{i-1}})^2$.

## 0x0102 Binary state compression

suppose we have a boolean array with length m, we can use a binary Integer to store it. This method is easy to calculate, save time and space. we can define a class to abstract it.

### Shortest Hamilton Path

#### 🚩 Problem Description

Given a weighted undirected graph with $n$ nodes ($n \le 20$), find the shortest path that starts at node $0$, ends at node $n-1$, and visits every node **exactly once**.

#### 📐 core method: DP

if we use brute-force, the complex will be n!, when n is 20, the machine will crash, it absently not gonna solve.
use **Dynamic Program**，we can reduce the complex to $O(n^2 2^n)$，about $4 \times 10^8$ times calculations.

##### 1. define state

we use a 2-D array `f[state][j]`：

* **`state`**: an Binary Integer, represent for **Set that all Node have been accessed**（if $k$-th is $1$，then $k$ has been accessed）。
* **`j`**: current**destination**。
* **`f[state][j]`**: the shortest distance in current state and destination j。

##### 2. transfer state

to access state $(state, j)$，last point must has $k$，and $k$ belong to $state$ not equal $j$。
$$f[state][j] = \min_{k \in \{state \setminus j\}} \{ f[state \setminus j][k] + weight[k][j] \}$$

* **bitwise operation `state \ j`**: `state ^ (1 << j)`

### 💻 Implementation

#### [JavaScript - Optimized with TypedArray]

In Javascript, $2^{20} \times 20$ data points will let V8 engine to crash, so that we need to use Int32Array to flatten 2D array memory leak and enhance CPU cache hit.

```javascript
/**
 * shortest Hamilton distance
 * @param {number} n - number of point
 * @param {number[][]} rawWeight - 2-D weight adjacent array
 * @returns {number}
 */
function solveHamilton(n, rawWeight) {
    const INF = 0x3f3f3f3f;
    const limit = 1 << n;
    
    // 1. flatten weight array
    const weight = new Int32Array(n * n);
    for(let i = 0; i < n; i++) 
        for(let j = 0; j < n; j++) weight[i * n + j] = rawWeight[i][j];

    // 2. use 1-D Int32Array instead of f[limit][n]
    const f = new Int32Array(limit * n).fill(INF);
    f[1 * n + 0] = 0; // first point

    // 3. state transfer
    for (let i = 1; i < limit; i++) {
        for (let j = 0; j < n; j++) {
            if ((i >> j) & 1) { // if i has j
                const pre = i ^ (1 << j); // get the state that before j
                if (pre === 0) continue; // if no point before j, continue.
                
                const curIdx = i * n + j;
                for (let k = 0; k < n; k++) {
                    if ((pre >> k) & 1) { // find last possible k
                        const cost = f[pre * n + k] + weight[k * n + j];
                        if (cost < f[curIdx]) f[curIdx] = cost;
                    }
                }
            }
        }
    }
    return f[(limit - 1) * n + (n - 1)];
}
