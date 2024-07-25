// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsApiGateway.Inputs
{

    /// <summary>
    /// LambdaAuthorizer provides the definition for a custom Authorizer for API Gateway.
    /// </summary>
    public sealed class AuthorizerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the authorization mechanism for the client. Typical values are "oauth2" or "custom".
        /// </summary>
        [Input("authType")]
        public string? AuthType { get; set; }

        /// <summary>
        /// The name for the Authorizer to be referenced as. This must be unique for each unique
        /// authorizer within the API. If no name if specified, a name will be generated for you.
        /// </summary>
        [Input("authorizerName")]
        public string? AuthorizerName { get; set; }

        /// <summary>
        /// The number of seconds during which the resulting IAM policy is cached. Default is 300s. You
        /// can set this value to 0 to disable caching. Max value is 3600s. Note - if you are sharing an
        /// authorizer across more than one route you will want to disable the cache or else it will
        /// cause problems for you.
        /// </summary>
        [Input("authorizerResultTtlInSeconds")]
        public double? AuthorizerResultTtlInSeconds { get; set; }

        /// <summary>
        /// The authorizerHandler specifies information about the authorizing Lambda.
        /// </summary>
        [Input("handler")]
        public Input<Pulumi.Aws.Lambda.Function>? Handler { get; set; }

        [Input("identitySource")]
        private List<string>? _identitySource;

        /// <summary>
        /// List of mapping expressions of the request parameters as the identity source. This indicates
        /// where in the request identity information is expected. Applicable for the authorizer of the
        /// "request" type only. Example: ["method.request.header.HeaderAuth1",
        /// "method.request.querystring.QueryString1"]
        /// </summary>
        public List<string> IdentitySource
        {
            get => _identitySource ?? (_identitySource = new List<string>());
            set => _identitySource = value;
        }

        /// <summary>
        /// A regular expression for validating the token as the incoming identity. It only invokes the
        /// authorizer's lambda if there is a match, else it will return a 401. This does not apply to
        /// REQUEST Lambda Authorizers. Example: "^x-[a-z]+".
        /// </summary>
        [Input("identityValidationExpression")]
        public string? IdentityValidationExpression { get; set; }

        [Input("methodsToAuthorize")]
        private List<string>? _methodsToAuthorize;

        /// <summary>
        /// For method authorization, you can define resource servers and custom scopes by specifying the
        /// "resource-server/scope". e.g. ["com.hamuta.movies/drama.view",
        /// "http://my.resource.com/file.read"] For more information on resource servers and custom
        /// scopes visit the AWS documentation -
        /// https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-define-resource-servers.html
        /// </summary>
        public List<string> MethodsToAuthorize
        {
            get => _methodsToAuthorize ?? (_methodsToAuthorize = new List<string>());
            set => _methodsToAuthorize = value;
        }

        /// <summary>
        /// Defines where in the request API Gateway should look for identity information. The value must
        /// be "header" or "query". If there are multiple identity sources, the value must be "header".
        /// </summary>
        [Input("parameterLocation")]
        public string? ParameterLocation { get; set; }

        /// <summary>
        /// parameterName is the name of the header or query parameter containing the authorization
        /// token. Must be "Unused" for multiple identity sources.
        /// </summary>
        [Input("parameterName", required: true)]
        public string ParameterName { get; set; } = null!;

        [Input("providerARNs")]
        private List<Input<string>>? _providerARNs;

        /// <summary>
        /// The ARNs of the Cognito User Pools to use.
        /// </summary>
        public List<Input<string>> ProviderARNs
        {
            get => _providerARNs ?? (_providerARNs = new List<Input<string>>());
            set => _providerARNs = value;
        }

        /// <summary>
        /// The type of the authorizer. This value must be one of the following:
        /// - "token", for an authorizer with the caller identity embedded in an authorization token
        /// - "request", for an authorizer with the caller identity contained in request parameters
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        public AuthorizerArgs()
        {
        }
        public static new AuthorizerArgs Empty => new AuthorizerArgs();
    }
}
