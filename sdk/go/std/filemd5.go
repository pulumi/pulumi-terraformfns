// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-std/sdk/go/std/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Reads the contents of a file into a string and returns the MD5 hash of it.
func Filemd5(ctx *pulumi.Context, args *Filemd5Args, opts ...pulumi.InvokeOption) (*Filemd5Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv Filemd5Result
	err := ctx.Invoke("std:index:filemd5", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type Filemd5Args struct {
	Input string `pulumi:"input"`
}

type Filemd5Result struct {
	Result string `pulumi:"result"`
}

func Filemd5Output(ctx *pulumi.Context, args Filemd5OutputArgs, opts ...pulumi.InvokeOption) Filemd5ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (Filemd5Result, error) {
			args := v.(Filemd5Args)
			r, err := Filemd5(ctx, &args, opts...)
			var s Filemd5Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(Filemd5ResultOutput)
}

type Filemd5OutputArgs struct {
	Input pulumi.StringInput `pulumi:"input"`
}

func (Filemd5OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Filemd5Args)(nil)).Elem()
}

type Filemd5ResultOutput struct{ *pulumi.OutputState }

func (Filemd5ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Filemd5Result)(nil)).Elem()
}

func (o Filemd5ResultOutput) ToFilemd5ResultOutput() Filemd5ResultOutput {
	return o
}

func (o Filemd5ResultOutput) ToFilemd5ResultOutputWithContext(ctx context.Context) Filemd5ResultOutput {
	return o
}

func (o Filemd5ResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v Filemd5Result) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(Filemd5ResultOutput{})
}
