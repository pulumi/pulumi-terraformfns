// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class JsondecodePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final JsondecodePlainArgs Empty = new JsondecodePlainArgs();

    @Import(name="input", required=true)
    private String input;

    public String input() {
        return this.input;
    }

    private JsondecodePlainArgs() {}

    private JsondecodePlainArgs(JsondecodePlainArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(JsondecodePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private JsondecodePlainArgs $;

        public Builder() {
            $ = new JsondecodePlainArgs();
        }

        public Builder(JsondecodePlainArgs defaults) {
            $ = new JsondecodePlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(String input) {
            $.input = input;
            return this;
        }

        public JsondecodePlainArgs build() {
            $.input = Objects.requireNonNull($.input, "expected parameter 'input' to be non-null");
            return $;
        }
    }

}
