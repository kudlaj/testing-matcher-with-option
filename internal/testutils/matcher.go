package testutils

import "github.com/golang/mock/gomock"

func CreateMyRequestMatcher(options ...Option) gomock.Matcher {
	matcher := &myRequestMatcher{}
	for _, opt := range options {
		// we use the apply method to set the values that we want to matchpa
		opt.apply(matcher)
	}
	return matcher
}
