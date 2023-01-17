// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class Sha512PlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final Sha512PlainArgs Empty = new Sha512PlainArgs();

    @Import(name="input", required=true)
    private String input;

    public String input() {
        return this.input;
    }

    private Sha512PlainArgs() {}

    private Sha512PlainArgs(Sha512PlainArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(Sha512PlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Sha512PlainArgs $;

        public Builder() {
            $ = new Sha512PlainArgs();
        }

        public Builder(Sha512PlainArgs defaults) {
            $ = new Sha512PlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(String input) {
            $.input = input;
            return this;
        }

        public Sha512PlainArgs build() {
            $.input = Objects.requireNonNull($.input, "expected parameter 'input' to be non-null");
            return $;
        }
    }

}
