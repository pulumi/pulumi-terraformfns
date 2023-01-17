// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.std.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Object;
import java.util.List;
import java.util.Objects;


public final class ContainsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final ContainsPlainArgs Empty = new ContainsPlainArgs();

    @Import(name="element", required=true)
    private Object element;

    public Object element() {
        return this.element;
    }

    @Import(name="input", required=true)
    private List<Object> input;

    public List<Object> input() {
        return this.input;
    }

    private ContainsPlainArgs() {}

    private ContainsPlainArgs(ContainsPlainArgs $) {
        this.element = $.element;
        this.input = $.input;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ContainsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ContainsPlainArgs $;

        public Builder() {
            $ = new ContainsPlainArgs();
        }

        public Builder(ContainsPlainArgs defaults) {
            $ = new ContainsPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder element(Object element) {
            $.element = element;
            return this;
        }

        public Builder input(List<Object> input) {
            $.input = input;
            return this;
        }

        public Builder input(Object... input) {
            return input(List.of(input));
        }

        public ContainsPlainArgs build() {
            $.element = Objects.requireNonNull($.element, "expected parameter 'element' to be non-null");
            $.input = Objects.requireNonNull($.input, "expected parameter 'input' to be non-null");
            return $;
        }
    }

}
