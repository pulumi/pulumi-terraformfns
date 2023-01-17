// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Double;
import java.util.Objects;

@CustomType
public final class MaxResult {
    private Double result;

    private MaxResult() {}
    public Double result() {
        return this.result;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(MaxResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Double result;
        public Builder() {}
        public Builder(MaxResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.result = defaults.result;
        }

        @CustomType.Setter
        public Builder result(Double result) {
            this.result = Objects.requireNonNull(result);
            return this;
        }
        public MaxResult build() {
            final var o = new MaxResult();
            o.result = result;
            return o;
        }
    }
}
