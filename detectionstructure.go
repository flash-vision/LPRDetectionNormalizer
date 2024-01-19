package lprdetectionnormalizer

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

// Define a struct for the desired output format
type TransformedMessage struct {
	SchemaVersion     string   `json:"SCHEMA_VERSION"`
	MessageUID        string   `json:"MESSAGE_UID"`
	MessageParentUID  string   `json:"MESSAGE_PARENT_UID"`
	RobotUID          string   `json:"ROBOT_UID"`
	EventTs           float64  `json:"EVENT_TS"`
	TzOffset          string   `json:"TZ_OFFSET"`
	VehicleConfidence float64  `json:"VEHICLE_CONFIDENCE,omitempty"`
	LpConfidence      float64  `json:"LP_CONFIDENCE,omitempty"`
	StateLabel        string   `json:"STATE_LABEL,omitempty"`
	StateConfidence   float64  `json:"STATE_CONFIDENCE,omitempty"`
	LptextLabel       string   `json:"LPTEXT_LABEL,omitempty"`
	LptextConfidence  float64  `json:"LPTEXT_CONFIDENCE,omitempty"`
	MakeLabel         string   `json:"MAKE_LABEL,omitempty"`
	MakeConfidence    float64  `json:"MAKE_CONFIDENCE,omitempty"`
	TypeLabel         string   `json:"TYPE_LABEL,omitempty"`
	TypeConfidence    float64  `json:"TYPE_CONFIDENCE,omitempty"`
	ColorLabel        string   `json:"COLOR_LABEL,omitempty"`
	ColorConfidence   float64  `json:"COLOR_CONFIDENCE,omitempty"`
	EviceLabel        string   `json:"EVICE_LABEL,omitempty"`
	EviceConfidence   float64  `json:"EVICE_CONFIDENCE,omitempty"`
	SignatureUri      string   `json:"SIGNATURE_URI,omitempty"`
	DbMatch           string   `json:"DB_MATCH,omitempty"`
	TravelDirection   string   `json:"TRAVEL_DIRECTION,omitempty"`
	TransitType       string   `json:"TRANSIT_TYPE,omitempty"`
	DetectionImages   []string `json:"DETECTION_IMAGES"`
	Image             string   `json:"IMAGE,omitempty"`
	LpcImage          string   `json:"LPC_IMAGE,omitempty"`
}
