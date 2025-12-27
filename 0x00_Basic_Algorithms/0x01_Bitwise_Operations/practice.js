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

console.log(fastPower(3n,5n,100n));
