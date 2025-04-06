package helpers

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// run:
// go test
// go test -v
// go test -r TestGreeting (run by prefix function name)
// go test ./... (run all inside folders)
func TestGreetingTo(t *testing.T) {
	result := GreetingTo("Angga")
	if result != "Hello Angga" {
		// unit testing failed
		panic("Result is not 'Hello Angga'") // will stop all unit test after
	}
}

func TestGreetingToEmpty(t *testing.T) {
	result := GreetingTo("")
	if result != "Hello there" {
		t.Fail() // mark testing as fail but continue the testing / assertion
		//t.FailNow() // mark testing as fail but stop the testing code / assertion
	}
	fmt.Println("end of TestGreetingToEmpty")
}

func TestGreetingToEmptyWithErrorMessage(t *testing.T) {
	result := GreetingTo("")
	if result != "Hello there" {
		// similar to Fail and FailNow but with message
		t.Error("Result is not 'Hello there'") // mark testing as fail but continue the testing / assertion
		//t.Fatal("Result is not 'Hello there'") // mark testing as fail but stop the testing code / assertion
	}
	fmt.Println("end of TestGreetingToEmptyWithErrorMessage")
}

func TestStringToSlug(t *testing.T) {
	slug := StringToSlug([]string{"Angga aRi", "wijaya"}, "kusuMa")
	assert.Equal(t, "angga-ari-wijaya-kusuma", slug, "Result is not angga-ari-wijaya-kusuma")
}

func TestSearchStringFound(t *testing.T) {
	assert.True(t, StringContains("Angga Ari", "ari"));
}

func TestSearchStringNotFound(t *testing.T) {
	require.False(t, StringContains("Angga Ari", "wijaya"));
	fmt.Println("end of TestSearchStringNotFound") // require will stop the testing when it fail (Fatal)
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Unit test skip in Mac OS")
	}
	result := GreetingTo("Angga")
	require.Equal(t, "Hello Angga", result)
}

// Should name "TestMain" to configure testing lifecycle
// Run once per package not per function of per test case
func TestMain(m *testing.M) {
	fmt.Println("Before unit test")

	m.Run()

	fmt.Println("After unit test")
}

func TestSubTest(t *testing.T) {
	// go test -run TestSubTest/Angga
	t.Run("Angga", func(t *testing.T) {
		result := GreetingTo("Angga")
		assert.Equal(t, "Hello Angga", result)
	})

	// go test -run TestSubTest/Ari
	t.Run("Ari", func(t *testing.T) {
		result := GreetingTo("Ari")
		assert.Equal(t, "Hello Ari", result)
	})
}

// testing by set of data
func TestGenerateSlugTable(t *testing.T) {
	tests := []struct {
		name string
		title string
		expected string
	} {
		{
			name: "slug:angga Ari",
			title: "angga Ari",
			expected: "angga-ari",
		},
		{
			name: "slug:New Content 15",
			title: "New Content 15",
			expected: "new-content-15",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, StringToSlug(test.title))
		})
	}
}

// go test -v -bench=. (running all test with benchmark)
// go test -v -bench=BenchmarkSlugGenerator (running all test with spcific benchmark)
// go test -v -run=NoMatchUnitTest -bench=. (running only benchmark, trick go by passing not found unit test)
// go test -v -bench=. ./...
func BenchmarkSlugGenerator(b *testing.B) {
	for b.Loop() {
		StringToSlug("Angga", "Ari", "Wijaya")
	}
}
func BenchmarkSlugGeneratorSingleString(b *testing.B) {
	for b.Loop() {
		StringToSlug("Angga Ari")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("This is long title single", func(b *testing.B) {
		for b.Loop() {
			StringToSlug("This is long title")
		}
	})
	b.Run("This is long title args", func(b *testing.B) {
		for b.Loop() {
			StringToSlug("This", "is", "long", "title")
		}
	})
}

func BenchmarkGenerateSlugTable(b *testing.B) {
	tests := []struct {
		name string
		title string
	} {
		{
			name: "slug:angga Ari",
			title: "angga Ari",
		},
		{
			name: "slug:New Content 15",
			title: "New Content 15",
		},
	}

	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for b.Loop() {
				StringToSlug(test.title)
			}
		})
	}
}