-- Update stored copy of descriptors. Include dump.desc-patch.json.
--
-- Usage:
--     rbxmk run --desc-latest update-desc.rbxmk.lua OUTPUT

assert(rbxmk.globalDesc, "must run with --desc-latest option")
local patch = fs.read(path.expand("$sd/dump-patch/dump.desc-patch.json"))
rbxmk.globalDesc:Patch(patch)

local output = ...
if type(output) ~= "string" then
	print(rbxmk.encodeFormat("desc.json", rbxmk.globalDesc))
	return
end

fs.write(output, rbxmk.globalDesc, "desc.json")
