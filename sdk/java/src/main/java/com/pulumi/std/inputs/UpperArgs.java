// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class UpperArgs extends com.pulumi.resources.InvokeArgs {

    public static final UpperArgs Empty = new UpperArgs();

    @Import(name="input", required=true)
    private Output<String> input;

    public Output<String> input() {
        return this.input;
    }

    private UpperArgs() {}

    private UpperArgs(UpperArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UpperArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UpperArgs $;

        public Builder() {
            $ = new UpperArgs();
        }

        public Builder(UpperArgs defaults) {
            $ = new UpperArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(Output<String> input) {
            $.input = input;
            return this;
        }

        public Builder input(String input) {
            return input(Output.of(input));
        }

        public UpperArgs build() {
            $.input = Objects.requireNonNull($.input, "expected parameter 'input' to be non-null");
            return $;
        }
    }

}