package metrics

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/margostino/openearth/common"
	"log"
	"net/http"
	"os"
)

var baseUrl = os.Getenv("GA_MP_URL")
var client = &http.Client{}

func send(payload *bytes.Buffer) {
	request, err := http.NewRequest("POST", baseUrl, payload)
	common.Check(err)
	q := request.URL.Query()
	q.Add("api_secret", os.Getenv("GA_MP_API_SECRET"))
	q.Add("measurement_id", os.Getenv("GA_MP_MEASUREMENT_ID"))
	request.URL.RawQuery = q.Encode()
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	common.Check(err)
	if response.StatusCode > 299 {
		log.Printf("Error publishing events to GA: %d\n", response.StatusCode)
	}
}

func encode(measure *Measure) *bytes.Buffer {
	postBody, _ := json.Marshal(measure)
	log.Printf("Measure: %s\n", string(postBody))
	payload := bytes.NewBuffer(postBody)
	return payload
}

func PublishRequest(request *http.Request) {
	measure := requestToMeasure(request)
	payload := encode(measure)
	send(payload)
}

func PublishQuery(ctx context.Context) {
	datasetContext, err := getDatasetContextFrom(ctx)
	common.Check(err)
	measure := contextToMeasure(datasetContext)
	payload := encode(measure)
	send(payload)
}
