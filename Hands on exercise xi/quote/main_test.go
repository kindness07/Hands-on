package quote

import "testing"

func TestSunAlsoTexts(t *testing.T) {
	s := SunAlsoTexts()
	if s != SunAlsoTexts() {
		t.Error("got", s, "expected", "the texts in SunAlso package")
	}
}

func ExampleSunAlsoTexts() {
	s := SunAlsoTexts()
	println(s)
	// Output: the texts in SunAlso package
}

func BenchmarkSunAlsoTexts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SunAlsoTexts()
	}
}
