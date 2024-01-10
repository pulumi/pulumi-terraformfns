// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-std/sdk/go/std/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Converts its argument to a boolean value. Only boolean values, null, and the exact strings
//
//	"true" and "false" can be converted to boolean. All other values will result in an error.
func Tobool(ctx *pulumi.Context, args *ToboolArgs, opts ...pulumi.InvokeOption) (*ToboolResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ToboolResult
	err := ctx.Invoke("std:index:tobool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ToboolArgs struct {
	Input interface{} `pulumi:"input"`
}

type ToboolResult struct {
	Result *bool `pulumi:"result"`
}

func ToboolOutput(ctx *pulumi.Context, args ToboolOutputArgs, opts ...pulumi.InvokeOption) ToboolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ToboolResult, error) {
			args := v.(ToboolArgs)
			r, err := Tobool(ctx, &args, opts...)
			var s ToboolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ToboolResultOutput)
}

type ToboolOutputArgs struct {
	Input pulumi.Input `pulumi:"input"`
}

func (ToboolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ToboolArgs)(nil)).Elem()
}

type ToboolResultOutput struct{ *pulumi.OutputState }

func (ToboolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ToboolResult)(nil)).Elem()
}

func (o ToboolResultOutput) ToToboolResultOutput() ToboolResultOutput {
	return o
}

func (o ToboolResultOutput) ToToboolResultOutputWithContext(ctx context.Context) ToboolResultOutput {
	return o
}

func (o ToboolResultOutput) Result() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ToboolResult) *bool { return v.Result }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ToboolResultOutput{})
}
