package mygoutil

func VOP(a []int, p []int) (v []int) {
	v = make([]int, len(a))
	for k := range a {
		if -1 == p[k] {
			v[k] = -1
			continue
		}
		v[k] = a[p[k]]
	}
	return v
}
