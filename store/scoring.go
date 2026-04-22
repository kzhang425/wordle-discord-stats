package store

func computeAverages(results []WordleResult) map[string]float64 {
	scores := map[string][]int{}
	for _, r := range results {
		if !r.Complete {
			continue
		}
		key := PlayerKey(r)
		scores[key] = append(scores[key], r.Score)
	}

	avgs := make(map[string]float64, len(scores))
	for key, ss := range scores {
		sum := 0
		for _, s := range ss {
			sum += s
		}
		avgs[key] = float64(sum) / float64(len(ss))
	}
	return avgs
}
