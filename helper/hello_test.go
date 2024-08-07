package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	m.Run()
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < 1000000; i++ {
		HelloWorld("Jonathan")
	}
}

func TestHello(t *testing.T) {
	result := HelloWorld("Jonathan")
	if result != "Hello Jonathan" {
		// Panic ini jelek karena menyelesaikan program
		// panic("Result is not Hello Jonathan")
		t.Error("Seharusnya Hello Jonathan")
	}
}

func TestHelloUsingAssert(t *testing.T) {
	result := HelloWorld("Nyewnyew")
	assert.Equal(t, "Hello Nyewnyew", result, "Result must be Hello Jonathan")
}
