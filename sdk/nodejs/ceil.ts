// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Returns the least integer value greater than or equal to the argument.
 */
export function ceil(args: CeilArgs, opts?: pulumi.InvokeOptions): Promise<CeilResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("std:index:ceil", {
        "input": args.input,
    }, opts);
}

export interface CeilArgs {
    input: number;
}

export interface CeilResult {
    readonly result: number;
}

export function ceilOutput(args: CeilOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<CeilResult> {
    return pulumi.output(args).apply(a => ceil(a, opts))
}

export interface CeilOutputArgs {
    input: pulumi.Input<number>;
}
