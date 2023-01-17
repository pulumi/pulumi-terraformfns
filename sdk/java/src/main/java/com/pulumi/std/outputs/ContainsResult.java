// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.util.Objects;

@CustomType
public final class ContainsResult {
    private Boolean result;

    private ContainsResult() {}
    public Boolean result() {
        return this.result;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ContainsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean result;
        public Builder() {}
        public Builder(ContainsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.result = defaults.result;
        }

        @CustomType.Setter
        public Builder result(Boolean result) {
            this.result = Objects.requireNonNull(result);
            return this;
        }
        public ContainsResult build() {
            final var o = new ContainsResult();
            o.result = result;
            return o;
        }
    }
}
