package algorithm

import (
	"log"
	"reflect"
	"testing"
)

func Test_LRU(t *testing.T) {
	testCase := []struct {
		describe string
		input    []int
		expect   []int
	}{
		{
			describe: "Test Case 1",
			input:    []int{1, 2, 3, 4, 5},
			expect:   []int{5, 4, 3, 2, 1},
		},
		{
			describe: "Test Case 2",
			input:    []int{1, 2, 3, 4, 5, 1, 1},
			expect:   []int{1, 5, 4, 3, 2},
		},
		{
			describe: "Test Case 1",
			input:    []int{1, 2, 3, 4, 5, 100, 5},
			expect:   []int{5, 100, 4, 3, 2},
		},
	}

	for _, tc := range testCase {
		lru := LRUCache{}
		for _, v := range tc.input {
			lru.refer(v)
		}
		if !reflect.DeepEqual(lru.q, tc.expect) {
			log.Fatalf(
				"%v: expect:%v != result:%v",
				tc.describe, tc.expect, lru.q,
			)
		}
	}
}
