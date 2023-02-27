package service

import (
	"encoding/json"
	"fmt"
	"github.com/margostino/openearth/common"
	"github.com/margostino/openearth/graph/model"
	"io/ioutil"
	"net/http"
)

var baseUrl = "https://www.unoosa.org/oosa/osoindex/waxs-search.json"

func CallOuterSpace(term string) ([]*model.OuterSpaceObject, error) {
	var startAt float64
	var found float64
	var response map[string]interface{}
	var outerSpaceObjects []*model.OuterSpaceObject

	found = 0
	startAt = 0
	for ok := true; ok; ok = (startAt == 0 || startAt < found) {
		var url = fmt.Sprintf("%s?criteria={\"startAt\":%f,\"match\":\"%s\"}", baseUrl, startAt, term)
		resp, err := http.Get(url)

		if !common.IsError(err, "when calling Outer Space Object") {
			bytes, err := ioutil.ReadAll(resp.Body)
			if !common.IsError(err, "when reading Outer Space Object response") {
				json.Unmarshal(bytes, &response)
				found = response["found"].(float64)
				results := response["results"].([]interface{})
				startAt += float64(len(results))
				for _, result := range results {
					result := result.(map[string]interface{})
					values := result["values"].(map[string]interface{})
					outerSpaceObject := &model.OuterSpaceObject{
						Name:       fmt.Sprintf("%s - %s", values["object.nameOfSpaceObjectIno_s1"].(string), values["object.nameOfSpaceObjectO_s1"].(string)),
						LaunchedAt: values["object.launch.dateOfLaunch_s1"].(string),
						DecayedAt:  values["object.status.dateOfDecay_s1"].(string),
						Status:     values["en#object.status.objectStatus_s1"].(string),
						Country:    values["object.launch.stateOfRegistry_s1"].(string),
						Function:   values["object.functionOfSpaceObject_s1"].(string),
					}
					outerSpaceObjects = append(outerSpaceObjects, outerSpaceObject)
				}
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
		resp.Body.Close()
	}

	return outerSpaceObjects, nil
}
