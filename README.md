# LPRDetectionNormalizer
 Converts a Flash Vision LPR detection message into a custom format.

# Overview

LPRDetectionNormalizer is a Go module designed to convert Flash Vision LPR detection messages into customizable formats. It offers flexibility in selecting specific fields from the original message and transforming them into a format suited for various application needs.

# Using the Normalizer:

First get the module...

> go get github.com/flash-vision/LPRDetectionNormalizer

Import the module in your code:

> import normalizer "github.com/flash-vision/LPRDetectionNormalizer"

```
func main() {
	// Define your JSON configuration as a string
	jsonConfig := `{
		"MessageUID": [{
			"MapToKey": "ID",
			"WhenNull": true,
			"FromOrdinal": -1
		}],
		"Detections": [{
			"MapToKey": "LPText",
			"WhenNull": false,
			"FromOrdinal": 6,
			"ChildField": "Label",
			"ChildFieldOrdinal": -1
		}, {
			"MapToKey": "LPTextConfidence",
			"WhenNull": false,
			"FromOrdinal": 6,
			"ChildField": "Confidence",
			"ChildFieldOrdinal": -1
		},
		{
			"MapToKey": "VehicleType",
			"WhenNull": false,
			"FromOrdinal": 3,
			"ChildField": "Label",
			"ChildFieldOrdinal": -1
		}, {
			"MapToKey": "VehicleTypeConfidence",
			"WhenNull": false,
			"FromOrdinal": 3,
			"ChildField": "Confidence",
			"ChildFieldOrdinal": -1
		}],
		"ImageData": [{
			"MapToKey": "FirstImage",
			"WhenNull": false,
			"FromOrdinal": 0,
			"ChildField": "DetectionImages",
			"ChildFieldOrdinal": 0
		}]
	}`

	// Initialize your MappingConfig
	customMap := make(normalizer.MappingConfig)

	// Unmarshal the JSON string into the customMap
	err := json.Unmarshal([]byte(jsonConfig), &customMap)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Assume you have a DetectionMessage instance ready to be transformed
	var msg normalizer.DetectionMessage
	jsonMessage := `{
		"schema_version": "v2",
		"message_uid": "d2260e19-946d-4373-a44c-f6ea52878558",
		"message_parent_uid": "",
		"robot_uid": "ROBOTID",
		"event_ts": 1705949638.1814127,
		"utc_offset": "utc-5",
		"detections": [
		  {
			"uid": "30b3bb13-2a71-463b-87cf-fd1a3c64ac80",
			"sensor_name": "32721016838860C6",
			"_type": "vehicle",
			"ts": 1705949638.1806557,
			"label": [
			  "0"
			],
			"confidence": [
			  0.9665083289146423
			]
		  },
		  {
			"uid": "fb4de9e3-8415-42b5-a0b1-85aee91e812e",
			"sensor_name": "32721016838860C6",
			"_type": "lp",
			"ts": 1705949638.1806781,
			"label": [
			  "0"
			],
			"confidence": [
			  0.96337890625
			]
		  },
		  {
			"uid": "8bf001bc-2645-4a2b-8ac9-750f160fee43",
			"sensor_name": "32721016838860C6",
			"_type": "make",
			"ts": 1705949638.1807373,
			"label": [
			  "toyota"
			],
			"confidence": [
			  0.7116
			]
		  },
		  {
			"uid": "bb89ce50-460f-419b-b1ff-27da00800a30",
			"sensor_name": "32721016838860C6",
			"_type": "vehicle_type",
			"ts": 1705949638.1807914,
			"label": [
			  "suv"
			],
			"confidence": [
			  0.941
			]
		  },
		  {
			"uid": "fe1f4ddc-ad3a-46b9-a2ab-ce785f8fd9f8",
			"sensor_name": "32721016838860C6",
			"_type": "color",
			"ts": 1705949638.1808448,
			"label": [
			  "gray"
			],
			"confidence": [
			  0.9884
			]
		  },
		  {
			"uid": "dcf71e55-0471-48b4-8acc-222d19c4e5df",
			"sensor_name": "32721016838860C6",
			"_type": "evice",
			"ts": 1705949638.180899,
			"label": [
			  "ICE"
			],
			"confidence": [
			  0.9358
			]
		  },
		  {
			"uid": "526a7301-8357-4060-9753-25a223dd58cb",
			"sensor_name": "32721016838860C6",
			"_type": "lptext",
			"ts": 1705949638.1809547,
			"label": [
			  "licenseplatetext"
			],
			"confidence": [
			  0.9735464875222194
			]
		  },
		  {
			"uid": "9f50eff8-90c2-4e97-b626-7bcdcaacba60",
			"sensor_name": "32721016838860C6",
			"_type": "lpstate",
			"ts": 1705949638.1810086,
			"label": [
			  "platestate"
			],
			"confidence": [
			  1
			]
		  },
		  {
			"uid": "5887da7a-9bfb-40af-99cf-efa79fe09aa5",
			"sensor_name": "32721016838860C6",
			"_type": "signature",
			"ts": 1705949638.1810622,
			"label": [
			  "SIGNATURE_FILE_URI"
			],
			"confidence": [
			  1
			]
		  },
		  {
			"uid": "449f0f6c-254e-474a-a79b-bbcd01e247e5",
			"sensor_name": "32721016838860C6",
			"_type": "db_match",
			"ts": 1705949638.181115,
			"label": [
			  "0"
			],
			"confidence": [
			  0
			]
		  },
		  {
			"uid": "8a9c5b19-2d2d-4f0e-bb6e-a396ed2204e7",
			"sensor_name": "32721016838860C6",
			"_type": "direction_of_travel",
			"ts": 1705949638.1812062,
			"label": [
			  "AWAY"
			],
			"confidence": [
			  1
			]
		  },
		  {
			"uid": "e0350a3c-7f71-4977-add3-7d84be1fb31b",
			"sensor_name": "32721016838860C6",
			"_type": "transit_type",
			"ts": 1705949638.1812718,
			"label": [
			  "EXIT"
			],
			"confidence": [
			  1
			]
		  }
		],
		"boundings": [
		  {
			"detection_uid": "30b3bb13-2a71-463b-87cf-fd1a3c64ac80",
			"bounding_type": "bounding_box",
			"_values": [
			  0.33471864461898804,
			  0.001,
			  0.7545456886291504,
			  0.7080800533294678
			]
		  },
		  {
			"detection_uid": "fb4de9e3-8415-42b5-a0b1-85aee91e812e",
			"bounding_type": "key_points",
			"_values": [
			  1964,
			  963.5,
			  2282,
			  963.5,
			  2272,
			  1102,
			  1954,
			  1112
			]
		  }
		],
		"image_data": [
		  {
			"detection_images": [
			  "IMAGE_URI_00-06.jpg",
			  "IMAGE_URI_01-06.jpg",
			  "IMAGE_URI_LPC.jpg"
			],
			"image_attributes": [
			  {
				"_name": "resolution",
				"_value": "1080p"
			  },
			  {
				"_name": "exposure",
				"_value": 3500
			  },
			  {
				"_name": "iso",
				"_value": 1600
			  },
			  {
				"_name": "focus",
				"_value": "0"
			  }
			]
		  }
		],
		"payload_data": [
		  {
			"_name": "last_encounter_ts",
			"_value": "-1"
		  }
		],
		"detections_count": 12,
		"image_count": 8,
		"boundings_count": 2,
		"image_attributes_count": 4,
		"payload_attributes_count": 1
	  }`
	err = json.Unmarshal([]byte(jsonMessage), &msg)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Transform the message
	transformed, err := normalizer.TransformCustomMessage(msg, customMap)
	if err != nil {
		log.Fatalf("Error transforming message: %v", err)
	}

	//marshal the transformed message to json

	// Print or use the transformed message
	fmt.Printf("Transformed Message: %+v\n", transformed)
}
```

# Configuration

The jsonConfig string in the example represents a configuration for mapping fields from the DetectionMessage to a custom format. Each entry in this configuration specifies how a field in the DetectionMessage should be mapped:

*"MapToKey"*: The new key name for the field in the transformed message.
*"WhenNull"*: A boolean indicating whether to skip mapping when the source field is null or not present.
*"FromOrdinal"*: Specifies the index of the element to map from an array field. Use -1 to map the entire array.
*"ChildField"* and "ChildFieldOrdinal": Used for nested fields within array elements.

This powerful feature allows for great flexibility in customizing the output format according to specific requirements.