package tern

func Q[T interface{}](clause bool, a T, b T) T {
	if clause {
		return a
	} else {
		return b
	}
}
