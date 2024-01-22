package lprdetectionnormalizer

type FieldMapping struct {
    MapToKey    string `json:"MapToKey"`
    WhenNull    bool   `json:"WhenNull"`
    FromOrdinal int    `json:"FromOrdinal"`
}

type MappingConfig map[string]FieldMapping
