package constants

type EmbedConfig struct {
	Color         int
	ToxicityLevel string
}

var EmbedConfigTypes = map[ToxicType]*EmbedConfig{
	ToxicTypeHigh: &EmbedConfig{
		Color:         10878976, // Rojo
		ToxicityLevel: "Alta",
	},
	ToxicTypeMedium: &EmbedConfig{
		Color:         12893718, // Amarillo
		ToxicityLevel: "Media",
	},
	ToxicTypeLow: &EmbedConfig{
		Color:         1491996, // Verde
		ToxicityLevel: "Baja",
	},
}
