package utils

import "strings"

const (
	MergeReplace MergeStrategy = iota
	MergeAppend
	MergeSkip
)

// MergeStrategy decide merge resolution when key conflict
type MergeStrategy int

// CombineMapString return combined maps without modifying inputs
// default MergeStrategy is using MergeReplace
func CombineMapString(a, b map[string]string, merge MergeStrategy) map[string]string {
	res := make(map[string]string, len(a)+len(b))
	for k, v := range a {
		res[k] = v
	}
	for k, v := range b {
		if res[k] != "" {
			switch merge {
			case MergeReplace:
				res[k] = v
			case MergeAppend:
				res[k] = stringArrayToArrayString(
					[]string{
						res[k],
						v,
					})
			case MergeSkip:
			default:
				res[k] = v
			}
		} else {
			res[k] = v
		}
	}
	return res
}

func stringArrayToArrayString(input []string) string {
	return `["` + strings.Join(input, `","`) + `"]`
}
