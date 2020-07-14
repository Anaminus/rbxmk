-- Test arguments
assert(select(1, ...) == true, "arg 1 must be true")
assert(select(2, ...) == false, "arg 2 must be false")
assert(select(3, ...) == nil, "arg 3 must be nil")
assert(select(4, ...) == 42, "arg 4 must be 42")
assert(select(5, ...) == 3.141592653589793, "arg 5 must be pi")
assert(select(6, ...) == -1e-8, "arg 6 must be -1e-8")
assert(select(7, ...) == "hello, world!", "arg 7 must be string")
assert(select(8, ...) == "hello\0world!", "arg 8 must be string with NUL")
