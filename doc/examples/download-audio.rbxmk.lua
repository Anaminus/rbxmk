-- Download list of audio assets. Input is a file containing a list of asset
-- IDs, one per line. Output is a path to a directory. Assumes mp3 format.
--
-- Usage:
--     rbxmk run download-audio.rbxmk.lua INPUT OUTPUT
--
-- Example:
--     rbxmk run download-audio.rbxmk.lua assetList.txt assets

local assetIds, outputDirectory = ...
fs.mkdir(outputDirectory, true)

local list = fs.read(assetIds, "bin")
for id in string.gmatch(list, "[^\r\n]+") do
	id = tonumber(id)
	if id then
		local ok, data = pcall(rbxassetid.read, {
			AssetId = id,
			Format = "bin",
		})
		if not ok then
			print(string.format("download %10d: %s", id, data))
		else
			-- WARNING: May not actually be mp3.
			local file = path.join(outputDirectory, id .. ".mp3")
			local ok, err = pcall(fs.write, file, data, "bin")
			if not ok then
				print(string.format("write    %10d: %s", id, err))
			else
				print(string.format("ok       %10d", id))
			end
		end
	end
end
