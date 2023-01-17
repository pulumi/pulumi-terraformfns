// Code generated by pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Reads the contents of a file and returns them as a base64-encoded string.
func Filebase64(ctx *pulumi.Context, args *Filebase64Args, opts ...pulumi.InvokeOption) (*Filebase64Result, error) {
	var rv Filebase64Result
	err := ctx.Invoke("std:index:filebase64", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type Filebase64Args struct {
	Input string `pulumi:"input"`
}

type Filebase64Result struct {
	Result string `pulumi:"result"`
}

func Filebase64Output(ctx *pulumi.Context, args Filebase64OutputArgs, opts ...pulumi.InvokeOption) Filebase64ResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (Filebase64Result, error) {
			args := v.(Filebase64Args)
			r, err := Filebase64(ctx, &args, opts...)
			var s Filebase64Result
			if r != nil {
				s = *r
			}
			return s, err
		}).(Filebase64ResultOutput)
}

type Filebase64OutputArgs struct {
	Input pulumi.StringInput `pulumi:"input"`
}

func (Filebase64OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Filebase64Args)(nil)).Elem()
}

type Filebase64ResultOutput struct{ *pulumi.OutputState }

func (Filebase64ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Filebase64Result)(nil)).Elem()
}

func (o Filebase64ResultOutput) ToFilebase64ResultOutput() Filebase64ResultOutput {
	return o
}

func (o Filebase64ResultOutput) ToFilebase64ResultOutputWithContext(ctx context.Context) Filebase64ResultOutput {
	return o
}

func (o Filebase64ResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v Filebase64Result) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(Filebase64ResultOutput{})
}