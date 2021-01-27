T.Fail(function() Instance.new() end                                , "expects first argument")
T.Pass(function() Instance.new("Foobar") end                        , "first argument is a string")
T.Pass(function() Instance.new("Foobar", Instance.new("Parent")) end, "second argument is an optional Instance")
T.Pass(function() Instance.new("Foobar", DataModel.new()) end       , "second argument can be a DataModel")
