#!/usr/bin/env bash

cd sync

# core=0
core="0-3"
f="Benchmark"

args="-test.bench=${f}  -test.run=^$"

cpu_cmd="nice -n -20 numactl -C ${core} -m 0"

cmds=(
# "export GODEBUG=memprofilerate=0"
"${cpu_cmd} go test ${args}"
"${cpu_cmd} go test ${args} -test.cpuprofile=cpu.pprof"
"${cpu_cmd} go test ${args} -test.trace=trace.pprof"
"${cpu_cmd} go test ${args} -test.blockprofile=block.pprof"
"${cpu_cmd} go test ${args} -test.mutexprofile=mutex.pprof"
"perf record -C ${core} -F 100 -B -g ${cpu_cmd} go test ${args}"
# "unset GODEBUG"
"${cpu_cmd} go test ${args} -test.cpuprofile=cpu.pprof -test.memprofile=mem.pprof -test.trace=trace.pprof -test.blockprofile=block.pprof -test.mutexprofile=mutex.pprof"
"${cpu_cmd} go test ${args} -test.memprofile=mem.pprof"
)

# benchmark
for ((i = 0; i < ${#cmds[@]}; i++)); do
  echo "${cmds[i]}"
  ${cmds[i]}
done

export fgprofile=fg.pprof
echo "fgprof ${cmds[0]}"
  ${cmds[0]}
unset fgprofile
