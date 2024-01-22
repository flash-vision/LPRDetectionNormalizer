package lprdetectionnormalizer
import (
	"fmt"
	"reflect"

)
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

        if mapping.FromOrdinal >= 0 {
            if fieldValue.Kind() == reflect.Slice && mapping.FromOrdinal < fieldValue.Len() {
                result[mapping.MapToKey] = fieldValue.Index(mapping.FromOrdinal).Interface()
            } else {
                return nil, fmt.Errorf("invalid FromOrdinal for field %s", field.Name)
            }
        } else {
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
