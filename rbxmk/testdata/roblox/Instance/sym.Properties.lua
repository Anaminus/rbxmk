local desc = fs.read(path.expand("$sd/../../dump.desc.json"))
local patch = fs.read(path.expand("$sd/../../dump.desc-patch.json"))
desc:Patch(patch)
local Enum = desc:EnumTypes()

-- Test XML without descriptors.
local d = fs.read(path.expand("$sd/decal.rbxmx")):Descend("Decal")
local p = d[sym.Properties]

T.Pass(typeof(p.AttributesSerialize) == "BinaryString", "xml AttributesSerialize type")
T.Pass(typeof(p.Color3) == "Color3", "xml Color3 type")
T.Pass(typeof(p.Face) == "token", "xml Face type")
T.Pass(typeof(p.Name) == "string", "xml Name type")
T.Pass(typeof(p.SourceAssetId) == "int64", "xml SourceAssetId type")
T.Pass(typeof(p.Tags) == "BinaryString", "xml Tags type")
T.Pass(typeof(p.Texture) == "Content", "xml Texture type")
T.Pass(typeof(p.Transparency) == "float", "xml Transparency type")

T.Pass(p.AttributesSerialize.Value == "", "xml AttributesSerialize value")
T.Pass(p.Color3 == Color3.new(1,1,1), "xml Color3 value")
T.Pass(p.Face.Value == 1, "xml Face value")
T.Pass(p.Name == "Decal", "xml Name value")
T.Pass(p.SourceAssetId.Value == -1, "xml SourceAssetId value")
T.Pass(p.Tags.Value == "", "xml Tags value")
T.Pass(p.Texture.Value == "rbxasset://textures/SpawnLocation.png", "xml Texture value")
T.Pass(p.Transparency.Value == 0, "xml Transparency value")

-- Test XML with descriptors.
d[sym.Desc] = desc
local p = d[sym.Properties]

T.Pass(typeof(p.AttributesSerialize) == "string", "xml desc AttributesSerialize type")
T.Pass(typeof(p.Color3) == "Color3", "xml desc Color3 type")
T.Pass(typeof(p.Face) == "EnumItem", "xml desc Face type")
T.Pass(typeof(p.Name) == "string", "xml desc Name type")
T.Pass(typeof(p.SourceAssetId) == "number", "xml desc SourceAssetId type")
T.Pass(typeof(p.Tags) == "string", "xml desc Tags type")
T.Pass(typeof(p.Texture) == "string", "xml desc Texture type")
T.Pass(typeof(p.Transparency) == "number", "xml desc Transparency type")

T.Pass(p.AttributesSerialize == "", "xml desc AttributesSerialize value")
T.Pass(p.Color3 == Color3.new(1,1,1), "xml desc Color3 value")
T.Pass(p.Face == Enum.NormalId.Top, "xml desc Face value")
T.Pass(p.Name == "Decal", "xml desc Name value")
T.Pass(p.SourceAssetId == -1, "xml desc SourceAssetId value")
T.Pass(p.Tags == "", "xml desc Tags value")
T.Pass(p.Texture == "rbxasset://textures/SpawnLocation.png", "xml desc Texture value")
T.Pass(p.Transparency == 0, "xml desc Transparency value")

-- Test binary without descriptors.
local d = fs.read(path.expand("$sd/decal.rbxm")):Descend("Decal")
local p = d[sym.Properties]

-- Binary encodes string-likes as string type.
T.Pass(typeof(p.AttributesSerialize) == "string", "binary AttributesSerialize type")
T.Pass(typeof(p.Color3) == "Color3", "binary Color3 type")
T.Pass(typeof(p.Face) == "token", "binary Face type")
T.Pass(typeof(p.Name) == "string", "binary Name type")
T.Pass(typeof(p.SourceAssetId) == "int64", "binary SourceAssetId type")
T.Pass(typeof(p.Tags) == "string", "binary Tags type")
T.Pass(typeof(p.Texture) == "string", "binary Texture type")
T.Pass(typeof(p.Transparency) == "float", "binary Transparency type")

T.Pass(p.AttributesSerialize == "", "binary AttributesSerialize value")
T.Pass(p.Color3 == Color3.new(1,1,1), "binary Color3 value")
T.Pass(p.Face.Value == 1, "binary Face value")
T.Pass(p.Name == "Decal", "binary Name value")
T.Pass(p.SourceAssetId.Value == -1, "binary SourceAssetId value")
T.Pass(p.Tags == "", "binary Tags value")
T.Pass(p.Texture == "rbxasset://textures/SpawnLocation.png", "binary Texture value")
T.Pass(p.Transparency.Value == 0, "binary Transparency value")

-- Test binary with descriptors.
d[sym.Desc] = desc
local p = d[sym.Properties]
-- for k,v in pairs(p) do print(k,typeof(v),v) end

T.Pass(typeof(p.AttributesSerialize) == "string", "binary desc AttributesSerialize type")
T.Pass(typeof(p.Color3) == "Color3", "binary desc Color3 type")
T.Pass(typeof(p.Face) == "EnumItem", "binary desc Face type")
T.Pass(typeof(p.Name) == "string", "binary desc Name type")
T.Pass(typeof(p.SourceAssetId) == "number", "binary desc SourceAssetId type")
T.Pass(typeof(p.Tags) == "string", "binary desc Tags type")
T.Pass(typeof(p.Texture) == "string", "binary desc Texture type")
T.Pass(typeof(p.Transparency) == "number", "binary desc Transparency type")

T.Pass(p.AttributesSerialize == "", "binary desc AttributesSerialize value")
T.Pass(p.Color3 == Color3.new(1,1,1), "binary desc Color3 value")
T.Pass(p.Face == Enum.NormalId.Top, "binary desc Face value")
T.Pass(p.Name == "Decal", "binary desc Name value")
T.Pass(p.SourceAssetId == -1, "binary desc SourceAssetId value")
T.Pass(p.Tags == "", "binary desc Tags value")
T.Pass(p.Texture == "rbxasset://textures/SpawnLocation.png", "binary desc Texture value")
T.Pass(p.Transparency == 0, "binary desc Transparency value")
