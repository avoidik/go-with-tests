package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func DummyWebsiteChecker(url string) bool {
	return (url[:5] == "http:" || url[:6] == "https:")
}

func SlowWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(t *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < 100; i++ {
		urls[i] = "http://url"
	}
	for i := 0; i < t.N; i++ {
		CheckWebsites(SlowWebsiteChecker, urls)
	}
}

func TestCheckWebsites(t *testing.T) {
	testUrls := []string{"brrn://localhost", "http://localhost", "https://localhost"}
	got := CheckWebsites(DummyWebsiteChecker, testUrls)
	want := map[string]bool{
		"brrn://localhost":  false,
		"http://localhost":  true,
		"https://localhost": true,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v got %v", want, got)
	}
}
