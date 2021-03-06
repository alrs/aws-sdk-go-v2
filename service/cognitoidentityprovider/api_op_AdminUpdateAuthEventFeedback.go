// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AdminUpdateAuthEventFeedbackInput struct {
	_ struct{} `type:"structure"`

	// The authentication event ID.
	//
	// EventId is a required field
	EventId *string `min:"1" type:"string" required:"true"`

	// The authentication event feedback value.
	//
	// FeedbackValue is a required field
	FeedbackValue FeedbackValueType `type:"string" required:"true" enum:"true"`

	// The user pool ID.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`

	// The user pool username.
	//
	// Username is a required field
	Username *string `min:"1" type:"string" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s AdminUpdateAuthEventFeedbackInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AdminUpdateAuthEventFeedbackInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AdminUpdateAuthEventFeedbackInput"}

	if s.EventId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EventId"))
	}
	if s.EventId != nil && len(*s.EventId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventId", 1))
	}
	if len(s.FeedbackValue) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("FeedbackValue"))
	}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}

	if s.Username == nil {
		invalidParams.Add(aws.NewErrParamRequired("Username"))
	}
	if s.Username != nil && len(*s.Username) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Username", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AdminUpdateAuthEventFeedbackOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AdminUpdateAuthEventFeedbackOutput) String() string {
	return awsutil.Prettify(s)
}

const opAdminUpdateAuthEventFeedback = "AdminUpdateAuthEventFeedback"

// AdminUpdateAuthEventFeedbackRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Provides feedback for an authentication event as to whether it was from a
// valid user. This feedback is used for improving the risk evaluation decision
// for the user pool as part of Amazon Cognito advanced security.
//
//    // Example sending a request using AdminUpdateAuthEventFeedbackRequest.
//    req := client.AdminUpdateAuthEventFeedbackRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/AdminUpdateAuthEventFeedback
func (c *Client) AdminUpdateAuthEventFeedbackRequest(input *AdminUpdateAuthEventFeedbackInput) AdminUpdateAuthEventFeedbackRequest {
	op := &aws.Operation{
		Name:       opAdminUpdateAuthEventFeedback,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AdminUpdateAuthEventFeedbackInput{}
	}

	req := c.newRequest(op, input, &AdminUpdateAuthEventFeedbackOutput{})

	return AdminUpdateAuthEventFeedbackRequest{Request: req, Input: input, Copy: c.AdminUpdateAuthEventFeedbackRequest}
}

// AdminUpdateAuthEventFeedbackRequest is the request type for the
// AdminUpdateAuthEventFeedback API operation.
type AdminUpdateAuthEventFeedbackRequest struct {
	*aws.Request
	Input *AdminUpdateAuthEventFeedbackInput
	Copy  func(*AdminUpdateAuthEventFeedbackInput) AdminUpdateAuthEventFeedbackRequest
}

// Send marshals and sends the AdminUpdateAuthEventFeedback API request.
func (r AdminUpdateAuthEventFeedbackRequest) Send(ctx context.Context) (*AdminUpdateAuthEventFeedbackResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AdminUpdateAuthEventFeedbackResponse{
		AdminUpdateAuthEventFeedbackOutput: r.Request.Data.(*AdminUpdateAuthEventFeedbackOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AdminUpdateAuthEventFeedbackResponse is the response type for the
// AdminUpdateAuthEventFeedback API operation.
type AdminUpdateAuthEventFeedbackResponse struct {
	*AdminUpdateAuthEventFeedbackOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AdminUpdateAuthEventFeedback request.
func (r *AdminUpdateAuthEventFeedbackResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
