package util

type AssertOpts struct {
	Msg string
}

type option func(*AssertOpts)

func AssertMsg(m string) option {
	return func(o *AssertOpts) {
		o.Msg = m
	}
}

func Assert(b bool, opts ...option) {
	a := &AssertOpts{
		Msg: "assertion failed",
	}

	for _, opt := range opts {
		opt(a)
	}

	if !b {
		panic(a.Msg)
	}
}
