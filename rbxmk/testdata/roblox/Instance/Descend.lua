local A = Instance.new("A"); A.Name = "A"
local B = Instance.new("B", A); B.Name = "B"
local C = Instance.new("C", B); C.Name = "C"
local D = Instance.new("D", C); D.Name = "D"

T.Fail(function() return A:Descend() == nil end , "no args => error")
T.Pass(A:Descend("Z") == nil                    , "Z == nil")
T.Pass(A:Descend("A") == nil                    , "A == nil")
T.Pass(A:Descend("B") == B                      , "B == B")
T.Pass(A:Descend("B", "C") == C                 , "B, C == C")
T.Pass(A:Descend("B", "C", "D") == D            , "B, C, D == D")
T.Pass(A:Descend("B", "D") == nil               , "B, D == nil")
T.Pass(A:Descend("C") == nil                    , "C == nil")
T.Pass(A:Descend("D") == nil                    , "D == nil")
