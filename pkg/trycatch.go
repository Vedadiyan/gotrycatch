package gotrycatch

type trial[TOutput any] struct {
	functions []func(arg any) (any, error)
	result    any
}

func Try[TOutput any](fn func() (any, error)) *trial[TOutput] {
	trial := &trial[TOutput]{
		functions: make([]func(arg any) (any, error), 0),
	}
	trial.functions = append(trial.functions, func(arg any) (any, error) {
		return fn()
	})
	return trial
}

func (trial trial[TOutput]) UnWrap() TOutput {
	if out, ok := trial.result.(TOutput); ok {
		return out
	}
	return *new(TOutput)
}

func (trial *trial[TOutput]) Catch() error {
	var _arg any
	for _, value := range trial.functions {
		res, err := value(_arg)
		if err != nil {
			return err
		}
		_arg = res
	}
	trial.result = _arg
	return nil
}

func Continue[TInput any, TOutput any](trial *trial[TOutput], fn func(arg TInput) (any, error)) {
	trial.functions = append(trial.functions, func(arg any) (any, error) {
		if arg == nil {
			_arg := new(TInput)
			res, err := fn(*_arg)
			return res, err
		}
		res, err := fn(arg.(TInput))
		return res, err
	})
}
