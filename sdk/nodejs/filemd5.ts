// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Reads the contents of a file into a string and returns the MD5 hash of it.
 */
export function filemd5(args: Filemd5Args, opts?: pulumi.InvokeOptions): Promise<Filemd5Result> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:filemd5", {
        "input": args.input,
    }, opts);
}

export interface Filemd5Args {
    input: string;
}

export interface Filemd5Result {
    readonly result: string;
}
/**
 * Reads the contents of a file into a string and returns the MD5 hash of it.
 */
export function filemd5Output(args: Filemd5OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<Filemd5Result> {
    return pulumi.output(args).apply((a: any) => filemd5(a, opts))
}

export interface Filemd5OutputArgs {
    input: pulumi.Input<string>;
}
