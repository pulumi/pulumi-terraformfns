// Code generated by pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Removes the specified prefix from the start of the given string, if present.
func TrimPrefix(ctx *pulumi.Context, args *TrimPrefixArgs, opts ...pulumi.InvokeOption) (*TrimPrefixResult, error) {
	var rv TrimPrefixResult
	err := ctx.Invoke("std:index:trimPrefix", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type TrimPrefixArgs struct {
	Input  string `pulumi:"input"`
	Prefix string `pulumi:"prefix"`
}

type TrimPrefixResult struct {
	Result string `pulumi:"result"`
}

func TrimPrefixOutput(ctx *pulumi.Context, args TrimPrefixOutputArgs, opts ...pulumi.InvokeOption) TrimPrefixResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (TrimPrefixResult, error) {
			args := v.(TrimPrefixArgs)
			r, err := TrimPrefix(ctx, &args, opts...)
			var s TrimPrefixResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(TrimPrefixResultOutput)
}

type TrimPrefixOutputArgs struct {
	Input  pulumi.StringInput `pulumi:"input"`
	Prefix pulumi.StringInput `pulumi:"prefix"`
}

func (TrimPrefixOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrimPrefixArgs)(nil)).Elem()
}

type TrimPrefixResultOutput struct{ *pulumi.OutputState }

func (TrimPrefixResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrimPrefixResult)(nil)).Elem()
}

func (o TrimPrefixResultOutput) ToTrimPrefixResultOutput() TrimPrefixResultOutput {
	return o
}

func (o TrimPrefixResultOutput) ToTrimPrefixResultOutputWithContext(ctx context.Context) TrimPrefixResultOutput {
	return o
}

func (o TrimPrefixResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v TrimPrefixResult) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TrimPrefixResultOutput{})
}
