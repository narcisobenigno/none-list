package results

type Result struct {
	fails map[string]string
}

func Success() Result {
	return newResult(map[string]string{})
}

func Failed(context, message string) Result {
	return newResult(map[string]string{context: message})
}

func newResult(fails map[string]string) Result {
	return Result{fails: fails}
}

func (r Result) Failed() bool {
	return len(r.fails) > 0
}
