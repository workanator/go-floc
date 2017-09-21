package floc

// ResultMask is the mask of possible results.
type ResultMask Result

const emptyResultMask = ResultMask(0)

// EmptyResultMask returns an empty result set.
func EmptyResultMask() ResultMask {
	return emptyResultMask
}

// NewResultMask constructs new instance from the mask given.
func NewResultMask(mask Result) ResultMask {
	return ResultMask(mask & usedBitsMask)
}

// Contains tests if the mask masks the result.
func (mask ResultMask) Contains(result Result) bool {
	return mask&ResultMask(result) == ResultMask(result)
}

// IsEmpty returns true if the mask does not mask any result.
func (mask ResultMask) IsEmpty() bool {
	return mask == 0
}
