package utils

type TypeNormalizer struct {
	Value string
}

func NormalizeType(value string) string {
	normalizer := TypeNormalizer{Value: value}
	return normalizer.Normalize()
}

func (normalizer *TypeNormalizer) Normalize() string {
	switch normalizer.Value {
	case "year":
		return "Int"
	case "string":
		return "String"
	case "any":
		return "Float"
	default:
		return "String"
	}
}
