package flex

type FlexParent struct {
	FlexBase
	pSum int
	children *[]Flex
}

func (fp FlexParent) coefficient(flex Flex) float32 {
	return float32(flex.Priority()) / float32(fp.pSum);
}
