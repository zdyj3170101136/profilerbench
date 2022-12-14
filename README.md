this benchmark bench different profiler type overhead cost.

it is not very accuracy.
#### pre benchmark
git clone https://github.com/zdyj3170101136/perfbench.git
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

禁止 turbo boost:
https://askubuntu.com/questions/619875/disabling-intel-turbo-boost-in-ubuntu

performance 模式（不支持）:
https://wsgzao.github.io/post/cpupower/
## benchmark
### json
#### description
使用 github.com/bytedance/sonic 的性能测试代码。

循坏 json marshal。

单核测试。
```go
func BenchmarkEncoder_Generic_StdLib(b *testing.B) {
	_, _ = json.Marshal(_GenericValue)
	b.SetBytes(int64(len(TwitterJson)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(_GenericValue)
	}
}
```
#### result
| Profiler          | Delta |
|-------------------|-------|
| cpu               | 23%   |
| mem               | 0%    |
| mutex             | 0%    |
| block             | 0%    |
| trace             | 14%   |
| all               | 33%   |
| perf cpu          | < 1%  |
| fgprof            | 24%   |

### sync
使用 go sdk sync 包中的压测代码。
#### chan
##### description
sync/chan_test.go
开启 procs 个 producer, consumer。

producer 和 consumer 通过 select chan 传递数据
##### result
| Profiler         | Delta |
|------------------|-------|
| cpu              | 4%    |
| mem              | 5%    |
| mutex            | 6%    |
| block            | 7%    |
| trace            | 9%    |
| all              | 9%    |
| perf cpu         | 6%    |
| fgprof           | 6%    |
