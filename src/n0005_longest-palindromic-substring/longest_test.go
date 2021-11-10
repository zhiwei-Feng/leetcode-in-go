package n0005

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Para struct {
	one string
}
type Ans struct {
	one string
}
type Question struct {
	p Para
	a Ans
}

func Test_longestPalindrome(t *testing.T) {
	ast := assert.New(t)

	qs := []Question{
		{p: Para{one: "babad"}, a: Ans{one: "bab"}},
		{p: Para{one: "cbbd"}, a: Ans{one: "bb"}},
		{p: Para{one: "a"}, a: Ans{one: "a"}},
		{p: Para{one: "ac"}, a: Ans{one: "a"}},
		{p: Para{one: "ccc"}, a: Ans{one: "ccc"}},
	}

	for _, q := range qs {
		ast.Equal(len(q.a.one), len(longestPalindrome(q.p.one)), q.a.one)
	}
}
