# testing-matcher-with-option

Have you ever found yourself in a situation where, while writing tests with Go, you have to pass a lot of information to the mock function?
This code code can help you to create you custom Matchers with one simple instruction. You can just test the attributes that are importat for your test and ignore the rest without having to create a custom Matcher.

## Usage

Creating Custom Matchers
One of the key features of this module is the ability to create custom matchers for gomock that can be used in unit tests. 
The MyMatcher struct and its associated functions provide a flexible and reusable way to create custom matchers that work with any type.

To create a matcher, you can use the CreateMyMatcher function and pass one or more Option arguments. An Option is an interface that defines a method called apply, which takes a pointer to a MyMatcher and modifies it in some way.

For example, the WithValue function is an Option that sets the value of a particular field in the Obj value. It takes three parameters: the obj value to match against, the name of the field to set, and the value to set it to.

You can use the Matches method of the MyMatcher struct to determine whether a given value matches the criteria specified by the matcher. If a field does not match, the WrongField and WrongValue fields of the matcher are set and the method returns false.
##  Example 1
```
mockService.EXPECT().MakeARequest(testutils.CreateMyMatcher(
			testutils.WithValue(model.MyRequest{}, "Name", "John Doe"),
			testutils.WithValue(model.MyRequest{}, "Id", "123"),
		)).Return(true)
```

##  Example 2
```
mockService.EXPECT().MakeARequest(testutils.CreateMyMatcher(
			testutils.WithValue(model.MyRequest{}, "Value1", 1),
			testutils.WithValue(model.MyRequest{}, "Id", "123"),
		)).Return(true)
```
