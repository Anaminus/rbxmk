-- Patch a descriptor file with a patch file.
--
-- Usage:
--     rbxmk run patch input [optionalOutput]

-- For example:
--     rbxmk run dump.desc-patch.json input.desc.json

-- To write to a different output:
--     rbxmk run dump.desc-patch.json input.desc.json output.desc.json
--
local patchFile, inputFile, outputFile = ...
local patch = fs.read(patchFile, "desc-patch.json")
local desc = fs.read(inputFile, "desc.json")
desc:Patch(patch)
outputFile = outputFile or inputFile
fs.write(outputFile, desc, "desc.json")
