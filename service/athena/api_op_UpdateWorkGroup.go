// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package athena

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateWorkGroupInput struct {
	_ struct{} `type:"structure"`

	// The workgroup configuration that will be updated for the given workgroup.
	ConfigurationUpdates *WorkGroupConfigurationUpdates `type:"structure"`

	// The workgroup description.
	Description *string `type:"string"`

	// The workgroup state that will be updated for the given workgroup.
	State WorkGroupState `type:"string" enum:"true"`

	// The specified workgroup that will be updated.
	//
	// WorkGroup is a required field
	WorkGroup *string `type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateWorkGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateWorkGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateWorkGroupInput"}

	if s.WorkGroup == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkGroup"))
	}
	if s.ConfigurationUpdates != nil {
		if err := s.ConfigurationUpdates.Validate(); err != nil {
			invalidParams.AddNested("ConfigurationUpdates", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateWorkGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateWorkGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateWorkGroup = "UpdateWorkGroup"

// UpdateWorkGroupRequest returns a request value for making API operation for
// Amazon Athena.
//
// Updates the workgroup with the specified name. The workgroup's name cannot
// be changed.
//
//    // Example sending a request using UpdateWorkGroupRequest.
//    req := client.UpdateWorkGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/athena-2017-05-18/UpdateWorkGroup
func (c *Client) UpdateWorkGroupRequest(input *UpdateWorkGroupInput) UpdateWorkGroupRequest {
	op := &aws.Operation{
		Name:       opUpdateWorkGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateWorkGroupInput{}
	}

	req := c.newRequest(op, input, &UpdateWorkGroupOutput{})

	return UpdateWorkGroupRequest{Request: req, Input: input, Copy: c.UpdateWorkGroupRequest}
}

// UpdateWorkGroupRequest is the request type for the
// UpdateWorkGroup API operation.
type UpdateWorkGroupRequest struct {
	*aws.Request
	Input *UpdateWorkGroupInput
	Copy  func(*UpdateWorkGroupInput) UpdateWorkGroupRequest
}

// Send marshals and sends the UpdateWorkGroup API request.
func (r UpdateWorkGroupRequest) Send(ctx context.Context) (*UpdateWorkGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateWorkGroupResponse{
		UpdateWorkGroupOutput: r.Request.Data.(*UpdateWorkGroupOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateWorkGroupResponse is the response type for the
// UpdateWorkGroup API operation.
type UpdateWorkGroupResponse struct {
	*UpdateWorkGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateWorkGroup request.
func (r *UpdateWorkGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
