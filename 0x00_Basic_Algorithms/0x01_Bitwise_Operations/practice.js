/**
 * calculate a^b mod p 1<= a,b,p <= 10^9
* @param {bigint} a
 * @param {bigint} b
 * @param {bigint} p
 * @returns {bigint}
 */
function fastPower(a ,b, p) {
    let ans = 1n % p;
    for(; b > 0n; b >>= 1n) {
        if (b & 1n) ans = (ans * a) % p;
        a = (a * a) % p;
    }
    return ans
}

// console.log(fastPower(3n,5n,100n)); 43n

/**
 * BinaryState - Class for Binary state compression
 */
class BinaryState {
    /**
     * Both support initial array or number
     * @param {Boolean[] | bigint | number} initial
     */
    constructor(initial = 0n) {
        if (Array.isArray(initial)) {
            this.states = BinaryState.fromArray(initial);
        } else {
            this.states = BigInt(initial);
        }
    }

    /**
     * Create compress state from boolean array
     * complex: O(N)
     * @param {Boolean[]} boolArray 
     */
    static fromArray(boolArray) {
        let res = 0n;
        for (let i = 0; i < boolArray.length; i++) {
            if (boolArray[i]) {
                res |= (1n << BigInt(i));
            }
        }
        return res;
    }

    get(k) {
        return Number((this.states >> BigInt(k)) & 1n);
    }

    setTrue(k) {
        this.states |= (1n << BigInt(k));
    }

    setFalse(k) {
        this.states &= ~(1n << BigInt(k));
    }

    toggle(k) {
        this.states ^= (1n << BigInt(k));
    }

    /**
     * Count how many 1 in current state
     * In state compression DP, we are frequently need to know the node which has been visited
     */
    countSetBits() {
        let count = 0;
        let temp = this.states;
        while (temp > 0n) {
            temp &= (temp - 1n); // remove the lowest 1.
            count++;
        }
        return count;
    }
}

// 使用示例
const initialState = [true, false, true, true]; // 代表二进制 1101 (十进制 13)
const bs = new BinaryState(initialState);
console.log(bs.states); // 13n
console.log(bs.get(2)); // 1

/**
 * 极致优化的 Hamilton 路径算法
 * @param {number} n - 点的数量
 * @param {number[][]} rawWeight - 原始二维邻接矩阵
 * @returns {number}
 */
function hamilton(n, rawWeight) {
    const INF = 0x3f3f3f3f;
    const limit = 1 << n;
    
    // for 2D array mapping Int32Array
    // a[i][j] = A[i * n + j]
    // transform 2-D weight array to 1-D Int32Array
    // reduce pointer jump
    const weight = new Int32Array(n * n);
    for (let r = 0; r < n; r++) {
        for (let c = 0; c < n; c++) {
            weight[r * n + c] = rawWeight[r][c];
        }
    }

    const f = new Int32Array(limit * n).fill(INF);
    f[1 * n + 0] = 0;

    for (let i = 1; i < limit; i++) {
        for (let j = 0; j < n; j++) {
            if ((i >> j) & 1) {
                const preState = i ^ (1 << j);
                // if have not to any other point when arrive j, continue
                if (preState === 0) continue;
                const currentIdx = i * n + j;
                for (let k = 0; k < n; k++) {
                    if ((preState >> k) & 1) {
                        // if have arrived k
                        const prevIdx = preState * n + k;
                        const weightIdx = k * n + j;
                        const cost = f[prevIdx] + weight[weightIdx];
                        if (cost < f[currentIdx]) {
                            f[currentIdx] = cost;
                        }
                    }
                }
            }
        }
    }

    return f[(limit - 1) * n + (n - 1)];
}

// --- 测试 ---
const n = 3;
const inputWeight = [
    [0, 10, 100],
    [10, 0, 10],
    [100, 10, 0]
];

console.time("Hamilton");
console.log(hamilton(n, inputWeight)); // 20
console.timeEnd("Hamilton");
