// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ParseintPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final ParseintPlainArgs Empty = new ParseintPlainArgs();

    @Import(name="base")
    private @Nullable Integer base;

    public Optional<Integer> base() {
        return Optional.ofNullable(this.base);
    }

    @Import(name="input", required=true)
    private String input;

    public String input() {
        return this.input;
    }

    private ParseintPlainArgs() {}

    private ParseintPlainArgs(ParseintPlainArgs $) {
        this.base = $.base;
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ParseintPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ParseintPlainArgs $;

        public Builder() {
            $ = new ParseintPlainArgs();
        }

        public Builder(ParseintPlainArgs defaults) {
            $ = new ParseintPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder base(@Nullable Integer base) {
            $.base = base;
            return this;
        }

        public Builder input(String input) {
            $.input = input;
            return this;
        }

        public ParseintPlainArgs build() {
            $.input = Objects.requireNonNull($.input, "expected parameter 'input' to be non-null");
            return $;
        }
    }

}
