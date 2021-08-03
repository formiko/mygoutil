package mygoutil

func LeftFirstGreater(arr []int) (ans []int) {
	ans = make([]int, len(arr))
	s := StackInt{}
	for k, v := range arr {
		for !s.empty() && arr[s.top()] <= v {
			s.pop()
		}
		if s.empty() {
			ans[k] = -1
		} else {
			ans[k] = s.top()
		}
		s.push(k)
	}
	return ans
}