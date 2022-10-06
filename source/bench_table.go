func benchIteratorSelector(b *testing.B, n int) {
	// ... setup here
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err := myExpensiveFunc()
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkIteratorSelector(b *testing.B) {
	for _, q := range []int{100, 1000, 10000, 100000} {
		b.Run(fmt.Sprintf("n=%d", q), func(b *testing.B) {
			benchIteratorSelector(b, q)
		})
	}
}
