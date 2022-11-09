// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Std
{
    public static class Join
    {
        /// <summary>
        /// Joins the list with the delimiter for a resultant string.
        /// </summary>
        public static Task<JoinResult> InvokeAsync(JoinArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<JoinResult>("std:index:join", args ?? new JoinArgs(), options.WithDefaults());

        /// <summary>
        /// Joins the list with the delimiter for a resultant string.
        /// </summary>
        public static Output<JoinResult> Invoke(JoinInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<JoinResult>("std:index:join", args ?? new JoinInvokeArgs(), options.WithDefaults());
    }


    public sealed class JoinArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        private List<string>? _input;
        public List<string> Input
        {
            get => _input ?? (_input = new List<string>());
            set => _input = value;
        }

        [Input("separator", required: true)]
        public string Separator { get; set; } = null!;

        public JoinArgs()
        {
        }
        public static new JoinArgs Empty => new JoinArgs();
    }

    public sealed class JoinInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("input", required: true)]
        private InputList<string>? _input;
        public InputList<string> Input
        {
            get => _input ?? (_input = new InputList<string>());
            set => _input = value;
        }

        [Input("separator", required: true)]
        public Input<string> Separator { get; set; } = null!;

        public JoinInvokeArgs()
        {
        }
        public static new JoinInvokeArgs Empty => new JoinInvokeArgs();
    }


    [OutputType]
    public sealed class JoinResult
    {
        public readonly string Result;

        [OutputConstructor]
        private JoinResult(string result)
        {
            Result = result;
        }
    }
}
