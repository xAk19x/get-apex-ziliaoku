package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "1.4" )

func hgO5Y20DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wMb89roPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xwKcDQ74Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OUJhnnk7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p0BiWs4BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uE6W1fWfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IGPiEzkRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GFSgBxf8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8dY7biohWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZtKRYEFFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q5hswsBuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BwTBHrljWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nk9aYUtnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S43gMzGqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0F4qrGJIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tzenzwVcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3cK6jopyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IZtOhNHPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K8HhhRNaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3rxSNS1lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hNsKfNYrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vshZVid8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qeI1rPXvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TGy5GOzgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OVbTwjc7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func edZQTSaTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mCkAx4ByWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BWsdEowWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gvGIhvSxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HPRlTfR0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func inFLYD1SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fq2EeJ1ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1uA0a3hFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kZEonf2UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r9IPXaNPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qDNSpaCSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dwLiE2jxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4JrmSA7YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 43FjMmenWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 63HA6IwBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nI1bsOhsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KnLSK9t0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E0IU6kDyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pgTjc09yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tnCptUTxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 87dfxxyUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f5gV46OtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E8SkrDeVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j5759CS1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Z7DgX8oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vLpDSROTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hFLmCDvUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iMjp8KlYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func INKLythtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u8BqKBmgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UZjTP8zVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qTLILGowWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LbyCHEC4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pyz5qY8gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VXh17xVOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YOC2OUj5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0OOGgYPIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XwXuSbFAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func chtGpTC7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jaisqYOJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m5mbuVe9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zvWMV3BBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VuIN2zPAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func taVtY4nwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L7aCHC8MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UnQsTPWRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kYZ5FFv7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 059ZWX6IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EjA5Z7tmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AA1OcMwqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5PDg9tjvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JDCFEVsSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MFcaKWfMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m32gzOJyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BXfnJNaiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fbE6YqzFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IDMUOICdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PEZFGcNBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ok0IHiIgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R1AdpRkHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ucOvxy8lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y9SwMkNeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s7674PskWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ic1nhuvOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lGLopWHGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5tPc7LgHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E8rETYTBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ljCdCEGdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func knqmYQfSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9pJOrUAUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6HrAX7G0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wd2G68tjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mSAOPmpaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XiFL0cd2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lzWQo2RoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VL5kLVWdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ubR07shEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aDMZz6qSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bQCmnXPTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W0abasacWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wQO28zISWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MCwfgOfzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5XOc8obOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cQZLhXD3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NN0OSgCJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zetkTjzcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mZNZ83WUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3cxelbbtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V9EPJLQBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nwUU4fzmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p8ehLO4LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VUo4ePddWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QpITeR3VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PqqmsYHWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O2M7fuiIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 15GxFOx0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H6w6M0FQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CPmpDSnTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kwj8uRRAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UlJTrPp3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CNGwdlmbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vQeR0XDXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lN3SqcFAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2PEVbGLrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AoBdQYgVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tgWxLtjlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4H62k9BPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UNVxi4pkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qSSVtncYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Hs9DGT1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HeMAEeEtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VpwmC9nPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZMCflA3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0DfBm6NXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fGjGjVJIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t3CXzDDFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JmhJPZcrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5urlornzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zXA2QQvIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RtC1gru5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5K5WR0xnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lj4Mi3GYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BTtvv0yxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QGWy30dqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QAgA7fwKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1rvJxpwBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B9hegP1ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V8M8ujudWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ip2tK7hsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K8GFLo2wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 83sxDjFLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5j01sjIoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TPPti0jCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5M6idqGNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3rGO6M03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8qKrVu0eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RZMonMgiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kikyYVklWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func me0nJcjCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Fbhg2OpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vUZfOAEGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9aEXzQv8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func siu9I3bSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C5tHvF8PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func buJoWgopWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yiNg81t0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JUu62sJJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X89UGbfqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7KmZCIGSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9WV1XZ0HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SQKM0IFMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mA9kOCarWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lfULOr2VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ewzxZzOKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oy9yFzHtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XYM3fBF1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hDneV3N5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cEvKoWTnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4E0gf7cIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AIWl04zYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oMCeEqMLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IpCxaxq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2o94aM5xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aXOkpHNEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l81QpvOMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r7p3FuqrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mHEVXr0FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func miEqb4YFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mZg8f5GgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bE2JBLFoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gh1rhMs8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ujpvReyKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PTSMnaPlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pguoTYt2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qii7DIxdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LYueXKKqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h4Owg4JKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mdAyQJNfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FpfimeRlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cU5cJAQ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XCGfHWlYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 65Tts67sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Um8C1eDcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kAuyDWQ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GEFDRh2VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NXjvsZKaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ai76juPjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NRRXGPXKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gFFsyTkYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8yQLVDGmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Yy8HP3CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IwZDPw5nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tSd3tOKcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MUTZJIQuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jfoYHZRMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YZUe0Sm5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HQPohAamWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AaMB9wpTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func smjC3zI5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fn7LLHSyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AfpDCxN8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 51GQ2PPKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zEBFZPWyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y2zSAoVOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cJgzMhOEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c09EyFhLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nn3652ZPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i0yrocFzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yoH0n71oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P2dDH4HcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5308pPCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fveX9bbeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eborasl3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WSaK93GKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7eXSDZKDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PPDGC7zhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tbBTIXniWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lfhkVDp3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XMDsrFdrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V814woh7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aBPhBigJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func menDCNhlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iBSXfsZLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ce0BfBChWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jq5L5sxbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yrqrYarDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pqkp0DgGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D0BXWA1YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ozYNaAYjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vctX1aNWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B3PgqQCqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CRR4Nz8DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kFDEdzHnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RIYt0Wl8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BJOc8t0uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m9sdiuwrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SI6FWDW0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gMhN9QwhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fYESX8tHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RvTzarxOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eLgSQw8jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JszfIaZMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k7UQfLxSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xYfXAVD5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U3OK9BcKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RLRIqikzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sw6sXA3VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JfxgeggTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kkiiM26cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KH36xXbHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ihg7un8MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EamcfFMCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q5qaWUcCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GNWztqXlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fu6oFjh4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tTSxuBNhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DdtKYD3cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p12AeDwuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LJ74sDkSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gwz3M6d1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L7aRnVk8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1mREnl8eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1wXgpVkXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VGzD5DSaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uRtc6j3XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WIJpDwAtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hTJJHjY0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FLCwbNkMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e1ZujfQoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jWGucdjfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vH9DHDkJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LMHcy9x3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C3mIaMuxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9JZv29WGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o1BfXmQIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x774QbXSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FzeCqEnGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9RYj89diWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uuh0bfTAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f78ptLxYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YgwMCaSSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7fEtRg22Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CtaLZRqNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vhlkCSSiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func agvwxHFoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cJBb0nMkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C58m9iVrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i9yv68p8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BK1XMLZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PAGAanPGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tysVHR63Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KNrabhaWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SuvCRnq0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V5muWU7zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eeMrOOzVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func of5NFTKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zJI1S4r4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tE18Icv8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GNHCgb5wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N0RmSkHzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LcMkNFKjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ZKqUKSZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kQX2OKUlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lv2i4m03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xTXz6r90Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KuFj7i0kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lrYLbfBbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 79GiZYCaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WakxgGl6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8FkrtdoSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dTvwpzAFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1u3Rx5ZfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DnGzJdhAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 98F8hcq6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 21bJkkUaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HakS5J0IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1kYCju5QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Asc2MYHPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1KmKbKUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SP0LVf3iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JWEVlRQtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tWlA11nvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OsfVvswdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HWAeMe3SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2QFRHlWvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 82B2uzUcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y0XFXHToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NmUBUSxSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vOYwMiLSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PvMGOD3KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ts7pJfAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8CqWCE2jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mVpxdBGuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7bd3UGhfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U5rdO0q2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E6eDmdE2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ig3ZuS2tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cf0966oRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9nJaC7fNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JFfehXILWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xaFkkrLFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TVfPwNhpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 16PugFLNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func afs15AplWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jh4cHE5xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E2PCGVHbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HBMmm3gtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IBXOOEbeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JSLOjpwpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MDqvWXFhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S3mbNopRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func umDJ9cvNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ylsqgPjpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HbtMkgZKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pB9qe6EqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2zDy76VzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t2CU4s5BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TheVkg6wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RlhQ8yNUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zc363krbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qS0scXHJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uG2b5bfQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oRVen4nLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5U3AYfHtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2g7jawgjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UFm5CgWLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QaoRxw4AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3AkcUY3GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iAriBnvuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 70uiwR97Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R5MAiWjfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5L7tUpoFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cIyVHbdtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func unUnl5BkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bK5qEDoNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zuzwka8QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 16chw9HvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4DLNarlqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wZ7peJBjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KS8hcmvOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vO6nloxaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bgUkJklRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P5MtqPv2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NsPppb5GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TCGVr5kYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MSuUMvlgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nifeMAlGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dBQ8XWOzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lintC0U9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Br1AG4A7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZLmnNAKWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p8lav4AuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hYKVV6HvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o5iTQbkVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qAyqPVKyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jvkqX8NgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HyGnm4AdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N11q843wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DarF5XuvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HjElJ5nnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZCbPWrqfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jnqpSoHcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func awhqwEWsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iTafQFfzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NFgJaUmFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xgnUTI6xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G9mmKvNVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JqkQIDbeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Obk84qhtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hdMc6pO6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7vBEiUGaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5p1iZXHYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mvyey6OpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L54Bz1e1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lbSFG5j7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MZp7NCezWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e2QzyMDNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SQ3JhMHKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w3W0cnSHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4pYhoCHqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TZPqqI4DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1JMm03BNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UPWx94vSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DXEfAq7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jStzqO0VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jJJll4TjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NlcX6UsSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g27y7ubiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ASSqVbMpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9pWPguSZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QiNZduHZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uD2K3BZCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xCocDqDEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cUWi2XWXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KKeVjOMQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lMkqfFRrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RZVjr7W3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XV87ciCNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g1DtVy8rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aiICwNIQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UCYuyAjEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func elXy5vvDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 38EUrCtwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DrnJhwKIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5cjDvZLhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PzK7eT5sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BCmYtdRFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TGX02WIjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ubfANApRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func suFBVkF3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JwJfiZy0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0lo0CGKpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7nqW9GFrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1dZCFqtEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pELKxVUkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SinVr4S6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZdroB3ItWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bWq0fgJDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kKPiLqEMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VPGuwulpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HvetLn4nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7RoEjv8SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X9XAfRxVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d9RZwJQBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PjbL3ZkOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 28drZrL9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K3y5CczTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func alcqbfjXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OqLrehhqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8BavMbuAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QagFjcEtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ebL1x3WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zsl02rM2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CnZEEuoZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WGLssaagWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WK8PkVu7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dMd6paELWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rgVc049nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o9Iw87FoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gdGRN9kwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZwTm5DdXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nX7ucaO2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7MRdp1URWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iqd8pp8rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ExvLMU1AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cil9PksbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5OI5rBqNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P03nsDdCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jOMiHpY5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ef8sQU37Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZdRC3F6eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Npl2OEzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gDNKfXvRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kx78DTeLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6rolc6HPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pYLdsIi4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func orituu5lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gBKmSMaVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 51pOk4oyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3OwYzurVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NBQXm6MQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Fhe6BYjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ABaz4EA3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MOCu1xwTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J4axGgADWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VVpqmsMYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kgaumfB9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E7ZBMdbSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fl4ooABuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fmiIyvaaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t7C3dSQUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g5Ty2bbwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oEDgds0LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IW1t9HGQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KILSL5xLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 42AJDpEhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IDSm5m3eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PeWVnNrHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YCMazB96Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z9IYIaReWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vZ5uK3izWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GVpGI1GDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZhZs7GHqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ASHa6WUDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MdlkJBeWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uPChqG1uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z8zANeKFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a4HEwCQBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func INFRZ6LpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X1H3C06NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fjSpSslnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fei55F2MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func McLkXVKcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5TMwhkr9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func exbawpslWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u70DFGgAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0JZLe2U6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bQ5gIlEWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e96FCqWwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CRZQhbytWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iiDSUNZ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y5AssJCuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ewbCRIE6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func um5SgysAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8T02EZKfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yjJe2o7OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oUyX5YkoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T4W4iKHZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FZkhjnHDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6UuPaXShWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uYvs37r5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VklGocMSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MhaM6yb6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VsEGPNGMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lAxQYKrKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N4mwxt8CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vlYcKmkcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tPpBVD9WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yziF58k3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FVzDcEB9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EzdnNUyIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SYEvWTqlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VibLmiSfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d9trsBXWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r1MDL3k3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FgTsg48eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3R720h0SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S9pLsuFgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BRjb68u4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tsi2ZbboWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YHTDhyNNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mmvI0dsSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EkNApgIzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QT1LcSdGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HFtuiwbUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RXW46JkuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sV1XogHyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZElWd3Q3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LqMEbwhAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EOrcBXV7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BUJxTG3LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7UieamqHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bgowGyXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oc8dP4bAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qKoRx8l7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fCvruC6XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gHIZfrHxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jmc47VlyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gLOPbvxCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B6SqQQW6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func InmgntXpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PCqShaCbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y9ctWtSOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QpVlgCEOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zvNzFFcxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MEJOH3wWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rDnVJYoHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bll8jNGjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1y3LDempWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MxVnQHJVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tN17AbWLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wWEA5OZTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vueBhheoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vXJebNJRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WXI9OX1VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TV2dvR59Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yPbD07moWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rhZ7ZPhkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q9PQNMpMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nO1doEXLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FRQDqNXKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NPfLnvywWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t0LdM19HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l4kS3fn3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GbfNNldxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 28ML0rJnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oQVR7AM0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 23nH3zg5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dc9g5cyaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8KwIMk0DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func srt0VBCVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LLVWkvPkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aKWouZfCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iv6CouOTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fuIA5bjSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m8g8YtS6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func riDCBCxzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IXKMka87Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nevkk6ANWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hHWmX3yaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 01xYe1ifWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qlrK8EjrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func biikgDY2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XwOWbEHVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yDKKGMREWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uIkWe5zvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hr9legq7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KqgMQl1ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l8tJrqzHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6WduYaKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zQIlJgSvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NcqIuLEEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PEDgGdSiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3US6RWObWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BwrqRSTsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kXweI54hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d5vD9YtyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UzmbxWx7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LgtsCXIeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 55VoziSdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EMRLISePWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i6RnoHRDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iJ2JXbZDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tw07YR2BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xNJTxOQMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LmygYhFHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R7orgxXiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gJGf6dLyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n4JiQ4w7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ddmpdtnhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SWF96eanWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YufMKt1ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bKGD71rCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yh4qrOhtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UvjUd0M5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DJvaYSqCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iEMmpUpFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KvoaRCcfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tXpn3GJSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mkGPOkrBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A80rKr9HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FzPN7Yg8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QahNkQtOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6sCp2OZhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OmNldKFlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2G1izDFDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qbjy9FEGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XqITrPptWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gzD7kwvdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WQ7zhrU5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MShjmks3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6oaPuBOrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U0BGbHJBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dBWcgulGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bjja0Yr3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UQjaeNGJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R7DQuZVCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RUvyOYakWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x0dkGO20Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uy6SBgfEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HkW0RUNWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IdujZJNGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xfbzc95EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Agp6zJWlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func itz08zXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d079LMqMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4pI0Nz6bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iWPyO6v1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QZ9bKSvBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GeFUnku3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5pQJm4BwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jCAq35LGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wbZYIy3fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aws8oEmgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o20W9X8PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mzTQ5bAFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gmq5ZswTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m77iDSoHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wPpdwy9NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5SUjcELvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nuJ6Cj5iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bu5E7NdUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iubzPySJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ttcGCfbNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hs0VtJruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4cUetIbbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3IotsU0QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QiM2yFWRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cO8EsTsaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lcbk0EaWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H2ixZBPkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PLWarZveWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 84h0VkqqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8zTqjNH4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QDHskomYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pNZ4Ug6fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fw1nStImWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V0soUoKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NRYVeuegWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8w2TMKUEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cuZLpTp6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sEUcsJqJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4kGA8l4MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2q1b4zHoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fGsTv5C4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SQS5kP6zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 28AlOBkNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ayQPnIKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CL9LGFIOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dzvpu0b9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aAG0ujcgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tW4ZDpgUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iIpX5OcGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e0uLGV1wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EBtqrseoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I6OSwyCCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pf38Ufs4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pHcRkZnNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func un8VJwvqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LhafcH7yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GZKGkH2hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bwbK474OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6buWal5vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 53TWv51AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pyGdZVYRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jcxl881DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OLllDyzLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2tNN9tgIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uDflIpupWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2z9bqKftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y7eTew5YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rDejpQxEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ztaRybHmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A4KkIyZMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C1wi7aKPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wj0xPwhcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F5RS81YPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k004Uc12Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WTss3sCZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rd9bEfPjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9o3Rq7eTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ViUrSZp4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JoQEAkAHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1LzHn2MtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eYGpbVyZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0HCZ7199Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bJLRv0BkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7gNdrOe7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jg24nfo4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RI6Ck8IpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OVDJlAR6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RLtmkuYOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rdoc4neHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SDJDG2GdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sGr3dOnAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kJHqDkD1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k1OYmIXzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QcBRvDBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lzIuRgLLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ikXinJyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bBRQmvK5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MZoxmhRqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HgZL4cpYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V7OGtRLqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g6kNXdJPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qmzAvFhOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sZZrIRLvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bo7sLZPJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yMvIYkpMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 03Y2DZPwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nUy9u5sIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X6d37CrKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Axy9LzDDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gVQ4ESkTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mUPHvhTbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y8005PoaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XsvbUiwHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zST7SqU6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IuP44RxTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6IdlN8ReWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mYBxdN2YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J73OFq3wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nL3cGxNKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func osuRarTBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ii5bVrBzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QdiHQ88tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u3LiDwrNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KM322NMTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func peYJO68dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lyRcqvo5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vccSy9wOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zqg4wSmUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dfwyuIWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1cDDDJEaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zazqgGJIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func btV3M8qRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IaXTDjjKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TElTonl1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3doTadsuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 61nLM5T0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tVg4o9t1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wgzfMXxBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U5P8MfRjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tqTWNcfvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LZpKnj2HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aS8xlIy6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0kTFTHniWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vm7wgNJzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NIY6CbqQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b801qd5tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vKY9UDjCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PKDwSVqgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KLXHOdaIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xNjIEufEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sgI48bX8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G0gkjgE6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YHoK73hUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xYsJxilWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m9B15aTcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bcoFKJQAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func at6hxEzWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LdeF0MYAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hfphlEZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tGtXsTWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qEWNd8R8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VSNIsZONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pB4UDJj6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RXAEnN8iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FcjaK4myWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1NrcjqPcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9E4AmOyrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lvwY3GmbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OWMsbwjdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zjEETxi8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0DkVToqoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sm2GOczuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EBkysON0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ocFNBP4SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IoTMoPzMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PEXjATsVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8v1aiN1PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FqdF4gjAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GWT5OCnoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TcNIcNhoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QAqwmRLUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tXSV6nwLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HqWdBt5IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T0MYmc9yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cmT0EOXOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nicFE3nQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0i2TIEIbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Mf256fAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nrVGkK5DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TU0El1PaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cc310NvOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K477oXhDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uxgH9ctVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yYBJJob8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bHOoYg6qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mVEkWaubWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eCfAjFhuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aW0p6F6GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q9B54tDiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iERV6d8aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xwxLIJQ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f8D1JWPRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b6nra5suWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BlWajTpJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8AfalQ7ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v6VFhUy1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bK7GXDHdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gbr70dNuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pFeW5m9DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KOIqj394Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2x4CpgfxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bhAA0pDpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SyyAntY1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Wzjp6BeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func elnPxlFkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fMWpjTD9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gAYVPQbsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EFQwFLYbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ghibiCPmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pngpye2jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7NXrsUAkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2QcawvMKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F9DcAzIcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func etr3DkSQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n5pcK2lxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kqSSmMb8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KOdgvvh6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gUjcXlgyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UsBenHpKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EgB4qWWCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o8vlZmTNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MwUV0CB6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AAa5yUnOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func luLkaiqNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2rrrBE8jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mop67mC2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LDKeEOiwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k5RhdcugWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2VpiqjI2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uBso14kiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ISWiB1uCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9tqUZN6eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pDbnhZkpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gOdH8KyDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MSYEiwz2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wiBqSX5oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2EGEwyWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OhUIffuZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tFVqMfpfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DysCtMa7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xgaVuk6WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A8f7hA2TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 42vDWzw7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RkEKTI6XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5gd89jPdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lBC8zrPrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oBaxcc2IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S86sYburWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MPC0UwdqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xIr8Fw2gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func peJNhOV8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oHs1NHRfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gtZ1rwD3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rGRGdWlsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vmZEBRAFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WX3c8LljWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DdzJWnSIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0w52BWjvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ddmrhsl5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MrbpXtKEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ptuVCVhrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ae5ur1GCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1tbwN0GSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vECSXKjAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tFyzvS0wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jn3AQ52DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SNJIQktJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQ3GxXpgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xRcxtAHkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RSwf31naWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A16PhCcHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 86vJijx6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F03Qf3QAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func usGzGYmTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MWzgQubaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xko118l9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oaLUgOJoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EOeVB0E6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9kBVGxwHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f368JvFeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qTpPPqAxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nq8OnE1gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MC5cLuYQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3q4T4HARWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xCT4WddLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6RCaAfzsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r0rd6uFaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QfNJJypMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LYC7mmusWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L98Rw4RlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W2UgA7diWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2MSBSTzbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 11RdH67jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VXCMLcRbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rkoX8zVlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wr2KxfxeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iMGmbF3AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0eufhTAbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bZDFQwvKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ykDGhevWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tKDuVQxdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cQGuHJuEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1z9aVJ4dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n909UatrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J6rNkkCHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VooDpnaSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3HL4GyNrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a5qwRTEIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PAXIWeXnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yXQWvylpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xUX6y8eWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2svrtuY0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yk8tWYZLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ba8O6YI5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VssrF7H9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a2RAAsjyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Y8vG0drWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MdKXuaejWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PtyDiqv5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rS5qGBNWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Hv27pwsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func auMbdKliWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yJ39csdyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JFSgsp8MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7BDplh4lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dFsvv25iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ZKLXPU5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c5YIwiP3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I86CRcQgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bvcD2IIRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XDZWMX4BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CEezBn0ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sN7o11NpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 00ODlGpMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KukLUmbSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

