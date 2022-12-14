#!/usr/bin/env bash

if [ -d "sonic" ]; then
  echo "already download sonic"
else
  # add fgpprof
  git clone http://github.com/zdyj3170101136/sonic
fi

cd sonic/encoder

# copy from sonic/run_bench.sh
export SONIC_NO_ASYNC_GC=1

# core=0
# f="BenchmarkEncoder_Generic_StdLib"
# procs=1
core="0"
f="BenchmarkEncoder_Parallel_Generic_StdLib"

args="-test.bench=${f} -test.benchtime=150000x  -test.run=^$"

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

export SONIC_NO_ASYNC_GC=0
