this benchmark bench different profiler type overhead cost.

it is not very accuracy.
#### USAGE
git clone https://github.com/zdyj3170101136/perfbench.git
cd perfbench && ./bench.sh
#### environment
go version go1.18.4 linux/amd64

machine:
```shell
# os release : 3.10.0-1062.18.1.el7.x86_64
# perf version : 3.10.0-1160.80.1.el7.x86_64.debug
# arch : x86_64
# nrcpus online : 16
# nrcpus avail : 16
# cpudesc : Intel(R) Xeon(R) Platinum 8269CY CPU @ 2.50GHz
# cpuid : GenuineIntel,6,85,7
# total memory : 32 GB
```
### benchmark
#### json
| Profiler          | Delta |
| ----------------- | ----- |
| cpu               | 23%   |
| cpu(4 core)       | 15%   |
| mem               | 0%    |
| mem(4 core)       | 0%    |
| mutex             | 0%    |
| mutex(4 core)     | 0%    |
| block             | 0%    |
| block(4 core)     | 0%    |
| trace             | 16%   |
| trace all(4 core) | 8%    |
| all               | 32%   |
| all(4 core)       | 23%   |
| perf cpu          | < 1%  |
| perf cpu(4 core)  | < 1%  |
| fgprof            | 23%   |
| fgprof            | 29%   |

