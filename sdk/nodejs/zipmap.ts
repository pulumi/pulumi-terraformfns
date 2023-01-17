// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Constructs a map from a list of keys and a corresponding list of values.
 */
export function zipmap(args: ZipmapArgs, opts?: pulumi.InvokeOptions): Promise<ZipmapResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("std:index:zipmap", {
        "keys": args.keys,
        "values": args.values,
    }, opts);
}

export interface ZipmapArgs {
    keys: string[];
    values: any[];
}

export interface ZipmapResult {
    readonly result: {[key: string]: any};
}
/**
 * Constructs a map from a list of keys and a corresponding list of values.
 */
export function zipmapOutput(args: ZipmapOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ZipmapResult> {
    return pulumi.output(args).apply((a: any) => zipmap(a, opts))
}

export interface ZipmapOutputArgs {
    keys: pulumi.Input<pulumi.Input<string>[]>;
    values: pulumi.Input<any[]>;
}
