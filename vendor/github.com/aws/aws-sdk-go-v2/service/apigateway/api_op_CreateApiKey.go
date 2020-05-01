// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request to create an ApiKey resource.
type CreateApiKeyInput struct {
	_ struct{} `type:"structure"`

	// An AWS Marketplace customer identifier , when integrating with the AWS SaaS
	// Marketplace.
	CustomerId *string `locationName:"customerId" type:"string"`

	// The description of the ApiKey.
	Description *string `locationName:"description" type:"string"`

	// Specifies whether the ApiKey can be used by callers.
	Enabled *bool `locationName:"enabled" type:"boolean"`

	// Specifies whether (true) or not (false) the key identifier is distinct from
	// the created API key value. This parameter is deprecated and should not be
	// used.
	GenerateDistinctId *bool `locationName:"generateDistinctId" type:"boolean"`

	// The name of the ApiKey.
	Name *string `locationName:"name" type:"string"`

	// DEPRECATED FOR USAGE PLANS - Specifies stages associated with the API key.
	StageKeys []StageKey `locationName:"stageKeys" type:"list"`

	// The key-value map of strings. The valid character set is [a-zA-Z+-=._:/].
	// The tag key can be up to 128 characters and must not start with aws:. The
	// tag value can be up to 256 characters.
	Tags map[string]string `locationName:"tags" type:"map"`

	// Specifies a value of the API key.
	Value *string `locationName:"value" type:"string"`
}

// String returns the string representation
func (s CreateApiKeyInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateApiKeyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.CustomerId != nil {
		v := *s.CustomerId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "customerId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Enabled != nil {
		v := *s.Enabled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enabled", protocol.BoolValue(v), metadata)
	}
	if s.GenerateDistinctId != nil {
		v := *s.GenerateDistinctId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "generateDistinctId", protocol.BoolValue(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StageKeys != nil {
		v := s.StageKeys

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "stageKeys", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.Value != nil {
		v := *s.Value

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "value", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// A resource that can be distributed to callers for executing Method resources
// that require an API key. API keys can be mapped to any Stage on any RestApi,
// which indicates that the callers with the API key can make requests to that
// stage.
//
// Use API Keys (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-api-keys.html)
type CreateApiKeyOutput struct {
	_ struct{} `type:"structure"`

	// The timestamp when the API Key was created.
	CreatedDate *time.Time `locationName:"createdDate" type:"timestamp"`

	// An AWS Marketplace customer identifier , when integrating with the AWS SaaS
	// Marketplace.
	CustomerId *string `locationName:"customerId" type:"string"`

	// The description of the API Key.
	Description *string `locationName:"description" type:"string"`

	// Specifies whether the API Key can be used by callers.
	Enabled *bool `locationName:"enabled" type:"boolean"`

	// The identifier of the API Key.
	Id *string `locationName:"id" type:"string"`

	// The timestamp when the API Key was last updated.
	LastUpdatedDate *time.Time `locationName:"lastUpdatedDate" type:"timestamp"`

	// The name of the API Key.
	Name *string `locationName:"name" type:"string"`

	// A list of Stage resources that are associated with the ApiKey resource.
	StageKeys []string `locationName:"stageKeys" type:"list"`

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]string `locationName:"tags" type:"map"`

	// The value of the API Key.
	Value *string `locationName:"value" type:"string"`
}

// String returns the string representation
func (s CreateApiKeyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateApiKeyOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CreatedDate != nil {
		v := *s.CreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.CustomerId != nil {
		v := *s.CustomerId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "customerId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Enabled != nil {
		v := *s.Enabled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "enabled", protocol.BoolValue(v), metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastUpdatedDate != nil {
		v := *s.LastUpdatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastUpdatedDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StageKeys != nil {
		v := s.StageKeys

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "stageKeys", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.Value != nil {
		v := *s.Value

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "value", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateApiKey = "CreateApiKey"

// CreateApiKeyRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Create an ApiKey resource.
//
// AWS CLI (https://docs.aws.amazon.com/cli/latest/reference/apigateway/create-api-key.html)
//
//    // Example sending a request using CreateApiKeyRequest.
//    req := client.CreateApiKeyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateApiKeyRequest(input *CreateApiKeyInput) CreateApiKeyRequest {
	op := &aws.Operation{
		Name:       opCreateApiKey,
		HTTPMethod: "POST",
		HTTPPath:   "/apikeys",
	}

	if input == nil {
		input = &CreateApiKeyInput{}
	}

	req := c.newRequest(op, input, &CreateApiKeyOutput{})
	return CreateApiKeyRequest{Request: req, Input: input, Copy: c.CreateApiKeyRequest}
}

// CreateApiKeyRequest is the request type for the
// CreateApiKey API operation.
type CreateApiKeyRequest struct {
	*aws.Request
	Input *CreateApiKeyInput
	Copy  func(*CreateApiKeyInput) CreateApiKeyRequest
}

// Send marshals and sends the CreateApiKey API request.
func (r CreateApiKeyRequest) Send(ctx context.Context) (*CreateApiKeyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateApiKeyResponse{
		CreateApiKeyOutput: r.Request.Data.(*CreateApiKeyOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateApiKeyResponse is the response type for the
// CreateApiKey API operation.
type CreateApiKeyResponse struct {
	*CreateApiKeyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateApiKey request.
func (r *CreateApiKeyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
