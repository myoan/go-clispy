package main

import "testing"

func TestScanner_EachChar(t *testing.T) {
	var ch byte
	s := NewScanner("hoge")
	ret := s.EachChar()
	if !ret {
		t.Errorf("return false")
		return
	}
	if string(s.ch) != "h" {
		t.Errorf("string not match: expect 'h', but actual: '%s'", string(ch))
		return
	}

	ret = s.EachChar()
	if !ret {
		t.Errorf("return false")
		return
	}
	if string(s.ch) != "o" {
		t.Errorf("string not match: expect 'o', but actual: '%s'", string(ch))
		return
	}

	ret = s.EachChar()
	if !ret {
		t.Errorf("return false")
		return
	}
	if string(s.ch) != "g" {
		t.Errorf("string not match: expect 'g', but actual: '%s'", string(ch))
		return
	}

	ret = s.EachChar()
	if !ret {
		t.Errorf("return false")
		return
	}
	if string(s.ch) != "e" {
		t.Errorf("string not match: expect 'e', but actual: '%s'", string(ch))
		return
	}

	ret = s.EachChar()
	if ret {
		t.Errorf("return true")
		return
	}
}
