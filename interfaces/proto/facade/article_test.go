package facade

import (
	"bytes"
	"encoding/json"
	"testing"
)

func BenchmarkJsonMarshal(b *testing.B) {
	b.ReportAllocs()
	data := make([]int8, 1024*1024)
	_, _ = json.Marshal(data)
}

func BenchmarkJsonEncoder(b *testing.B) {
	b.ReportAllocs()
	data := make([]int8, 1024*1024)
	buf := bytes.NewBuffer(nil)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(data)
}
