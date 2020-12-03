package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		// The Pool's New function should generally only return pointer
		// types, since a pointer can be put into the return interface
		// value without an allocation:
		return new(bytes.Buffer)
	},
}

// timeNow is a fake version of time.Now for tests.
func timeNow() time.Time {
	return time.Unix(1136214245, 0)
}

func Log(w io.Writer, key, val string) {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	// Replace this with time.Now() in a real logger.
	b.WriteString(timeNow().UTC().Format(time.RFC3339))
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	w.Write(b.Bytes())
	bufPool.Put(b)
}

var byteBufPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 1024)
	},
}

// 一个临时对象池
// 保存和复用临时对象，减少内存分配，降低GC压力
// 没有就new一个新的，有就直接拿，使用完在put回去
func main() {
	Log(os.Stdout, "path", "/search?q=flowers")

	buf := byteBufPool.Get().([]byte)

	copy(buf, []byte("abcaaaaaa"))

	byteBufPool.Put(buf[:1024])

	buf1 := byteBufPool.Get().([]byte)[:1024]
	fmt.Println("buf1", string(buf1))
	fmt.Println("buf len1", len(buf1[:2]))
}
