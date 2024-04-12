package substrings

import (
	"github.com/Grishun/problems/task/slices/pkg/slices"
	"sort"
	"strings"
)

func InArray(array1 []string, array2 []string) (res []string) {

	for _, str := range array2 {
		for _, substr := range array1 {
			if strings.Contains(str, substr) {
				res = append(res, substr)
				break
			}
		}
	}
	if len(res) == 0 {
		return nil
	}

	sort.SliceStable(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	return slices.RemoveRepeating(res)
}
