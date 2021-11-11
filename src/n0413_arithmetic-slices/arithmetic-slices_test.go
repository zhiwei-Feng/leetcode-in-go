package n0413

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Para struct {
	one []int
}
type Ans struct {
	one int
}

type Question struct {
	p Para
	a Ans
}

func Test_numberOfArithmeticSlices(t *testing.T) {
	ast := assert.New(t)

	qs := []Question{
		{p: Para{one: []int{1}}, a: Ans{one: 0}},
		{p: Para{one: []int{1, 2, 3, 4}}, a: Ans{one: 3}},
	}

	for _, q := range qs {
		ast.Equal(q.a.one, numberOfArithmeticSlices(q.p.one))
	}
}
