package core

type Script struct {
	Src    string
	Defer  bool
	Async  bool
	Module bool
}

type ScriptOption func(*Script)

func (h *Html) AddScript(src string, opts ...ScriptOption) {
	s := Script{Src: src}

	for _, opt := range opts {
		opt(&s)
	}

	h.Head.Scripts = append(h.Head.Scripts, s)
}

func WithDefer() ScriptOption {
	return func(s *Script) {
		s.Defer = true
	}
}

func WithAsync() ScriptOption {
	return func(s *Script) {
		s.Async = true
	}
}

func WithModule() ScriptOption {
	return func(s *Script) {
		s.Module = true
	}
}
