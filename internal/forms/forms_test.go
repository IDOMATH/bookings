package forms

import (
	"net/url"
	"testing"
)

func TestForm_Has(t *testing.T) {

	postedData := url.Values{}
	form := New(postedData)

	has := form.Has("whatever")
	if has {
		t.Error("shows form has field when it does not")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error("shows form does not have field when it does")
	}
}

func TestForm_MinLength(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)

	form.MinLength("x", 10)
	if form.Valid() {
		t.Error("form shows min length achieved for non-existent field")
	}

	isError := form.Errors.Get("x")
	if isError == "" {
		t.Error("should have gotten an error, but did not.")
	}
	postedValues = url.Values{}
	postedValues.Add("some_field", "some value")
	form = New(postedValues)

	form.MinLength("some_field", 100)
	if form.Valid() {
		t.Error("form shows min length achieved when data is shorter")
	}

	postedValues = url.Values{}
	postedValues.Add("another_field", "abc123")
	form = New(postedValues)

	form.MinLength("another_field", 1)
	if !form.Valid() {
		t.Error("form shows min length is not met when it is")
	}

	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("should not have gotten an error, but did")
	}

}

func TestForm_IsEmail(t *testing.T) {
	postedValues := url.Values{}
	form := New(postedValues)

	form.IsEmail("x")
	if form.Valid() {
		t.Error("form shows valid email for non-existent field")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "me@here.com")
	form = New(postedValues)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got an invalid email when we should not have.")
	}

	postedValues = url.Values{}
	postedValues.Add("email", "inavlid@email")
	form = New(postedValues)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("got a valid email when we should not have.")
	}
}
