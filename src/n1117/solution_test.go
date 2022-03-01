package n1117

import "testing"

func TestPrint(t *testing.T) {
	type args struct {
		water string
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{"OOHHHH"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Print(tt.args.water)
		})
	}
}
