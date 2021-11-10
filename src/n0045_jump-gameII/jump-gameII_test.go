package n0045

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

func Test_jump(t *testing.T) {
	ast := assert.New(t)

	qs := []Question{
		{p: Para{one: []int{2, 3, 1, 1, 4}}, a: Ans{one: 2}},
		{p: Para{one: []int{2, 3, 0, 1, 4}}, a: Ans{one: 2}},
		{p: Para{one: []int{0}}, a: Ans{one: 0}},
		{p: Para{one: []int{1, 2, 3}}, a: Ans{one: 2}},
	}

	for _, q := range qs {
		ast.Equal(q.a.one, jump(q.p.one))
	}
}
