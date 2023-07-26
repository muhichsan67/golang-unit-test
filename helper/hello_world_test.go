package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Benchmark Performance
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Ichsan")
	}
}

// Sub Benchmark
func BenchmarkHelloWordSub(b *testing.B) {
	b.Run("Ichsan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Ichsan")
		}
	})

	b.Run("Fathurrochman", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Fathurrochman")
		}
	})
}

// Table Benchmark
func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{name: "HelloWorld(Muhammad)", request: "Muhammad"},
		{name: "HelloWorld(Ichsan)", request: "Ichsan"},
		{name: "HelloWorld(Fathurrochman)", request: "Fathurrochman"},
	}

	for _, benchmarks := range benchmarks {
		b.Run(benchmarks.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmarks.request)
			}
		})
	}
}

// Menggagalkan UT
// t.Error(args) = Log error + fail()
// t.Fatal(args) = Log error + fail now()

// TestMain = Semua UT didalam package ini di eksekusi
func TestMain(m *testing.M) {
	fmt.Println("Before Unit Test")

	m.Run()

	fmt.Println("After Unit Test")
}

func TestSubTest(t *testing.T) {
	t.Run("Budi", func(t *testing.T) {
		result := HelloWorld("Budi")
		require.Equal(t, "Hello Budi", result, "Result must be Hello Budi")
	})
	t.Run("Joko", func(t *testing.T) {
		result := HelloWorld("Joko")
		require.Equal(t, "Hello Joko", result, "Result must be Hello Joko")
	})
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{name: "HelloWorld(Muhammad)", request: "Muhammad", expected: "Hello Muhammad"},
		{name: "HelloWorld(Ichsan)", request: "Ichsan", expected: "Hello Ichsan"},
		{name: "HelloWorld(Fathurrochman)", request: "Fathurrochman", expected: "Hello Fathurrochman"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, "Result must be equal!")
		})
	}
}

func TestHelloWorld(t *testing.T) {
	nama := "Ichsan"
	result := HelloWorld(nama)

	if result != "Hello "+nama {
		t.Error("Result must be hello ichsan!")
	}
	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldIchsan(t *testing.T) {
	nama := "Ichsan"
	result := HelloWorld(nama)
	if result != "Hello "+nama {
		t.Fatal("Result must be hello ichsan")
	}
	fmt.Println("TestHelloWorldIchsan Done")
}

func TestHelloWorldAssert(t *testing.T) {
	nama := "Ichsan"
	result := HelloWorld(nama)

	assert.Equal(t, "Hello Ichsan", result, "Result must be hello íchsan")
	fmt.Println("Dieksekusi")
}

func TestHelloWorldRequire(t *testing.T) {
	nama := "Ichsan"
	result := HelloWorld(nama)

	require.Equal(t, "Hello Ichsan", result, "Result must be hello íchsan")
	fmt.Println("Dieksekusi")
}

func TestHelloWorldSkip(t *testing.T) {
	nama := "Ichsan"
	if runtime.GOOS == "darwin" {
		t.Skip("Unit test tidak bisa dijalankan di windows")
	}
	result := HelloWorld(nama)

	require.Equal(t, "Hello Ichsan", result, "Result must be hello íchsan")
	fmt.Println("Dieksekusi")
}
