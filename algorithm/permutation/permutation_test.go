package permutation

import (
	"reflect"
	"testing"
)

func Test_NextPermutation_Int(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		a        []int
		want     []int
		wantBool bool
	}{
		{
			name:     "{}",
			a:        []int{},
			want:     []int{},
			wantBool: false,
		},
		{
			name:     "{1}",
			a:        []int{1},
			want:     []int{1},
			wantBool: false,
		},
		{
			name:     "{1,2}",
			a:        []int{1, 2},
			want:     []int{2, 1},
			wantBool: true,
		},
		{
			name:     "{1,2,3}",
			a:        []int{1, 2, 3},
			want:     []int{1, 3, 2},
			wantBool: true,
		},
		{
			name:     "{3,5,1}",
			a:        []int{3, 5, 1},
			want:     []int{5, 1, 3},
			wantBool: true,
		},
		{
			name:     "{1,2,3,4,5}",
			a:        []int{1, 2, 3, 4, 5},
			want:     []int{1, 2, 3, 5, 4},
			wantBool: true,
		},
		{
			name:     "{5,4,3,2,1}",
			a:        []int{5, 4, 3, 2, 1},
			want:     []int{1, 2, 3, 4, 5},
			wantBool: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := Next(test.a)
			if got != test.wantBool {
				t.Errorf("got %v, want %v", got, test.wantBool)
			}

			if !reflect.DeepEqual(test.a, test.want) {
				t.Errorf("got %v, want %v", test.a, test.want)
			}
		})
	}
}

func Test_PrevPermutation_Int(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		a        []int
		want     []int
		wantBool bool
	}{
		{
			name:     "{}",
			a:        []int{},
			want:     []int{},
			wantBool: false,
		},
		{
			name:     "{1}",
			a:        []int{1},
			want:     []int{1},
			wantBool: false,
		},
		{
			name:     "{2, 1}",
			a:        []int{2, 1},
			want:     []int{1, 2},
			wantBool: true,
		},
		{
			name:     "{3,2,1}",
			a:        []int{3, 2, 1},
			want:     []int{3, 1, 2},
			wantBool: true,
		},
		{
			name:     "{3,5,1}",
			a:        []int{3, 5, 1},
			want:     []int{3, 1, 5},
			wantBool: true,
		},
		{
			name:     "{1,2,3,4,5}",
			a:        []int{1, 2, 3, 4, 5},
			want:     []int{5, 4, 3, 2, 1},
			wantBool: false,
		},
		{
			name:     "{5,4,3,2,1}",
			a:        []int{5, 4, 3, 2, 1},
			want:     []int{5, 4, 3, 1, 2},
			wantBool: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := Prev(test.a)
			if got != test.wantBool {
				t.Errorf("got %v, want %v", got, test.wantBool)
			}

			if !reflect.DeepEqual(test.a, test.want) {
				t.Errorf("got %v, want %v", test.a, test.want)
			}
		})
	}

}
