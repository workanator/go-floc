package floc

import (
	"bytes"
	"fmt"
)

// ResultMask is the mask of possible results.
type ResultMask Result

const emptyResultMask = ResultMask(0)

// EmptyResultMask returns empty result set.
func EmptyResultMask() ResultMask {
	return emptyResultMask
}

// NewResultMask constructs new instance from the mask given.
func NewResultMask(mask Result) ResultMask {
	return ResultMask(mask & usedBitsMask)
}

// IsMasked tests if the result is masked.
func (mask ResultMask) IsMasked(result Result) bool {
	return mask&ResultMask(result) == ResultMask(result)
}

// IsEmpty returns true if no result is masked.
func (mask ResultMask) IsEmpty() bool {
	return mask == 0
}

func (mask ResultMask) String() string {
	buf := &bytes.Buffer{}

	fmt.Fprint(buf, "[")
	empty := true
	for _, result := range []Result{None, Completed, Canceled, Failed} {
		if mask.IsMasked(result) {
			if empty {
				fmt.Fprint(buf, result.String())
			} else {
				fmt.Fprint(buf, ",", result.String())
			}
			empty = false
		}
	}
	fmt.Fprint(buf, "]")

	return buf.String()
}
