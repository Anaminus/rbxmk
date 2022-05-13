-- Read a model asset URL from the clipboard, then copy the content of the model
-- to the clipboard, ready to be pasted into Studio.
--
-- Usage:
--     rbxmk run copy-model.rbxmk.lua
--
-- Use studio cookies:
--     rbxmk run copy-model.rbxmk.lua true

local useStudioCookies = ...
local url = clipboard.read("txt")
local id = tonumber(url:match("%d+"))
if not id then
	error(string.format("failed to parse asset ID from URL %q", url), 0)
end
local options = {
	AssetId = id,
	Format = "rbxm",
}
if useStudioCookies then
	options.Cookies = Cookie.from("studio")
end
local asset = rbxassetid.read(options)
clipboard.write(asset, "rbxm")
