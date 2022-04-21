package util

func Reverse[S ~[]E, E any](s S) S {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func IntersectNotNull[S ~[]E, E comparable](xs S, ys S) bool {
	for _, x := range xs {
		for _, y := range ys {
			if x == y {
				return true
			}
		}
	}

	return false
}
