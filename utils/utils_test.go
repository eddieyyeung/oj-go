package utils

import (
	"testing"
)

func TestGenerateNode(t *testing.T) {
	var case1 = []*int{
		PointerInt(1),
		PointerInt(2),
		PointerInt(3),
		PointerInt(4),
		PointerInt(5)}
	node := GenerateNode(case1)
	logger.Sugar().Infow("case1", "node", node)

	var case2 = []*int{
		PointerInt(4),
		PointerInt(5),
		PointerInt(2),
		nil,
		nil,
		PointerInt(3),
		PointerInt(1),
	}
	node = GenerateNode(case2)
	logger.Sugar().Infow("case2", "node", node)
}
