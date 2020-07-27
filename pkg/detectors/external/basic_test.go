package external

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DACUS1995/anomaly-detector/pkg/dataset"
	"github.com/stretchr/testify/assert"
)

func TestBasicDetector(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := []byte(`{
			"class_id": 0,
			"class_name": "nominal"
		}`)
		w.Write(response)
	}))
	defer ts.Close()

	detector, _ := NewBasicDetector(ts.URL)
	result := detector.Detect(&dataset.SimpleDatapoint{
		Values: []float32{0.75953972, 2.67458352, -2.19266376, -0.68166648, 0.94098628, -0.98737533, -1.58778538, 1.03491316, 0.94223391, -0.47893457, -0.7477086},
	})

	assert.Equal(t, false, result.IsAnomaly)
}
