package kmp

import (
	"reflect"
	"testing"
)

func TestKmpSearch(t *testing.T) {
	type test struct {
		input   string
		pattern string
		want    []int
	}
	tests := map[string]test{
		"simple":    {input: "what is mind? no matter; what is matter? never mind.", pattern: "mind", want: []int{8, 47}},
		"short str": {input: "abaacababcac", pattern: "ababc", want: []int{5}},
		"long str":  {input: "bacbababadababacambabacaddababacasdsd", pattern: "ababaca", want: []int{10, 26}},
		"chinese":   {input: "黑化肥发灰会挥发,灰化肥挥发会发黑", pattern: "化肥", want: []int{3, 28}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := KmpSearch(tc.input, tc.pattern)
			if !reflect.DeepEqual(got, tc.want) {
				t.Error(tc.want, got)
			}
		})
	}
}

func BenchmarkKmpSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		KmpSearch("bacbababadababacambabacaddababacasdsd", "ababaca")
	}
}
