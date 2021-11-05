package n0198

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

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []Question{
		{p: Para{one: []int{2, 3, 2}}, a: Ans{one: 4}},
		{p: Para{one: []int{1, 2, 3, 1}}, a: Ans{one: 4}},
		{p: Para{one: []int{1, 2, 3}}, a: Ans{one: 4}},
		{p: Para{one: []int{2, 7, 9, 3, 1}}, a: Ans{one: 12}},
	}

	for _, q := range qs {
		ast.Equal(q.a.one, rob(q.p.one))
	}
}
