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
    'LookupResult',
    'AwaitableLookupResult',
    'lookup',
    'lookup_output',
]

@pulumi.output_type
class LookupResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, dict):
            raise TypeError("Expected argument 'result' to be a dict")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> Any:
        return pulumi.get(self, "result")


class AwaitableLookupResult(LookupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return LookupResult(
            result=self.result)


def lookup(default: Optional[Any] = None,
           key: Optional[str] = None,
           map: Optional[Mapping[str, Any]] = None,
           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableLookupResult:
    """
    Performs a dynamic lookup into a map variable.
    """
    __args__ = dict()
    __args__['default'] = default
    __args__['key'] = key
    __args__['map'] = map
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:lookup', __args__, opts=opts, typ=LookupResult).value

    return AwaitableLookupResult(
        result=__ret__.result)


@_utilities.lift_output_func(lookup)
def lookup_output(default: Optional[pulumi.Input[Optional[Any]]] = None,
                  key: Optional[pulumi.Input[str]] = None,
                  map: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[LookupResult]:
    """
    Performs a dynamic lookup into a map variable.
    """
    ...
