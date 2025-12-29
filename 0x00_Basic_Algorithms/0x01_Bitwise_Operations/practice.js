/**
 * 📘 AlgorithmKit - 0x01 Bitwise Operations Module
 */
const AlgorithmKit = {
    // =========================================================
    // 1. solve a^b mod p (Fast Power)
    // =========================================================
    Power: {
        calculate(a, b, p) {
            let a_big = BigInt(a), b_big = BigInt(b), p_big = BigInt(p);
            let ans = 1n % p_big;
            for (; b_big > 0n; b_big >>= 1n) {
                if (b_big & 1n) ans = (ans * a_big) % p_big;
                a_big = (a_big * a_big) % p_big;
            }
            return ans;
        },
        test() {
            console.log("--- 🧪 Testing Power (a^b % p) ---");
            const res = this.calculate(3, 5, 100);
            console.log(`Result: 3^5 % 100 = ${res}n ${res === 43n ? '✅' : '❌'}\n`);
        }
    },

    // =========================================================
    // 2. (Binary State Compression)
    // =========================================================
    BinaryState: {
        create(initial) {
            /**
             * BinaryState Class - abstract bitwise logic
             */
            return new class {
                constructor(init) {
                    this.states = Array.isArray(init) ? this._fromArray(init) : BigInt(init);
                }
                _fromArray(boolArray) {
                    let res = 0n;
                    for (let i = 0; i < boolArray.length; i++) {
                        if (boolArray[i]) res |= (1n << BigInt(i));
                    }
                    return res;
                }
                get(k) { return Boolean((this.states >> BigInt(k)) & 1n); }
                setTrue(k) {
                    this.states |= (1n << BigInt(k));
                }
                setFalse(k) {
                    this.states &= ~(1n << BigInt(k));
                }
                toggle(k) {
                    this.states ^= (1n << BigInt(k));
                }
                countSetBits() {
                    let count = 0, temp = this.states;
                    while (temp > 0n) { temp &= (temp - 1n); count++; }
                    return count;
                }
            }(initial);
        },
        test() {
            console.log("--- 🧪 Testing BinaryState Wrapper ---");
            const bs = this.create([true, false, true, true]); // 1101 (13)
            const count = bs.countSetBits();
            console.log(`State: ${bs.states}n, Count Bits: ${count} ${count === 3 ? '✅' : '❌'}\n`);
        }
    },

    // =========================================================
    // 3. Shortest Hamilton Distance (DP)
    // In Javascript, 
    // 2^20 x 20 data points will let V8 engine to crash, 
    // so that we need to use Int32Array to 
    // flatten 2D array memory leak and enhance CPU cache hit.
    // =========================================================
    Hamilton: {
        solve(n, rawWeight) {
            const INF = 0x3f3f3f3f;
            const limit = 1 << n;
            const weight = new Int32Array(n * n);
            for (let r = 0; r < n; r++) {
                for (let c = 0; c < n; c++) weight[r * n + c] = rawWeight[r][c];
            }
            const f = new Int32Array(limit * n).fill(INF);
            f[1 * n + 0] = 0;

            for (let i = 1; i < limit; i++) {
                for (let j = 0; j < n; j++) {
                    if ((i >> j) & 1) {
                        const pre = i ^ (1 << j);
                        if (pre === 0) continue;
                        for (let k = 0; k < n; k++) {
                            if ((pre >> k) & 1) {
                                const cost = f[pre * n + k] + weight[k * n + j];
                                if (cost < f[i * n + j]) f[i * n + j] = cost;
                            }
                        }
                    }
                }
            }
            return f[(limit - 1) * n + (n - 1)];
        },
        test() {
            console.log("--- 🧪 Testing Hamilton Shortest Path ---");
            const n = 3;
            const matrix = [[0, 10, 100], [10, 0, 10], [100, 10, 0]];
            console.time("Hamilton Timer");
            const res = this.solve(n, matrix);
            console.timeEnd("Hamilton Timer");
            console.log(`Shortest Path: ${res} ${res === 20 ? '✅' : '❌'}\n`);
        }
    },

    // =========================================================
    // 4. wake up difficult
    // =========================================================
    SleepyDragon: {
        solve(n, m, ops) {
            let val = 0, ans = 0;
            // from highest bit greedy
            for (let bit = 29; bit >= 0; bit--) {
                let res0 = this._check(bit, 0, ops);
                let res1 = this._check(bit, 1, ops);
                
                // 1 if current res higher and less than m when add last val
                // 0 for other case
                if (val + (1 << bit) <= m && res1 > res0) {
                    val += (1 << bit);
                    ans |= (res1 << bit);
                } else {
                    ans |= (res0 << bit);
                }
            }
            return ans;
        },
        _check(bit, now, ops) {
            for (let op of ops) {
                let x = (op.val >> bit) & 1;
                if (op.type === 'AND') now &= x;
                else if (op.type === 'OR') now |= x;
                else if (op.type === 'XOR') now ^= x;
            }
            return now;
        },
        test() {
            console.log("--- 🧪 Testing Sleepy Dragon (Greedy) ---");
            const ops = [
                { type: 'AND', val: 5 }, // 0101
                { type: 'OR', val: 6 },  // 0110
                { type: 'XOR', val: 7 }  // 0111
            ];
            const res = this.solve(3, 10, ops);
            console.log(`Max Result: ${res} ✅\n`);
        }
    },
};

AlgorithmKit.Power.test();      
AlgorithmKit.Hamilton.test();     
AlgorithmKit.SleepyDragon.test();