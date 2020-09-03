package main

var (
	states []map[byte]int
)

func isNumber(s string) bool {
	states = make([]map[byte]int, 9)

	state0 := make(map[byte]int, 0)
	state0[' '] = 0 // [初始]->空格->[初始]
	state0['s'] = 1 // [初始]->符号->[E前-小数点前-数字开始]
	state0['d'] = 2 // [初始]->数字->[E前-小数点前-数字]
	state0['.'] = 4 // [初始]->点->[E前-小数点后-数字开始], 如: .5e6
	states[0] = state0

	state1 := make(map[byte]int, 0)
	state1['d'] = 2 // [E前-小数点前-数字开始]->数字->[E前-小数点前-数字]
	state1['.'] = 4 // [E前-小数点前-数字开始]->点->[E前-小数点后-数字开始]
	states[1] = state1

	state2 := make(map[byte]int, 0)
	state2['d'] = 2 // [E前-小数点前-数字]->数字->[E前-小数点前-数字]
	state2['.'] = 3 // [E前-小数点前-数字]->点->[E前-小数点后-数字], 如: 1.e6
	state2['e'] = 5 // [E前-小数点前-数字]->e->[E后-数字开始]
	state2[' '] = 8 // [E前-小数点前-数字]->空格->[结束]
	states[2] = state2

	state3 := make(map[byte]int, 0)
	state3['d'] = 3 // [E前-小数点后-数字]->数字->[E前-小数点后-数字]
	state3['e'] = 5 // [E前-小数点后-数字]->e->[E后-开始]
	state3[' '] = 8 // [E前-小数点后-数字]->空格->[结束]
	states[3] = state3

	state4 := make(map[byte]int, 0)
	state4['d'] = 3 // [E前-小数点后-数字开始]->数字->[E前-小数点后-数字]
	states[4] = state4

	state5 := make(map[byte]int, 0)
	state5['s'] = 6 // [E后-开始]->符号->[E后-数字开始]
	state5['d'] = 7 // [E后-开始]->数字->[E后-数字]
	states[5] = state5

	state6 := make(map[byte]int, 0)
	state6['d'] = 7 // [E后-数字开始]->数字->[E后-数字]
	states[6] = state6

	state7 := make(map[byte]int, 0)
	state7['d'] = 7 // [E后-数字]->数字->[E后-数字]
	state7[' '] = 8 // [E后-数字]->空格->[结束]
	states[7] = state7

	state8 := make(map[byte]int, 0)
	state8[' '] = 8 // [结束]->空格->[结束]
	states[8] = state8

	var p int
	var t byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			t = 'd'
		} else if c == '+' || c == '-' {
			t = 's'
		} else if c == 'e' || c == 'E' {
			t = 'e'
		} else if c == '.' || c == ' ' {
			t = c
		} else {
			t = '?'
		}
		if _, ok := states[p][t]; !ok {
			return false
		}
		p = states[p][t]
	}
	return p == 2 || p == 3 || p == 7 || p == 8
}

func main() {

}
