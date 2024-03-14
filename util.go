package streams

func drain[T any](ch <-chan T) {
	for range ch {
	}
}
