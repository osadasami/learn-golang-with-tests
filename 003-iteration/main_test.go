package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a", 3)
	want := "aaa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}
