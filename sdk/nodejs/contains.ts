// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Returns true if a list contains the given element and returns false otherwise.
 */
export function contains(args: ContainsArgs, opts?: pulumi.InvokeOptions): Promise<ContainsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("std:index:contains", {
        "element": args.element,
        "input": args.input,
    }, opts);
}

export interface ContainsArgs {
    element: any;
    input: any[];
}

export interface ContainsResult {
    readonly result: boolean;
}

export function containsOutput(args: ContainsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ContainsResult> {
    return pulumi.output(args).apply(a => contains(a, opts))
}

export interface ContainsOutputArgs {
    element: any;
    input: pulumi.Input<any[]>;
}
