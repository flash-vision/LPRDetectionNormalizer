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
type MappingConfig map[string]FieldMapping

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

    for i := 0; i < val.NumField(); i++ {
        field := val.Type().Field(i)
        mapping, ok := config[field.Name]
        if !ok {
            continue
        }

        fieldValue := val.Field(i)
        if mapping.WhenNull && isFieldNilOrEmpty(fieldValue) {
            continue
        }

        if fieldValue.Kind() == reflect.Slice {
            if mapping.FromOrdinal >= 0 && mapping.FromOrdinal < fieldValue.Len() {
                // Handle array/slice field
                element := fieldValue.Index(mapping.FromOrdinal)
                if mapping.ChildField != "" {
                    // Extract specific child field from the element
                    childFieldVal := element.FieldByName(mapping.ChildField)
                    if mapping.ChildFieldOrdinal >= 0 && childFieldVal.Kind() == reflect.Slice && mapping.ChildFieldOrdinal < childFieldVal.Len() {
                        result[mapping.MapToKey] = childFieldVal.Index(mapping.ChildFieldOrdinal).Interface()
                    } else if mapping.ChildFieldOrdinal == -1 {
                        result[mapping.MapToKey] = childFieldVal.Interface()
                    } else {
                        return nil, fmt.Errorf("invalid ChildFieldOrdinal for field %s", mapping.ChildField)
                    }
                } else {
                    result[mapping.MapToKey] = element.Interface()
                }
            } else if mapping.FromOrdinal == -1 {
                // If FromOrdinal is -1, map the entire array/slice
                result[mapping.MapToKey] = fieldValue.Interface()
            }
        } else {
            // Handle non-array/slice field
            result[mapping.MapToKey] = fieldValue.Interface()
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
