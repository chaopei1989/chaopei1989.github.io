package main

import "log"

type TaskCouple struct {
	left int
	next int
}

func leastInterval(tasks []byte, n int) int {
	// 每次取冷却结束的耗时最多的任务
	// 1. 建立二元组 (剩余任务数, 下次执行时间)
	taskCouples := make([]*TaskCouple, 26)
	for i := 0; i < len(tasks); i++ {
		task := taskCouples[tasks[i]-'A']
		if task == nil {
			task = new(TaskCouple)
			taskCouples[tasks[i]-'A'] = task
		}
		task.left++
	}
	// 2. tick start
	t := 0
	for ; ; t++ {
		// 2.1 select
		leftTaskCount := 0
		selected := -1
		for i, tc := range taskCouples {
			if tc == nil {
				continue
			}
			leftTaskCount++
			if tc.next <= t {
				// 可运行
				if selected < 0 {
					selected = i
				} else {
					if taskCouples[selected].left < tc.left {
						selected = i
					}
				}
			}
		}
		if leftTaskCount == 0 {
			break
		}
		// 2.2 exe
		if selected >= 0 {
			taskCouples[selected].left--
			taskCouples[selected].next = t+n+1

			if taskCouples[selected].left <= 0 {
				taskCouples[selected] = nil
			}
		}
	}
	return t
}

func main() {
	//leastInterval([]byte{'A','A','A','B','B','B'}, 2)
	a := 1
	b := &a
	c := &b
	var d ***int = &c
	log.Println(&d)

}