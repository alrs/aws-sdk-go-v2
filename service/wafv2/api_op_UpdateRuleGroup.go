// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafv2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateRuleGroupInput struct {
	_ struct{} `type:"structure"`

	// A description of the rule group that helps with identification. You cannot
	// change the description of a rule group after you create it.
	Description *string `min:"1" type:"string"`

	// A unique identifier for the rule group. This ID is returned in the responses
	// to create and list commands. You provide it to operations like update and
	// delete.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`

	// A token used for optimistic locking. AWS WAF returns a token to your get
	// and list requests, to mark the state of the entity at the time of the request.
	// To make changes to the entity associated with the token, you provide the
	// token to operations like update and delete. AWS WAF uses the token to ensure
	// that no changes have been made to the entity since you last retrieved it.
	// If a change has been made, the update fails with a WAFOptimisticLockException.
	// If this happens, perform another get, and use the new token returned by that
	// operation.
	//
	// LockToken is a required field
	LockToken *string `min:"1" type:"string" required:"true"`

	// The name of the rule group. You cannot change the name of a rule group after
	// you create it.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The Rule statements used to identify the web requests that you want to allow,
	// block, or count. Each rule includes one top-level statement that AWS WAF
	// uses to identify matching web requests, and parameters that govern how AWS
	// WAF handles them.
	Rules []Rule `type:"list"`

	// Specifies whether this is for an AWS CloudFront distribution or for a regional
	// application. A regional application can be an Application Load Balancer (ALB)
	// or an API Gateway stage.
	//
	// To work with CloudFront, you must also specify the Region US East (N. Virginia)
	// as follows:
	//
	//    * CLI - Specify the Region when you use the CloudFront scope: --scope=CLOUDFRONT
	//    --region=us-east-1.
	//
	//    * API and SDKs - For all calls, use the Region endpoint us-east-1.
	//
	// Scope is a required field
	Scope Scope `type:"string" required:"true" enum:"true"`

	// Defines and enables Amazon CloudWatch metrics and web request sample collection.
	//
	// VisibilityConfig is a required field
	VisibilityConfig *VisibilityConfig `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateRuleGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateRuleGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateRuleGroupInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}

	if s.LockToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("LockToken"))
	}
	if s.LockToken != nil && len(*s.LockToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LockToken", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if len(s.Scope) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Scope"))
	}

	if s.VisibilityConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("VisibilityConfig"))
	}
	if s.Rules != nil {
		for i, v := range s.Rules {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Rules", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.VisibilityConfig != nil {
		if err := s.VisibilityConfig.Validate(); err != nil {
			invalidParams.AddNested("VisibilityConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateRuleGroupOutput struct {
	_ struct{} `type:"structure"`

	// A token used for optimistic locking. AWS WAF returns this token to your update
	// requests. You use NextLockToken in the same manner as you use LockToken.
	NextLockToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateRuleGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateRuleGroup = "UpdateRuleGroup"

// UpdateRuleGroupRequest returns a request value for making API operation for
// AWS WAFV2.
//
//
// This is the latest version of AWS WAF, named AWS WAFV2, released in November,
// 2019. For information, including how to migrate your AWS WAF resources from
// the prior release, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
//
// Updates the specified RuleGroup.
//
// A rule group defines a collection of rules to inspect and control web requests
// that you can use in a WebACL. When you create a rule group, you define an
// immutable capacity limit. If you update a rule group, you must stay within
// the capacity. This allows others to reuse the rule group with confidence
// in its capacity requirements.
//
//    // Example sending a request using UpdateRuleGroupRequest.
//    req := client.UpdateRuleGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/wafv2-2019-07-29/UpdateRuleGroup
func (c *Client) UpdateRuleGroupRequest(input *UpdateRuleGroupInput) UpdateRuleGroupRequest {
	op := &aws.Operation{
		Name:       opUpdateRuleGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateRuleGroupInput{}
	}

	req := c.newRequest(op, input, &UpdateRuleGroupOutput{})

	return UpdateRuleGroupRequest{Request: req, Input: input, Copy: c.UpdateRuleGroupRequest}
}

// UpdateRuleGroupRequest is the request type for the
// UpdateRuleGroup API operation.
type UpdateRuleGroupRequest struct {
	*aws.Request
	Input *UpdateRuleGroupInput
	Copy  func(*UpdateRuleGroupInput) UpdateRuleGroupRequest
}

// Send marshals and sends the UpdateRuleGroup API request.
func (r UpdateRuleGroupRequest) Send(ctx context.Context) (*UpdateRuleGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateRuleGroupResponse{
		UpdateRuleGroupOutput: r.Request.Data.(*UpdateRuleGroupOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateRuleGroupResponse is the response type for the
// UpdateRuleGroup API operation.
type UpdateRuleGroupResponse struct {
	*UpdateRuleGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateRuleGroup request.
func (r *UpdateRuleGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
