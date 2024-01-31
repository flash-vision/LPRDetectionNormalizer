package lprdetectionnormalizer
import (
    "fmt"
	"reflect"

)
type FieldMapping struct {
	MapToKey    string `json:"MapToKey"`
	WhenNull    bool   `json:"WhenNull"`
	FromOrdinal int    `json:"FromOrdinal"`
	ChildField  string `json:"ChildField"`
	// ChildFieldOrdinal is only used when ChildField is set
	ChildFieldOrdinal int `json:"ChildFieldOrdinal"`
}

type FieldMappings []FieldMapping
type MappingConfig map[string]FieldMappings

// Define structs for the original Kafka message
type DetectionMessage struct {
	SchemaVersion          string        `json:"schema_version"`
	MessageUID             string        `json:"message_uid"`
	MessageParentUID       string        `json:"message_parent_uid"`
	RobotUID               string        `json:"robot_uid"`
	EventTs                float64       `json:"event_ts"`
	UtcOffset              string        `json:"utc_offset"`
	Detections             []Detection   `json:"detections"`
	Boundings              []Bounding    `json:"boundings"`
	ImageData              []ImageData   `json:"image_data"`
	PayloadData            []PayloadData `json:"payload_data"`
	DetectionsCount        int           `json:"detections_count"`
	ImageCount             int           `json:"image_count"`
	BoundingsCount         int           `json:"boundings_count"`
	ImageAttributesCount   int           `json:"image_attributes_count"`
	PayloadAttributesCount int           `json:"payload_attributes_count"`
}

type Detection struct {
	Uid        string    `json:"uid"`
	SensorName string    `json:"sensor_name"`
	Type       string    `json:"_type"`
	Ts         float64   `json:"ts"`
	Label      []string  `json:"label"`
	Confidence []float64 `json:"confidence"`
}

type Bounding struct {
	DetectionUID string    `json:"detection_uid"`
	BoundingType string    `json:"bounding_type"`
	Values       []float64 `json:"_values"`
}

type ImageData struct {
	DetectionImages []string         `json:"detection_images"`
	ImageAttributes []ImageAttribute `json:"image_attributes"`
}

type ImageAttribute struct {
	Name  string      `json:"_name"`
	Value interface{} `json:"_value"`
}

type PayloadData struct {
	Name  string `json:"_name"`
	Value string `json:"_value"`
}

func TransformCustomMessage(msg DetectionMessage, config MappingConfig) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	val := reflect.ValueOf(msg)

	for fieldName, mappings := range config {
		fieldVal := val.FieldByName(fieldName)
		if !fieldVal.IsValid() {
			return nil, fmt.Errorf("field %s not found in DetectionMessage", fieldName)
		}

		for _, mapping := range mappings {
			if mapping.WhenNull && isFieldNilOrEmpty(fieldVal) {
				continue
			}

			if fieldVal.Kind() == reflect.Slice && fieldVal.Len() > 0 {
				// Handle case when FromOrdinal is -1 (selecting the last element)
				fromOrdinal := mapping.FromOrdinal
				if fromOrdinal == -1 {
					fromOrdinal = fieldVal.Len() - 1
				}

				if fromOrdinal >= 0 && fromOrdinal < fieldVal.Len() {
					element := fieldVal.Index(fromOrdinal)
					if mapping.ChildField != "" {
						childFieldVal := element.FieldByName(mapping.ChildField)
						if childFieldVal.IsValid() {
							childFieldOrdinal := mapping.ChildFieldOrdinal
							// Handle case when ChildFieldOrdinal is -1 (selecting the last element)
							if childFieldOrdinal == -1 && childFieldVal.Kind() == reflect.Slice {
								childFieldOrdinal = childFieldVal.Len() - 1
							}

							if childFieldOrdinal >= 0 && childFieldOrdinal < childFieldVal.Len() {
								result[mapping.MapToKey] = childFieldVal.Index(childFieldOrdinal).Interface()
							}
						} else {
							return nil, fmt.Errorf("child field %s not found in %s", mapping.ChildField, fieldName)
						}
					} else {
						result[mapping.MapToKey] = element.Interface()
					}
				}
			} else {
				result[mapping.MapToKey] = fieldVal.Interface()
			}
		}
	}

	return result, nil
}
func isFieldNilOrEmpty(field reflect.Value) bool {
	switch field.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Interface:
		return field.IsNil()
	case reflect.String:
		return field.Len() == 0
	default:
		return false
	}
}