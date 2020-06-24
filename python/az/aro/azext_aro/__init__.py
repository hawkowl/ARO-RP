# Copyright (c) Microsoft Corporation.
# Licensed under the Apache License 2.0.

import yaml
import pathlib

from ._client_factory import cf_aro
from ._params import load_arguments
from .commands import load_command_table
from azure.cli.core import AzCommandsLoader
from azure.cli.core.commands import CliCommandType
from knack.help_files import helps

_MODULE = __name__
_HERE = pathlib.Path(__file__).resolve().parent


class AroCommandsLoader(AzCommandsLoader):
    def __init__(self, cli_ctx=None):
        aro_custom = CliCommandType(operations_tmpl=_MODULE + '.custom#{}',
                                    client_factory=cf_aro)
        super(AroCommandsLoader, self).__init__(cli_ctx=cli_ctx,
                                                custom_command_type=aro_custom)

    def load_command_table(self, args):
        load_command_table(self, args)
        return self.command_table

    def load_arguments(self, command):
        load_arguments(self, command)


with _HERE.joinpath("help.yaml").open('r') as f:
    helps.update(yaml.safe_load(f))


COMMAND_LOADER_CLS = AroCommandsLoader
