// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package std

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-std/sdk/go/std/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Does a search and replace on the given string.
// All instances of search are replaced with the value of replace.
// If search is wrapped in forward slashes, it is treated as a regular expression.
// If using a regular expression, replace can reference subcaptures in the regular expression by
// using $n where n is the index or name of the subcapture. If using a regular expression,
// the syntax conforms to the re2 regular expression syntax.
func Replace(ctx *pulumi.Context, args *ReplaceArgs, opts ...pulumi.InvokeOption) (*ReplaceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ReplaceResult
	err := ctx.Invoke("std:index:replace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ReplaceArgs struct {
	Replace string `pulumi:"replace"`
	Search  string `pulumi:"search"`
	Text    string `pulumi:"text"`
}

type ReplaceResult struct {
	Result string `pulumi:"result"`
}

func ReplaceOutput(ctx *pulumi.Context, args ReplaceOutputArgs, opts ...pulumi.InvokeOption) ReplaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ReplaceResult, error) {
			args := v.(ReplaceArgs)
			r, err := Replace(ctx, &args, opts...)
			var s ReplaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ReplaceResultOutput)
}

type ReplaceOutputArgs struct {
	Replace pulumi.StringInput `pulumi:"replace"`
	Search  pulumi.StringInput `pulumi:"search"`
	Text    pulumi.StringInput `pulumi:"text"`
}

func (ReplaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplaceArgs)(nil)).Elem()
}

type ReplaceResultOutput struct{ *pulumi.OutputState }

func (ReplaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplaceResult)(nil)).Elem()
}

func (o ReplaceResultOutput) ToReplaceResultOutput() ReplaceResultOutput {
	return o
}

func (o ReplaceResultOutput) ToReplaceResultOutputWithContext(ctx context.Context) ReplaceResultOutput {
	return o
}

func (o ReplaceResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v ReplaceResult) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ReplaceResultOutput{})
}
