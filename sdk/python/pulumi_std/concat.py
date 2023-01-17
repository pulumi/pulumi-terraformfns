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
    'ConcatResult',
    'AwaitableConcatResult',
    'concat',
    'concat_output',
]

@pulumi.output_type
class ConcatResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, list):
            raise TypeError("Expected argument 'result' to be a list")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> Sequence[Any]:
        return pulumi.get(self, "result")


class AwaitableConcatResult(ConcatResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ConcatResult(
            result=self.result)


def concat(input: Optional[Sequence[Sequence[Any]]] = None,
           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableConcatResult:
    """
    Combines two or more lists into a single list.
    """
    __args__ = dict()
    __args__['input'] = input
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:concat', __args__, opts=opts, typ=ConcatResult).value

    return AwaitableConcatResult(
        result=__ret__.result)


@_utilities.lift_output_func(concat)
def concat_output(input: Optional[pulumi.Input[Sequence[Sequence[Any]]]] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ConcatResult]:
    """
    Combines two or more lists into a single list.
    """
    ...