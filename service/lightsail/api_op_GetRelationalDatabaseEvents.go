// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetRelationalDatabaseEventsInput struct {
	_ struct{} `type:"structure"`

	// The number of minutes in the past from which to retrieve events. For example,
	// to get all events from the past 2 hours, enter 120.
	//
	// Default: 60
	//
	// The minimum is 1 and the maximum is 14 days (20160 minutes).
	DurationInMinutes *int64 `locationName:"durationInMinutes" type:"integer"`

	// The token to advance to the next page of results from your request.
	//
	// To get a page token, perform an initial GetRelationalDatabaseEvents request.
	// If your results are paginated, the response will return a next page token
	// that you can specify as the page token in a subsequent request.
	PageToken *string `locationName:"pageToken" type:"string"`

	// The name of the database from which to get events.
	//
	// RelationalDatabaseName is a required field
	RelationalDatabaseName *string `locationName:"relationalDatabaseName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRelationalDatabaseEventsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRelationalDatabaseEventsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRelationalDatabaseEventsInput"}

	if s.RelationalDatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RelationalDatabaseName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetRelationalDatabaseEventsOutput struct {
	_ struct{} `type:"structure"`

	// The token to advance to the next page of resutls from your request.
	//
	// A next page token is not returned if there are no more results to display.
	//
	// To get the next page of results, perform another GetRelationalDatabaseEvents
	// request and specify the next page token using the pageToken parameter.
	NextPageToken *string `locationName:"nextPageToken" type:"string"`

	// An object describing the result of your get relational database events request.
	RelationalDatabaseEvents []RelationalDatabaseEvent `locationName:"relationalDatabaseEvents" type:"list"`
}

// String returns the string representation
func (s GetRelationalDatabaseEventsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetRelationalDatabaseEvents = "GetRelationalDatabaseEvents"

// GetRelationalDatabaseEventsRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns a list of events for a specific database in Amazon Lightsail.
//
//    // Example sending a request using GetRelationalDatabaseEventsRequest.
//    req := client.GetRelationalDatabaseEventsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetRelationalDatabaseEvents
func (c *Client) GetRelationalDatabaseEventsRequest(input *GetRelationalDatabaseEventsInput) GetRelationalDatabaseEventsRequest {
	op := &aws.Operation{
		Name:       opGetRelationalDatabaseEvents,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRelationalDatabaseEventsInput{}
	}

	req := c.newRequest(op, input, &GetRelationalDatabaseEventsOutput{})

	return GetRelationalDatabaseEventsRequest{Request: req, Input: input, Copy: c.GetRelationalDatabaseEventsRequest}
}

// GetRelationalDatabaseEventsRequest is the request type for the
// GetRelationalDatabaseEvents API operation.
type GetRelationalDatabaseEventsRequest struct {
	*aws.Request
	Input *GetRelationalDatabaseEventsInput
	Copy  func(*GetRelationalDatabaseEventsInput) GetRelationalDatabaseEventsRequest
}

// Send marshals and sends the GetRelationalDatabaseEvents API request.
func (r GetRelationalDatabaseEventsRequest) Send(ctx context.Context) (*GetRelationalDatabaseEventsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRelationalDatabaseEventsResponse{
		GetRelationalDatabaseEventsOutput: r.Request.Data.(*GetRelationalDatabaseEventsOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRelationalDatabaseEventsResponse is the response type for the
// GetRelationalDatabaseEvents API operation.
type GetRelationalDatabaseEventsResponse struct {
	*GetRelationalDatabaseEventsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRelationalDatabaseEvents request.
func (r *GetRelationalDatabaseEventsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
