// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package efs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type PutBackupPolicyInput struct {
	_ struct{} `type:"structure"`

	// The backup policy included in the PutBackupPolicy request.
	//
	// BackupPolicy is a required field
	BackupPolicy *BackupPolicy `type:"structure" required:"true"`

	// Specifies which EFS file system to update the backup policy for.
	//
	// FileSystemId is a required field
	FileSystemId *string `location:"uri" locationName:"FileSystemId" type:"string" required:"true"`
}

// String returns the string representation
func (s PutBackupPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutBackupPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutBackupPolicyInput"}

	if s.BackupPolicy == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupPolicy"))
	}

	if s.FileSystemId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FileSystemId"))
	}
	if s.BackupPolicy != nil {
		if err := s.BackupPolicy.Validate(); err != nil {
			invalidParams.AddNested("BackupPolicy", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBackupPolicyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.BackupPolicy != nil {
		v := s.BackupPolicy

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "BackupPolicy", v, metadata)
	}
	if s.FileSystemId != nil {
		v := *s.FileSystemId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FileSystemId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type PutBackupPolicyOutput struct {
	_ struct{} `type:"structure"`

	// Describes the file system's backup policy, indicating whether automatic backups
	// are turned on or off..
	BackupPolicy *BackupPolicy `type:"structure"`
}

// String returns the string representation
func (s PutBackupPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBackupPolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BackupPolicy != nil {
		v := s.BackupPolicy

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "BackupPolicy", v, metadata)
	}
	return nil
}

const opPutBackupPolicy = "PutBackupPolicy"

// PutBackupPolicyRequest returns a request value for making API operation for
// Amazon Elastic File System.
//
// Updates the file system's backup policy. Use this action to start or stop
// automatic backups of the file system.
//
//    // Example sending a request using PutBackupPolicyRequest.
//    req := client.PutBackupPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticfilesystem-2015-02-01/PutBackupPolicy
func (c *Client) PutBackupPolicyRequest(input *PutBackupPolicyInput) PutBackupPolicyRequest {
	op := &aws.Operation{
		Name:       opPutBackupPolicy,
		HTTPMethod: "PUT",
		HTTPPath:   "/2015-02-01/file-systems/{FileSystemId}/backup-policy",
	}

	if input == nil {
		input = &PutBackupPolicyInput{}
	}

	req := c.newRequest(op, input, &PutBackupPolicyOutput{})

	return PutBackupPolicyRequest{Request: req, Input: input, Copy: c.PutBackupPolicyRequest}
}

// PutBackupPolicyRequest is the request type for the
// PutBackupPolicy API operation.
type PutBackupPolicyRequest struct {
	*aws.Request
	Input *PutBackupPolicyInput
	Copy  func(*PutBackupPolicyInput) PutBackupPolicyRequest
}

// Send marshals and sends the PutBackupPolicy API request.
func (r PutBackupPolicyRequest) Send(ctx context.Context) (*PutBackupPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutBackupPolicyResponse{
		PutBackupPolicyOutput: r.Request.Data.(*PutBackupPolicyOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutBackupPolicyResponse is the response type for the
// PutBackupPolicy API operation.
type PutBackupPolicyResponse struct {
	*PutBackupPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutBackupPolicy request.
func (r *PutBackupPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
