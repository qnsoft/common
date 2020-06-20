package qn_hash_test

import (
	"testing"

	"github.com/qnsoft/common/encoding/qn_hash"
	"github.com/qnsoft/common/test/qn_test"
)

var (
	strBasic = []byte("This is the test string for hash.")
)

func Test_BKDRHash(t *testing.T) {
	var x uint32 = 200645773
	qn_test.C(t, func(t *qn_test.T) {
		j := qn_hash.BKDRHash(strBasic)
		t.Assert(j, x)
	})
}

func Test_BKDRHash64(t *testing.T) {
	var x uint64 = 4214762819217104013
	qn_test.C(t, func(t *qn_test.T) {
		j := qn_hash.BKDRHash64(strBasic)
		t.Assert(j, x)
	})
}

func Test_SDBMHash(t *testing.T) {
	var x uint32 = 1069170245
	qn_test.C(t, func(t *qn_test.T) {
		j := qn_hash.SDBMHash(strBasic)
		t.Assert(j, x)
	})
}

func Test_SDBMHash64(t *testing.T) {
	var x uint64 = 9881052176572890693
	qn_test.C(t, func(t *qn_test.T) {
		j := qn_hash.SDBMHash64(strBasic)
		t.Assert(j, x)
	})
}

func Test_RSHash(t *testing.T) {
	var x uint32 = 1944033799
	qn_test.C(t, func(t *qn_test.T) {
		j := qn_hash.RSHash(strBasic)
		t.Assert(j, x)
	})
}

func Test_RSHash64(t *testing.T) {
	var x uint64 = 13439708950444349959
	qn_test.C(t, func(t *qn_test.T) {
		j := qn_hash.RSHash64(strBasic)
		t.Assert(j, x)
	})
}

func Test_JSHash(t *testing.T) {
	var x uint32 = 498688898
	qn_test.C(t, func(t *qn_test.T) {
		j := qn_hash.JSHash(strBasic)
		t.Assert(j, x)
	})
}

func Test_JSHash64(t *testing.T) {
	var x uint64 = 13410163655098759877
	qn_test.C(t, func(t *qn_test.T) {
		j := qn_hash.JSHash64(strBasic)
		t.Assert(j, x)
	})
}

func Test_PJWHash(t *testing.T) {
	var x uint32 = 7244206
	qn_test.C(t, func(t *qn_test.T) {
		j := qn_hash.PJWHash(strBasic)
		t.Assert(j, x)
	})
}

func Test_PJWHash64(t *testing.T) {
	var x uint64 = 31150
	qn_test.C(t, func(t *qn_test.T) {
		j := qn_hash.PJWHash64(strBasic)
		t.Assert(j, x)
	})
}

func Test_ELFHash(t *testing.T) {
	var x uint32 = 7244206
	qn_test.C(t, func(t *qn_test.T) {
		j := qn_hash.ELFHash(strBasic)
		t.Assert(j, x)
	})
}

func Test_ELFHash64(t *testing.T) {
	var x uint64 = 31150
	qn_test.C(t, func(t *qn_test.T) {
		j := qn_hash.ELFHash64(strBasic)
		t.Assert(j, x)
	})
}

func Test_DJBHash(t *testing.T) {
	var x uint32 = 959862602
	qn_test.C(t, func(t *qn_test.T) {
		j := qn_hash.DJBHash(strBasic)
		t.Assert(j, x)
	})
}

func Test_DJBHash64(t *testing.T) {
	var x uint64 = 2519720351310960458
	qn_test.C(t, func(t *qn_test.T) {
		j := qn_hash.DJBHash64(strBasic)
		t.Assert(j, x)
	})
}

func Test_APHash(t *testing.T) {
	var x uint32 = 3998202516
	qn_test.C(t, func(t *qn_test.T) {
		j := qn_hash.APHash(strBasic)
		t.Assert(j, x)
	})
}

func Test_APHash64(t *testing.T) {
	var x uint64 = 2531023058543352243
	qn_test.C(t, func(t *qn_test.T) {
		j := qn_hash.APHash64(strBasic)
		t.Assert(j, x)
	})
}
