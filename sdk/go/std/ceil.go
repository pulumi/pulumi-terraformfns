// Code generated by pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the least integer value greater than or equal to the argument.
func Ceil(ctx *pulumi.Context, args *CeilArgs, opts ...pulumi.InvokeOption) (*CeilResult, error) {
	var rv CeilResult
	err := ctx.Invoke("std:index:ceil", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type CeilArgs struct {
	Input float64 `pulumi:"input"`
}

type CeilResult struct {
	Result float64 `pulumi:"result"`
}

func CeilOutput(ctx *pulumi.Context, args CeilOutputArgs, opts ...pulumi.InvokeOption) CeilResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (CeilResult, error) {
			args := v.(CeilArgs)
			r, err := Ceil(ctx, &args, opts...)
			var s CeilResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(CeilResultOutput)
}

type CeilOutputArgs struct {
	Input pulumi.Float64Input `pulumi:"input"`
}

func (CeilOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CeilArgs)(nil)).Elem()
}

type CeilResultOutput struct{ *pulumi.OutputState }

func (CeilResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CeilResult)(nil)).Elem()
}

func (o CeilResultOutput) ToCeilResultOutput() CeilResultOutput {
	return o
}

func (o CeilResultOutput) ToCeilResultOutputWithContext(ctx context.Context) CeilResultOutput {
	return o
}

func (o CeilResultOutput) Result() pulumi.Float64Output {
	return o.ApplyT(func(v CeilResult) float64 { return v.Result }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(CeilResultOutput{})
}
