T.Pass(math.clamp(0,2,4) == 2)
T.Pass(math.clamp(1,2,4) == 2)
T.Pass(math.clamp(2,2,4) == 2)
T.Pass(math.clamp(3,2,4) == 3)
T.Pass(math.clamp(4,2,4) == 4)
T.Pass(math.clamp(5,2,4) == 4)
T.Pass(math.clamp(6,2,4) == 4)
T.Fail(function() math.clamp(0,4,2) end)

T.Pass(math.log(100)/math.log(10) == 2)
T.Pass(math.log(100, 10) == 2)

T.Pass(math.round(10.75) == 11)
T.Pass(math.round(10.5) == 11)
T.Pass(math.round(10.25) == 10)
T.Pass(math.round(-10.25) == -10)
T.Pass(math.round(-10.5) == -11)
T.Pass(math.round(-10.75) == -11)

T.Pass(math.sign(math.pi) == 1)
T.Pass(math.sign(0) == 0)
T.Pass(math.sign(-math.pi) == -1)