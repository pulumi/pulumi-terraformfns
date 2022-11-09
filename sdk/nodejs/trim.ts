// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Removes the specified set of characters from the start and end of the given string.
 */
export function trim(args: TrimArgs, opts?: pulumi.InvokeOptions): Promise<TrimResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("std:index:trim", {
        "cutset": args.cutset,
        "input": args.input,
    }, opts);
}

export interface TrimArgs {
    cutset: string;
    input: string;
}

export interface TrimResult {
    readonly result: string;
}

export function trimOutput(args: TrimOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<TrimResult> {
    return pulumi.output(args).apply(a => trim(a, opts))
}

export interface TrimOutputArgs {
    cutset: pulumi.Input<string>;
    input: pulumi.Input<string>;
}
