package core

type Script struct {
	Src     string
	Content string
	Defer   bool
	Async   bool
	Module  bool
}

type ScriptOption func(*Script)

func (h *Html) AddScript(src string, opts ...ScriptOption) {
	s := Script{Src: src}

	for _, opt := range opts {
		opt(&s)
	}

	h.Scripts = append(h.Scripts, s)
}

func (h *Html) RawScript(content string, opts ...ScriptOption) {
	s := Script{Content: content}

	for _, opt := range opts {
		opt(&s)
	}

	h.Scripts = append(h.Scripts, s)
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

func (s Script) Render() string {
	if s.Src != "" {
		result := "<script src=\"" + s.Src + "\""
		if s.Defer {
			result += " defer"
		}
		if s.Async {
			result += " async"
		}
		if s.Module {
			result += " type=\"module\""
		}
		result += "></script>"
		return result
	}

	return "<script>" + s.Content + "</script>"
}
