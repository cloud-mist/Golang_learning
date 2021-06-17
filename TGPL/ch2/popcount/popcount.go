package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// E3：改写为循环
func PopCountE3(x uint64) int {
	res := 0

	for i := 0; i < 8; i++ {
		res += int(pc[byte(x>>(i*8))])
	}
	return res
}

// E4：移位和舍弃位来增加当前位的数，
func PopCountE4(x uint64) int {
	res := 0

	for x != 0 {
		res += int(x & 1)
		x >>= 1
	}
	return res
}

// E5:
func PopCountE5(x uint64) int {
	res := 0

	for x != 0 {
		res++
		x = x & (x - 1) // 将最低的一个非0位清零
	}
	return res
}
