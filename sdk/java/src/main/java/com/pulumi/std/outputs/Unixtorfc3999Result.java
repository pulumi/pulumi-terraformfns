// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.util.Objects;

@CustomType
public final class Unixtorfc3999Result {
    private Integer result;

    private Unixtorfc3999Result() {}
    public Integer result() {
        return this.result;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(Unixtorfc3999Result defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer result;
        public Builder() {}
        public Builder(Unixtorfc3999Result defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.result = defaults.result;
        }

        @CustomType.Setter
        public Builder result(Integer result) {
            this.result = Objects.requireNonNull(result);
            return this;
        }
        public Unixtorfc3999Result build() {
            final var o = new Unixtorfc3999Result();
            o.result = result;
            return o;
        }
    }
}