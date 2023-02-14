package metrics

type Event struct {
	Name   string                 `json:"name"`
	Params map[string]interface{} `json:"params"`
}

type Measure struct {
	ClientId           string  `json:"client_id"`
	NonPersonalizedAds bool    `json:"non_personalized_ads"`
	Events             []Event `json:"events"`
}

type DatasetContext struct {
	Dataset   string
	Variables []string
}
