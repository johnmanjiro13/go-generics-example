package main

func Map[T any, U any](s []T, f func(T) U) []U {
	dst := make([]U, len(s), len(s))
	for i, n := range s {
		dst[i] = f(n)
	}
	return dst
}
