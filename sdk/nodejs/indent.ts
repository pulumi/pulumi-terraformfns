// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Adds a given number of spaces after each newline character in the given string.
 */
export function indent(args: IndentArgs, opts?: pulumi.InvokeOptions): Promise<IndentResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("std:index:indent", {
        "input": args.input,
        "spaces": args.spaces,
    }, opts);
}

export interface IndentArgs {
    input: string;
    spaces: number;
}

export interface IndentResult {
    readonly result: string;
}

export function indentOutput(args: IndentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<IndentResult> {
    return pulumi.output(args).apply(a => indent(a, opts))
}

export interface IndentOutputArgs {
    input: pulumi.Input<string>;
    spaces: pulumi.Input<number>;
}
