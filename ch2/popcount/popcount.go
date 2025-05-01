package popcount

var pc[256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// exercise 2.3
func PopCountByAccumulate(x uint64) int {
	count := 0
	for i := range 8 {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}

// exercise 2.4: Test most right bit of x
func PopCountByBitCount(x uint64) int {
	count := 0
	for x > 0 {
		count += int(x & 1)
		x >>= 1
	}
	return count
}
