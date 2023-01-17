// Code generated by pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Removes the specified prefix from the start of the given string, if present.
func Trimprefix(ctx *pulumi.Context, args *TrimprefixArgs, opts ...pulumi.InvokeOption) (*TrimprefixResult, error) {
	var rv TrimprefixResult
	err := ctx.Invoke("std:index:trimprefix", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type TrimprefixArgs struct {
	Input  string `pulumi:"input"`
	Prefix string `pulumi:"prefix"`
}

type TrimprefixResult struct {
	Result string `pulumi:"result"`
}

func TrimprefixOutput(ctx *pulumi.Context, args TrimprefixOutputArgs, opts ...pulumi.InvokeOption) TrimprefixResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (TrimprefixResult, error) {
			args := v.(TrimprefixArgs)
			r, err := Trimprefix(ctx, &args, opts...)
			var s TrimprefixResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(TrimprefixResultOutput)
}

type TrimprefixOutputArgs struct {
	Input  pulumi.StringInput `pulumi:"input"`
	Prefix pulumi.StringInput `pulumi:"prefix"`
}

func (TrimprefixOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrimprefixArgs)(nil)).Elem()
}

type TrimprefixResultOutput struct{ *pulumi.OutputState }

func (TrimprefixResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrimprefixResult)(nil)).Elem()
}

func (o TrimprefixResultOutput) ToTrimprefixResultOutput() TrimprefixResultOutput {
	return o
}

func (o TrimprefixResultOutput) ToTrimprefixResultOutputWithContext(ctx context.Context) TrimprefixResultOutput {
	return o
}

func (o TrimprefixResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v TrimprefixResult) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TrimprefixResultOutput{})
}
