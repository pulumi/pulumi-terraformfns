# coding=utf-8
# *** WARNING: this file was generated by pulumi. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'SubstrResult',
    'AwaitableSubstrResult',
    'substr',
    'substr_output',
]

@pulumi.output_type
class SubstrResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, str):
            raise TypeError("Expected argument 'result' to be a str")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> str:
        return pulumi.get(self, "result")


class AwaitableSubstrResult(SubstrResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return SubstrResult(
            result=self.result)


def substr(input: Optional[str] = None,
           length: Optional[int] = None,
           offset: Optional[int] = None,
           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableSubstrResult:
    """
    Extracts a substring from the given string.
    """
    __args__ = dict()
    __args__['input'] = input
    __args__['length'] = length
    __args__['offset'] = offset
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:substr', __args__, opts=opts, typ=SubstrResult).value

    return AwaitableSubstrResult(
        result=__ret__.result)


@_utilities.lift_output_func(substr)
def substr_output(input: Optional[pulumi.Input[str]] = None,
                  length: Optional[pulumi.Input[int]] = None,
                  offset: Optional[pulumi.Input[int]] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[SubstrResult]:
    """
    Extracts a substring from the given string.
    """
    ...