// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-std/sdk/go/std/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Adds a duration to a timestamp, returning a new timestamp.
//
//	Timestamps are represented as strings using RFC 3339 "Date and time format" syntax.
//	'timestamp' must be a string adhering this syntax, i.e. "2017-11-22T00:00:00Z".
//	'duration' is a string representation of a time difference, comprised of sequences of
//	numbers and unit pairs, i.e. "3.5h" or "2h15m".
//	Accepted units are "ns", "us" or "µs", "ms", "s", "m", and "h". The first number may be negative
//	to provide a negative duration, i.e. "-2h15m".
func Timeadd(ctx *pulumi.Context, args *TimeaddArgs, opts ...pulumi.InvokeOption) (*TimeaddResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv TimeaddResult
	err := ctx.Invoke("std:index:timeadd", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type TimeaddArgs struct {
	Duration  string `pulumi:"duration"`
	Timestamp string `pulumi:"timestamp"`
}

type TimeaddResult struct {
	Result string `pulumi:"result"`
}

func TimeaddOutput(ctx *pulumi.Context, args TimeaddOutputArgs, opts ...pulumi.InvokeOption) TimeaddResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (TimeaddResult, error) {
			args := v.(TimeaddArgs)
			r, err := Timeadd(ctx, &args, opts...)
			var s TimeaddResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(TimeaddResultOutput)
}

type TimeaddOutputArgs struct {
	Duration  pulumi.StringInput `pulumi:"duration"`
	Timestamp pulumi.StringInput `pulumi:"timestamp"`
}

func (TimeaddOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeaddArgs)(nil)).Elem()
}

type TimeaddResultOutput struct{ *pulumi.OutputState }

func (TimeaddResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeaddResult)(nil)).Elem()
}

func (o TimeaddResultOutput) ToTimeaddResultOutput() TimeaddResultOutput {
	return o
}

func (o TimeaddResultOutput) ToTimeaddResultOutputWithContext(ctx context.Context) TimeaddResultOutput {
	return o
}

func (o TimeaddResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v TimeaddResult) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TimeaddResultOutput{})
}
