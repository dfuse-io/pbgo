#!/bin/bash
# Copyright 2019 dfuse Platform Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Protobuf definitions
PROTO=${PROTO:-"$ROOT/../proto"}

function main() {
  current_dir="`pwd`"
  trap "cd \"$current_dir\"" EXIT
  pushd "$ROOT/pb" &> /dev/null

  generate "dfuse/blockmeta/v1/blockmeta.proto"
  generate "dfuse/bstream/v1/bstream.proto"
  generate "dfuse/graphql/v1/graphql.proto"
  generate "dfuse/headinfo/v1/headinfo.proto"
  generate "dfuse/merger/v1/merger.proto"
  generate "dfuse/search/v1/search.proto"
  generate "grpc/health/v1/health.proto"

  echo "generate.sh - `date` - `whoami`" > $ROOT/last_generate.txt
  echo "dfuse-io/proto revision: `GIT_DIR=$PROTO/.git git rev-parse HEAD`" >> $ROOT/last_generate.txt
}

function generate() {
    protoc -I$PROTO $1 --go_out=plugins=grpc,paths=source_relative:.
}

main "$@"