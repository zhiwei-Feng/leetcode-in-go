package offer0067

import "testing"

func Test_strToInt(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {"case1", args{"42"}, 42},
		// {"case2", args{"   -42"}, -42},
		// {"case3", args{"4193 with words"}, 4193},
		// {"case4", args{"words and 987"}, 0},
		// {"case5", args{"-91283472332"}, -2147483648},
		// {"case6", args{"9223372036854775808"}, 2147483647},
		{"case7", args{"-5-"}, -5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strToInt(tt.args.str); got != tt.want {
				t.Errorf("strToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
