package n0062

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Para struct {
	one int
	two int
}

type Ans struct {
	one int
}

type Question struct {
	p Para
	a Ans
}

func Test_uniquePaths(t *testing.T) {
	ast := assert.New(t)

	qs := []Question{
		{p: Para{one: 3, two: 7}, a: Ans{one: 28}},
		{p: Para{one: 3, two: 2}, a: Ans{one: 3}},
	}

	for _, q := range qs {
		ast.Equal(q.a.one, uniquePaths(q.p.one, q.p.two))
	}
}
