// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * For two lists values and keys of equal length,
 * returns all elements from values where the corresponding element from keys exists in the searchset list.
 */
export function matchkeys(args: MatchkeysArgs, opts?: pulumi.InvokeOptions): Promise<MatchkeysResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:matchkeys", {
        "searchList": args.searchList,
        "values": args.values,
    }, opts);
}

export interface MatchkeysArgs {
    searchList: string[];
    values: string[];
}

export interface MatchkeysResult {
    readonly result: string[];
}
/**
 * For two lists values and keys of equal length,
 * returns all elements from values where the corresponding element from keys exists in the searchset list.
 */
export function matchkeysOutput(args: MatchkeysOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<MatchkeysResult> {
    return pulumi.output(args).apply((a: any) => matchkeys(a, opts))
}

export interface MatchkeysOutputArgs {
    searchList: pulumi.Input<pulumi.Input<string>[]>;
    values: pulumi.Input<pulumi.Input<string>[]>;
}
