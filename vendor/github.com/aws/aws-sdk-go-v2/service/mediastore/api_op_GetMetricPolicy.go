// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediastore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetMetricPolicyInput struct {
	_ struct{} `type:"structure"`

	// The name of the container that is associated with the metric policy.
	//
	// ContainerName is a required field
	ContainerName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetMetricPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMetricPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetMetricPolicyInput"}

	if s.ContainerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContainerName"))
	}
	if s.ContainerName != nil && len(*s.ContainerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ContainerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetMetricPolicyOutput struct {
	_ struct{} `type:"structure"`

	// The metric policy that is associated with the specific container.
	//
	// MetricPolicy is a required field
	MetricPolicy *MetricPolicy `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetMetricPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetMetricPolicy = "GetMetricPolicy"

// GetMetricPolicyRequest returns a request value for making API operation for
// AWS Elemental MediaStore.
//
// Returns the metric policy for the specified container.
//
//    // Example sending a request using GetMetricPolicyRequest.
//    req := client.GetMetricPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/GetMetricPolicy
func (c *Client) GetMetricPolicyRequest(input *GetMetricPolicyInput) GetMetricPolicyRequest {
	op := &aws.Operation{
		Name:       opGetMetricPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetMetricPolicyInput{}
	}

	req := c.newRequest(op, input, &GetMetricPolicyOutput{})
	return GetMetricPolicyRequest{Request: req, Input: input, Copy: c.GetMetricPolicyRequest}
}

// GetMetricPolicyRequest is the request type for the
// GetMetricPolicy API operation.
type GetMetricPolicyRequest struct {
	*aws.Request
	Input *GetMetricPolicyInput
	Copy  func(*GetMetricPolicyInput) GetMetricPolicyRequest
}

// Send marshals and sends the GetMetricPolicy API request.
func (r GetMetricPolicyRequest) Send(ctx context.Context) (*GetMetricPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetMetricPolicyResponse{
		GetMetricPolicyOutput: r.Request.Data.(*GetMetricPolicyOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetMetricPolicyResponse is the response type for the
// GetMetricPolicy API operation.
type GetMetricPolicyResponse struct {
	*GetMetricPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetMetricPolicy request.
func (r *GetMetricPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
