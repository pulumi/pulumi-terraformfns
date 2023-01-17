// Code generated by pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Constructs a map from a list of keys and a corresponding list of values.
func Zipmap(ctx *pulumi.Context, args *ZipmapArgs, opts ...pulumi.InvokeOption) (*ZipmapResult, error) {
	var rv ZipmapResult
	err := ctx.Invoke("std:index:zipmap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ZipmapArgs struct {
	Keys   []string      `pulumi:"keys"`
	Values []interface{} `pulumi:"values"`
}

type ZipmapResult struct {
	Result map[string]interface{} `pulumi:"result"`
}

func ZipmapOutput(ctx *pulumi.Context, args ZipmapOutputArgs, opts ...pulumi.InvokeOption) ZipmapResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ZipmapResult, error) {
			args := v.(ZipmapArgs)
			r, err := Zipmap(ctx, &args, opts...)
			var s ZipmapResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ZipmapResultOutput)
}

type ZipmapOutputArgs struct {
	Keys   pulumi.StringArrayInput `pulumi:"keys"`
	Values pulumi.ArrayInput       `pulumi:"values"`
}

func (ZipmapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ZipmapArgs)(nil)).Elem()
}

type ZipmapResultOutput struct{ *pulumi.OutputState }

func (ZipmapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ZipmapResult)(nil)).Elem()
}

func (o ZipmapResultOutput) ToZipmapResultOutput() ZipmapResultOutput {
	return o
}

func (o ZipmapResultOutput) ToZipmapResultOutputWithContext(ctx context.Context) ZipmapResultOutput {
	return o
}

func (o ZipmapResultOutput) Result() pulumi.MapOutput {
	return o.ApplyT(func(v ZipmapResult) map[string]interface{} { return v.Result }).(pulumi.MapOutput)
}

func init() {
	pulumi.RegisterOutputType(ZipmapResultOutput{})
}