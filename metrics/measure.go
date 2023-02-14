package metrics

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
	"net/http"
)

func sanitizeParam(value string) string {
	if value == "" {
		value = "empty"
	}
	return value
}

func requestToMeasure(request *http.Request) *Measure {
	ip := sanitizeParam(request.Header.Get("X-Real-Ip"))
	city := sanitizeParam(request.Header.Get("X-Vercel-Ip-City"))
	country := sanitizeParam(request.Header.Get("X-Vercel-Ip-Country"))
	path := sanitizeParam(request.URL.EscapedPath())

	event := Event{
		Name:   "request",
		Params: make(map[string]interface{}),
	}
	event.Params["country"] = country
	event.Params["city"] = city
	event.Params["path"] = path

	events := make([]Event, 0)
	events = append(events, event)

	measure := &Measure{
		ClientId:           ip,
		NonPersonalizedAds: false,
		Events:             events,
	}

	return measure
}
func contextToMeasure(datasetContext *DatasetContext) *Measure {
	var variables = make([]map[string]string, 0)
	for _, name := range datasetContext.Variables {
		var variable = make(map[string]string)
		variable["name"] = name
		variables = append(variables, variable)
	}

	event := Event{
		Name:   "query",
		Params: make(map[string]interface{}),
	}
	event.Params["dataset"] = datasetContext.Dataset
	event.Params["variables_count"] = len(datasetContext.Variables)
	event.Params["items"] = variables

	events := make([]Event, 0)
	events = append(events, event)

	measure := &Measure{
		ClientId:           "fetcher",
		NonPersonalizedAds: false,
		Events:             events,
	}

	return measure
}

func getDatasetContextFrom(ctx context.Context) (*DatasetContext, error) {
	fieldContext := graphql.GetFieldContext(ctx)
	if fieldContext != nil && len(fieldContext.Path()) > 0 {
		var datasetContext = &DatasetContext{
			Dataset:   fieldContext.Path().String(),
			Variables: make([]string, 0),
		}
		for _, selection := range fieldContext.Field.SelectionSet {
			field := selection.(*ast.Field)
			datasetContext.Variables = append(datasetContext.Variables, field.Name)
		}

		return datasetContext, nil
	}
	return nil, errors.New("field context does not exist in resolver")
}
