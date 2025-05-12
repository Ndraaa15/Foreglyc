package enum

type GlucoseStatus uint8

const (
	GlucoseStatusUnknown              GlucoseStatus = 0
	GlucoseStatusHypoglycemiaAccute   GlucoseStatus = 1
	GlucoseStatusHypoglycemiaChronic  GlucoseStatus = 2
	GlucoseStatusHyperglycemiaChronic GlucoseStatus = 3
	GlucoseStatusHyperglycemiaAccute  GlucoseStatus = 4
	GlucoseStatusNormal               GlucoseStatus = 5
)

var (
	GlucoseStatusMap = map[GlucoseStatus]string{
		GlucoseStatusHyperglycemiaAccute:  "Hyperglycemia Accute",
		GlucoseStatusHyperglycemiaChronic: "Hyperglycemia Chronic",
		GlucoseStatusHypoglycemiaAccute:   "Hypoglycemia Accute",
		GlucoseStatusHypoglycemiaChronic:  "Hypoglycemia Chronic",
		GlucoseStatusNormal:               "Normal",
	}
)

func (c GlucoseStatus) String() string {
	return GlucoseStatusMap[c]
}

func ValueOfGlucoseStatus(value string) GlucoseStatus {
	for k, v := range GlucoseStatusMap {
		if v == value {
			return k
		}
	}
	return GlucoseStatusUnknown
}

func (c GlucoseStatus) IsValid() bool {
	_, ok := GlucoseStatusMap[c]
	return ok
}
