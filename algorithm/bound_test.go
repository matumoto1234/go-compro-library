package algorithm

import "testing"

func Test_LowerBound_Int(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a    []int
		v    int
		want int
	}{
		{
			name: "a:{1, 2, 3, 4, 5} v:3",
			a:    []int{1, 2, 3, 4, 5},
			v:    3,
			want: 2,
		},
		{
			name: "a:{1, 2, 3, 4, 5} v:0",
			a:    []int{1, 2, 3, 4, 5},
			v:    0,
			want: 0,
		},
		{
			name: "a:{1, 2, 3, 4, 5} v:6",
			a:    []int{1, 2, 3, 4, 5},
			v:    6,
			want: 5,
		},
		{
			name: "a:{1, 2, 3, 4, 5} v:2",
			a:    []int{1, 2, 3, 4, 5},
			v:    2,
			want: 1,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := LowerBound(test.a, test.v)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}

func Test_UpperBound_Int(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a    []int
		v    int
		want int
	}{
		{
			name: "a:{1, 2, 3, 4, 5} v:3",
			a:    []int{1, 2, 3, 4, 5},
			v:    3,
			want: 3,
		},
		{
			name: "a:{1, 2, 3, 4, 5} v:0",
			a:    []int{1, 2, 3, 4, 5},
			v:    0,
			want: 0,
		},
		{
			name: "a:{1, 2, 3, 4, 5} v:1",
			a:    []int{1, 2, 3, 4, 5},
			v:    1,
			want: 1,
		},
		{

			name: "a:{1, 2, 3, 4, 5} v:6",
			a:    []int{1, 2, 3, 4, 5},
			v:    6,
			want: 5,
		},
		{

			name: "a:{1, 2, 3, 4, 5} v:5",
			a:    []int{1, 2, 3, 4, 5},
			v:    5,
			want: 5,
		},
		{
			name: "a:{1, 2, 3, 4, 5} v:2",
			a:    []int{1, 2, 3, 4, 5},
			v:    2,
			want: 2,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := UpperBound(test.a, test.v)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
