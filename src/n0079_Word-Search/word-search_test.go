package n0079

import "testing"
import "github.com/stretchr/testify/assert"

type Para struct {
	one [][]byte
	two string
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
		{
			p: Para{
				one: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				two: "ABCCED",
			},
			a: Ans{
				one: true,
			},
		},
		{
			p: Para{
				one: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				two: "SEE",
			},
			a: Ans{
				one: true,
			},
		},
		{
			p: Para{
				one: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				two: "ABCB",
			},
			a: Ans{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, exist(p.one, p.two), "input %v", p)
	}
}
