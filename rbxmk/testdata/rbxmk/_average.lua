-- Returns the average of the given arguments.
local args = {...}
local average = 0
for _, v in ipairs(args) do
	average = average + v
end
return average/#args
