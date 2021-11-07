package n0055

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Para struct {
	one []int
}

type Ans struct {
	one bool
}

type Question struct {
	p Para
	a Ans
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []Question{
		{p: Para{one: []int{2, 3, 1, 1, 4}}, a: Ans{one: true}},
		{p: Para{one: []int{3, 2, 1, 0, 4}}, a: Ans{one: false}},
		{p: Para{one: []int{0}}, a: Ans{one: true}},
		{p: Para{one: []int{2, 0, 0}}, a: Ans{one: true}},
	}

	for _, q := range qs {
		ast.Equal(q.a.one, canJump(q.p.one))
	}
}
