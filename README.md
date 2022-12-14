this benchmark bench different profiler type overhead cost.

it is not very accuracy.
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
单核测试。
#### description
使用 github.com/bytedance/sonic 的性能测试代码。

循坏 json marshal。
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
| perf cpu          | 0%    |
| fgprof            | 24%   |

### rpc
使用 github.com/bytedance/kitex-benchmark 中测试 grpc 的代码。

服务端 4 个 core。

客户端 12 个 core，100 个连接，通过 unary 总共发送两千万条 1kb 的消息。

#### result
| Profiler         | Delta |
|------------------|-------|
| cpu              | -10%  |
| mem              | 0%    |
| mutex            | 0%    |
| block            | 3%    |
| trace            | 36%   |
| all              | 30%   |
| perf cpu         | 0%    |
| fgprof           | 6%    |

### sync
使用 go sdk sync 包中的压测代码。

八核测试。
#### chan
##### description
sync/chan_test.go
分别开启 procs 个 producer 和 consumer。

producer 和 consumer 通过 select chan 总共发送 1 亿条 msg。
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
