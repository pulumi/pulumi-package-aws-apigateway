// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-apigateway/sdk/v2/go/apigateway/utilities"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigateway"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The RestAPI component offers a simple interface for creating a fully functional API Gateway REST API. The
// REST API can define any number of routes, each of which maps a path and HTTP method to one of (1) an event
// hander route that invokes a Lambda Function (2) a local path route which uploads local files into an S3 bucket
// and serves them or (3) an integration target such as an HTTP proxy or service integration.
type RestAPI struct {
	pulumi.ResourceState

	// The underlying RestAPI resource.
	Api apigateway.RestApiOutput `pulumi:"api"`
	// The underlying RestAPIPolicy resource.
	ApiPolicy apigateway.RestApiPolicyOutput `pulumi:"apiPolicy"`
	// The underlying Deployment resource.
	Deployment apigateway.DeploymentOutput `pulumi:"deployment"`
	// The underlying Stage resource.
	Stage apigateway.StageOutput `pulumi:"stage"`
	// The URL where the Rest API is exposed.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewRestAPI registers a new resource with the given unique name, arguments, and options.
func NewRestAPI(ctx *pulumi.Context,
	name string, args *RestAPIArgs, opts ...pulumi.ResourceOption) (*RestAPI, error) {
	if args == nil {
		args = &RestAPIArgs{}
	}

	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RestAPI
	err := ctx.RegisterRemoteComponentResource("aws-apigateway:index:RestAPI", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type restAPIArgs struct {
	// The source for the apikey. This can either be a HEADER or AUTHORIZER. If `apiKeyRequired` is
	// set to true on a route, and this is not defined the value will default to HEADER.
	ApiKeySource *APIKeySource `pulumi:"apiKeySource"`
	// List of binary media types supported by the REST API. By default, the REST API supports only UTF-8-encoded text payloads.
	// If importing an OpenAPI specification via the body argument, this corresponds to the x-amazon-apigateway-binary-media-types extension.
	// If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	BinaryMediaTypes []string `pulumi:"binaryMediaTypes"`
	// Description of the REST API.
	Description *string `pulumi:"description"`
	// Whether clients can invoke your API by using the default execute-api endpoint. By default, clients can invoke
	// your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that
	// clients use a custom domain name to invoke your API, disable the default endpoint. Defaults to false.
	DisableExecuteApiEndpoint *bool `pulumi:"disableExecuteApiEndpoint"`
	// Define custom gateway responses for the API. This can be used to properly enable
	// CORS for Lambda Authorizers.
	GatewayResponses map[string]SwaggerGatewayResponse `pulumi:"gatewayResponses"`
	// Request Validator specifies the validator to use at the API level. Note method level validators
	// override this.
	RequestValidator *RequestValidator `pulumi:"requestValidator"`
	// Routes to use to initialize the APIGateway.  These will be used to create the Swagger
	// specification for the API.
	//
	// Either `swaggerString` or `routes` must be specified.
	Routes []Route `pulumi:"routes"`
	// The stage name for your API. This will get added as a base path to your API url.
	StageName *string `pulumi:"stageName"`
	// Bucket to use for placing resources for static resources.  If not provided a default one will
	// be created on your behalf if any `StaticRoute`s are provided.
	StaticRoutesBucket *s3.Bucket `pulumi:"staticRoutesBucket"`
	// A Swagger specification already in string form to use to initialize the APIGateway.  Note
	// that you must manually provide permission for any route targets to be invoked by API Gateway
	// when using `swaggerString`.
	//
	// Either `swaggerString` or `routes` must be specified.
	SwaggerString *string `pulumi:"swaggerString"`
	// 'Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present,
	// tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a RestAPI resource.
type RestAPIArgs struct {
	// The source for the apikey. This can either be a HEADER or AUTHORIZER. If `apiKeyRequired` is
	// set to true on a route, and this is not defined the value will default to HEADER.
	ApiKeySource *APIKeySource
	// List of binary media types supported by the REST API. By default, the REST API supports only UTF-8-encoded text payloads.
	// If importing an OpenAPI specification via the body argument, this corresponds to the x-amazon-apigateway-binary-media-types extension.
	// If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
	BinaryMediaTypes []pulumi.StringInput
	// Description of the REST API.
	Description pulumi.StringPtrInput
	// Whether clients can invoke your API by using the default execute-api endpoint. By default, clients can invoke
	// your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that
	// clients use a custom domain name to invoke your API, disable the default endpoint. Defaults to false.
	DisableExecuteApiEndpoint pulumi.BoolPtrInput
	// Define custom gateway responses for the API. This can be used to properly enable
	// CORS for Lambda Authorizers.
	GatewayResponses map[string]SwaggerGatewayResponseInput
	// Request Validator specifies the validator to use at the API level. Note method level validators
	// override this.
	RequestValidator *RequestValidator
	// Routes to use to initialize the APIGateway.  These will be used to create the Swagger
	// specification for the API.
	//
	// Either `swaggerString` or `routes` must be specified.
	Routes []RouteInput
	// The stage name for your API. This will get added as a base path to your API url.
	StageName pulumi.StringPtrInput
	// Bucket to use for placing resources for static resources.  If not provided a default one will
	// be created on your behalf if any `StaticRoute`s are provided.
	StaticRoutesBucket s3.BucketInput
	// A Swagger specification already in string form to use to initialize the APIGateway.  Note
	// that you must manually provide permission for any route targets to be invoked by API Gateway
	// when using `swaggerString`.
	//
	// Either `swaggerString` or `routes` must be specified.
	SwaggerString pulumi.StringPtrInput
	// 'Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present,
	// tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (RestAPIArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*restAPIArgs)(nil)).Elem()
}

type RestAPIInput interface {
	pulumi.Input

	ToRestAPIOutput() RestAPIOutput
	ToRestAPIOutputWithContext(ctx context.Context) RestAPIOutput
}

func (*RestAPI) ElementType() reflect.Type {
	return reflect.TypeOf((**RestAPI)(nil)).Elem()
}

func (i *RestAPI) ToRestAPIOutput() RestAPIOutput {
	return i.ToRestAPIOutputWithContext(context.Background())
}

func (i *RestAPI) ToRestAPIOutputWithContext(ctx context.Context) RestAPIOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestAPIOutput)
}

type RestAPIOutput struct{ *pulumi.OutputState }

func (RestAPIOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RestAPI)(nil)).Elem()
}

