package main

type CT int

const (
	EQUAL     = 0
	UNCONVERT = -1
	CONVERT   = 1
)

func convertType(left, right string, length int) CT {
	// 默认长度一致
	count := 0
	for i := 0; i < length; i++ {
		if left[i] != right[i] {
			count++
		}
	}
	if count == 0 {
		return EQUAL
	}
	if count == 1 {
		return CONVERT
	}
	return UNCONVERT
}

type node struct {
	level int
	word string
	next map[string]*node
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	set := make(map[string]*node, 0)
	//for _, word := range wordList {
	//	set[word] = null
	//}
	tmp := new(node)
	tmp.word = beginWord
	tmp.next = make(map[string]*node, 0)
	// 1. 构建图
	buildNode(tmp, wordList, set)
	// 2. 图的广度遍历
	m := make(map[string]struct{}, 0)
	queue := make([]*node, 0)
	tmp.level = 1
	queue = append(queue, tmp)
	for {
		if len(queue) == 0 {
			break
		}
		curr := queue[0]
		queue = queue[1:]
		for k, v := range curr.next {
			if _, ok := m[k]; ok {
				continue
			}
			v.level = curr.level + 1
			if endWord == k {
				return v.level
			}
			m[k] = struct{}{}
			queue = append(queue, v)
		}
	}
	return 0
}

func buildNode(tmp *node, wordList []string, set map[string]*node) {
	for _, r := range wordList {
		if cv := convertType(tmp.word, r, len(tmp.word)); cv == CONVERT {
			if newTmp, ok := set[r]; !ok {
				newTmp := new(node)
				newTmp.word = r
				newTmp.next = make(map[string]*node, 0)
				newTmp.next[tmp.word] = tmp
				tmp.next[r] = newTmp
				set[r] = newTmp
				buildNode(newTmp, wordList, set)
			} else {
				newTmp.next[tmp.word] = tmp
				tmp.next[r] = newTmp
			}
		} else if cv == UNCONVERT {
			// ignore
		} else { // EQUAL
			// ignore
		}
	}
}
