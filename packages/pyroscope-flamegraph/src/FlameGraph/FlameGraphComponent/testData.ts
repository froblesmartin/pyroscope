import { Units } from '../../../../webapp/javascript/util/format';

const TestData = {
  empty: {
    names: [],
    levels: [],
    numTicks: 0,
    sampleRate: 0,
    units: Units.Samples,
    spyName: '',
    format: 'single' as const,
    version: 0,
  },
  SimpleTree: {
    topLevel: 0,
    rangeMin: 0,
    format: 'single' as const,
    numTicks: 988,
    sampleRate: 100,
    names: [
      'total',
      'runtime.main',
      'main.slowFunction',
      'main.work',
      'main.main',
      'main.fastFunction',
    ],
    levels: [
      [0, 988, 0, 0],
      [0, 988, 0, 1],
      [0, 214, 0, 5, 214, 3, 2, 4, 217, 771, 0, 2],
      [0, 214, 214, 3, 216, 1, 1, 5, 217, 771, 771, 3],
    ],

    rangeMax: 1,
    units: Units.Samples,
    fitMode: 'HEAD',

    spyName: 'gospy',
    version: 1,
  },
  ComplexTree: {
    names: [
      'total',
      'runtime.mcall',
      'runtime.park_m',
      'runtime.schedule',
      'runtime.unlockWithRank',
      'runtime.unlock2',
      'runtime.resetspinning',
      'runtime.wakep',
      'runtime.startm',
      'runtime.notewakeup',
      'runtime.futexwakeup',
      'runtime.futex',
      'runtime.pMask.read',
      'runtime.findrunnable',
      'runtime.stopm',
      'runtime.mPark',
      'runtime.notesleep',
      'runtime.futexsleep',
      'runtime.runqget',
      'runtime.pidleput',
      'runtime.updateTimerPMask',
      'runtime.netpoll',
      'runtime.read',
      'runtime.epollwait',
      'runtime.checkTimers',
      'runtime.runtimer',
      'runtime.runOneTimer',
      'time.sendTime',
      'runtime.selectnbsend',
      'runtime.chansend',
      'runtime.send',
      'runtime.goready',
      'runtime.goready.func1',
      'runtime.ready',
      'runtime.siftdownTimer',
      'runtime.bgscavenge.func1',
      'runtime.wakeScavenger',
      'runtime.injectglist.func1',
      'runtime.gcBgMarkWorker',
      'runtime.systemstack',
      'runtime.gcBgMarkWorker.func2',
      'runtime.gcDrain',
      'runtime.scanobject',
      'runtime.spanOf',
      'runtime.pageIndexOf',
      'runtime.markBits.isMarked',
      'runtime.greyobject',
      'runtime.findObject',
      'runtime.heapBitsForAddr',
      'runtime.heapBits.bits',
      'runtime.(*gcWork).tryGet',
      'runtime.putempty',
      'runtime.(*lfstack).push',
      'runtime.(*lfstack).pop',
      'runtime.(*gcWork).balance',
      'runtime.handoff',
      'runtime.memmove',
      'runtime.(*gcWork).tryGetFast',
      'runtime.bgsweep',
      'runtime.sweepone',
      'runtime.(*mspan).sweep',
      'runtime.(*mheap).freeSpan',
      'runtime.(*mheap).freeSpan.func1',
      'runtime.(*mheap).freeSpanLocked',
      'runtime.(*pageAlloc).free',
      'runtime.bgscavenge',
      'runtime.bgscavenge.func2',
      'runtime.(*pageAlloc).scavenge',
      'runtime.(*pageAlloc).scavengeOne',
      'runtime.(*pageAlloc).scavengeRangeLocked',
      'runtime.sysUnused',
      'runtime.madvise',
      'runtime.resettimer',
      'runtime.modtimer',
      'runtime.wakeNetPoller',
      'runtime.write',
      'runtime.write1',
      'runtime._System',
      'runtime.gogo',
      'github.com/pyroscope-io/pyroscope/pkg/storage/cache.New.func2',
      'github.com/pyroscope-io/pyroscope/pkg/storage/cache.(*Cache).saveToDisk',
      'github.com/pyroscope-io/pyroscope/pkg/storage.treeCodec.Serialize',
      'github.com/pyroscope-io/pyroscope/pkg/storage/tree.(*Tree).SerializeTruncate',
      'github.com/pyroscope-io/pyroscope/pkg/storage/tree.(*Tree).minValue',
      'github.com/pyroscope-io/pyroscope/pkg/storage/tree.(*Tree).iterateWithTotal',
      'github.com/pyroscope-io/pyroscope/pkg/storage/tree.(*Tree).minValue.func1',
      'github.com/pyroscope-io/pyroscope/pkg/structs/cappedarr.(*CappedArray).Push',
      'github.com/pyroscope-io/pyroscope/pkg/storage/dict.(*Dict).Put',
      'github.com/pyroscope-io/pyroscope/pkg/storage/dict.(*trieNode).findNodeAt',
      'github.com/pyroscope-io/pyroscope/pkg/util/varint.Writer.Write',
      'bytes.(*Buffer).tryGrowByReslice',
      'bytes.(*Buffer).Write',
      'bytes.(*Buffer).grow',
      'runtime.makeslice',
      'github.com/pyroscope-io/pyroscope/pkg/storage.segmentCodec.Serialize',
      'github.com/pyroscope-io/pyroscope/pkg/storage/segment.(*Segment).Serialize',
      'runtime.typedslicecopy',
      'runtime.growslice',
      'runtime.mallocgc',
      'runtime.heapBitsSetType',
      'runtime.(*mcache).nextFree',
      'runtime.(*mcache).refill',
      'runtime.(*mcentral).uncacheSpan',
      'runtime.(*spanSet).push',
      'runtime.(*headTailIndex).incTail',
      'runtime.(*mcentral).cacheSpan',
      'runtime.(*mcentral).grow',
      'runtime.(*mheap).alloc',
      'runtime.memclrNoHeapPointers',
      'runtime.getMCache',
      'github.com/valyala/bytebufferpool.(*ByteBuffer).Write',
      'github.com/pyroscope-io/pyroscope/pkg/agent/upstream/direct.(*Direct).uploadLoop',
      'github.com/pyroscope-io/pyroscope/pkg/agent/upstream/direct.(*Direct).safeUpload',
      'github.com/pyroscope-io/pyroscope/pkg/agent/upstream/direct.(*Direct).uploadProfile',
      'github.com/pyroscope-io/pyroscope/pkg/structs/transporttrie.(*Trie).Iterate',
      'github.com/pyroscope-io/pyroscope/pkg/storage.IngestionObserver.Put',
      'github.com/pyroscope-io/pyroscope/pkg/storage.(*Storage).Put',
      'github.com/pyroscope-io/pyroscope/pkg/storage/segment.(*Segment).Put',
      'github.com/pyroscope-io/pyroscope/pkg/storage/segment.(*streeNode).put',
      'github.com/pyroscope-io/pyroscope/pkg/storage/segment.(*Segment).Put.func1',
      'github.com/pyroscope-io/pyroscope/pkg/storage.(*Storage).Put.func1',
      'github.com/pyroscope-io/pyroscope/pkg/storage/tree.(*Tree).Merge',
      'github.com/pyroscope-io/pyroscope/pkg/storage/tree.(*treeNode).insert',
      'sort.Search',
      'bytes.Compare',
      'cmpbody',
      'github.com/pyroscope-io/pyroscope/pkg/agent.(*ProfileSession).takeSnapshots',
      'runtime.selectgo',
      'runtime.sellock',
      'runtime.lockWithRank',
      'runtime.lock2',
      'github.com/pyroscope-io/pyroscope/pkg/agent/gospy.(*GoSpy).Snapshot',
      'runtime/pprof.writeHeap',
      'runtime/pprof.writeHeapInternal',
      'runtime/pprof.writeHeapProto',
      'runtime/pprof.newProfileBuilder',
      'runtime/pprof.(*profileBuilder).readMapping',
      'runtime/pprof.parseProcSelfMaps',
      'runtime/pprof.elfBuildID',
      'os.(*File).pread',
      'internal/poll.(*FD).Pread',
      'syscall.Pread',
      'syscall.Syscall6',
      'runtime/pprof.(*profileBuilder).pbSample',
      'runtime/pprof.(*protobuf).endMessage',
      'runtime/pprof.(*profileBuilder).flush',
      'compress/flate.(*Writer).Write',
      'compress/flate.(*compressor).write',
      'compress/flate.(*compressor).encSpeed',
      'compress/flate.(*huffmanBitWriter).writeBlockDynamic',
      'compress/flate.(*huffmanBitWriter).writeTokens',
      'compress/flate.(*huffmanBitWriter).writeCode',
      'runtime/pprof.(*profileBuilder).appendLocsForStack',
      'runtime/pprof.(*profileBuilder).stringIndex',
      'runtime.mapassign_faststr',
      'runtime.hashGrow',
      'runtime.makeBucketArray',
      'runtime.newarray',
      'runtime/pprof.(*profileBuilder).emitLocation',
      'runtime.mapassign_fast64',
      'runtime.growWork_fast64',
      'runtime.bucketShift',
      'runtime.mapaccess2_fast64',
      'runtime.memhash64',
      'runtime.GC',
      'runtime.osyield',
      'github.com/valyala/bytebufferpool.(*ByteBuffer).WriteString',
      'github.com/pyroscope-io/pyroscope/pkg/convert.(*Profile).Get',
      'github.com/pyroscope-io/pyroscope/pkg/convert.(*Profile).findFunctionName',
      'github.com/pyroscope-io/pyroscope/pkg/convert.(*Profile).findLocation',
      'github.com/pyroscope-io/pyroscope/pkg/convert.(*Profile).findLocation.func1',
      'github.com/pyroscope-io/pyroscope/pkg/convert.(*Profile).findFunction',
      'github.com/pyroscope-io/pyroscope/pkg/convert.(*Profile).findFunction.func1',
      'github.com/pyroscope-io/pyroscope/pkg/agent/gospy.(*GoSpy).Snapshot.func3',
      'github.com/pyroscope-io/pyroscope/pkg/agent.(*ProfileSession).takeSnapshots.func1',
      'github.com/pyroscope-io/pyroscope/pkg/structs/transporttrie.(*Trie).Insert',
      'github.com/pyroscope-io/pyroscope/pkg/structs/transporttrie.(*trieNode).findNodeAt',
      'runtime.newobject',
      'runtime.releasem',
      'github.com/pyroscope-io/pyroscope/pkg/agent/gospy.getHeapProfile',
      'io/ioutil.ReadAll',
      'io.ReadAll',
      'compress/gzip.(*Reader).Read',
      'compress/flate.(*decompressor).Read',
      'compress/flate.(*dictDecoder).availWrite',
      'compress/flate.(*decompressor).huffmanBlock',
      'compress/flate.(*decompressor).huffSym',
      'github.com/pyroscope-io/pyroscope/pkg/convert.ParsePprof',
      'google.golang.org/protobuf/proto.Unmarshal',
      'google.golang.org/protobuf/proto.UnmarshalOptions.unmarshal',
      'google.golang.org/protobuf/internal/impl.(*MessageInfo).unmarshal',
      'google.golang.org/protobuf/internal/impl.(*MessageInfo).unmarshalPointer',
      'google.golang.org/protobuf/internal/impl.pointer.AppendPointerSlice',
      'runtime.(*spanSet).pop',
      'google.golang.org/protobuf/internal/impl.consumeStringSliceValidateUTF8',
      'runtime.slicebytetostring',
      'github.com/pyroscope-io/pyroscope/pkg/agent.(*ProfileSession).reset',
      'github.com/pyroscope-io/pyroscope/pkg/agent.(*ProfileSession).uploadTries',
      'github.com/pyroscope-io/pyroscope/pkg/structs/transporttrie.(*Trie).Diff',
      'github.com/pyroscope-io/pyroscope/pkg/structs/transporttrie.(*Trie).Diff.func1',
      'runtime.nextFreeFast',
      'github.com/pyroscope-io/pyroscope/pkg/agent.(*ProfileSession).isDueForReset',
      'time.Now',
      'runtime.walltime',
      'github.com/dgraph-io/badger/v2.(*levelsController).runCompactor',
      'runtime.gopark',
      'github.com/dgraph-io/badger/v2.(*levelsController).pickCompactLevels',
      'github.com/dgraph-io/badger/v2.(*compactStatus).overlapsWith',
      'github.com/dgraph-io/badger/v2.(*levelsController).isLevel0Compactable',
      'sync.(*RWMutex).RLock',
      'github.com/dgraph-io/badger/v2.(*DB).doWrites.func1',
      'github.com/dgraph-io/badger/v2.(*DB).writeRequests',
      'github.com/dgraph-io/badger/v2.(*valueLog).write',
      'github.com/dgraph-io/badger/v2.(*valueLog).write.func2',
      'github.com/dgraph-io/badger/v2.(*valueLog).write.func1',
      'os.(*File).write',
      'syscall.Write',
      'syscall.write',
      'syscall.Syscall',
      'github.com/dgraph-io/badger/v2.(*logFile).encodeEntry',
      'bytes.makeSlice',
      'runtime.(*mcache).allocLarge',
    ],
    levels: [
      [0, 178, 0, 0],
      [
        0, 2, 0, 210, 2, 3, 0, 204, 5, 39, 0, 126, 44, 3, 0, 111, 47, 15, 0, 79,
        62, 1, 0, 77, 63, 4, 0, 65, 67, 3, 0, 58, 70, 56, 0, 38, 126, 52, 0, 1,
      ],
      [
        0, 2, 0, 211, 2, 1, 0, 208, 3, 1, 0, 206, 4, 1, 0, 127, 5, 1, 0, 201, 6,
        2, 0, 196, 8, 34, 0, 131, 42, 2, 0, 127, 44, 3, 0, 112, 47, 15, 0, 80,
        62, 1, 1, 78, 63, 1, 0, 72, 64, 3, 0, 39, 67, 3, 2, 59, 70, 56, 0, 39,
        126, 52, 0, 2,
      ],
      [
        0, 2, 0, 212, 2, 1, 1, 209, 3, 1, 1, 207, 4, 1, 1, 205, 5, 1, 0, 202, 6,
        2, 0, 197, 8, 5, 0, 179, 13, 17, 0, 167, 30, 1, 0, 166, 31, 4, 0, 164,
        35, 7, 0, 132, 42, 1, 0, 129, 43, 1, 1, 128, 44, 3, 0, 113, 47, 10, 0,
        94, 57, 5, 0, 81, 63, 1, 0, 73, 64, 3, 0, 66, 69, 1, 0, 60, 70, 56, 0,
        40, 126, 52, 0, 3,
      ],
      [
        0, 1, 0, 219, 1, 1, 0, 213, 5, 1, 1, 203, 6, 2, 0, 198, 8, 3, 0, 187,
        11, 2, 0, 180, 13, 7, 0, 173, 20, 10, 1, 168, 30, 1, 1, 56, 31, 4, 1,
        59, 35, 7, 0, 133, 42, 1, 1, 130, 44, 2, 0, 115, 46, 1, 0, 114, 47, 10,
        1, 95, 57, 5, 0, 82, 63, 1, 0, 74, 64, 3, 0, 67, 69, 1, 0, 61, 70, 1, 1,
        57, 71, 55, 6, 41, 126, 46, 7, 13, 172, 1, 1, 12, 173, 4, 1, 6, 177, 1,
        0, 4,
      ],
      [
        0, 1, 0, 91, 1, 1, 0, 214, 6, 2, 0, 176, 8, 3, 0, 188, 11, 2, 0, 181,
        13, 7, 0, 174, 21, 5, 0, 171, 26, 4, 0, 169, 32, 3, 0, 60, 35, 7, 0,
        134, 44, 2, 0, 116, 46, 1, 0, 97, 48, 1, 0, 89, 49, 7, 2, 97, 56, 1, 0,
        96, 57, 4, 0, 87, 61, 1, 0, 83, 63, 1, 0, 75, 64, 3, 0, 68, 69, 1, 0,
        39, 77, 9, 0, 54, 86, 3, 0, 50, 89, 3, 3, 49, 92, 3, 3, 48, 95, 31, 16,
        42, 133, 10, 5, 24, 143, 16, 3, 21, 159, 1, 0, 19, 160, 1, 1, 18, 161,
        11, 1, 14, 174, 3, 0, 7, 177, 1, 1, 5,
      ],
      [
        0, 1, 0, 92, 1, 1, 0, 215, 6, 2, 0, 199, 8, 3, 0, 189, 11, 2, 0, 182,
        13, 7, 1, 175, 21, 5, 3, 123, 26, 4, 2, 123, 32, 3, 0, 61, 35, 3, 0,
        152, 38, 3, 0, 143, 41, 1, 0, 135, 44, 2, 0, 117, 46, 1, 0, 98, 48, 1,
        1, 110, 51, 1, 1, 109, 52, 4, 0, 98, 56, 1, 1, 56, 57, 4, 0, 88, 61, 1,
        0, 84, 63, 1, 1, 76, 64, 3, 0, 69, 69, 1, 0, 62, 77, 9, 0, 55, 86, 1, 1,
        53, 87, 2, 0, 51, 111, 5, 5, 47, 116, 3, 3, 46, 119, 4, 4, 45, 123, 2,
        2, 44, 125, 1, 1, 43, 138, 5, 1, 25, 146, 12, 12, 23, 158, 1, 1, 22,
        159, 1, 1, 20, 162, 10, 0, 15, 174, 3, 0, 8,
      ],
      [
        0, 1, 0, 220, 1, 1, 0, 216, 6, 2, 0, 177, 8, 3, 0, 190, 11, 2, 0, 183,
        14, 6, 4, 176, 24, 2, 2, 172, 28, 2, 2, 170, 32, 3, 0, 39, 35, 1, 0,
        162, 36, 1, 0, 158, 37, 1, 0, 153, 38, 2, 0, 145, 40, 1, 1, 144, 41, 1,
        0, 136, 44, 2, 0, 118, 46, 1, 0, 100, 52, 2, 0, 100, 54, 1, 1, 48, 55,
        1, 1, 99, 57, 3, 1, 89, 60, 1, 1, 56, 61, 1, 0, 85, 64, 3, 0, 70, 69, 1,
        0, 63, 77, 8, 8, 53, 85, 1, 1, 56, 87, 2, 2, 52, 139, 3, 0, 26, 142, 1,
        0, 4, 162, 10, 0, 16, 174, 3, 0, 9,
      ],
      [
        0, 1, 0, 93, 1, 1, 0, 217, 6, 1, 1, 98, 7, 1, 1, 200, 8, 3, 0, 191, 11,
        1, 0, 185, 12, 1, 1, 184, 18, 1, 1, 56, 19, 1, 0, 177, 32, 3, 0, 129,
        35, 1, 1, 163, 36, 1, 0, 159, 37, 1, 0, 154, 38, 2, 0, 146, 41, 1, 0,
        137, 44, 2, 0, 119, 46, 1, 0, 101, 52, 2, 0, 101, 58, 1, 0, 91, 59, 1,
        1, 90, 61, 1, 0, 86, 64, 3, 3, 71, 69, 1, 1, 64, 139, 1, 0, 35, 140, 1,
        1, 34, 141, 1, 0, 27, 142, 1, 1, 5, 162, 10, 1, 17, 174, 3, 0, 10,
      ],
      [
        0, 1, 0, 98, 1, 1, 1, 218, 8, 2, 0, 194, 10, 1, 0, 192, 11, 1, 1, 186,
        19, 1, 1, 178, 32, 3, 1, 130, 36, 1, 0, 160, 37, 1, 0, 155, 38, 2, 0,
        147, 41, 1, 0, 138, 44, 2, 0, 120, 46, 1, 0, 105, 52, 1, 0, 105, 53, 1,
        0, 102, 58, 1, 0, 92, 61, 1, 1, 56, 139, 1, 0, 36, 141, 1, 0, 28, 163,
        9, 9, 11, 174, 3, 3, 11,
      ],
      [
        0, 1, 0, 221, 8, 2, 1, 195, 10, 1, 0, 97, 33, 2, 2, 165, 36, 1, 1, 161,
        37, 1, 0, 156, 38, 2, 0, 148, 41, 1, 0, 139, 44, 2, 0, 121, 46, 1, 0,
        106, 52, 1, 0, 106, 53, 1, 0, 103, 58, 1, 1, 93, 139, 1, 0, 37, 141, 1,
        0, 29,
      ],
      [
        0, 1, 0, 107, 9, 1, 1, 56, 10, 1, 0, 98, 37, 1, 0, 157, 38, 2, 0, 149,
        41, 1, 0, 140, 44, 2, 0, 122, 46, 1, 0, 107, 52, 1, 0, 107, 53, 1, 1,
        104, 139, 1, 0, 8, 141, 1, 0, 30,
      ],
      [
        0, 1, 1, 108, 10, 1, 0, 100, 37, 1, 0, 98, 38, 2, 1, 150, 41, 1, 0, 141,
        44, 2, 0, 123, 46, 1, 1, 108, 52, 1, 1, 108, 139, 1, 0, 9, 141, 1, 0,
        31,
      ],
      [
        10, 1, 0, 101, 37, 1, 1, 99, 39, 1, 1, 151, 41, 1, 1, 142, 44, 2, 0,
        124, 139, 1, 0, 10, 141, 1, 0, 32,
      ],
      [10, 1, 0, 105, 44, 2, 2, 125, 139, 1, 1, 11, 141, 1, 0, 33],
      [10, 1, 1, 193, 141, 1, 0, 7],
      [141, 1, 0, 8],
      [141, 1, 1, 9],
    ],
    numTicks: 178,
    maxSelf: 16,
    spyName: 'gospy',
    sampleRate: 100,
    units: Units.Samples,
    format: 'single' as const,
    version: 1,
  },
  DiffTree: {
    names: [
      'total',
      'runtime.main',
      'main.main',
      'main.becomesAdded',
      'main.becomesSlower',
      'main.work',
      'runtime.asyncPreempt',
      'main.becomesFaster',
      'github.com/pyroscope-io/pyroscope/pkg/agent/upstream/remote.(*Remote).handleJobs',
      'github.com/pyroscope-io/pyroscope/pkg/agent/upstream/remote.(*Remote).safeUpload',
      'net/http.(*Client).Do',
      'net/http.(*Client).do',
      'net/http.(*Client).send',
      'net/http.send',
      'net/http.(*Transport).RoundTrip',
      'net/http.setupRewindBody',
      'runtime.newobject',
      'runtime.mallocgc',
      'runtime.arenaIndex',
      'github.com/pyroscope-io/pyroscope/pkg/agent.(*ProfileSession).takeSnapshots',
      'runtime.heapBitsSetType',
    ],
    levels: [
      [0, 991, 0, 0, 987, 0, 0],
      [0, 0, 0, 0, 1, 0, 19, 0, 0, 0, 1, 1, 0, 8, 0, 991, 0, 2, 985, 0, 1],
      [
        0, 0, 0, 0, 1, 0, 16, 0, 0, 0, 1, 1, 0, 9, 0, 217, 0, 2, 229, 0, 3, 217,
        165, 0, 231, 147, 0, 7, 382, 603, 1, 378, 604, 0, 4, 985, 6, 6, 982, 5,
        4, 2,
      ],
      [
        0, 0, 0, 0, 1, 0, 17, 0, 0, 0, 1, 1, 0, 10, 0, 217, 217, 2, 229, 229, 5,
        217, 165, 165, 231, 147, 147, 5, 383, 602, 601, 378, 604, 604, 5, 991,
        0, 0, 986, 1, 1, 3,
      ],
      [0, 0, 0, 0, 1, 1, 20, 0, 0, 0, 1, 1, 0, 11, 984, 1, 1, 982, 0, 0, 6],
      [0, 0, 0, 1, 1, 0, 12],
      [0, 0, 0, 1, 1, 0, 13],
      [0, 0, 0, 1, 1, 0, 14],
      [0, 0, 0, 1, 1, 0, 15],
      [0, 0, 0, 1, 1, 0, 16],
      [0, 0, 0, 1, 1, 0, 17],
      [0, 0, 0, 1, 1, 1, 18],
    ],
    numTicks: 1978,
    maxSelf: 604,
    spyName: 'gospy',
    sampleRate: 100,
    units: Units.Samples,
    format: 'double' as const,
    leftTicks: 991,
    rightTicks: 987,
    version: 1,
  },
};

export default TestData;
