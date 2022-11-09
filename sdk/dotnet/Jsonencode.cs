// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Jsonencode
    {
        /// <summary>
        /// Returns a JSON-encoded representation of the given value, 
        /// which can contain arbitrarily-nested lists and maps. 
        /// Note that if the value is a string then its value will be placed in quotes.
        /// </summary>
        public static Task<JsonencodeResult> InvokeAsync(JsonencodeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<JsonencodeResult>("std:index:jsonencode", args ?? new JsonencodeArgs(), options.WithDefaults());

        /// <summary>
        /// Returns a JSON-encoded representation of the given value, 
        /// which can contain arbitrarily-nested lists and maps. 
        /// Note that if the value is a string then its value will be placed in quotes.
        /// </summary>
        public static Output<JsonencodeResult> Invoke(JsonencodeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<JsonencodeResult>("std:index:jsonencode", args ?? new JsonencodeInvokeArgs(), options.WithDefaults());
    }


    public sealed class JsonencodeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public object Input { get; set; } = null!;

        public JsonencodeArgs()
        {
        }
        public static new JsonencodeArgs Empty => new JsonencodeArgs();
    }

    public sealed class JsonencodeInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public Input<object> Input { get; set; } = null!;

        public JsonencodeInvokeArgs()
        {
        }
        public static new JsonencodeInvokeArgs Empty => new JsonencodeInvokeArgs();
    }


    [OutputType]
    public sealed class JsonencodeResult
    {
        public readonly string Result;

        [OutputConstructor]
        private JsonencodeResult(string result)
        {
            Result = result;
        }
    }
}
