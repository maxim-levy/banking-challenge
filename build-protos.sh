#!/usr/bin/env bash
capnp compile -I "$GOPATH/src/server" -ogo protos/*.capnp
capnp compile -I "$GOPATH/src/server" -ogo protos/account/*.capnp
capnp compile -I "$GOPATH/src/server" -ogo protos/transfer/*.capnp
