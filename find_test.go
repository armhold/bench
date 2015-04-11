package bench

import "testing"

// path is a text file path.
const path = "./data.txt"

// word is a word to find.
const word = "aa"

// want is an expected number.
const want = "1:0,1:10,6:0,6:1"

func TestFind(t *testing.T) {
	got, err := Find(path, word)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("Find(%q, %q) => %q, want %q", path, word, got, want)
	}
}

func TestFind_emptyWord(t *testing.T) {
	if _, err := Find(path, ""); err == nil {
		t.Error("some kind of error should be returned")
	}
}

func TestFind_emptyResult(t *testing.T) {
	got, err := Find(path, "not_exist_word")
	if err != nil {
		t.Fatal(err)
	}
	if got != "" {
		t.Error("empty value should be returned")
	}
}

func Test_kmpBuildTable_ABCDABD(t *testing.T) {
	W := "ABCDABD"
	T := kmpBuildTable(W)

	expected := []int{-1, 0, 0, 0, 0, 1, 2}

	if len(expected) != len(T) {
		t.Errorf("kmpBuildTable(%q) => %q, want %q", W, T, expected)
	}

	for i, v := range T {
		if v != expected[i] {
			t.Errorf("kmpBuildTable(%q) => %v, want %v", W, T, expected)
			break
		}
	}
}

func Test_kmpBuildTable_PARTICIPATE(t *testing.T) {
	W := "PARTICIPATE IN PARACHUTE"
	T := kmpBuildTable(W)

	expected := []int{-1, 0, 0, 0, 0, 0, 0, 0, 1, 2, 0, 0, 0, 0, 0, 0, 1, 2, 3, 0, 0, 0, 0, 0}

	if len(expected) != len(T) {
		t.Errorf("kmpBuildTable(%q) => %q, want %q", W, T, expected)
	}

	for i, v := range T {
		if v != expected[i] {
			t.Errorf("kmpBuildTable(%q) => %v, want %v", W, T, expected)
			break
		}
	}
}

func BenchmarkFind(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Find(path, word)
	}
}
