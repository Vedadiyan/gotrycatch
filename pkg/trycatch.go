package gotrycatch

type trial struct {
	functions []func(arg any) (any, error)
	result    any
}

func Try[TOutput any](fn func() (any, error)) *trial {
	trial := &trial{
		functions: make([]func(arg any) (any, error), 0),
	}
	trial.functions = append(trial.functions, func(arg any) (any, error) {
		return fn()
	})
	return trial
}

func UnWrap[TOutput any](trial *trial) TOutput {
	if out, ok := trial.result.(TOutput); ok {
		return out
	}
	return *new(TOutput)
}

func (trial *trial) Catch() error {
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

func (trial *trial) Then(fn func(arg any) (any, error)) {
	trial.functions = append(trial.functions, func(arg any) (any, error) {
		res, err := fn(arg)
		return res, err
	})
}
