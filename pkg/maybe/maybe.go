package maybe

type Maybe[T any] struct {
	exists bool
	item   T
}

func ApplyFunctor[T1, T2, TResult any](application Maybe[func(T1, T2) TResult], first Maybe[T1]) Maybe[func(T2) TResult] {
	return Some(func(second T2) TResult {
		return application.item(first.item, second)
	})
}

func Apply[T2, TResult any](application Maybe[func(T2) TResult], second Maybe[T2]) Maybe[TResult] {
	return Some(application.item(second.item))
}

func Some[T any](item T) Maybe[T] {
	return Maybe[T]{
		exists: true,
		item:   item,
	}
}
