package trie

import "testing"

type opKind int

const (
	insert opKind = iota
	search
	startsWith
)

type op struct {
	kind opKind
	arg  string
	want bool
}

func (o op) String() string {
	switch o.kind {
	case insert:
		return "Insert(" + o.arg + ")"
	case search:
		return "Search(" + o.arg + ")"
	default:
		return "StartsWith(" + o.arg + ")"
	}
}

func TestTrie(t *testing.T) {
	tests := []struct {
		name string
		ops  []op
	}{
		{
			name: "classic example",
			ops: []op{
				{kind: insert, arg: "apple"},
				{kind: search, arg: "apple", want: true},
				{kind: search, arg: "app", want: false},
				{kind: startsWith, arg: "app", want: true},
				{kind: insert, arg: "app"},
				{kind: search, arg: "app", want: true},
			},
		},
		{
			name: "empty trie",
			ops: []op{
				{kind: search, arg: "a", want: false},
				{kind: startsWith, arg: "a", want: false},
			},
		},
		{
			name: "word longer than stored word",
			ops: []op{
				{kind: insert, arg: "app"},
				{kind: search, arg: "apple", want: false},
				{kind: startsWith, arg: "apple", want: false},
			},
		},
		{
			name: "prefix equal to whole word",
			ops: []op{
				{kind: insert, arg: "go"},
				{kind: startsWith, arg: "go", want: true},
				{kind: search, arg: "go", want: true},
			},
		},
		{
			name: "branching words share prefix",
			ops: []op{
				{kind: insert, arg: "car"},
				{kind: insert, arg: "cat"},
				{kind: insert, arg: "cart"},
				{kind: search, arg: "car", want: true},
				{kind: search, arg: "cat", want: true},
				{kind: search, arg: "cart", want: true},
				{kind: search, arg: "ca", want: false},
				{kind: search, arg: "cars", want: false},
				{kind: startsWith, arg: "ca", want: true},
				{kind: startsWith, arg: "cart", want: true},
				{kind: startsWith, arg: "co", want: false},
			},
		},
		{
			name: "duplicate insert is idempotent",
			ops: []op{
				{kind: insert, arg: "dog"},
				{kind: insert, arg: "dog"},
				{kind: search, arg: "dog", want: true},
				{kind: search, arg: "do", want: false},
			},
		},
		{
			name: "single letter words",
			ops: []op{
				{kind: insert, arg: "a"},
				{kind: search, arg: "a", want: true},
				{kind: search, arg: "b", want: false},
				{kind: startsWith, arg: "a", want: true},
			},
		},
		{
			name: "full alphabet range",
			ops: []op{
				{kind: insert, arg: "abcdefghijklmnopqrstuvwxyz"},
				{kind: insert, arg: "zebra"},
				{kind: search, arg: "abcdefghijklmnopqrstuvwxyz", want: true},
				{kind: startsWith, arg: "abcdefghijklm", want: true},
				{kind: search, arg: "zebra", want: true},
				{kind: startsWith, arg: "zz", want: false},
			},
		},
		{
			name: "new trie does not share state with earlier tries",
			ops: []op{
				{kind: search, arg: "apple", want: false},
				{kind: startsWith, arg: "ap", want: false},
				{kind: search, arg: "zebra", want: false},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tr := NewTrie()
			if tr == nil {
				t.Fatal("NewTrie() returned nil")
			}
			for i, o := range tc.ops {
				var got bool
				switch o.kind {
				case insert:
					tr.Insert(o.arg)
					continue
				case search:
					got = tr.Search(o.arg)
				case startsWith:
					got = tr.StartsWith(o.arg)
				}
				if got != o.want {
					t.Errorf("op #%d %v = %v, want %v", i, o, got, o.want)
				}
			}
		})
	}
}
