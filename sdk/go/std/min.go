// Code generated by pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the smallest of the floats.
func Min(ctx *pulumi.Context, args *MinArgs, opts ...pulumi.InvokeOption) (*MinResult, error) {
	var rv MinResult
	err := ctx.Invoke("std:index:min", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type MinArgs struct {
	Input []float64 `pulumi:"input"`
}

type MinResult struct {
	Result float64 `pulumi:"result"`
}

func MinOutput(ctx *pulumi.Context, args MinOutputArgs, opts ...pulumi.InvokeOption) MinResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (MinResult, error) {
			args := v.(MinArgs)
			r, err := Min(ctx, &args, opts...)
			var s MinResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(MinResultOutput)
}

type MinOutputArgs struct {
	Input pulumi.Float64ArrayInput `pulumi:"input"`
}

func (MinOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MinArgs)(nil)).Elem()
}

type MinResultOutput struct{ *pulumi.OutputState }

func (MinResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MinResult)(nil)).Elem()
}

func (o MinResultOutput) ToMinResultOutput() MinResultOutput {
	return o
}

func (o MinResultOutput) ToMinResultOutputWithContext(ctx context.Context) MinResultOutput {
	return o
}

func (o MinResultOutput) Result() pulumi.Float64Output {
	return o.ApplyT(func(v MinResult) float64 { return v.Result }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(MinResultOutput{})
}
