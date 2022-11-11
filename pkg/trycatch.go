package gotrycatch

import "fmt"

type trial struct {
	fn            func() (any, error)
	thenFunctions []func(arg any) (any, error)
	result        any
	err           error
}

func Try(fn func() (any, error)) *trial {
	trial := &trial{
		thenFunctions: make([]func(arg any) (any, error), 0),
	}
	trial.fn = fn
	return trial
}

func UnWrap[TOutput any](trial *trial) (*TOutput, error) {
	trial.run()
	if trial.err != nil {
		return nil, trial.err
	}
	if out, ok := trial.result.(TOutput); ok {
		return &out, nil
	}
	return nil, fmt.Errorf("invalid cast")
}

func (trial *trial) run() {
	defer func() {
		if recover := recover(); recover != nil {
			trial.err = fmt.Errorf("%v", recover)
		}
	}()
	_arg, err := trial.fn()
	if err != nil {
		trial.err = err
		return
	}
	for _, value := range trial.thenFunctions {
		res, err := value(_arg)
		if err != nil {
			trial.err = err
			return
		}
		_arg = res
	}
	trial.result = _arg
}

func (trial *trial) Then(fn func(arg any) (any, error)) {
	trial.thenFunctions = append(trial.thenFunctions, func(arg any) (any, error) {
		res, err := fn(arg)
		return res, err
	})
}
