// Code generated by pulumi DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Takes an IP address range in CIDR notation and returns the address-formatted subnet mask format
// that some systems expect for IPv4 interfaces.
// For example, cidrnetmask("10.0.0.0/8") returns 255.0.0.0.
// Not applicable to IPv6 networks since CIDR notation is the only valid notation for IPv6.
func Cidrnetmask(ctx *pulumi.Context, args *CidrnetmaskArgs, opts ...pulumi.InvokeOption) (*CidrnetmaskResult, error) {
	var rv CidrnetmaskResult
	err := ctx.Invoke("std:index:cidrnetmask", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type CidrnetmaskArgs struct {
	Input string `pulumi:"input"`
}

type CidrnetmaskResult struct {
	Result string `pulumi:"result"`
}

func CidrnetmaskOutput(ctx *pulumi.Context, args CidrnetmaskOutputArgs, opts ...pulumi.InvokeOption) CidrnetmaskResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (CidrnetmaskResult, error) {
			args := v.(CidrnetmaskArgs)
			r, err := Cidrnetmask(ctx, &args, opts...)
			var s CidrnetmaskResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(CidrnetmaskResultOutput)
}

type CidrnetmaskOutputArgs struct {
	Input pulumi.StringInput `pulumi:"input"`
}

func (CidrnetmaskOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CidrnetmaskArgs)(nil)).Elem()
}

type CidrnetmaskResultOutput struct{ *pulumi.OutputState }

func (CidrnetmaskResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CidrnetmaskResult)(nil)).Elem()
}

func (o CidrnetmaskResultOutput) ToCidrnetmaskResultOutput() CidrnetmaskResultOutput {
	return o
}

func (o CidrnetmaskResultOutput) ToCidrnetmaskResultOutputWithContext(ctx context.Context) CidrnetmaskResultOutput {
	return o
}

func (o CidrnetmaskResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v CidrnetmaskResult) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CidrnetmaskResultOutput{})
}
