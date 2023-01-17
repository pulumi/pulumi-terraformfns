// Code generated by pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Determines if the input string starts with the suffix.
func Startswith(ctx *pulumi.Context, args *StartswithArgs, opts ...pulumi.InvokeOption) (*StartswithResult, error) {
	var rv StartswithResult
	err := ctx.Invoke("std:index:startswith", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type StartswithArgs struct {
	Input  string `pulumi:"input"`
	Prefix string `pulumi:"prefix"`
}

type StartswithResult struct {
	Result bool `pulumi:"result"`
}

func StartswithOutput(ctx *pulumi.Context, args StartswithOutputArgs, opts ...pulumi.InvokeOption) StartswithResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (StartswithResult, error) {
			args := v.(StartswithArgs)
			r, err := Startswith(ctx, &args, opts...)
			var s StartswithResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(StartswithResultOutput)
}

type StartswithOutputArgs struct {
	Input  pulumi.StringInput `pulumi:"input"`
	Prefix pulumi.StringInput `pulumi:"prefix"`
}

func (StartswithOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StartswithArgs)(nil)).Elem()
}

type StartswithResultOutput struct{ *pulumi.OutputState }

func (StartswithResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StartswithResult)(nil)).Elem()
}

func (o StartswithResultOutput) ToStartswithResultOutput() StartswithResultOutput {
	return o
}

func (o StartswithResultOutput) ToStartswithResultOutputWithContext(ctx context.Context) StartswithResultOutput {
	return o
}

func (o StartswithResultOutput) Result() pulumi.BoolOutput {
	return o.ApplyT(func(v StartswithResult) bool { return v.Result }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(StartswithResultOutput{})
}
