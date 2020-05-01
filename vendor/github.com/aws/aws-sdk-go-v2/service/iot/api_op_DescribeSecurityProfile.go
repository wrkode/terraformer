// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeSecurityProfileInput struct {
	_ struct{} `type:"structure"`

	// The name of the security profile whose information you want to get.
	//
	// SecurityProfileName is a required field
	SecurityProfileName *string `location:"uri" locationName:"securityProfileName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeSecurityProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSecurityProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeSecurityProfileInput"}

	if s.SecurityProfileName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecurityProfileName"))
	}
	if s.SecurityProfileName != nil && len(*s.SecurityProfileName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SecurityProfileName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeSecurityProfileInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.SecurityProfileName != nil {
		v := *s.SecurityProfileName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "securityProfileName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeSecurityProfileOutput struct {
	_ struct{} `type:"structure"`

	// A list of metrics whose data is retained (stored). By default, data is retained
	// for any metric used in the profile's behaviors, but it is also retained for
	// any metric specified here.
	//
	// Note: This API field is deprecated. Please use DescribeSecurityProfileResponse$additionalMetricsToRetainV2
	// instead.
	AdditionalMetricsToRetain []string `locationName:"additionalMetricsToRetain" deprecated:"true" type:"list"`

	// A list of metrics whose data is retained (stored). By default, data is retained
	// for any metric used in the profile's behaviors, but it is also retained for
	// any metric specified here.
	AdditionalMetricsToRetainV2 []MetricToRetain `locationName:"additionalMetricsToRetainV2" type:"list"`

	// Where the alerts are sent. (Alerts are always sent to the console.)
	AlertTargets map[string]AlertTarget `locationName:"alertTargets" type:"map"`

	// Specifies the behaviors that, when violated by a device (thing), cause an
	// alert.
	Behaviors []Behavior `locationName:"behaviors" type:"list"`

	// The time the security profile was created.
	CreationDate *time.Time `locationName:"creationDate" type:"timestamp"`

	// The time the security profile was last modified.
	LastModifiedDate *time.Time `locationName:"lastModifiedDate" type:"timestamp"`

	// The ARN of the security profile.
	SecurityProfileArn *string `locationName:"securityProfileArn" type:"string"`

	// A description of the security profile (associated with the security profile
	// when it was created or updated).
	SecurityProfileDescription *string `locationName:"securityProfileDescription" type:"string"`

	// The name of the security profile.
	SecurityProfileName *string `locationName:"securityProfileName" min:"1" type:"string"`

	// The version of the security profile. A new version is generated whenever
	// the security profile is updated.
	Version *int64 `locationName:"version" type:"long"`
}

// String returns the string representation
func (s DescribeSecurityProfileOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeSecurityProfileOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AdditionalMetricsToRetain != nil {
		v := s.AdditionalMetricsToRetain

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "additionalMetricsToRetain", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.AdditionalMetricsToRetainV2 != nil {
		v := s.AdditionalMetricsToRetainV2

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "additionalMetricsToRetainV2", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.AlertTargets != nil {
		v := s.AlertTargets

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "alertTargets", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetFields(k1, v1)
		}
		ms0.End()

	}
	if s.Behaviors != nil {
		v := s.Behaviors

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "behaviors", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.CreationDate != nil {
		v := *s.CreationDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "creationDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.LastModifiedDate != nil {
		v := *s.LastModifiedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastModifiedDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.SecurityProfileArn != nil {
		v := *s.SecurityProfileArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "securityProfileArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SecurityProfileDescription != nil {
		v := *s.SecurityProfileDescription

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "securityProfileDescription", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SecurityProfileName != nil {
		v := *s.SecurityProfileName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "securityProfileName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.Int64Value(v), metadata)
	}
	return nil
}

const opDescribeSecurityProfile = "DescribeSecurityProfile"

// DescribeSecurityProfileRequest returns a request value for making API operation for
// AWS IoT.
//
// Gets information about a Device Defender security profile.
//
//    // Example sending a request using DescribeSecurityProfileRequest.
//    req := client.DescribeSecurityProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeSecurityProfileRequest(input *DescribeSecurityProfileInput) DescribeSecurityProfileRequest {
	op := &aws.Operation{
		Name:       opDescribeSecurityProfile,
		HTTPMethod: "GET",
		HTTPPath:   "/security-profiles/{securityProfileName}",
	}

	if input == nil {
		input = &DescribeSecurityProfileInput{}
	}

	req := c.newRequest(op, input, &DescribeSecurityProfileOutput{})
	return DescribeSecurityProfileRequest{Request: req, Input: input, Copy: c.DescribeSecurityProfileRequest}
}

// DescribeSecurityProfileRequest is the request type for the
// DescribeSecurityProfile API operation.
type DescribeSecurityProfileRequest struct {
	*aws.Request
	Input *DescribeSecurityProfileInput
	Copy  func(*DescribeSecurityProfileInput) DescribeSecurityProfileRequest
}

// Send marshals and sends the DescribeSecurityProfile API request.
func (r DescribeSecurityProfileRequest) Send(ctx context.Context) (*DescribeSecurityProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSecurityProfileResponse{
		DescribeSecurityProfileOutput: r.Request.Data.(*DescribeSecurityProfileOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSecurityProfileResponse is the response type for the
// DescribeSecurityProfile API operation.
type DescribeSecurityProfileResponse struct {
	*DescribeSecurityProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSecurityProfile request.
func (r *DescribeSecurityProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
