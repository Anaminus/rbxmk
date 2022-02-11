package main

import (
	"strings"
)

// repeatedString is a string flag that can be specified multiple times.
type repeatedString []string

func (s repeatedString) String() string {
	return strings.Join(s, ",")
}

func (s *repeatedString) Set(v string) error {
	*s = append(*s, v)
	return nil
}

type funcFlag func(string) error

func (f funcFlag) Set(s string) error { return f(s) }
func (f funcFlag) String() string     { return "" }
func (f funcFlag) Type() string       { return "function" } // Purpose of this method unclear.
