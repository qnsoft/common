package qn_pool_test

import (
	"os"
	"testing"

	"github.com/qnsoft/common/os/gfpool"
	"github.com/qnsoft/common/os/qn_file"
	"github.com/qnsoft/common/os/qn_time"
	"github.com/qnsoft/common/test/qn_test"
)

func Test_ConcurrentOS(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		path := qn_file.TempDir(qn_time.TimestampNanoStr())
		defer qn_file.Remove(path)
		f1, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		t.Assert(err, nil)
		defer f1.Close()

		f2, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		t.Assert(err, nil)
		defer f2.Close()

		for i := 0; i < 100; i++ {
			_, err = f1.Write([]byte("@1234567890#"))
			t.Assert(err, nil)
		}
		for i := 0; i < 100; i++ {
			_, err = f2.Write([]byte("@1234567890#"))
			t.Assert(err, nil)
		}

		for i := 0; i < 1000; i++ {
			_, err = f1.Write([]byte("@1234567890#"))
			t.Assert(err, nil)
		}
		for i := 0; i < 1000; i++ {
			_, err = f2.Write([]byte("@1234567890#"))
			t.Assert(err, nil)
		}
		t.Assert(qn.str.Count(qn_file.GetContents(path), "@1234567890#"), 2200)
	})

	qn_test.C(t, func(t *qn_test.T) {
		path := qn_file.TempDir(qn_time.TimestampNanoStr())
		defer qn_file.Remove(path)
		f1, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		t.Assert(err, nil)
		defer f1.Close()

		f2, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		t.Assert(err, nil)
		defer f2.Close()

		for i := 0; i < 1000; i++ {
			_, err = f1.Write([]byte("@1234567890#"))
			t.Assert(err, nil)
		}
		for i := 0; i < 1000; i++ {
			_, err = f2.Write([]byte("@1234567890#"))
			t.Assert(err, nil)
		}
		t.Assert(qn.str.Count(qn_file.GetContents(path), "@1234567890#"), 2000)
	})
	qn_test.C(t, func(t *qn_test.T) {
		path := qn_file.TempDir(qn_time.TimestampNanoStr())
		defer qn_file.Remove(path)
		f1, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		t.Assert(err, nil)
		defer f1.Close()

		f2, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		t.Assert(err, nil)
		defer f2.Close()

		s1 := ""
		for i := 0; i < 1000; i++ {
			s1 += "@1234567890#"
		}
		_, err = f2.Write([]byte(s1))
		t.Assert(err, nil)

		s2 := ""
		for i := 0; i < 1000; i++ {
			s2 += "@1234567890#"
		}
		_, err = f2.Write([]byte(s2))
		t.Assert(err, nil)

		t.Assert(qn.str.Count(qn_file.GetContents(path), "@1234567890#"), 2000)
	})
	// DATA RACE
	//qn_test.C(t, func(t *qn_test.T) {
	//	path := qn_file.TempDir(qn_time.TimestampNanoStr())
	//	defer qn_file.Remove(path)
	//	f1, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
	//	t.Assert(err, nil)
	//	defer f1.Close()
	//
	//	f2, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
	//	t.Assert(err, nil)
	//	defer f2.Close()
	//
	//	wg := sync.WaitGroup{}
	//	ch := make(chan struct{})
	//	for i := 0; i < 1000; i++ {
	//		wg.Add(1)
	//		go func() {
	//			defer wg.Done()
	//			<-ch
	//			_, err = f1.Write([]byte("@1234567890#"))
	//			t.Assert(err, nil)
	//		}()
	//	}
	//	for i := 0; i < 1000; i++ {
	//		wg.Add(1)
	//		go func() {
	//			defer wg.Done()
	//			<-ch
	//			_, err = f2.Write([]byte("@1234567890#"))
	//			t.Assert(err, nil)
	//		}()
	//	}
	//	close(ch)
	//	wg.Wait()
	//	t.Assert(qn.str.Count(qn_file.GetContents(path), "@1234567890#"), 2000)
	//})
}

func Test_ConcurrentGFPool(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		path := qn_file.TempDir(qn_time.TimestampNanoStr())
		defer qn_file.Remove(path)
		f1, err := gfpool.Open(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		t.Assert(err, nil)
		defer f1.Close()

		f2, err := gfpool.Open(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
		t.Assert(err, nil)
		defer f2.Close()

		for i := 0; i < 1000; i++ {
			_, err = f1.Write([]byte("@1234567890#"))
			t.Assert(err, nil)
		}
		for i := 0; i < 1000; i++ {
			_, err = f2.Write([]byte("@1234567890#"))
			t.Assert(err, nil)
		}
		t.Assert(qn.str.Count(qn_file.GetContents(path), "@1234567890#"), 2000)
	})
	// DATA RACE
	//qn_test.C(t, func(t *qn_test.T) {
	//	path := qn_file.TempDir(qn_time.TimestampNanoStr())
	//	defer qn_file.Remove(path)
	//	f1, err := gfpool.Open(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
	//	t.Assert(err, nil)
	//	defer f1.Close()
	//
	//	f2, err := gfpool.Open(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
	//	t.Assert(err, nil)
	//	defer f2.Close()
	//
	//	wg := sync.WaitGroup{}
	//	ch := make(chan struct{})
	//	for i := 0; i < 1000; i++ {
	//		wg.Add(1)
	//		go func() {
	//			defer wg.Done()
	//			<-ch
	//			_, err = f1.Write([]byte("@1234567890#"))
	//			t.Assert(err, nil)
	//		}()
	//	}
	//	for i := 0; i < 1000; i++ {
	//		wg.Add(1)
	//		go func() {
	//			defer wg.Done()
	//			<-ch
	//			_, err = f2.Write([]byte("@1234567890#"))
	//			t.Assert(err, nil)
	//		}()
	//	}
	//	close(ch)
	//	wg.Wait()
	//	t.Assert(qn.str.Count(qn_file.GetContents(path), "@1234567890#"), 2000)
	//})
}
