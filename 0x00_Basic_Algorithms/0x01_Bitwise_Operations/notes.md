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

#### 📐 core method: DP with bitwise

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

### Wake Up Difficult Greedy

#### Problem Description

An Boss defend matrix has build by n defend door, each door's attribute contain one operation $op_i$ and one parameter $t_i$, operation must be one of OR, XOR, AND, parameter is positive integer. the player attack will be x, after one gate will go to $x op_i t_i$. the final damage $x_0$ will be the value x through all gate. The Initial attack only can be a integer in range [0, m]. now we want to know the suitable initial attack which make the biggest attack.

#### core method Greedy with bitwise

The primary characteristic of bitwise operations is that no carry occurs in binary. Therefore, each bit operates independently. We can systematically evaluate each bit from high to low, deciding whether to set it to 0 or 1. (Greedy approach: Since we seek the maximum value, 1000 is naturally more valuable than 0111.) The condition for setting a bit to 1 is that the sum of the already-set higher-order bits plus the current bit shifted left by k does not exceed the required value m.

## 0x0103 Lowbit Calculation

### What is lowbit(n)?

Lowbit is defined as the value formed by the “least significant bit 1 followed by zeros” in binary representation for a non-negative integer n. For example: if n is 10, $(1010)_2$ the last one and remain 0 is 10, so lowbit(10) is 2.

### What is the formula?

First, invert n. At this point, the kth bit becomes 0, and the following bits become 1. Let n = n + 1. Now, the most significant bit to the k+1st bit are exactly the opposite of the original values. Therefore, n & (~n+1) has only k as 1, with all others being 0. In two's complement representation, ~n = -1 - n, so:
$$
lowbit(n) = n \& (~n + 1) = n \& (-n)
$$
