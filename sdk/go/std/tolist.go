// Code generated by pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Converts its argument to a list value.
func Tolist(ctx *pulumi.Context, args *TolistArgs, opts ...pulumi.InvokeOption) (*TolistResult, error) {
	var rv TolistResult
	err := ctx.Invoke("std:index:tolist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type TolistArgs struct {
	Input []interface{} `pulumi:"input"`
}

type TolistResult struct {
	Result []interface{} `pulumi:"result"`
}

func TolistOutput(ctx *pulumi.Context, args TolistOutputArgs, opts ...pulumi.InvokeOption) TolistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (TolistResult, error) {
			args := v.(TolistArgs)
			r, err := Tolist(ctx, &args, opts...)
			var s TolistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(TolistResultOutput)
}

type TolistOutputArgs struct {
	Input pulumi.ArrayInput `pulumi:"input"`
}

func (TolistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TolistArgs)(nil)).Elem()
}

type TolistResultOutput struct{ *pulumi.OutputState }

func (TolistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TolistResult)(nil)).Elem()
}

func (o TolistResultOutput) ToTolistResultOutput() TolistResultOutput {
	return o
}

func (o TolistResultOutput) ToTolistResultOutputWithContext(ctx context.Context) TolistResultOutput {
	return o
}

func (o TolistResultOutput) Result() pulumi.ArrayOutput {
	return o.ApplyT(func(v TolistResult) []interface{} { return v.Result }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(TolistResultOutput{})
}
