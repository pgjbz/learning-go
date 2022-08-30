package main

import (
	"reflect"
	"testing"
	"time"
)

func TestCheckWebsites(t *testing.T) {
	websites := []string {
		"http://foo.com",
		"https://google.com",
		"https://github.com",
	}
	want := map[string]bool {
		"http://foo.com": false,
		"https://google.com": true,
		"https://github.com": true,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v got %v", want, got)
	}
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

func mockWebsiteChecker(url string) bool {
	if url == "http://foo.com" {
		return false
	}
	return true
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}