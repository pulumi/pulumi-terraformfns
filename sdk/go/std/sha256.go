// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-std/sdk/go/std/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns a hexadecimal representation of the SHA-256 hash of the given string.
func Sha256(ctx *pulumi.Context, args *Sha256Args, opts ...pulumi.InvokeOption) (*Sha256Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv Sha256Result
	err := ctx.Invoke("std:index:sha256", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type Sha256Args struct {
	Input string `pulumi:"input"`
}

type Sha256Result struct {
	Result string `pulumi:"result"`
}

func Sha256Output(ctx *pulumi.Context, args Sha256OutputArgs, opts ...pulumi.InvokeOption) Sha256ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (Sha256Result, error) {
			args := v.(Sha256Args)
			r, err := Sha256(ctx, &args, opts...)
			var s Sha256Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(Sha256ResultOutput)
}

type Sha256OutputArgs struct {
	Input pulumi.StringInput `pulumi:"input"`
}

func (Sha256OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sha256Args)(nil)).Elem()
}

type Sha256ResultOutput struct{ *pulumi.OutputState }

func (Sha256ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sha256Result)(nil)).Elem()
}

func (o Sha256ResultOutput) ToSha256ResultOutput() Sha256ResultOutput {
	return o
}

func (o Sha256ResultOutput) ToSha256ResultOutputWithContext(ctx context.Context) Sha256ResultOutput {
	return o
}

func (o Sha256ResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v Sha256Result) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(Sha256ResultOutput{})
}
