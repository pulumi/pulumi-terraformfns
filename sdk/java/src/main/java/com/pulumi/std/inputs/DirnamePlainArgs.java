// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class DirnamePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final DirnamePlainArgs Empty = new DirnamePlainArgs();

    @Import(name="input", required=true)
    private String input;

    public String input() {
        return this.input;
    }

    private DirnamePlainArgs() {}

    private DirnamePlainArgs(DirnamePlainArgs $) {
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DirnamePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DirnamePlainArgs $;

        public Builder() {
            $ = new DirnamePlainArgs();
        }

        public Builder(DirnamePlainArgs defaults) {
            $ = new DirnamePlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder input(String input) {
            $.input = input;
            return this;
        }

        public DirnamePlainArgs build() {
            $.input = Objects.requireNonNull($.input, "expected parameter 'input' to be non-null");
            return $;
        }
    }

}
