package gostepper

type Stepper[TOutput any] struct {
	functions []func(arg any) (any, error)
	result    any
}

func NewStepper[TOutput any]() *Stepper[TOutput] {
	stepper := &Stepper[TOutput]{
		functions: make([]func(arg any) (any, error), 0),
	}
	return stepper
}

func (stepper Stepper[TOutput]) UnWrap() TOutput {
	if out, ok := stepper.result.(TOutput); ok {
		return out
	}
	return *new(TOutput)
}

func (stepper *Stepper[TOutput]) Exec() error {
	var _arg any
	for _, value := range stepper.functions {
		res, err := value(_arg)
		if err != nil {
			return err
		}
		_arg = res
	}
	stepper.result = _arg
	return nil
}

func Step[TInput any, TOutput any](stepper *Stepper[TOutput], fn func(arg TInput) (any, error)) {
	stepper.functions = append(stepper.functions, func(arg any) (any, error) {
		if arg == nil {
			_arg := new(TInput)
			res, err := fn(*_arg)
			return res, err
		}
		res, err := fn(arg.(TInput))
		return res, err
	})
}
