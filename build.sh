#!/usr/bin/env bash

set -e
set -o pipefail

projects=$(find proto/cloud -name buf.gen.yaml -type f | xargs dirname | xargs basename)
currPwd=${PWD}
workspace=$(realpath proto/cloud)

buildProjects=$1

buildProject() {
  template=$1
  projPath=$2

  echo "buf generate --template ${template} --path ${projPath}"
  cd "${workspace}"
  buf generate --template "${template}" --path "${projPath}"
  cd "${currPwd}"
}

for proj in ${projects}; do
  template="${proj}/buf.gen.yaml"
  projPath="${proj}"
  if [[ ${#buildProjects} -gt 0 ]]; then
    echo "${buildProjects}" | egrep "(^|,)${proj}(,|$)" || continue
    buildProject "${template}" "${projPath}"

  else
    buildProject "${template}" "${projPath}"
  fi
done