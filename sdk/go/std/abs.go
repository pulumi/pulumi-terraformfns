// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-std/sdk/go/std/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the absolute value of a given float.
// Example: abs(1) returns 1, and abs(-1) would also return 1, whereas abs(-3.14) would return 3.14.
func Abs(ctx *pulumi.Context, args *AbsArgs, opts ...pulumi.InvokeOption) (*AbsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv AbsResult
	err := ctx.Invoke("std:index:abs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type AbsArgs struct {
	Input float64 `pulumi:"input"`
}

type AbsResult struct {
	Result float64 `pulumi:"result"`
}

func AbsOutput(ctx *pulumi.Context, args AbsOutputArgs, opts ...pulumi.InvokeOption) AbsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (AbsResult, error) {
			args := v.(AbsArgs)
			r, err := Abs(ctx, &args, opts...)
			var s AbsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(AbsResultOutput)
}

type AbsOutputArgs struct {
	Input pulumi.Float64Input `pulumi:"input"`
}

func (AbsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AbsArgs)(nil)).Elem()
}

type AbsResultOutput struct{ *pulumi.OutputState }

func (AbsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AbsResult)(nil)).Elem()
}

func (o AbsResultOutput) ToAbsResultOutput() AbsResultOutput {
	return o
}

func (o AbsResultOutput) ToAbsResultOutputWithContext(ctx context.Context) AbsResultOutput {
	return o
}

func (o AbsResultOutput) Result() pulumi.Float64Output {
	return o.ApplyT(func(v AbsResult) float64 { return v.Result }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(AbsResultOutput{})
}
