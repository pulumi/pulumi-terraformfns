// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class Sha1PlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final Sha1PlainArgs Empty = new Sha1PlainArgs();

    @Import(name="input", required=true)
    private String input;

    public String input() {
        return this.input;
    }

    private Sha1PlainArgs() {}

    private Sha1PlainArgs(Sha1PlainArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(Sha1PlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Sha1PlainArgs $;

        public Builder() {
            $ = new Sha1PlainArgs();
        }

        public Builder(Sha1PlainArgs defaults) {
            $ = new Sha1PlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(String input) {
            $.input = input;
            return this;
        }

        public Sha1PlainArgs build() {
            $.input = Objects.requireNonNull($.input, "expected parameter 'input' to be non-null");
            return $;
        }
    }

}