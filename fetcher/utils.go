package fetcher

import "github.com/margostino/openearth/common"

func matchStringFor(expected *string, current string) bool {
	return expected == nil || (expected != nil && *expected == current)
}

func containsString(expected *string, current string) bool {
	return expected == nil || (expected != nil && common.NewString(current).ToLower().Contains(*expected))
}
