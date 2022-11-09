// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Formats a string according to the given format. The syntax for the format is standard sprintf syntax.
 */
export function format(args: FormatArgs, opts?: pulumi.InvokeOptions): Promise<FormatResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("std:index:format", {
        "args": args.args,
        "input": args.input,
    }, opts);
}

export interface FormatArgs {
    args: any[];
    input: string;
}

export interface FormatResult {
    readonly result: string;
}

export function formatOutput(args: FormatOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<FormatResult> {
    return pulumi.output(args).apply(a => format(a, opts))
}

export interface FormatOutputArgs {
    args: pulumi.Input<any[]>;
    input: pulumi.Input<string>;
}
