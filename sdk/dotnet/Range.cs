// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Range
    {
        /// <summary>
        /// Generates a list of numbers using a start value, a limit value, and a step value.
        /// Start and step may be omitted, in which case start defaults to zero and step defaults to either one or negative one
        /// depending on whether limit is greater than or less than start.
        /// </summary>
        public static Task<RangeResult> InvokeAsync(RangeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<RangeResult>("std:index:range", args ?? new RangeArgs(), options.WithDefaults());

        /// <summary>
        /// Generates a list of numbers using a start value, a limit value, and a step value.
        /// Start and step may be omitted, in which case start defaults to zero and step defaults to either one or negative one
        /// depending on whether limit is greater than or less than start.
        /// </summary>
        public static Output<RangeResult> Invoke(RangeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<RangeResult>("std:index:range", args ?? new RangeInvokeArgs(), options.WithDefaults());
    }


    public sealed class RangeArgs : global::Pulumi.InvokeArgs
    {
        [Input("limit", required: true)]
        public double Limit { get; set; }

        [Input("start")]
        public double? Start { get; set; }

        [Input("step")]
        public double? Step { get; set; }

        public RangeArgs()
        {
        }
        public static new RangeArgs Empty => new RangeArgs();
    }

    public sealed class RangeInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("limit", required: true)]
        public Input<double> Limit { get; set; } = null!;

        [Input("start")]
        public Input<double>? Start { get; set; }

        [Input("step")]
        public Input<double>? Step { get; set; }

        public RangeInvokeArgs()
        {
        }
        public static new RangeInvokeArgs Empty => new RangeInvokeArgs();
    }


    [OutputType]
    public sealed class RangeResult
    {
        public readonly ImmutableArray<double> Result;

        [OutputConstructor]
        private RangeResult(ImmutableArray<double> result)
        {
            Result = result;
        }
    }
}
