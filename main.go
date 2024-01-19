package lprdetectionnormalizer

import (
	"encoding/json"
	"fmt"
	"strings"
)



func TransformMessage(msg DetectionMessage) TransformedMessage {
	transformed := TransformedMessage{
		SchemaVersion:    strings.ToUpper(msg.SchemaVersion),
		MessageUID:       strings.ToUpper(msg.MessageUID),
		MessageParentUID: strings.ToUpper(msg.MessageParentUID),
		RobotUID:         strings.ToUpper(msg.RobotUID),
		EventTs:          msg.EventTs,
		TzOffset:         strings.ToUpper(strings.Replace(msg.UtcOffset, "utc", "UTC-", 1)),
		DetectionImages:  make([]string, 0),
	}

	for _, det := range msg.Detections {
		switch det.Type {
		case "vehicle":
			transformed.VehicleConfidence = det.Confidence[0]
		case "lp":
			transformed.LpConfidence = det.Confidence[0]
		case "lpstate":
			transformed.StateLabel = strings.ToUpper(det.Label[0])
			transformed.StateConfidence = det.Confidence[0]
		case "lptext":
			transformed.LptextLabel = strings.ToUpper(det.Label[0])
			transformed.LptextConfidence = det.Confidence[0]
		case "make":
			transformed.MakeLabel = strings.ToUpper(det.Label[0])
			transformed.MakeConfidence = det.Confidence[0]
		case "vehicle_type":
			transformed.TypeLabel = strings.ToUpper(det.Label[0])
			transformed.TypeConfidence = det.Confidence[0]
		case "color":
			transformed.ColorLabel = strings.ToUpper(det.Label[0])
			transformed.ColorConfidence = det.Confidence[0]
		case "evice":
			transformed.EviceLabel = strings.ToUpper(det.Label[0])
			transformed.EviceConfidence = det.Confidence[0]
		case "signature":
			transformed.SignatureUri = det.Label[0]
		case "db_match":
			transformed.DbMatch = det.Label[0]
		case "direction_of_travel":
			transformed.TravelDirection = det.Label[0]
		case "transit_type":
			transformed.TransitType = det.Label[0]
		}
	}

	for _, imgData := range msg.ImageData {
		for _, img := range imgData.DetectionImages {
			transformed.DetectionImages = append(transformed.DetectionImages, img)
			if strings.Contains(img, "__01-") {
				transformed.Image = img
			}
			if strings.Contains(img, "__LPC.jpg") {
				transformed.LpcImage = img
			}
		}
	}

	return transformed
}
