// Code generated by pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns a list of the values of the map.
func Values(ctx *pulumi.Context, args *ValuesArgs, opts ...pulumi.InvokeOption) (*ValuesResult, error) {
	var rv ValuesResult
	err := ctx.Invoke("std:index:values", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ValuesArgs struct {
	Input map[string]interface{} `pulumi:"input"`
}

type ValuesResult struct {
	Result []interface{} `pulumi:"result"`
}

func ValuesOutput(ctx *pulumi.Context, args ValuesOutputArgs, opts ...pulumi.InvokeOption) ValuesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ValuesResult, error) {
			args := v.(ValuesArgs)
			r, err := Values(ctx, &args, opts...)
			var s ValuesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ValuesResultOutput)
}

type ValuesOutputArgs struct {
	Input pulumi.MapInput `pulumi:"input"`
}

func (ValuesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ValuesArgs)(nil)).Elem()
}

type ValuesResultOutput struct{ *pulumi.OutputState }

func (ValuesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ValuesResult)(nil)).Elem()
}

func (o ValuesResultOutput) ToValuesResultOutput() ValuesResultOutput {
	return o
}

func (o ValuesResultOutput) ToValuesResultOutputWithContext(ctx context.Context) ValuesResultOutput {
	return o
}

func (o ValuesResultOutput) Result() pulumi.ArrayOutput {
	return o.ApplyT(func(v ValuesResult) []interface{} { return v.Result }).(pulumi.ArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ValuesResultOutput{})
}
