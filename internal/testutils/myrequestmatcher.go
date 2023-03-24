package testutils

import "testingmatcher-with-option/internal/model"

type Option interface {
	apply(*myRequestMatcher)
}

type myRequestMatcher struct {
	id       string
	name     string
	datetime string
	value1   int
	value2   int
	value3   int
	value4   int
	value5   int
	value6   int
}

type optionFunc func(*myRequestMatcher)

func (f optionFunc) apply(m *myRequestMatcher) {
	f(m)
}

func WithId(id string) Option {
	return optionFunc(func(m *myRequestMatcher) {
		m.id = id
	})
}

func WithName(name string) Option {
	return optionFunc(func(m *myRequestMatcher) {
		m.name = name
	})
}

func WithDatetime(datetime string) Option {
	return optionFunc(func(m *myRequestMatcher) {
		m.datetime = datetime
	})
}

func WithValue1(value1 int) Option {
	return optionFunc(func(m *myRequestMatcher) {
		m.value1 = value1
	})
}

func WithValue2(value2 int) Option {
	return optionFunc(func(m *myRequestMatcher) {
		m.value2 = value2
	})
}

func WithValue3(value3 int) Option {
	return optionFunc(func(m *myRequestMatcher) {
		m.value3 = value3
	})
}

func WithValue4(value4 int) Option {
	return optionFunc(func(m *myRequestMatcher) {
		m.value4 = value4
	})
}

func WithValue5(value5 int) Option {
	return optionFunc(func(m *myRequestMatcher) {
		m.value5 = value5
	})
}

func WithValue6(value6 int) Option {
	return optionFunc(func(m *myRequestMatcher) {
		m.value6 = value6
	})
}

func (m *myRequestMatcher) Matches(x interface{}) bool {
	if req, ok := x.(model.MyRequest); ok {
		if m.id != "" && req.Id != m.id {
			return false
		}
		if m.name != "" && req.Name != m.name {
			return false
		}
		if m.datetime != "" && req.Datetime != m.datetime {
			return false
		}
		if m.value1 != 0 && req.Value1 != m.value1 {
			return false
		}
		if m.value2 != 0 && req.Value2 != m.value2 {
			return false
		}
		if m.value3 != 0 && req.Value3 != m.value3 {
			return false
		}
		if m.value4 != 0 && req.Value4 != m.value4 {
			return false
		}
		if m.value5 != 0 && req.Value5 != m.value5 {
			return false
		}
		if m.value6 != 0 && req.Value6 != m.value6 {
			return false
		}
		return true
	}
	return false
}

func (m *myRequestMatcher) String() string {
	return "is a MyRequest with custom conditions"
}
