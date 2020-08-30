package insertsort

import "errors"

func sort(items []int) ([]int, error) {
	if len(items) == 0 {
		return nil, errors.New("Items can't be emtpy")
	}

	for j := 1; j < len(items); j++ {
		key := items[j]

		i := j - 1
		for i >= 0 && items[i] > key {
			items[i+1] = items[i]
			i--
		}
		items[i+1] = key
	}

	return items, nil
}
