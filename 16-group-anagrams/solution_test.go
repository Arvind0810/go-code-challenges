package groupanagrams

import (
	"reflect"
	"slices"
	"testing"
)

func normalize(groups [][]string) [][]string {
	out := make([][]string, 0, len(groups))
	for _, group := range groups {
		sorted := slices.Clone(group)
		slices.Sort(sorted)
		out = append(out, sorted)
	}
	slices.SortFunc(out, slices.Compare)
	return out
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: "classic example",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{name: "single empty string", strs: []string{""}, want: [][]string{{""}}},
		{name: "single word", strs: []string{"a"}, want: [][]string{{"a"}}},
		{name: "no input", strs: []string{}, want: nil},
		{name: "no anagrams", strs: []string{"abc", "def", "ghi"}, want: [][]string{{"abc"}, {"def"}, {"ghi"}}},
		{name: "duplicates stay together", strs: []string{"ab", "ba", "ab"}, want: [][]string{{"ab", "ab", "ba"}}},
		{name: "same letters different counts", strs: []string{"aab", "abb", "bba", "baa"}, want: [][]string{{"aab", "baa"}, {"abb", "bba"}}},
		{name: "empty strings group", strs: []string{"", "b", ""}, want: [][]string{{"", ""}, {"b"}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := normalize(GroupAnagrams(slices.Clone(tc.strs)))
			want := normalize(tc.want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("GroupAnagrams(%q) = %q, want %q", tc.strs, got, want)
			}
		})
	}
}
