// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cassandra

import (
	"github.com/uber/cadence/common/config"
	"github.com/uber/cadence/common/log"
	"github.com/uber/cadence/common/persistence/nosql/nosqlplugin"
	"github.com/uber/cadence/common/persistence/nosql/nosqlplugin/cassandra/gocql"
)

// cdb represents a logical connection to Cassandra database
type cdb struct {
	logger  log.Logger
	client  gocql.Client
	session gocql.Session
	cfg     *config.NoSQL
}

var _ nosqlplugin.DB = (*cdb)(nil)

// newCassandraDBFromSession returns a DB from a session
func newCassandraDBFromSession(cfg *config.NoSQL, session gocql.Session, logger log.Logger) *cdb {
	return &cdb{
		client:  gocql.GetRegisteredClient(),
		session: session,
		logger:  logger,
		cfg:     cfg,
	}
}

func (db *cdb) Close() {
	if db.session != nil {
		db.session.Close()
	}
}

func (db *cdb) PluginName() string {
	return PluginName
}

func (db *cdb) IsNotFoundError(err error) bool {
	return db.client.IsNotFoundError(err)
}

func (db *cdb) IsTimeoutError(err error) bool {
	return db.client.IsTimeoutError(err)
}

func (db *cdb) IsThrottlingError(err error) bool {
	return db.client.IsThrottlingError(err)
}

func (db *cdb) IsDBUnavailableError(err error) bool {
	return db.client.IsDBUnavailableError(err)
}
