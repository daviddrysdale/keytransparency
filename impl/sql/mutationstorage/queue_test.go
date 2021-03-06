// Copyright 2018 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mutationstorage

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/google/keytransparency/core/api/v1/keytransparency_go_proto"
	_ "github.com/mattn/go-sqlite3"
)

func newForTest(ctx context.Context, t *testing.T, logIDs ...int64) *Mutations {
	m, err := New(newDB(t))
	if err != nil {
		t.Fatalf("Failed to create Mutations: %v", err)
	}
	if err := m.AddLogs(ctx, directoryID, logIDs...); err != nil {
		t.Fatalf("AddLogs(): %v", err)
	}
	return m
}

func TestRandLog(t *testing.T) {
	ctx := context.Background()

	for _, tc := range []struct {
		desc     string
		send     []int64
		wantCode codes.Code
		wantLogs map[int64]bool
	}{
		{desc: "no rows", wantCode: codes.NotFound, wantLogs: map[int64]bool{}},
		{desc: "one row", send: []int64{10}, wantLogs: map[int64]bool{10: true}},
		{desc: "second", send: []int64{1, 2, 3}, wantLogs: map[int64]bool{
			1: true,
			2: true,
			3: true,
		}},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			m := newForTest(ctx, t, tc.send...)
			logs := make(map[int64]bool)
			for i := 0; i < 10*len(tc.wantLogs); i++ {
				logID, err := m.randLog(ctx, directoryID)
				if got, want := status.Code(err), tc.wantCode; got != want {
					t.Errorf("randLog(): %v, want %v", got, want)
				}
				if err != nil {
					break
				}
				logs[logID] = true
			}
			if got, want := logs, tc.wantLogs; !cmp.Equal(got, want) {
				t.Errorf("logs: %v, want %v", got, want)
			}
		})
	}
}

func TestSend(t *testing.T) {
	ctx := context.Background()

	m := newForTest(ctx, t, 1, 2)
	update := []byte("bar")
	ts1 := time.Now()
	ts2 := ts1.Add(time.Duration(1))
	ts3 := ts2.Add(time.Duration(1))

	// Test cases are cumulative. Earlier test caes setup later test cases.
	for _, tc := range []struct {
		desc     string
		ts       time.Time
		wantCode codes.Code
	}{
		// Enforce timestamp uniqueness.
		{desc: "First", ts: ts2},
		{desc: "Second", ts: ts2, wantCode: codes.Aborted},
		// Enforce a monotonically increasing timestamp
		{desc: "Old", ts: ts1, wantCode: codes.Aborted},
		{desc: "New", ts: ts3},
	} {
		err := m.send(ctx, directoryID, 1, update, tc.ts)
		if got, want := status.Code(err), tc.wantCode; got != want {
			t.Errorf("%v: send(): %v, got: %v, want %v", tc.desc, err, got, want)
		}
	}
}

func TestWatermark(t *testing.T) {
	ctx := context.Background()
	logIDs := []int64{1, 2}
	m := newForTest(ctx, t, logIDs...)
	update := []byte("bar")

	startTS := time.Now()
	for ts := startTS; ts.Before(startTS.Add(10)); ts = ts.Add(1) {
		for _, logID := range logIDs {
			if err := m.send(ctx, directoryID, logID, update, ts); err != nil {
				t.Fatalf("m.send(%v): %v", logID, err)
			}
		}
	}

	start := startTS.UnixNano()
	for _, tc := range []struct {
		desc      string
		logID     int64
		start     int64
		batchSize int32
		count     int32
		want      int64
	}{
		{desc: "log1 max", logID: 1, batchSize: 100, want: start + 9, count: 10},
		{desc: "log2 max", logID: 2, batchSize: 100, want: start + 9, count: 10},
		{desc: "batch0", logID: 1, batchSize: 0},
		{desc: "batch0start55", logID: 1, start: 55, batchSize: 0, want: 55},
		{desc: "batch5", logID: 1, batchSize: 5, want: start + 4, count: 5},
		{desc: "start1", logID: 1, start: start + 2, batchSize: 5, want: start + 7, count: 5},
		{desc: "start8", logID: 1, start: start + 8, batchSize: 5, want: start + 9, count: 1},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			count, got, err := m.HighWatermark(ctx, directoryID, tc.logID, tc.start, tc.batchSize)
			if err != nil {
				t.Errorf("highWatermark(): %v", err)
			}
			if got != tc.want {
				t.Errorf("highWatermark(%v) high: %v, want %v", tc.start, got, tc.want)
			}
			if count != tc.count {
				t.Errorf("highWatermark(%v) count: %v, want %v", tc.start, count, tc.count)
			}
		})
	}
}

func TestReadLog(t *testing.T) {
	ctx := context.Background()
	logID := int64(5)
	m := newForTest(ctx, t, logID)
	for i := byte(0); i < 10; i++ {
		entry := &pb.EntryUpdate{Mutation: &pb.Entry{Index: []byte{i}}}
		if err := m.Send(ctx, directoryID, entry); err != nil {
			t.Fatalf("Send(): %v", err)
		}
	}

	for _, tc := range []struct {
		batchSize int32
		count     int
	}{
		{batchSize: 0, count: 0},
		{batchSize: 1, count: 1},
		{batchSize: 1, count: 1},
		{batchSize: 100, count: 10},
	} {
		rows, err := m.ReadLog(ctx, directoryID, logID, 0, time.Now().UnixNano(), tc.batchSize)
		if err != nil {
			t.Fatalf("ReadLog(): %v", err)
		}
		if got, want := len(rows), tc.count; got != want {
			t.Fatalf("ReadLog(): len: %v, want %v", got, want)
		}
		for i, r := range rows {
			if got, want := r.Mutation.GetIndex()[0], byte(i); got != want {
				t.Errorf("ReadLog()[%v]: %v, want %v", i, got, want)
			}
		}
	}
}
