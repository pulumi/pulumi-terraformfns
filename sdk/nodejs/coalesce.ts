// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Returns the first non-empty value from the given arguments.
 */
export function coalesce(args: CoalesceArgs, opts?: pulumi.InvokeOptions): Promise<CoalesceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:coalesce", {
        "input": args.input,
    }, opts);
}

export interface CoalesceArgs {
    input: string[];
}

export interface CoalesceResult {
    readonly result: string;
}
/**
 * Returns the first non-empty value from the given arguments.
 */
export function coalesceOutput(args: CoalesceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<CoalesceResult> {
    return pulumi.output(args).apply((a: any) => coalesce(a, opts))
}

export interface CoalesceOutputArgs {
    input: pulumi.Input<pulumi.Input<string>[]>;
}
