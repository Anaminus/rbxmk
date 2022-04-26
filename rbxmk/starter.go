//go:build !interactive_commands

package main

import "context"

func Starter() func(context.Context) error {
	return Program.ExecuteContext
}
