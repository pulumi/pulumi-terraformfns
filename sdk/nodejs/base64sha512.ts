// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Returns a base64-encoded representation of raw SHA-512 sum of the given string.
 * This is not equivalent of base64encode(sha512(string)) since sha512() returns hexadecimal representation.
 */
export function base64sha512(args: Base64sha512Args, opts?: pulumi.InvokeOptions): Promise<Base64sha512Result> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("std:index:base64sha512", {
        "input": args.input,
    }, opts);
}

export interface Base64sha512Args {
    input: string;
}

export interface Base64sha512Result {
    readonly result: string;
}

export function base64sha512Output(args: Base64sha512OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<Base64sha512Result> {
    return pulumi.output(args).apply(a => base64sha512(a, opts))
}

export interface Base64sha512OutputArgs {
    input: pulumi.Input<string>;
}
