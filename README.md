This is a demostrator for https://github.com/golang/go/issues/69625 and https://github.com/jfreymuth/pulse/pull/40.

It uses jfreymuth/pulse and my own patrix and digimodes libraries to play a sentece in
morse code. The `time` package is not used directly, but maybe somewhere in the libraries.

The panic occurs after a random time, usually in under one minute. Please be patient.

- https://github.com/jfreymuth/pulse
- https://github.com/ftl/digimodes
- https://github.com/ftl/patrix

## Stacktrace

This is the full stacktrace of one execution.

```
unexpected fault address 0x0
fatal error: fault
[signal SIGSEGV: segmentation violation code=0x80 addr=0x0 pc=0x456652]

goroutine 102 gp=0xc000203500 m=10 mp=0xc000181c08 [running]:
runtime.throw({0x532f15?, 0x0?})
	/usr/lib/go/src/runtime/panic.go:1067 +0x48 fp=0xc0000ebc30 sp=0xc0000ebc00 pc=0x46b4c8
runtime.sigpanic()
	/usr/lib/go/src/runtime/signal_unix.go:931 +0x26c fp=0xc0000ebc90 sp=0xc0000ebc30 pc=0x46ccec
runtime.(*timer).maybeRunChan(0xc0001af620?)
	/usr/lib/go/src/runtime/time.go:1300 +0x12 fp=0xc0000ebcc0 sp=0xc0000ebc90 pc=0x456652
runtime.selectgo(0xc0000ebe20, 0xc0000ebe18, 0x0?, 0x0, 0x0?, 0x0)
	/usr/lib/go/src/runtime/select.go:177 +0x13f fp=0xc0000ebde8 sp=0xc0000ebcc0 pc=0x449cff
github.com/ftl/digimodes/cw.(*Modulator).nextAction(0x53e558, 0x4025f0b17e4b78dd)
	/home/florian/go/pkg/mod/github.com/ftl/digimodes@v0.1.2/cw/mod.go:191 +0x6b fp=0xc0000ebe70 sp=0xc0000ebde8 pc=0x4ba04b
github.com/ftl/digimodes/cw.(*Modulator).Modulate(0x53e558, 0x4025f0b17e4b78dd, 0xc00003a008?, 0xc0001af728?, 0x0)
	/home/florian/go/pkg/mod/github.com/ftl/digimodes@v0.1.2/cw/mod.go:178 +0x76 fp=0xc0000ebe98 sp=0xc0000ebe70 pc=0x4b9f36
github.com/ftl/patrix/osc.(*Oscillator).Synth32(0xc0001aa000, {0xc0002d6000, 0x800, 0x0?})
	/home/florian/go/pkg/mod/github.com/ftl/patrix@v0.2.0/osc/oscillator.go:51 +0xfb fp=0xc0000ebed8 sp=0xc0000ebe98 pc=0x4f04fb
github.com/ftl/patrix/osc.(*Oscillator).Synth32-fm({0xc0002d6000?, 0xc000203500?, 0x4e89bd?})
	<autogenerated>:1 +0x31 fp=0xc0000ebf08 sp=0xc0000ebed8 pc=0x4f0c51
github.com/jfreymuth/pulse.Float32Reader.Read(0xc0002185b0?, {0xc0002d6000?, 0xc000306000?, 0x2000?})
	/home/florian/go/pkg/mod/github.com/jfreymuth/pulse@v0.1.1/format.go:91 +0x5b fp=0xc0000ebf60 sp=0xc0000ebf08 pc=0x4ef47b
github.com/jfreymuth/pulse.(*PlaybackStream).run(0xc000203180)
	/home/florian/go/pkg/mod/github.com/jfreymuth/pulse@v0.1.1/playback.go:88 +0x12b fp=0xc0000ebfc8 sp=0xc0000ebf60 pc=0x4efbcb
github.com/jfreymuth/pulse.(*Client).NewPlayback.gowrap1()
	/home/florian/go/pkg/mod/github.com/jfreymuth/pulse@v0.1.1/playback.go:77 +0x25 fp=0xc0000ebfe0 sp=0xc0000ebfc8 pc=0x4efa65
runtime.goexit({})
	/usr/lib/go/src/runtime/asm_amd64.s:1700 +0x1 fp=0xc0000ebfe8 sp=0xc0000ebfe0 pc=0x4728e1
created by github.com/jfreymuth/pulse.(*Client).NewPlayback in goroutine 1
	/home/florian/go/pkg/mod/github.com/jfreymuth/pulse@v0.1.1/playback.go:77 +0x436

goroutine 1 gp=0xc0000061c0 m=nil [select]:
runtime.gopark(0xc0000f7dd0?, 0x2?, 0xe0?, 0x7b?, 0xc0000f7d44?)
	/usr/lib/go/src/runtime/proc.go:424 +0xce fp=0xc0000f7be0 sp=0xc0000f7bc0 pc=0x46b5ee
runtime.selectgo(0xc0000f7dd0, 0xc0000f7d40, 0xc00001e800?, 0x0, 0x40?, 0x1)
	/usr/lib/go/src/runtime/select.go:335 +0x7a5 fp=0xc0000f7d08 sp=0xc0000f7be0 pc=0x44a365
github.com/jfreymuth/pulse/proto.(*Client).Request(0xc0002185b0, {0x5647e8, 0xc000204d98}, {0x0, 0x0})
	/home/florian/go/pkg/mod/github.com/jfreymuth/pulse@v0.1.1/proto/client.go:95 +0x465 fp=0xc0000f7e18 sp=0xc0000f7d08 pc=0x4e8765
github.com/jfreymuth/pulse.(*PlaybackStream).Drain(...)
	/home/florian/go/pkg/mod/github.com/jfreymuth/pulse@v0.1.1/playback.go:152
github.com/ftl/patrix/pa.(*Oscillator).Stop(0xc00019a168, {0xc000112240?, 0xc0000f7ea0?})
	/home/florian/go/pkg/mod/github.com/ftl/patrix@v0.2.0/pa/pa.go:56 +0x8f fp=0xc0000f7e58 sp=0xc0000f7e18 pc=0x4f0baf
main.play({0x5650a0, 0xc0000a6000}, {0x538e43, 0x1f})
	/home/florian/repo/faultdemo/main.go:43 +0x205 fp=0xc0000f7ec0 sp=0xc0000f7e58 pc=0x4f1045
main.main()
	/home/florian/repo/faultdemo/main.go:22 +0x115 fp=0xc0000f7f50 sp=0xc0000f7ec0 pc=0x4f0db5
runtime.main()
	/usr/lib/go/src/runtime/proc.go:272 +0x28b fp=0xc0000f7fe0 sp=0xc0000f7f50 pc=0x43848b
runtime.goexit({})
	/usr/lib/go/src/runtime/asm_amd64.s:1700 +0x1 fp=0xc0000f7fe8 sp=0xc0000f7fe0 pc=0x4728e1

goroutine 2 gp=0xc000006c40 m=nil [force gc (idle), 1 minutes]:
runtime.gopark(0x0?, 0x0?, 0x0?, 0x0?, 0x0?)
	/usr/lib/go/src/runtime/proc.go:424 +0xce fp=0xc00005afa8 sp=0xc00005af88 pc=0x46b5ee
runtime.goparkunlock(...)
	/usr/lib/go/src/runtime/proc.go:430
runtime.forcegchelper()
	/usr/lib/go/src/runtime/proc.go:337 +0xb3 fp=0xc00005afe0 sp=0xc00005afa8 pc=0x4387d3
runtime.goexit({})
	/usr/lib/go/src/runtime/asm_amd64.s:1700 +0x1 fp=0xc00005afe8 sp=0xc00005afe0 pc=0x4728e1
created by runtime.init.7 in goroutine 1
	/usr/lib/go/src/runtime/proc.go:325 +0x1a

goroutine 3 gp=0xc000007180 m=nil [GC sweep wait]:
runtime.gopark(0x1?, 0x0?, 0x0?, 0x0?, 0x0?)
	/usr/lib/go/src/runtime/proc.go:424 +0xce fp=0xc00005b780 sp=0xc00005b760 pc=0x46b5ee
runtime.goparkunlock(...)
	/usr/lib/go/src/runtime/proc.go:430
runtime.bgsweep(0xc000088000)
	/usr/lib/go/src/runtime/mgcsweep.go:317 +0xdf fp=0xc00005b7c8 sp=0xc00005b780 pc=0x4234ff
runtime.gcenable.gowrap1()
	/usr/lib/go/src/runtime/mgc.go:204 +0x25 fp=0xc00005b7e0 sp=0xc00005b7c8 pc=0x417dc5
runtime.goexit({})
	/usr/lib/go/src/runtime/asm_amd64.s:1700 +0x1 fp=0xc00005b7e8 sp=0xc00005b7e0 pc=0x4728e1
created by runtime.gcenable in goroutine 1
	/usr/lib/go/src/runtime/mgc.go:204 +0x66

goroutine 4 gp=0xc000007340 m=nil [GC scavenge wait]:
runtime.gopark(0x10000?, 0x562c78?, 0x0?, 0x0?, 0x0?)
	/usr/lib/go/src/runtime/proc.go:424 +0xce fp=0xc00005bf78 sp=0xc00005bf58 pc=0x46b5ee
runtime.goparkunlock(...)
	/usr/lib/go/src/runtime/proc.go:430
runtime.(*scavengerState).park(0x628220)
	/usr/lib/go/src/runtime/mgcscavenge.go:425 +0x49 fp=0xc00005bfa8 sp=0xc00005bf78 pc=0x420ee9
runtime.bgscavenge(0xc000088000)
	/usr/lib/go/src/runtime/mgcscavenge.go:658 +0x59 fp=0xc00005bfc8 sp=0xc00005bfa8 pc=0x421479
runtime.gcenable.gowrap2()
	/usr/lib/go/src/runtime/mgc.go:205 +0x25 fp=0xc00005bfe0 sp=0xc00005bfc8 pc=0x417d65
runtime.goexit({})
	/usr/lib/go/src/runtime/asm_amd64.s:1700 +0x1 fp=0xc00005bfe8 sp=0xc00005bfe0 pc=0x4728e1
created by runtime.gcenable in goroutine 1
	/usr/lib/go/src/runtime/mgc.go:205 +0xa5

goroutine 5 gp=0xc000007c00 m=nil [finalizer wait, 1 minutes]:
runtime.gopark(0x0?, 0x0?, 0x0?, 0x0?, 0x0?)
	/usr/lib/go/src/runtime/proc.go:424 +0xce fp=0xc00005c620 sp=0xc00005c600 pc=0x46b5ee
runtime.runfinq()
	/usr/lib/go/src/runtime/mfinal.go:193 +0x107 fp=0xc00005c7e0 sp=0xc00005c620 pc=0x416e47
runtime.goexit({})
	/usr/lib/go/src/runtime/asm_amd64.s:1700 +0x1 fp=0xc00005c7e8 sp=0xc00005c7e0 pc=0x4728e1
created by runtime.createfing in goroutine 1
	/usr/lib/go/src/runtime/mfinal.go:163 +0x3d

goroutine 6 gp=0xc000007dc0 m=nil [chan receive]:
runtime.gopark(0x53e180?, 0xc00001a1b0?, 0xe?, 0x7?, 0x2?)
	/usr/lib/go/src/runtime/proc.go:424 +0xce fp=0xc00005a718 sp=0xc00005a6f8 pc=0x46b5ee
runtime.chanrecv(0xc00001a150, 0x0, 0x1)
	/usr/lib/go/src/runtime/chan.go:639 +0x41c fp=0xc00005a790 sp=0xc00005a718 pc=0x407e3c
runtime.chanrecv1(0x438620?, 0xc00005a776?)
	/usr/lib/go/src/runtime/chan.go:489 +0x12 fp=0xc00005a7b8 sp=0xc00005a790 pc=0x4079f2
runtime.unique_runtime_registerUniqueMapCleanup.func1(...)
	/usr/lib/go/src/runtime/mgc.go:1781
runtime.unique_runtime_registerUniqueMapCleanup.gowrap1()
	/usr/lib/go/src/runtime/mgc.go:1784 +0x2f fp=0xc00005a7e0 sp=0xc00005a7b8 pc=0x41abef
runtime.goexit({})
	/usr/lib/go/src/runtime/asm_amd64.s:1700 +0x1 fp=0xc00005a7e8 sp=0xc00005a7e0 pc=0x4728e1
created by unique.runtime_registerUniqueMapCleanup in goroutine 1
	/usr/lib/go/src/runtime/mgc.go:1779 +0x96

goroutine 7 gp=0xc0000ac000 m=nil [select, 1 minutes, locked to thread]:
runtime.gopark(0xc00005cfa8?, 0x2?, 0xa8?, 0xce?, 0xc00005cf94?)
	/usr/lib/go/src/runtime/proc.go:424 +0xce fp=0xc00005ce30 sp=0xc00005ce10 pc=0x46b5ee
runtime.selectgo(0xc00005cfa8, 0xc00005cf90, 0x0?, 0x0, 0x0?, 0x1)
	/usr/lib/go/src/runtime/select.go:335 +0x7a5 fp=0xc00005cf58 sp=0xc00005ce30 pc=0x44a365
runtime.ensureSigM.func1()
	/usr/lib/go/src/runtime/signal_unix.go:1077 +0x19f fp=0xc00005cfe0 sp=0xc00005cf58 pc=0x463b1f
runtime.goexit({})
	/usr/lib/go/src/runtime/asm_amd64.s:1700 +0x1 fp=0xc00005cfe8 sp=0xc00005cfe0 pc=0x4728e1
created by runtime.ensureSigM in goroutine 1
	/usr/lib/go/src/runtime/signal_unix.go:1060 +0xc8

goroutine 8 gp=0xc0000ac1c0 m=5 mp=0xc000100008 [syscall, 1 minutes]:
runtime.notetsleepg(0x648400, 0xffffffffffffffff)
	/usr/lib/go/src/runtime/lock_futex.go:246 +0x29 fp=0xc00005d7a0 sp=0xc00005d778 pc=0x40d509
os/signal.signal_recv()
	/usr/lib/go/src/runtime/sigqueue.go:152 +0x29 fp=0xc00005d7c0 sp=0xc00005d7a0 pc=0x46ce89
os/signal.loop()
	/usr/lib/go/src/os/signal/signal_unix.go:23 +0x13 fp=0xc00005d7e0 sp=0xc00005d7c0 pc=0x4b6393
runtime.goexit({})
	/usr/lib/go/src/runtime/asm_amd64.s:1700 +0x1 fp=0xc00005d7e8 sp=0xc00005d7e0 pc=0x4728e1
created by os/signal.Notify.func1.1 in goroutine 1
	/usr/lib/go/src/os/signal/signal.go:151 +0x1f

goroutine 18 gp=0xc000202000 m=nil [chan receive, 1 minutes]:
runtime.gopark(0x0?, 0x0?, 0x0?, 0x0?, 0x0?)
	/usr/lib/go/src/runtime/proc.go:424 +0xce fp=0xc0000566e0 sp=0xc0000566c0 pc=0x46b5ee
runtime.chanrecv(0xc0000a8000, 0x0, 0x1)
	/usr/lib/go/src/runtime/chan.go:639 +0x41c fp=0xc000056758 sp=0xc0000566e0 pc=0x407e3c
runtime.chanrecv1(0x0?, 0x0?)
	/usr/lib/go/src/runtime/chan.go:489 +0x12 fp=0xc000056780 sp=0xc000056758 pc=0x4079f2
main.handleCancelation(0xc0000a8000, 0xc0000160b0)
	/home/florian/repo/faultdemo/main.go:52 +0x35 fp=0xc0000567c0 sp=0xc000056780 pc=0x4f10f5
main.main.gowrap1()
	/home/florian/repo/faultdemo/main.go:19 +0x25 fp=0xc0000567e0 sp=0xc0000567c0 pc=0x4f0e05
runtime.goexit({})
	/usr/lib/go/src/runtime/asm_amd64.s:1700 +0x1 fp=0xc0000567e8 sp=0xc0000567e0 pc=0x4728e1
created by main.main in goroutine 1
	/home/florian/repo/faultdemo/main.go:19 +0xfa

goroutine 33 gp=0xc000202380 m=nil [GC worker (idle), 1 minutes]:
runtime.gopark(0x0?, 0x0?, 0x0?, 0x0?, 0x0?)
	/usr/lib/go/src/runtime/proc.go:424 +0xce fp=0xc0001b0738 sp=0xc0001b0718 pc=0x46b5ee
runtime.gcBgMarkWorker(0xc00020aa10)
	/usr/lib/go/src/runtime/mgc.go:1412 +0xe9 fp=0xc0001b07c8 sp=0xc0001b0738 pc=0x419f09
runtime.gcBgMarkStartWorkers.gowrap1()
	/usr/lib/go/src/runtime/mgc.go:1328 +0x25 fp=0xc0001b07e0 sp=0xc0001b07c8 pc=0x419de5
runtime.goexit({})
	/usr/lib/go/src/runtime/asm_amd64.s:1700 +0x1 fp=0xc0001b07e8 sp=0xc0001b07e0 pc=0x4728e1
created by runtime.gcBgMarkStartWorkers in goroutine 1
	/usr/lib/go/src/runtime/mgc.go:1328 +0x105

goroutine 32 gp=0xc0000ac8c0 m=nil [GC worker (idle), 1 minutes]:
runtime.gopark(0x46e9b9?, 0xc00019f890?, 0x90?, 0x77?, 0x48d9e7?)
	/usr/lib/go/src/runtime/proc.go:424 +0xce fp=0xc000057738 sp=0xc000057718 pc=0x46b5ee
runtime.gcBgMarkWorker(0xc00020aa10)
	/usr/lib/go/src/runtime/mgc.go:1412 +0xe9 fp=0xc0000577c8 sp=0xc000057738 pc=0x419f09
runtime.gcBgMarkStartWorkers.gowrap1()
	/usr/lib/go/src/runtime/mgc.go:1328 +0x25 fp=0xc0000577e0 sp=0xc0000577c8 pc=0x419de5
runtime.goexit({})
	/usr/lib/go/src/runtime/asm_amd64.s:1700 +0x1 fp=0xc0000577e8 sp=0xc0000577e0 pc=0x4728e1
created by runtime.gcBgMarkStartWorkers in goroutine 1
	/usr/lib/go/src/runtime/mgc.go:1328 +0x105

goroutine 100 gp=0xc0000aca80 m=nil [select]:
runtime.gopark(0xc0001aefb0?, 0x2?, 0x0?, 0x0?, 0xc0001aefa4?)
	/usr/lib/go/src/runtime/proc.go:424 +0xce fp=0xc0001aee48 sp=0xc0001aee28 pc=0x46b5ee
runtime.selectgo(0xc0001aefb0, 0xc0001aefa0, 0x0?, 0x0, 0x0?, 0x1)
	/usr/lib/go/src/runtime/select.go:335 +0x7a5 fp=0xc0001aef70 sp=0xc0001aee48 pc=0x44a365
github.com/ftl/digimodes/cw.(*Modulator).AbortWhenDone.func1()
	/home/florian/go/pkg/mod/github.com/ftl/digimodes@v0.1.2/cw/mod.go:74 +0x53 fp=0xc0001aefe0 sp=0xc0001aef70 pc=0x4b9673
runtime.goexit({})
	/usr/lib/go/src/runtime/asm_amd64.s:1700 +0x1 fp=0xc0001aefe8 sp=0xc0001aefe0 pc=0x4728e1
created by github.com/ftl/digimodes/cw.(*Modulator).AbortWhenDone in goroutine 1
	/home/florian/go/pkg/mod/github.com/ftl/digimodes@v0.1.2/cw/mod.go:73 +0x67

goroutine 50 gp=0xc000202540 m=nil [GC worker (idle), 1 minutes]:
runtime.gopark(0x46e9b9?, 0xc0002044e0?, 0x90?, 0x7f?, 0x48d9e7?)
	/usr/lib/go/src/runtime/proc.go:424 +0xce fp=0xc000057f38 sp=0xc000057f18 pc=0x46b5ee
runtime.gcBgMarkWorker(0xc00020aa10)
	/usr/lib/go/src/runtime/mgc.go:1412 +0xe9 fp=0xc000057fc8 sp=0xc000057f38 pc=0x419f09
runtime.gcBgMarkStartWorkers.gowrap1()
	/usr/lib/go/src/runtime/mgc.go:1328 +0x25 fp=0xc000057fe0 sp=0xc000057fc8 pc=0x419de5
runtime.goexit({})
	/usr/lib/go/src/runtime/asm_amd64.s:1700 +0x1 fp=0xc000057fe8 sp=0xc000057fe0 pc=0x4728e1
created by runtime.gcBgMarkStartWorkers in goroutine 1
	/usr/lib/go/src/runtime/mgc.go:1328 +0x105

goroutine 31 gp=0xc000184540 m=nil [GC worker (idle), 1 minutes]:
runtime.gopark(0x0?, 0x0?, 0x0?, 0x0?, 0x0?)
	/usr/lib/go/src/runtime/proc.go:424 +0xce fp=0xc000058f38 sp=0xc000058f18 pc=0x46b5ee
runtime.gcBgMarkWorker(0xc00020aa10)
	/usr/lib/go/src/runtime/mgc.go:1412 +0xe9 fp=0xc000058fc8 sp=0xc000058f38 pc=0x419f09
runtime.gcBgMarkStartWorkers.gowrap1()
	/usr/lib/go/src/runtime/mgc.go:1328 +0x25 fp=0xc000058fe0 sp=0xc000058fc8 pc=0x419de5
runtime.goexit({})
	/usr/lib/go/src/runtime/asm_amd64.s:1700 +0x1 fp=0xc000058fe8 sp=0xc000058fe0 pc=0x4728e1
created by runtime.gcBgMarkStartWorkers in goroutine 1
	/usr/lib/go/src/runtime/mgc.go:1328 +0x105

goroutine 51 gp=0xc0002028c0 m=nil [GC worker (idle), 1 minutes]:
runtime.gopark(0x0?, 0x0?, 0x0?, 0x0?, 0x50a220?)
	/usr/lib/go/src/runtime/proc.go:424 +0xce fp=0xc0001b0f38 sp=0xc0001b0f18 pc=0x46b5ee
runtime.gcBgMarkWorker(0xc00020aa10)
	/usr/lib/go/src/runtime/mgc.go:1412 +0xe9 fp=0xc0001b0fc8 sp=0xc0001b0f38 pc=0x419f09
runtime.gcBgMarkStartWorkers.gowrap1()
	/usr/lib/go/src/runtime/mgc.go:1328 +0x25 fp=0xc0001b0fe0 sp=0xc0001b0fc8 pc=0x419de5
runtime.goexit({})
	/usr/lib/go/src/runtime/asm_amd64.s:1700 +0x1 fp=0xc0001b0fe8 sp=0xc0001b0fe0 pc=0x4728e1
created by runtime.gcBgMarkStartWorkers in goroutine 1
	/usr/lib/go/src/runtime/mgc.go:1328 +0x105

goroutine 52 gp=0xc000202a80 m=nil [GC worker (idle), 1 minutes]:
runtime.gopark(0x0?, 0x0?, 0x0?, 0x0?, 0x0?)
	/usr/lib/go/src/runtime/proc.go:424 +0xce fp=0xc000059738 sp=0xc000059718 pc=0x46b5ee
runtime.gcBgMarkWorker(0xc00020aa10)
	/usr/lib/go/src/runtime/mgc.go:1412 +0xe9 fp=0xc0000597c8 sp=0xc000059738 pc=0x419f09
runtime.gcBgMarkStartWorkers.gowrap1()
	/usr/lib/go/src/runtime/mgc.go:1328 +0x25 fp=0xc0000597e0 sp=0xc0000597c8 pc=0x419de5
runtime.goexit({})
	/usr/lib/go/src/runtime/asm_amd64.s:1700 +0x1 fp=0xc0000597e8 sp=0xc0000597e0 pc=0x4728e1
created by runtime.gcBgMarkStartWorkers in goroutine 1
	/usr/lib/go/src/runtime/mgc.go:1328 +0x105

goroutine 53 gp=0xc000202c40 m=nil [GC worker (idle), 1 minutes]:
runtime.gopark(0x629cedda5281?, 0x1?, 0x27?, 0x2b?, 0x0?)
	/usr/lib/go/src/runtime/proc.go:424 +0xce fp=0xc000059f38 sp=0xc000059f18 pc=0x46b5ee
runtime.gcBgMarkWorker(0xc00020aa10)
	/usr/lib/go/src/runtime/mgc.go:1412 +0xe9 fp=0xc000059fc8 sp=0xc000059f38 pc=0x419f09
runtime.gcBgMarkStartWorkers.gowrap1()
	/usr/lib/go/src/runtime/mgc.go:1328 +0x25 fp=0xc000059fe0 sp=0xc000059fc8 pc=0x419de5
runtime.goexit({})
	/usr/lib/go/src/runtime/asm_amd64.s:1700 +0x1 fp=0xc000059fe8 sp=0xc000059fe0 pc=0x4728e1
created by runtime.gcBgMarkStartWorkers in goroutine 1
	/usr/lib/go/src/runtime/mgc.go:1328 +0x105

goroutine 54 gp=0xc000202e00 m=nil [GC worker (idle)]:
runtime.gopark(0x629cedda5d65?, 0x0?, 0x0?, 0x0?, 0x0?)
	/usr/lib/go/src/runtime/proc.go:424 +0xce fp=0xc000272738 sp=0xc000272718 pc=0x46b5ee
runtime.gcBgMarkWorker(0xc00020aa10)
	/usr/lib/go/src/runtime/mgc.go:1412 +0xe9 fp=0xc0002727c8 sp=0xc000272738 pc=0x419f09
runtime.gcBgMarkStartWorkers.gowrap1()
	/usr/lib/go/src/runtime/mgc.go:1328 +0x25 fp=0xc0002727e0 sp=0xc0002727c8 pc=0x419de5
runtime.goexit({})
	/usr/lib/go/src/runtime/asm_amd64.s:1700 +0x1 fp=0xc0002727e8 sp=0xc0002727e0 pc=0x4728e1
created by runtime.gcBgMarkStartWorkers in goroutine 1
	/usr/lib/go/src/runtime/mgc.go:1328 +0x105

goroutine 101 gp=0xc000202700 m=nil [IO wait]:
runtime.gopark(0xc0001c21ac?, 0xc00006ced8?, 0x65?, 0xb8?, 0xb?)
	/usr/lib/go/src/runtime/proc.go:424 +0xce fp=0xc00006cca0 sp=0xc00006cc80 pc=0x46b5ee
runtime.netpollblock(0x47ecb8?, 0x405306?, 0x0?)
	/usr/lib/go/src/runtime/netpoll.go:575 +0xf7 fp=0xc00006ccd8 sp=0xc00006cca0 pc=0x431037
internal/poll.runtime_pollWait(0x7a5bc5b80e40, 0x72)
	/usr/lib/go/src/runtime/netpoll.go:351 +0x85 fp=0xc00006ccf8 sp=0xc00006ccd8 pc=0x46a8e5
internal/poll.(*pollDesc).wait(0xc0001ba400?, 0xc000209689?, 0x0)
	/usr/lib/go/src/internal/poll/fd_poll_runtime.go:84 +0x27 fp=0xc00006cd20 sp=0xc00006ccf8 pc=0x4a4c87
internal/poll.(*pollDesc).waitRead(...)
	/usr/lib/go/src/internal/poll/fd_poll_runtime.go:89
internal/poll.(*FD).Read(0xc0001ba400, {0xc000209689, 0x177, 0x177})
	/usr/lib/go/src/internal/poll/fd_unix.go:165 +0x27a fp=0xc00006cdb8 sp=0xc00006cd20 pc=0x4a55ba
net.(*netFD).Read(0xc0001ba400, {0xc000209689?, 0xc00006ceb0?, 0x2000?})
	/usr/lib/go/src/net/fd_posix.go:55 +0x25 fp=0xc00006ce00 sp=0xc00006cdb8 pc=0x4d4125
net.(*conn).Read(0xc000002028, {0xc000209689?, 0x517f20?, 0xc00006ce01?})
	/usr/lib/go/src/net/net.go:189 +0x45 fp=0xc00006ce48 sp=0xc00006ce00 pc=0x4dc3a5
net.(*UnixConn).Read(0x6?, {0xc000209689?, 0x0?, 0x566980?})
	<autogenerated>:1 +0x25 fp=0xc00006ce78 sp=0xc00006ce48 pc=0x4e7205
github.com/jfreymuth/pulse/proto.(*ProtocolReader).fill(0xc0002185b0, 0x4)
	/home/florian/go/pkg/mod/github.com/jfreymuth/pulse@v0.1.1/proto/reader.go:39 +0x143 fp=0xc00006cec0 sp=0xc00006ce78 pc=0x4eae03
github.com/jfreymuth/pulse/proto.(*ProtocolReader).uint32(0xc0002185b0)
	/home/florian/go/pkg/mod/github.com/jfreymuth/pulse@v0.1.1/proto/reader.go:76 +0x25 fp=0xc00006cee8 sp=0xc00006cec0 pc=0x4eb045
github.com/jfreymuth/pulse/proto.(*Client).readLoop(0xc0002185b0)
	/home/florian/go/pkg/mod/github.com/jfreymuth/pulse@v0.1.1/proto/client.go:122 +0x29 fp=0xc00006cfc8 sp=0xc00006cee8 pc=0x4e8a69
github.com/jfreymuth/pulse/proto.(*Client).Open.gowrap1()
	/home/florian/go/pkg/mod/github.com/jfreymuth/pulse@v0.1.1/proto/client.go:54 +0x25 fp=0xc00006cfe0 sp=0xc00006cfc8 pc=0x4e82c5
runtime.goexit({})
	/usr/lib/go/src/runtime/asm_amd64.s:1700 +0x1 fp=0xc00006cfe8 sp=0xc00006cfe0 pc=0x4728e1
created by github.com/jfreymuth/pulse/proto.(*Client).Open in goroutine 1
	/home/florian/go/pkg/mod/github.com/jfreymuth/pulse@v0.1.1/proto/client.go:54 +0x125
```