func (o RestAPIOutput) ToRestAPIOutput() RestAPIOutput {
	return o
}

func (o RestAPIOutput) ToRestAPIOutputWithContext(ctx context.Context) RestAPIOutput {
	return o
}

// The underlying RestAPI resource.
func (o RestAPIOutput) Api() apigateway.RestApiOutput {
	return o.ApplyT(func(v *RestAPI) apigateway.RestApiOutput { return v.Api }).(apigateway.RestApiOutput)
}

// The underlying RestAPIPolicy resource.
func (o RestAPIOutput) ApiPolicy() apigateway.RestApiPolicyOutput {
	return o.ApplyT(func(v *RestAPI) apigateway.RestApiPolicyOutput { return v.ApiPolicy }).(apigateway.RestApiPolicyOutput)
}

// The underlying Deployment resource.
func (o RestAPIOutput) Deployment() apigateway.DeploymentOutput {
	return o.ApplyT(func(v *RestAPI) apigateway.DeploymentOutput { return v.Deployment }).(apigateway.DeploymentOutput)
}

// The underlying Stage resource.
func (o RestAPIOutput) Stage() apigateway.StageOutput {
	return o.ApplyT(func(v *RestAPI) apigateway.StageOutput { return v.Stage }).(apigateway.StageOutput)
}

// The URL where the Rest API is exposed.
func (o RestAPIOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *RestAPI) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RestAPIInput)(nil)).Elem(), &RestAPI{})
	pulumi.RegisterOutputType(RestAPIOutput{})
}
