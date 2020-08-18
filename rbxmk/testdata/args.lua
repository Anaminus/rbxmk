-- Test arguments
T.Pass(select(1, ...) == true             , "arg 1 is true")
T.Pass(select(2, ...) == false            , "arg 2 is false")
T.Pass(select(3, ...) == nil              , "arg 3 is nil")
T.Pass(select(4, ...) == 42               , "arg 4 is 42")
T.Pass(select(5, ...) == 3.141592653589793, "arg 5 is pi")
T.Pass(select(6, ...) == -1e-8            , "arg 6 is -1e-8")
T.Pass(select(7, ...) == "hello, world!"  , "arg 7 is string")
T.Pass(select(8, ...) == "hello\0world!"  , "arg 8 is string with NUL")
