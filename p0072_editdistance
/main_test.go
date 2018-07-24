// 85. Maximal Rectangle
// https://leetcode.com/problems/maximal-rectangle/

package main

import (
	"reflect"
	"testing"
)

type TestCase struct {
	word1, word2 string
	output       int
}

func TestMinDistance(t *testing.T) {
	cases := []TestCase{
		{"", "", 0},
		{"a", "", 1},
		{"", "a", 1},
		{"a", "a", 0},
		{"a", "b", 1},
		{"ab", "b", 1},
		{"ab", "c", 2},
		{"aba", "b", 2},
		{"aa", "b", 2},
		{"a", "ab", 1},
		{"a", "ba", 1},
		{"abba", "ab", 2},
		{"abba", "bb", 2},
		{"abba", "ba", 2},
		{"ab", "ac", 1},
		{"ab", "cb", 1},
		{"ab", "bc", 2},
		{"ab", "cd", 2},
		{"bab", "cab", 1},
		{"abab", "cab", 2},
		{"baba", "cab", 2},
		{"baba", "abab", 2},
		{"ab", "cc", 2},
		{"abcd", "bcda", 2},
		{"bcda", "abcd", 2},
		{"horse", "or", 3},
		{"horse", "ros", 3},
		{"intention", "execution", 5},
	}
	for _, c := range cases {
		if output := minDistance(c.word1, c.word2); !reflect.DeepEqual(output, c.output) {
			t.Errorf("minDistance(\"%s\", \"%s\") == %v != %v\n", c.word1, c.word2, output, c.output)
		}
	}
}

func BenchmarkMinDistance(b *testing.B) {
	minDistance("abracadabra", "creature")
	minDistance("moreefficientmethodwouldneverrepeatthesamedistancecalculation", "distanceofallpossibleprefixesmightbestoredinanarray")
	minDistance("moreefficientmethodwouldneverrepeatthesamedistancecalculationmoreefficientmethodwouldneverrepeatthesamedistancecalculationmoreefficientmethodwouldneverrepeatthesamedistancecalculationmoreefficientmethodwouldneverrepeatthesamedistancecalculationmoreefficientmethodwouldneverrepeatthesamedistancecalculationmoreefficientmethodwouldneverrepeatthesamedistancecalculationmoreefficientmethodwouldneverrepeatthesamedistancecalculationmoreefficientmethodwouldneverrepeatthesamedistancecalculationmoreefficientmethodwouldneverrepeatthesamedistancecalculationmoreefficientmethodwouldneverrepeatthesamedistancecalculationmoreefficientmethodwouldneverrepeatthesamedistancecalculationmoreefficientmethodwouldneverrepeatthesamedistancecalculationmoreefficientmethodwouldneverrepeatthesamedistancecalculationmoreefficientmethodwouldneverrepeatthesamedistancecalculationmoreefficientmethodwouldneverrepeatthesamedistancecalculationmoreefficientmethodwouldneverrepeatthesamedistancecalculationmoreefficientmethodwouldneverrepeatthesamedistancecalculationmoreefficientmethodwouldneverrepeatthesamedistancecalculationmoreefficientmethodwouldneverrepeatthesamedistancecalculationmoreefficientmethodwouldneverrepeatthesamedistancecalculationmoreefficientmethodwouldneverrepeatthesamedistancecalculationmoreefficientmethodwouldneverrepeatthesamedistancecalculationmoreefficientmethodwouldneverrepeatthesamedistancecalculationmoreefficientmethodwouldneverrepeatthesamedistancecalculation", "distanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarraydistanceofallpossibleprefixesmightbestoredinanarray")
}
