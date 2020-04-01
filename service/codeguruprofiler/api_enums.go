// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codeguruprofiler

type AggregationPeriod string

// Enum values for AggregationPeriod
const (
	AggregationPeriodP1d  AggregationPeriod = "P1D"
	AggregationPeriodPt1h AggregationPeriod = "PT1H"
	AggregationPeriodPt5m AggregationPeriod = "PT5M"
)

func (enum AggregationPeriod) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AggregationPeriod) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type OrderBy string

// Enum values for OrderBy
const (
	OrderByTimestampAscending  OrderBy = "TimestampAscending"
	OrderByTimestampDescending OrderBy = "TimestampDescending"
)

func (enum OrderBy) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum OrderBy) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
