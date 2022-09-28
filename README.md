Vpplink govpp apigen plugin
===========================

This repository contains a govpp binapigenerator plugin. It is built as a goplugin, producing a ``.so`` that will be dynamically loaded as part of the api generation process if passing its path to [binapigen](https://github.com/FDio/govpp/tree/master/binapigen).

A typical usecase can be found in ``examples/`` where we build this plugin and inject it in binapigen.
