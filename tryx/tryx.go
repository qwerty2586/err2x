package tryx

func To(f func() error, errx error) {
	err := f()
	if err != nil {
		panic(errx)
	}
}

func To1[T any](f func() (T, error), errx error) T {
	t, err := f()
	if err != nil {
		panic(errx)
	}
	return t
}

func To2[T, U any](f func() (T, U, error), errx error) (T, U) {
	t, u, err := f()
	if err != nil {
		panic(errx)
	}
	return t, u
}

func To3[T, U, V any](f func() (T, U, V, error), errx error) (T, U, V) {
	t, u, v, err := f()
	if err != nil {
		panic(errx)
	}
	return t, u, v
}
