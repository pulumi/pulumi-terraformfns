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
    'ElementResult',
    'AwaitableElementResult',
    'element',
    'element_output',
]

@pulumi.output_type
class ElementResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, dict):
            raise TypeError("Expected argument 'result' to be a dict")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> Any:
        return pulumi.get(self, "result")


class AwaitableElementResult(ElementResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ElementResult(
            result=self.result)


def element(index: Optional[int] = None,
            input: Optional[Sequence[Any]] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableElementResult:
    """
    Returns the element at the specified index.
    """
    __args__ = dict()
    __args__['index'] = index
    __args__['input'] = input
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:element', __args__, opts=opts, typ=ElementResult).value

    return AwaitableElementResult(
        result=__ret__.result)


@_utilities.lift_output_func(element)
def element_output(index: Optional[pulumi.Input[int]] = None,
                   input: Optional[pulumi.Input[Sequence[Any]]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ElementResult]:
    """
    Returns the element at the specified index.
    """
    ...
