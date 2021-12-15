package testpool
import (
	"time"
	"testing"
	"context"
	"leaktest"
)
/*
   subject : Leaktest Example These tests fail, because they leak a goroutine.   https://github.com/fortytw2/leaktest
*/
// Default "Check" will poll for 5 seconds to check that all
// goroutines are cleaned up
func TestPool(t *testing.T) {
    defer leaktest.Check(t)()

    go func() {
        for {
            time.Sleep(time.Second)
        }
    }()
}

// Helper function to timeout after X duration
func TestPoolTimeout(t *testing.T) {
    defer leaktest.CheckTimeout(t, time.Second)()

    go func() {
        for {
            time.Sleep(time.Second)
        }
    }()
}

// Use Go 1.7+ context.Context for cancellation
func TestPoolContext(t *testing.T) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    defer leaktest.CheckContext(ctx, t)()

    go func() {
        for {
            time.Sleep(time.Second)
        }
    }()
}
// Akira@MBP 04tests % go test -v
// === RUN   TestPool
//     leaktest.go:132: leaktest: timed out checking goroutines
//     leaktest.go:150: leaktest: leaked goroutine: goroutine 19 [sleep]:
//         time.Sleep(0x3b9aca00)
//         	/usr/local/Cellar/go/1.16.6/libexec/src/runtime/time.go:193 +0xd2
//         practice/04tests.TestPool.func1()
//         	/Users/Akira/go/src/practice/04tests/test02_test.go:18 +0x2a
//         created by practice/04tests.TestPool
//         	/Users/Akira/go/src/practice/04tests/test02_test.go:16 +0x85
// --- FAIL: TestPool (5.04s)
// === RUN   TestPoolTimeout
//     leaktest.go:132: leaktest: timed out checking goroutines
//     leaktest.go:150: leaktest: leaked goroutine: goroutine 24 [sleep]:
//         time.Sleep(0x3b9aca00)
//         	/usr/local/Cellar/go/1.16.6/libexec/src/runtime/time.go:193 +0xd2
//         practice/04tests.TestPoolTimeout.func1()
//         	/Users/Akira/go/src/practice/04tests/test02_test.go:29 +0x2a
//         created by practice/04tests.TestPoolTimeout
//         	/Users/Akira/go/src/practice/04tests/test02_test.go:27 +0x7c
// --- FAIL: TestPoolTimeout (1.06s)
// === RUN   TestPoolContext
//     leaktest.go:132: leaktest: timed out checking goroutines
//     leaktest.go:150: leaktest: leaked goroutine: goroutine 49 [sleep]:
//         time.Sleep(0x3b9aca00)
//         	/usr/local/Cellar/go/1.16.6/libexec/src/runtime/time.go:193 +0xd2
//         practice/04tests.TestPoolContext.func1()
//         	/Users/Akira/go/src/practice/04tests/test02_test.go:42 +0x2a
//         created by practice/04tests.TestPoolContext
//         	/Users/Akira/go/src/practice/04tests/test02_test.go:40 +0xb7
// --- FAIL: TestPoolContext (1.03s)
// FAIL
// exit status 1
// FAIL	practice/04tests	7.405s
// Akira@MBP 04tests %
