local desc = fs.read(path.expand("$sd/../../dump.desc.json"))
local patch = fs.read(path.expand("$sd/../../dump.desc-patch.json"))
desc:Patch(patch)
local Enum = desc:EnumTypes()

-- Test XML without descriptors.
local d = fs.read(path.expand("$sd/decal.rbxmx")):Descend("Decal")
local p = d[sym.Properties]

T.Pass(rbxmk.propType(d, "AttributesSerialize") == "BinaryString", "xml instance AttributesSerialize type")
T.Pass(rbxmk.propType(d, "Color3") == "Color3", "xml instance Color3 type")
T.Pass(rbxmk.propType(d, "Face") == "token", "xml instance Face type")
T.Pass(rbxmk.propType(d, "Name") == "string", "xml instance Name type")
T.Pass(rbxmk.propType(d, "SourceAssetId") == "int64", "xml instance SourceAssetId type")
T.Pass(rbxmk.propType(d, "Tags") == "BinaryString", "xml instance Tags type")
T.Pass(rbxmk.propType(d, "Texture") == "Content", "xml instance Texture type")
T.Pass(rbxmk.propType(d, "Transparency") == "float", "xml instance Transparency type")

T.Pass(typeof(p.AttributesSerialize) == "string", "xml table AttributesSerialize type")
T.Pass(typeof(p.Color3) == "Color3", "xml table Color3 type")
T.Pass(typeof(p.Face) == "number", "xml table Face type")
T.Pass(typeof(p.Name) == "string", "xml table Name type")
T.Pass(typeof(p.SourceAssetId) == "number", "xml table SourceAssetId type")
T.Pass(typeof(p.Tags) == "string", "xml table Tags type")
T.Pass(typeof(p.Texture) == "string", "xml table Texture type")
T.Pass(typeof(p.Transparency) == "number", "xml table Transparency type")

T.Pass(p.AttributesSerialize == "", "xml AttributesSerialize value")
T.Pass(p.Color3 == Color3.new(1,1,1), "xml Color3 value")
T.Pass(p.Face == 1, "xml Face value")
T.Pass(p.Name == "Decal", "xml Name value")
T.Pass(p.SourceAssetId == -1, "xml SourceAssetId value")
T.Pass(p.Tags == "", "xml Tags value")
T.Pass(p.Texture == "rbxasset://textures/SpawnLocation.png", "xml Texture value")
T.Pass(p.Transparency == 0, "xml Transparency value")

-- Test XML with descriptors.
d[sym.Desc] = desc
local p = d[sym.Properties]

T.Pass(typeof(p.AttributesSerialize) == "string", "xml desc table AttributesSerialize type")
T.Pass(typeof(p.Color3) == "Color3", "xml desc table Color3 type")
T.Pass(typeof(p.Face) == "EnumItem", "xml desc table Face type")
T.Pass(typeof(p.Name) == "string", "xml desc table Name type")
T.Pass(typeof(p.SourceAssetId) == "number", "xml desc table SourceAssetId type")
T.Pass(typeof(p.Tags) == "string", "xml desc table Tags type")
T.Pass(typeof(p.Texture) == "string", "xml desc table Texture type")
T.Pass(typeof(p.Transparency) == "number", "xml desc table Transparency type")

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
T.Pass(rbxmk.propType(d,"AttributesSerialize") == "string", "binary instance AttributesSerialize type")
T.Pass(rbxmk.propType(d,"Color3") == "Color3", "binary instance Color3 type")
T.Pass(rbxmk.propType(d,"Face") == "token", "binary instance Face type")
T.Pass(rbxmk.propType(d,"Name") == "string", "binary instance Name type")
T.Pass(rbxmk.propType(d,"SourceAssetId") == "int64", "binary instance SourceAssetId type")
T.Pass(rbxmk.propType(d,"Tags") == "string", "binary instance Tags type")
T.Pass(rbxmk.propType(d,"Texture") == "string", "binary instance Texture type")
T.Pass(rbxmk.propType(d,"Transparency") == "float", "binary instance Transparency type")

T.Pass(typeof(p.AttributesSerialize) == "string", "binary table AttributesSerialize type")
T.Pass(typeof(p.Color3) == "Color3", "binary table Color3 type")
T.Pass(typeof(p.Face) == "number", "binary table Face type")
T.Pass(typeof(p.Name) == "string", "binary table Name type")
T.Pass(typeof(p.SourceAssetId) == "number", "binary table SourceAssetId type")
T.Pass(typeof(p.Tags) == "string", "binary table Tags type")
T.Pass(typeof(p.Texture) == "string", "binary table Texture type")
T.Pass(typeof(p.Transparency) == "number", "binary table Transparency type")

T.Pass(p.AttributesSerialize == "", "binary AttributesSerialize value")
T.Pass(p.Color3 == Color3.new(1,1,1), "binary Color3 value")
T.Pass(p.Face == 1, "binary Face value")
T.Pass(p.Name == "Decal", "binary Name value")
T.Pass(p.SourceAssetId == -1, "binary SourceAssetId value")
T.Pass(p.Tags == "", "binary Tags value")
T.Pass(p.Texture == "rbxasset://textures/SpawnLocation.png", "binary Texture value")
T.Pass(p.Transparency == 0, "binary Transparency value")

-- Test binary with descriptors.
d[sym.Desc] = desc
local p = d[sym.Properties]
-- for k,v in pairs(p) do print(k,typeof(v),v) end

T.Pass(typeof(p.AttributesSerialize) == "string", "binary desc table AttributesSerialize type")
T.Pass(typeof(p.Color3) == "Color3", "binary desc table Color3 type")
T.Pass(typeof(p.Face) == "EnumItem", "binary desc table Face type")
T.Pass(typeof(p.Name) == "string", "binary desc table Name type")
T.Pass(typeof(p.SourceAssetId) == "number", "binary desc table SourceAssetId type")
T.Pass(typeof(p.Tags) == "string", "binary desc table Tags type")
T.Pass(typeof(p.Texture) == "string", "binary desc table Texture type")
T.Pass(typeof(p.Transparency) == "number", "binary desc table Transparency type")

T.Pass(p.AttributesSerialize == "", "binary desc AttributesSerialize value")
T.Pass(p.Color3 == Color3.new(1,1,1), "binary desc Color3 value")
T.Pass(p.Face == Enum.NormalId.Top, "binary desc Face value")
T.Pass(p.Name == "Decal", "binary desc Name value")
T.Pass(p.SourceAssetId == -1, "binary desc SourceAssetId value")
T.Pass(p.Tags == "", "binary desc Tags value")
T.Pass(p.Texture == "rbxasset://textures/SpawnLocation.png", "binary desc Texture value")
T.Pass(p.Transparency == 0, "binary desc Transparency value")
