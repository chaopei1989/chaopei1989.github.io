package main

func judgeCircle(moves string) bool {
	ud := 0
	lr := 0
	for i := 0; i < len(moves); i++ {
		switch moves[i] {
		case 'U':
			ud--
			break
		case 'D':
			ud++
			break
		case 'L':
			lr--
			break
		case 'R':
			lr++
			break
		}
	}
	return ud == 0 && lr == 0
}

func main() {

}
