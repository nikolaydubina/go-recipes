func TestPoolContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	defer leaktest.CheckContext(ctx, t)()

	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()
}