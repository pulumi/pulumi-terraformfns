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
    'ContainsResult',
    'AwaitableContainsResult',
    'contains',
    'contains_output',
]

@pulumi.output_type
class ContainsResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, bool):
            raise TypeError("Expected argument 'result' to be a bool")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> bool:
        return pulumi.get(self, "result")


class AwaitableContainsResult(ContainsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ContainsResult(
            result=self.result)


def contains(element: Optional[Any] = None,
             input: Optional[Sequence[Any]] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableContainsResult:
    """
    Returns true if a list contains the given element and returns false otherwise.
    """
    __args__ = dict()
    __args__['element'] = element
    __args__['input'] = input
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:contains', __args__, opts=opts, typ=ContainsResult).value

    return AwaitableContainsResult(
        result=__ret__.result)


@_utilities.lift_output_func(contains)
def contains_output(element: Optional[Any] = None,
                    input: Optional[pulumi.Input[Sequence[Any]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ContainsResult]:
    """
    Returns true if a list contains the given element and returns false otherwise.
    """
    ...