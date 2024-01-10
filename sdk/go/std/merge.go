// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-std/sdk/go/std/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the union of 2 or more maps. The maps are consumed in the order provided,
// and duplicate keys overwrite previous entries.
func Merge(ctx *pulumi.Context, args *MergeArgs, opts ...pulumi.InvokeOption) (*MergeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv MergeResult
	err := ctx.Invoke("std:index:merge", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type MergeArgs struct {
	Input []map[string]interface{} `pulumi:"input"`
}

type MergeResult struct {
	Result map[string]interface{} `pulumi:"result"`
}

func MergeOutput(ctx *pulumi.Context, args MergeOutputArgs, opts ...pulumi.InvokeOption) MergeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (MergeResult, error) {
			args := v.(MergeArgs)
			r, err := Merge(ctx, &args, opts...)
			var s MergeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(MergeResultOutput)
}

type MergeOutputArgs struct {
	Input pulumi.MapArrayInput `pulumi:"input"`
}

func (MergeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MergeArgs)(nil)).Elem()
}

type MergeResultOutput struct{ *pulumi.OutputState }

func (MergeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MergeResult)(nil)).Elem()
}

func (o MergeResultOutput) ToMergeResultOutput() MergeResultOutput {
	return o
}

func (o MergeResultOutput) ToMergeResultOutputWithContext(ctx context.Context) MergeResultOutput {
	return o
}

func (o MergeResultOutput) Result() pulumi.MapOutput {
	return o.ApplyT(func(v MergeResult) map[string]interface{} { return v.Result }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterOutputType(MergeResultOutput{})
}
