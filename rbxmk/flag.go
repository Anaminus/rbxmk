package main

type funcFlag func(string) error

func (f funcFlag) Set(s string) error { return f(s) }
func (f funcFlag) String() string     { return "" }
func (f funcFlag) Type() string       { return "function" } // Purpose of this method unclear.
