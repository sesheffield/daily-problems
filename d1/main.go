package main

func main() {

}

func pairOfSums(l []int64, k int64) [][]int64 {
	m := make(map[int64]bool)
	for _, v := range l {
		m[v] = true
	}

	results := make([][]int64, 0)
	counter := make(map[int64]bool)
	for _, v := range l {
		if _, ok := m[k-v]; ok {
			if _, ok := counter[v]; !ok {
				if _, ok := counter[k-v]; !ok {
					results = append(results, []int64{v, k - v})
					counter[k-v] = false
					counter[v] = false
				}
			}
		}
	}
	return results
}
