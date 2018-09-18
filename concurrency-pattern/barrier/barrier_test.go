package barrier

import "testing"
import "strings"

func TestBarrier(t *testing.T) {
	t.Run("Correct endpoints", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers",
			"http://httpbin.org/User-Agent"}

		result := captureEndpoints(endpoints...)
		notEncoding := !strings.Contains(result, "Accept-Encoding")
		isAgent := strings.Contains(result, "User-Agent")

		if notEncoding || isAgent {
			t.Fail()
		}

		t.Log(result)
	})

	/* 	t.Run("One endpoint incorrect", func(t *testing.T) {
	   		endpoints := []string{"http://malformed-url", "http://httpbin.org/User-Agent"}
	   	})

	   	t.Run("Very short timeout", func(t *testing.T) {
	   		endpoints := []string{"http://httpbin.org/headers", "http://httpbin.org/User-Agent"}
	   	}) */
}
