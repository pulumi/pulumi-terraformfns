// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Object;
import java.util.List;
import java.util.Objects;

@CustomType
public final class ValuesResult {
    private List<Object> result;

    private ValuesResult() {}
    public List<Object> result() {
        return this.result;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ValuesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<Object> result;
        public Builder() {}
        public Builder(ValuesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.result = defaults.result;
        }

        @CustomType.Setter
        public Builder result(List<Object> result) {
            this.result = Objects.requireNonNull(result);
            return this;
        }
        public Builder result(Object... result) {
            return result(List.of(result));
        }
        public ValuesResult build() {
            final var o = new ValuesResult();
            o.result = result;
            return o;
        }
    }
}
