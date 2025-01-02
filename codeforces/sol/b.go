package sol

func B(n int, a []int) string {
	sol := "NO"
	sum := 0
	for _, v := range a {
		sum += v
	}
	len := len(a)
	if sum%len != 0 {
		return sol
	}
	t := sum / len
	for i := 1; i < n - 1; i++ {
		if a[i-1] > t {
			for a[i-1] > t{
				fromPrev(i, a)
			}
		} else if a[i+1] > t {
			for a[i+1] > t {
				fromNext(i, a)
			}
		}
	}
	for i := 1; i < n; i++ {
		if a[i] != a[i-1] {
			return sol
		}
	}
	return "YES"
}

func fromNext(i int, a []int) {
	a[i-1] = a[i-1] + 1
	a[i+1] = a[i+1] - 1
}

func fromPrev(i int, a []int) {
	a[i-1] = a[i-1] - 1
	a[i+1] = a[i+1] + 1
}
