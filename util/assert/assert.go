package assert

type Opts struct {
	Msg string
}

type option func(*Opts)

func Msg(m string) option {
	return func(o *Opts) {
		o.Msg = m
	}
}

func Do(b bool, opts ...option) {
	a := &Opts{
		Msg: "assertion failed",
	}

	for _, opt := range opts {
		opt(a)
	}

	if !b {
		panic(a.Msg)
	}
}
