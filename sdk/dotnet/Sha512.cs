// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Sha512
    {
        /// <summary>
        /// Returns a hexadecimal representation of the SHA-512 hash of the given string.
        /// </summary>
        public static Task<Sha512Result> InvokeAsync(Sha512Args args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<Sha512Result>("std:index:sha512", args ?? new Sha512Args(), options.WithDefaults());

        /// <summary>
        /// Returns a hexadecimal representation of the SHA-512 hash of the given string.
        /// </summary>
        public static Output<Sha512Result> Invoke(Sha512InvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<Sha512Result>("std:index:sha512", args ?? new Sha512InvokeArgs(), options.WithDefaults());
    }


    public sealed class Sha512Args : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public string Input { get; set; } = null!;

        public Sha512Args()
        {
        }
        public static new Sha512Args Empty => new Sha512Args();
    }

    public sealed class Sha512InvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        public Input<string> Input { get; set; } = null!;

        public Sha512InvokeArgs()
        {
        }
        public static new Sha512InvokeArgs Empty => new Sha512InvokeArgs();
    }


    [OutputType]
    public sealed class Sha512Result
    {
        public readonly string Result;

        [OutputConstructor]
        private Sha512Result(string result)
        {
            Result = result;
        }
    }
}
