// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

import * as pulumiAws from "@pulumi/aws";

/**
 * The RestAPI component offers a simple interface for creating a fully functional API Gateway REST API. The
 * REST API can define any number of routes, each of which maps a path and HTTP method to one of (1) an event
 * hander route that invokes a Lambda Function (2) a local path route which uploads local files into an S3 bucket
 * and serves them or (3) an integration target such as an HTTP proxy or service integration.
 */
export class RestAPI extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'aws-apigateway:index:RestAPI';

    /**
     * Returns true if the given object is an instance of RestAPI.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RestAPI {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RestAPI.__pulumiType;
    }

    /**
     * The underlying RestAPI resource.
     */
    public /*out*/ readonly api!: pulumi.Output<pulumiAws.apigateway.RestApi>;
    /**
     * The underlying RestAPIPolicy resource.
     */
    public /*out*/ readonly apiPolicy!: pulumi.Output<pulumiAws.apigateway.RestApiPolicy | undefined>;
    /**
     * The underlying Deployment resource.
     */
    public /*out*/ readonly deployment!: pulumi.Output<pulumiAws.apigateway.Deployment>;
    /**
     * The underlying Stage resource.
     */
    public /*out*/ readonly stage!: pulumi.Output<pulumiAws.apigateway.Stage>;
    /**
     * The URL where the Rest API is exposed.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a RestAPI resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RestAPIArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["apiKeySource"] = args ? args.apiKeySource : undefined;
            resourceInputs["binaryMediaTypes"] = args ? args.binaryMediaTypes : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["disableExecuteApiEndpoint"] = args ? args.disableExecuteApiEndpoint : undefined;
            resourceInputs["gatewayResponses"] = args ? args.gatewayResponses : undefined;
            resourceInputs["requestValidator"] = args ? args.requestValidator : undefined;
            resourceInputs["routes"] = args ? args.routes : undefined;
            resourceInputs["stageName"] = args ? args.stageName : undefined;
            resourceInputs["staticRoutesBucket"] = args ? args.staticRoutesBucket : undefined;
            resourceInputs["swaggerString"] = args ? args.swaggerString : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["api"] = undefined /*out*/;
            resourceInputs["apiPolicy"] = undefined /*out*/;
            resourceInputs["deployment"] = undefined /*out*/;
            resourceInputs["stage"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        } else {
            resourceInputs["api"] = undefined /*out*/;
            resourceInputs["apiPolicy"] = undefined /*out*/;
            resourceInputs["deployment"] = undefined /*out*/;
            resourceInputs["stage"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RestAPI.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a RestAPI resource.
 */
export interface RestAPIArgs {
    /**
     * The source for the apikey. This can either be a HEADER or AUTHORIZER. If `apiKeyRequired` is
     * set to true on a route, and this is not defined the value will default to HEADER.
     */
    apiKeySource?: enums.APIKeySource;
    /**
     * List of binary media types supported by the REST API. By default, the REST API supports only UTF-8-encoded text payloads. 
     * If importing an OpenAPI specification via the body argument, this corresponds to the x-amazon-apigateway-binary-media-types extension. 
     * If the argument value is provided and is different than the OpenAPI value, the argument value will override the OpenAPI value.
     */
    binaryMediaTypes?: pulumi.Input<string>[];
    /**
     * Description of the REST API.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether clients can invoke your API by using the default execute-api endpoint. By default, clients can invoke
     * your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that
     * clients use a custom domain name to invoke your API, disable the default endpoint. Defaults to false.
     */
    disableExecuteApiEndpoint?: pulumi.Input<boolean>;
    /**
     * Define custom gateway responses for the API. This can be used to properly enable
     * CORS for Lambda Authorizers.
     */
    gatewayResponses?: {[key: string]: pulumi.Input<inputs.SwaggerGatewayResponseArgs>};
    /**
     * Request Validator specifies the validator to use at the API level. Note method level validators
     * override this.
     */
    requestValidator?: enums.RequestValidator;
    /**
     * Routes to use to initialize the APIGateway.  These will be used to create the Swagger
     * specification for the API.
     *
     * Either `swaggerString` or `routes` must be specified.
     */
    routes?: inputs.RouteArgs[];
    /**
     * The stage name for your API. This will get added as a base path to your API url.
     */
    stageName?: pulumi.Input<string>;
    /**
     * Bucket to use for placing resources for static resources.  If not provided a default one will
     * be created on your behalf if any `StaticRoute`s are provided.
     */
    staticRoutesBucket?: pulumi.Input<pulumiAws.s3.Bucket>;
    /**
     * A Swagger specification already in string form to use to initialize the APIGateway.  Note
     * that you must manually provide permission for any route targets to be invoked by API Gateway
     * when using `swaggerString`.
     *
     * Either `swaggerString` or `routes` must be specified.
     */
    swaggerString?: pulumi.Input<string>;
    /**
     * 'Map of tags to assign to the resource. If configured with a provider `defaultTags` configuration block present,
     * tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
