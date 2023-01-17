// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Determines if the input string starts with the suffix.
 */
export function startswith(args: StartswithArgs, opts?: pulumi.InvokeOptions): Promise<StartswithResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:startswith", {
        "input": args.input,
        "prefix": args.prefix,
    }, opts);
}

export interface StartswithArgs {
    input: string;
    prefix: string;
}

export interface StartswithResult {
    readonly result: boolean;
}
/**
 * Determines if the input string starts with the suffix.
 */
export function startswithOutput(args: StartswithOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<StartswithResult> {
    return pulumi.output(args).apply((a: any) => startswith(a, opts))
}

export interface StartswithOutputArgs {
    input: pulumi.Input<string>;
    prefix: pulumi.Input<string>;
}