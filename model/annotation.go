package model

import (
	"encoding/json"
	"time"
)

// Annotation associates an event that explains latency with a timestamp.
type Annotation struct {
	Timestamp time.Time
	Value     string
}

// MarshalJSON implements custom JSON encoding
func (a *Annotation) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Timestamp int64  `json:"timestamp"`
		Value     string `json:"value"`
	}{
		Timestamp: a.Timestamp.Round(time.Microsecond).UnixNano() / 1e3,
		Value:     a.Value,
	})
}

// UnmarshalJSON implements custom JSON decoding
func (a *Annotation) UnmarshalJSON(b []byte) error {
	type Alias Annotation
	annotation := &struct {
		TimeStamp int64 `json:"timestamp,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(b, &annotation); err != nil {
		return err
	}
	a.Timestamp = time.Unix(0, annotation.TimeStamp*1e3)
	return nil
}
