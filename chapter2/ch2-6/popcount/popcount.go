package popcount

//0~256までの数を2進数であわらした時の1が設定されているビット数
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount1(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount2(x uint64) int {
	var i uint
	b := 0
	for i = 0; i <= 7; i++ {
		b += int(pc[byte(x>>(i*8))])
	}
	return b
}

func PopCount3(x uint64) int {
	var i uint
	b := 0
	for i = 0; i < 64; i++ {
		if (x>>i)&1 == 1 {
			b++
		}
	}
	return b
}

func PopCount4(x uint64) int {
	b := 0
	for x > 0 {
		x &= x - 1
		b++
	}
	return b
}
