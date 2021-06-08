package pocdemo1

import (
	"encoding/xml"
	"fmt"
)

// FooRequest a simple request
type FooRequest struct {
	XMLName xml.Name `xml:"fooRequest"`
	Foo     string   `json:"foo"`
}

// FooResponse a simple response
type FooResponse struct {
	Foo string `json:"foo"`
	Bar string `json:"bar"`
}

func PrintRequest(rq FooRequest) FooResponse {

	resp := FooResponse{
		Foo: "bar",
		Bar: "hello " + rq.Foo,
	}

	fmt.Println(resp)

	return resp
}
