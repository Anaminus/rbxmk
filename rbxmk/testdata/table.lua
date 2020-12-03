local t = {1,2,3}
table.clear(t)
T.Pass(#t == 0, "table.clear clears table")

local t = table.create(10, 1)
T.Pass(#t == 10, "table create creates table with 10 elements")
for i = 1, #t do
	T.Pass(t[i] == 1, "value at index "..i.." is 1")
end

T.Pass(table.find({"a","b","c","b"}, "b") == 2, "table.find finds b at index 2")
T.Pass(table.find({"a","b","c","b"}, "b", 3) == 4, "table.find finds b at index 4")
T.Pass(table.find({"a","b","c"}, "d") == nil, "table.find does not find d")

local t = table.move({1,2,3,4,5}, 2, 4, 3, {4,5,6})
T.Pass(table.concat(t," ") == "4 5 2 3 4")

local t = table.pack(1,2,3,nil,5)
T.Pass(t[1] == 1, "table.pack expects 1 at index 1")
T.Pass(t[2] == 2, "table.pack expects 2 at index 2")
T.Pass(t[3] == 3, "table.pack expects 3 at index 3")
T.Pass(t[4] == nil, "table.pack expects nil at index 4")
T.Pass(t[5] == 5, "table.pack expects 5 at index 5")
T.Pass(t.n == 5)

local a,b,c,d = table.unpack({1,2,3,4,5}, 2,4)
T.Pass(a == 2, "table.unpack expects 2 as value 1")
T.Pass(b == 3, "table.unpack expects 3 as value 2")
T.Pass(c == 4, "table.unpack expects 4 as value 3")
T.Pass(d == nil, "table.unpack expects nil as value 4")
