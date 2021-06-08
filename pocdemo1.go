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

type UserRequest struct {
	XMLName    xml.Name `xml:"get-user-request"`
	Permission int      `json:"permission"`
}

type UserResponse struct {
	UID      string
	Username string
}

type StatusResponse struct {
	Status  string
	Message string
	Data    []UserResponse
}

func PrintRequest(rq FooRequest) FooResponse {

	resp := FooResponse{
		Foo: "bar",
		Bar: "hello " + rq.Foo,
	}

	fmt.Println(resp)

	return resp
}

func GetUserRequest(rq UserRequest) StatusResponse {

	if rq.Permission != 0 {
		return StatusResponse{Status: "fail", Message: "permission denied"}
	}

	resp := StatusResponse{
		Status:  "success",
		Message: "success",
		Data: []UserResponse{
			{UID: "1111", Username: "TEST1"},
			{UID: "2222", Username: "TEST2"},
			{UID: "3333", Username: "TEST3"},
		},
	}

	fmt.Println(resp)

	return resp
}
