// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class ReplacePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final ReplacePlainArgs Empty = new ReplacePlainArgs();

    @Import(name="replace", required=true)
    private String replace;

    public String replace() {
        return this.replace;
    }

    @Import(name="search", required=true)
    private String search;

    public String search() {
        return this.search;
    }

    @Import(name="text", required=true)
    private String text;

    public String text() {
        return this.text;
    }

    private ReplacePlainArgs() {}

    private ReplacePlainArgs(ReplacePlainArgs $) {
        this.replace = $.replace;
        this.search = $.search;
        this.text = $.text;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ReplacePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ReplacePlainArgs $;

        public Builder() {
            $ = new ReplacePlainArgs();
        }

        public Builder(ReplacePlainArgs defaults) {
            $ = new ReplacePlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder replace(String replace) {
            $.replace = replace;
            return this;
        }

        public Builder search(String search) {
            $.search = search;
            return this;
        }

        public Builder text(String text) {
            $.text = text;
            return this;
        }

        public ReplacePlainArgs build() {
            $.replace = Objects.requireNonNull($.replace, "expected parameter 'replace' to be non-null");
            $.search = Objects.requireNonNull($.search, "expected parameter 'search' to be non-null");
            $.text = Objects.requireNonNull($.text, "expected parameter 'text' to be non-null");
            return $;
        }
    }

}
