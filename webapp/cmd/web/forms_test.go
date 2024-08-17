package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Get(t *testing.T) {}

func TestForm_Has(t *testing.T) {
	form := NewForm(nil)
	has := form.Has("whatever")
	if has {
		t.Error("form shows has field when it should not")
	}

	postedData := url.Values{}
	postedData.Add("key", "value")
	form = NewForm(postedData)
	has = form.Has("key")
	if !has {
		t.Error("shows form does not have field when it should")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := NewForm(r.PostForm)

	// 請求內必須要有 a b c
	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields are missing")
	}

	postedData := url.Values{}
	// 請求參數有 a b c
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r, _ = http.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData

	form = NewForm(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows post does not have required")
	}
}

func TestForm_Check(t *testing.T) {
	form := NewForm(nil)
	form.Check(false, "password", "password is required")
	if form.Valid() {
		t.Error("Valid() return false, and it should be true when calling Check()")
	}
}

func TestForm_ErrorGet(t *testing.T) {
	form := NewForm(nil)
	form.Check(false, "password", "password is required")
	s := form.Errors.Get("password")
	if len(s) == 0 {
		t.Error("should have an error returned, but do not")
	}

	s = form.Errors.Get("whatever")
	if len(s) != 0 {
		t.Error("should not have an error, but got one")
	}
}
