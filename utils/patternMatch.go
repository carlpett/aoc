package utils

// Adapted from github.com/jd78/gopatternmatching

type match struct {
	input     string
	isMatched bool
}

type matchResult struct {
	input     string
	output    string
	isMatched bool
}

// Match condition and run an action that does not return a type
func Match(input string) match {
	return match{input, false}
}

// when param is "func(input string) bool" used to match a condition
// a param is "func()" used to run an action
func (m match) When(f func(input string) bool, a func()) match {
	if m.isMatched {
		return m
	}

	if f(m.input) {
		a()
		m.isMatched = true
	}

	return m
}

// val param is "string" used to exact match the input with the passed condition
// a param is "func()" used to run an action
func (m match) WhenValue(val string, a func()) match {
	if m.isMatched {
		return m
	}

	if m.input == val {
		a()
		m.isMatched = true
	}

	return m
}

// OtherwiseThrow throws if the pattern is not matched, optionally used at the end of the pattern matching
func (m match) OtherwiseThrow() {
	if !m.isMatched {
		panic("pattern not matched")
	}
}

// ResultMatch matches conditions and run a function that returns a type
func ResultMatch(input string) matchResult {
	return matchResult{input, "", false}
}

// when param is "func(input string) bool" used to match a condition
// a param is "func() string" used to run an action
func (m matchResult) When(f func(input string) bool, a func() string) matchResult {
	if m.isMatched {
		return m
	}

	if f(m.input) {
		m.output = a()
		m.isMatched = true
	}

	return m
}

// val param is "string" used to exact match the input with the passed condition
// a param is "func() string" used to run an action
func (m matchResult) WhenValue(val string, a func() string) matchResult {
	if m.isMatched {
		return m
	}

	if m.input == val {
		m.output = a()
		m.isMatched = true
	}

	return m
}

// Get the result from the pattern or throws if not matched
func (m matchResult) Result() string {
	if m.output == "" {
		panic("pattern not matched")
	}
	return m.output
}

// Get the result from the pattern or the passed default value
func (m matchResult) ResultOrDefault(def string) string {
	if m.output == "" {
		return def
	}
	return m.output
}
