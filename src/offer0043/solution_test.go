package offer0043

import "testing"

func Test_countDigitOne(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		// {"1", args{12}, 5},
		// {"2", args{13}, 6},
		{"3", args{10}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countDigitOne(tt.args.n); gotAns != tt.wantAns {
				t.Errorf("countDigitOne() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
