package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	k            int
	head, output *ListNode
}

func TestReverseKGroup(t *testing.T) {
	type N = ListNode
	cases := []TestCase{
		{0, nil, nil},
		{1, nil, nil},
		{0, &N{1, nil}, &N{1, nil}},
		{1, &N{1, nil}, &N{1, nil}},
		{2, &N{1, nil}, &N{1, nil}},
		{0, &N{1, &N{2, nil}}, &N{1, &N{2, nil}}},
		{1, &N{1, &N{2, nil}}, &N{1, &N{2, nil}}},
		{2, &N{1, &N{2, nil}}, &N{2, &N{1, nil}}},
		{3, &N{1, &N{2, nil}}, &N{1, &N{2, nil}}},
		{2,
			&N{1, &N{2, &N{3, &N{4, &N{5, nil}}}}},
			&N{2, &N{1, &N{4, &N{3, &N{5, nil}}}}},
		},
		{3,
			&N{1, &N{2, &N{3, &N{4, &N{5, nil}}}}},
			&N{3, &N{2, &N{1, &N{4, &N{5, nil}}}}},
		},
		{4,
			&N{1, &N{2, &N{3, &N{4, &N{5, nil}}}}},
			&N{4, &N{3, &N{2, &N{1, &N{5, nil}}}}},
		},
		{5,
			&N{1, &N{2, &N{3, &N{4, &N{5, nil}}}}},
			&N{5, &N{4, &N{3, &N{2, &N{1, nil}}}}},
		},
		{6,
			&N{1, &N{2, &N{3, &N{4, &N{5, nil}}}}},
			&N{1, &N{2, &N{3, &N{4, &N{5, nil}}}}},
		},
	}
	for _, c := range cases {
		if output := reverseKGroup(c.head, c.k); !reflect.DeepEqual(output, c.output) {
			t.Errorf("reverseKGroup(%s, %d) == %v != %v\n", c.head, c.k, output, c.output)
		}
	}
}

func BenchmarkReverseKGroup(b *testing.B) {
	root := &N{Val: 0}
	for i, n := 1, root; i < 1000000; i++ {
		n.Next = &N{Val: i}
	}

	b.ResetTimer()
	reverseKGroup(root, 100)
}
