package lprdetectionnormalizer
import (

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
        for _, mapping := range mappings {
            fieldValue := val.FieldByName(fieldName)
            if mapping.WhenNull && isFieldNilOrEmpty(fieldValue) {
                continue
            }

            if fieldValue.Kind() == reflect.Slice {
                // Handle slice/array fields
                // ... Rest of your logic for handling slice/array fields
            } else {
                // Handle non-slice/array fields
                result[mapping.MapToKey] = fieldValue.Interface()
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
