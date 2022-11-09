// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Combines two or more lists into a single list.
 */
export function concat(args: ConcatArgs, opts?: pulumi.InvokeOptions): Promise<ConcatResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("std:index:concat", {
        "input": args.input,
    }, opts);
}

export interface ConcatArgs {
    input: any[][];
}

export interface ConcatResult {
    readonly result: any[];
}

export function concatOutput(args: ConcatOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ConcatResult> {
    return pulumi.output(args).apply(a => concat(a, opts))
}

export interface ConcatOutputArgs {
    input: pulumi.Input<pulumi.Input<any[]>[]>;
}
