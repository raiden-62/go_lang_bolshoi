package benchmarks

import (
	"Go_lang/internal/pkg/storage"
	"strconv"
	"testing"
)

func BenchmarkGet(b *testing.B) {

	s, err := storage.NewStorage()
	if err != nil {
		b.Errorf("new storage: %v", err)
	}
	for i := 0; i < b.N; i++ {
		s.Set(strconv.Itoa(i), strconv.Itoa(i))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Get(strconv.Itoa(i))
	}
}
