-- Test arguments
T.Pass("arg 1 is true",
	select(1, ...) == true)
T.Pass("arg 2 is false",
	select(2, ...) == false)
T.Pass("arg 3 is nil",
	select(3, ...) == nil)
T.Pass("arg 4 is 42",
	select(4, ...) == 42)
T.Pass("arg 5 is pi",
	select(5, ...) == 3.141592653589793)
T.Pass("arg 6 is -1e-8",
	select(6, ...) == -1e-8)
T.Pass("arg 7 is string",
	select(7, ...) == "hello, world!")
T.Pass("arg 8 is string with NUL",
	select(8, ...) == "hello\0world!")
