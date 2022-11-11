package gotrycatch

type Trial[TOutput any] struct {
	functions []func(arg any) (any, error)
	result    any
}

func New[TOutput any]() *Trial[TOutput] {
	trial := &Trial[TOutput]{
		functions: make([]func(arg any) (any, error), 0),
	}
	return trial
}

func (trial Trial[TOutput]) UnWrap() TOutput {
	if out, ok := trial.result.(TOutput); ok {
		return out
	}
	return *new(TOutput)
}

func (trial *Trial[TOutput]) Catch() error {
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

func Try[TInput any, TOutput any](trial *Trial[TOutput], fn func(arg TInput) (any, error)) {
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
