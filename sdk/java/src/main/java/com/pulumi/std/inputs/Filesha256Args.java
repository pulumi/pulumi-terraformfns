// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class Filesha256Args extends com.pulumi.resources.InvokeArgs {

    public static final Filesha256Args Empty = new Filesha256Args();

    @Import(name="input", required=true)
    private Output<String> input;

    public Output<String> input() {
        return this.input;
    }

    private Filesha256Args() {}

    private Filesha256Args(Filesha256Args $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(Filesha256Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Filesha256Args $;

        public Builder() {
            $ = new Filesha256Args();
        }

        public Builder(Filesha256Args defaults) {
            $ = new Filesha256Args(Objects.requireNonNull(defaults));
        }

        public Builder input(Output<String> input) {
            $.input = input;
            return this;
        }

        public Builder input(String input) {
            return input(Output.of(input));
        }

        public Filesha256Args build() {
            $.input = Objects.requireNonNull($.input, "expected parameter 'input' to be non-null");
            return $;
        }
    }

}