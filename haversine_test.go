//Package haversine_test for test
package haversine_test

/*
go test -v
go test -bench .
go test -bench . -cpu 1,2,4
go test -bench . -benchtime 5s
go test -cover
*/
import (
	"github.com/suifengtec/haversine"
	"os"
	"testing"
)

var testsCases = []struct {
	from           haversine.Point
	to             haversine.Point
	expectedMeters float64
	expectedKM     float64
	expectedMile   float64
}{
	{
		haversine.Point{Lat: 22.55, Lon: 43.12},
		haversine.Point{Lat: 13.45, Lon: 100.28},
		6.094544408786774e+06,
		6094.544408786774,
		3786.251258825624,
	},
	{
		haversine.Point{Lat: 51.510357, Lon: -0.116773},
		haversine.Point{Lat: 38.889931, Lon: -77.009003},
		5.897658288856054e+06,
		5897.658288856054,
		3663.9352546369896,
	},
}

func TestDistance(t *testing.T) {
	for _, input := range testsCases {
		m, km, mi := input.from.Distance(input.to)

		if input.expectedMeters != m {
			t.Errorf("FAIL: 从 %v 到 %v 的以m为单位的距离应为: %v ,但计算结果为 %v",
				input.from,
				input.to,
				input.expectedMeters,
				m,
			)
		}
		if input.expectedKM != km {
			t.Errorf("FAIL: 从 %v 到 %v 的以KM为单位的距离应为: %v ,但计算结果为 %v",
				input.from,
				input.to,
				input.expectedKM,
				km,
			)
		}
		if input.expectedMile != mi {
			t.Errorf("FAIL: 从 %v 到 %v 的以Mile为单位的距离应为: %v ,但计算结果为 %v",
				input.from,
				input.to,
				input.expectedMile,
				mi,
			)
		}
	}
}

func BenchmarkDistance(b *testing.B) {

	here := haversine.Point{Lat: 22.55, Lon: 43.12}
	there := haversine.Point{Lat: 13.45, Lon: 100.28}

	for i := 0; i < b.N; i++ {

		here.Distance(there)
	}

}

func heap() []byte {

	return make([]byte, 1024*10)
}
func Benchmarkheap(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {

		_ = heap()
	}

}

func TestMain(m *testing.M) {

	id := m.Run()
	os.Exit(id)

}
