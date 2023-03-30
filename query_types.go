package prometheus

import (
	"net/url"
	"time"
)

// Query - interface for query types
type Query interface {
	values() url.Values
}

// InstantQuery - query for a single point in time
type InstantQuery struct {
	QueryExpression string
	Timeout         time.Duration
	Time            time.Time
}

func (q InstantQuery) values() url.Values {
	if q.Time.IsZero() {
		q.Time = time.Now()
	}
	if q.Timeout == 0 {
		q.Timeout = 10 * time.Second
	}

	data := url.Values{}
	data.Set("query", q.QueryExpression)
	data.Set("time", q.Time.Format(time.RFC3339))
	data.Set("timeout", q.Timeout.String())
	return data
}

// RangeQuery - query for a range of time
type RangeQuery struct {
	QueryExpression string
	Timeout         time.Duration
	Start           time.Time
	End             time.Time
	Step            time.Duration
}

func (q RangeQuery) values() url.Values {
	if q.Timeout == 0 {
		q.Timeout = 10 * time.Second
	}

	data := url.Values{}
	data.Set("query", q.QueryExpression)
	data.Set("start", q.Start.Format(time.RFC3339))
	data.Set("end", q.End.Format(time.RFC3339))
	data.Set("step", q.Step.String())
	data.Set("timeout", q.Timeout.String())
	return data
}
