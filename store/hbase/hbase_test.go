// Copyright 2015 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package hbasekv

import (
	"testing"

	. "github.com/pingcap/check"
)

func TestT(t *testing.T) {
	TestingT(t)
}

var _ = Suite(&testHBaseSuite{})

type testHBaseSuite struct {
}

func (t *testHBaseSuite) TestParsePath(c *C) {
	tbl := []struct {
		dsn    string
		ok     bool
		zks    []string
		oracle string
		table  string
	}{
		{"hbase://z,k,zk/tbl", true, []string{"z", "k", "zk"}, "", "tbl"},
		{"hbase://z:80,k:80/tbl?tso=127.0.0.1:1234", true, []string{"z:80", "k:80"}, "127.0.0.1:1234", "tbl"},
		{"goleveldb://zk/tbl", false, nil, "", ""},
		{"hbase://zk/path/tbl", false, nil, "", ""},
		{"hbase:///zk/tbl", false, nil, "", ""},
	}

	for _, t := range tbl {
		zks, oracle, table, err := parsePath(t.dsn)
		if t.ok {
			c.Assert(err, IsNil, Commentf("dsn=%v", t.dsn))
			c.Assert(zks, DeepEquals, t.zks, Commentf("dsn=%v", t.dsn))
			c.Assert(oracle, Equals, t.oracle, Commentf("dsn=%v", t.dsn))
			c.Assert(table, Equals, t.table, Commentf("dsn=%v", t.dsn))
		} else {
			c.Assert(err, NotNil, Commentf("dsn=%v", t.dsn))
		}
	}
}
