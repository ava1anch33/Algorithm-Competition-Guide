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

---

### 💻 Implementation

The algorithm iteratively checks the LSB (Least Significant Bit) using `b & 1` and shifts `b` rightward.

#### [c++]

```cpp
int power(int a, int b, int p) {
    int = 1 % p;
    for (; b; b >>= 1) {
        if (b & 1) ans = (long long) ans * a % p;
        a = (long long) a * a % p;
    }
    return ans;
}
```
