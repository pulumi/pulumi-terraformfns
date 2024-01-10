// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-std/sdk/go/std/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Removes the specified set of characters from the start and end of the given string.
func Trim(ctx *pulumi.Context, args *TrimArgs, opts ...pulumi.InvokeOption) (*TrimResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv TrimResult
	err := ctx.Invoke("std:index:trim", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type TrimArgs struct {
	Cutset string `pulumi:"cutset"`
	Input  string `pulumi:"input"`
}

type TrimResult struct {
	Result string `pulumi:"result"`
}

func TrimOutput(ctx *pulumi.Context, args TrimOutputArgs, opts ...pulumi.InvokeOption) TrimResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (TrimResult, error) {
			args := v.(TrimArgs)
			r, err := Trim(ctx, &args, opts...)
			var s TrimResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(TrimResultOutput)
}

type TrimOutputArgs struct {
	Cutset pulumi.StringInput `pulumi:"cutset"`
	Input  pulumi.StringInput `pulumi:"input"`
}

func (TrimOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrimArgs)(nil)).Elem()
}

type TrimResultOutput struct{ *pulumi.OutputState }

func (TrimResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrimResult)(nil)).Elem()
}

func (o TrimResultOutput) ToTrimResultOutput() TrimResultOutput {
	return o
}

func (o TrimResultOutput) ToTrimResultOutputWithContext(ctx context.Context) TrimResultOutput {
	return o
}

func (o TrimResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v TrimResult) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TrimResultOutput{})
}
