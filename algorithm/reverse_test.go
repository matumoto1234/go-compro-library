package algorithm

import (
	"reflect"
	"testing"
)

func Test_Reverse_Int(t *testing.T) {
	t.Parallel()

	type Test struct {
		name string
		a    []int
		want []int
	}

	tests := []Test{
		{
			name: "{}",
			a:    []int{},
			want: []int{},
		},
		{
			name: "{1}",
			a:    []int{1},
			want: []int{1},
		},
		{
			name: "{1,2,3,4,5}",
			a:    []int{1, 2, 3, 4, 5},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "{1,2,3,4,5,6}",
			a:    []int{1, 2, 3, 4, 5, 6},
			want: []int{6, 5, 4, 3, 2, 1},
		},
		{
			name: "{3,1,4,1,5,9,2,6,5,3,5}",
			a:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
			want: []int{5, 3, 5, 6, 2, 9, 5, 1, 4, 1, 3},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			Reverse(test.a)
			if !reflect.DeepEqual(test.a, test.want) {
				t.Errorf("got %v, want %v", test.a, test.want)
			}
		})
	}
}

func Test_Reverse_String(t *testing.T) {
	t.Parallel()

	type Test struct {
		name string
		a    []string
		want []string
	}

	tests := []Test{
		{
			name: "{}",
			a:    []string{},
			want: []string{},
		},
		{
			name: `{"1"}`,
			a:    []string{"1"},
			want: []string{"1"},
		},
		{
			name: `{"1","2","3","4","5"}`,
			a:    []string{"1", "2", "3", "4", "5"},
			want: []string{"5", "4", "3", "2", "1"},
		},
		{
			name: `{"1","2","3","4","5","6"}`,
			a:    []string{"1", "2", "3", "4", "5", "6"},
			want: []string{"6", "5", "4", "3", "2", "1"},
		},
		{
			name: `{"3","1","4","1","5","9","2","6","5","3","5""}`,
			a:    []string{"3", "1", "4", "1", "5", "9", "2", "6", "5", "3", "5"},
			want: []string{"5", "3", "5", "6", "2", "9", "5", "1", "4", "1", "3"},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			Reverse(test.a)
			if !reflect.DeepEqual(test.a, test.want) {
				t.Errorf("got %v, want %v", test.a, test.want)
			}
		})
	}
}

func Test_Reverse_Struct(t *testing.T) {
	t.Parallel()

	type MyStruct struct {
		a int
		b string
	}

	type Test struct {
		name string
		a    []MyStruct
		want []MyStruct
	}

	tests := []Test{
		{
			name: `{{1,"a"},{2,"b"},{3,"c"}}`,
			a: []MyStruct{
				{a: 1, b: "a"}, {a: 2, b: "b"}, {a: 3, b: "c"},
			},
			want: []MyStruct{
				{a: 3, b: "c"}, {a: 2, b: "b"}, {a: 1, b: "a"},
			},
		},
		{
			name: `{{5,"a"},{2,"b"},{8,"c"},{9,"d"}}`,
			a: []MyStruct{
				{a: 5, b: "a"}, {a: 2, b: "b"}, {a: 8, b: "c"}, {a: 9, b: "d"},
			},
			want: []MyStruct{
				{a: 9, b: "d"}, {a: 8, b: "c"}, {a: 2, b: "b"}, {a: 5, b: "a"},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			Reverse(test.a)
			if !reflect.DeepEqual(test.a, test.want) {
				t.Errorf("got %v, want %v", test.a, test.want)
			}
		})
	}

}
