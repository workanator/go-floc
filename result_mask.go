package floc

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
