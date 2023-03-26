package testutils

import (
	"fmt"
	"reflect"

	"github.com/golang/mock/gomock"
)

type Option[T any] interface {
	apply(*MyMatcher[T])
}

type MyMatcher[T any] struct {
	Id       string
	Name     string
	Datetime string
	Value1   int
	Value2   int
	Value3   int
	Value4   int
	Value5   int
	Value6   int
}

type optionFunc[T any] func(*MyMatcher[T])

func (f optionFunc[T]) apply(m *MyMatcher[T]) {
	f(m)
}

func WithValue[T any](obj T, attributeName string, value any) Option[T] {
	return optionFunc[T](func(m *MyMatcher[T]) {
		SetField(m, attributeName, value)
	})
}

func (m *MyMatcher[T]) Matches(x interface{}) bool {
	if req, ok := x.(T); ok {
		v := reflect.ValueOf(*m)
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			if !field.IsZero() && field.Interface() != reflect.ValueOf(req).Field(i).Interface() {
				return false
			}
		}
		return true
	}
	return false
}

func (m *MyMatcher[T]) String() string {
	return "is a MyRequest with custom conditions"
}

func SetField(obj interface{}, fieldName string, value interface{}) error {
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() != reflect.Ptr {
		return fmt.Errorf("obj must be a pointer")
	}

	objValue = objValue.Elem()
	if objValue.Kind() != reflect.Struct {
		return fmt.Errorf("obj must be a pointer to a struct")
	}

	field := objValue.FieldByName(fieldName)
	if !field.IsValid() {
		return fmt.Errorf("field %s not found", fieldName)
	}

	if !field.CanSet() {
		return fmt.Errorf("field %s cannot be set", fieldName)
	}

	valueValue := reflect.ValueOf(value)
	if field.Type() != valueValue.Type() {
		return fmt.Errorf("field %s has type %s, but received value of type %s", fieldName, field.Type(), valueValue.Type())
	}

	field.Set(valueValue)
	return nil
}

func CreateMyMatcher[T any](obj T, options ...Option[T]) gomock.Matcher {
	matcher := &MyMatcher[T]{}
	for _, opt := range options {
		// we use the apply method to set the values that we want to matchpa
		opt.apply(matcher)
	}
	return matcher
}
