package reflect

import (
	"github.com/anaminus/rbxmk"
)

func All() []func() rbxmk.Type {
	return []func() rbxmk.Type{
		Array,
		Axes,
		BinaryString,
		Bool,
		BrickColor,
		CFrame,
		Color3,
		Color3uint8,
		ColorSequence,
		ColorSequenceKeypoint,
		Content,
		Dictionary,
		Double,
		Faces,
		Float,
		Instance,
		Instances,
		Int,
		Int64,
		Nil,
		Number,
		NumberRange,
		NumberSequence,
		NumberSequenceKeypoint,
		PhysicalProperties,
		ProtectedString,
		Ray,
		Rect,
		Region3,
		Region3int16,
		SharedString,
		String,
		Symbol,
		Table,
		Tuple,
		UDim,
		UDim2,
		Variant,
		Vector2,
		Vector2int16,
		Vector3,
		Vector3int16,
	}
}
