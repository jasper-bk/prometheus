// Copyright 2021 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package textparse

import (
	"fmt"
	"io"

	"github.com/matttproud/golang_protobuf_extensions/pbutil"

	"github.com/prometheus/prometheus/pkg/exemplar"
	"github.com/prometheus/prometheus/pkg/histogram"
	"github.com/prometheus/prometheus/pkg/labels"

	dto "github.com/prometheus/prometheus/prompb/io/prometheus/client"
)

// ProtobufParser is a very inefficient way of unmarshaling the old Prometheus
// protobuf format and then present it as it if were parsed by a
// Prometheus-2-style text parser. This is only done so that we can easily plug
// in the protobuf format into Prometheus 2. For future use (with the final
// format that will be used for sparse histograms), we probably need to rewrite
// the whole parsing part again.
type ProtobufParser struct {
}

func NewProtobufParser(b []byte) Parser {
	for {
		mf := &dto.MetricFamily{}
		if _, err = pbutil.ReadDelimited(resp.Body, mf); err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("reading metric family protocol buffer failed: %v", err)
		}
		ch <- mf
	}
	return &ProtobufParser{}
}

// Series returns the bytes of a series with a simple float64 as a
// value, the timestamp if set, and the value of the current sample.
func (p *ProtobufParser) Series() ([]byte, *int64, float64) {
	// TODO
	return nil, nil, 0
}

// Histogram returns the bytes of a series with a sparse histogram as a
// value, the timestamp if set, and the sparse histogram in the current
// sample.
func (p *ProtobufParser) Histogram() ([]byte, *int64, histogram.SparseHistogram) {
	// TODO
	return nil, nil, histogram.SparseHistogram{}
}

// Help returns the metric name and help text in the current entry.
// Must only be called after Next returned a help entry.
// The returned byte slices become invalid after the next call to Next.
func (p *ProtobufParser) Help() ([]byte, []byte) {
	// TODO
	return nil, nil
}

// Type returns the metric name and type in the current entry.
// Must only be called after Next returned a type entry.
// The returned byte slices become invalid after the next call to Next.
func (p *ProtobufParser) Type() ([]byte, MetricType) {
	// TODO
	return nil, ""
}

// Unit returns the metric name and unit in the current entry.
// Must only be called after Next returned a unit entry.
// The returned byte slices become invalid after the next call to Next.
func (p *ProtobufParser) Unit() ([]byte, []byte) {
	// TODO
	return nil, nil
}

// Comment returns the text of the current comment.
// Must only be called after Next returned a comment entry.
// The returned byte slice becomes invalid after the next call to Next.
func (p *ProtobufParser) Comment() []byte {
	// TODO
	return nil
}

// Metric writes the labels of the current sample into the passed labels.
// It returns the string from which the metric was parsed.
func (p *ProtobufParser) Metric(l *labels.Labels) string {
	// TODO
	return ""
}

// Exemplar always returns false because exemplars aren't supported yet by the
// protobuf format.
func (p *ProtobufParser) Exemplar(l *exemplar.Exemplar) bool {
	return false
}

// Next advances the parser to the next "sample" (emulating the behavior of a
// text format parser). It returns io.EOF if no samples were read.
func (p *ProtobufParser) Next() (Entry, error) {
	// TODO
	return EntryInvalid, io.EOF
}
