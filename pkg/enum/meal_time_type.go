package enum

type MealTimeType uint8

const (
	MealTimeTypeUnknown        MealTimeType = 0
	MealTimeTypeBreakfast      MealTimeType = 1
	MealTimeTypeMorningSnack   MealTimeType = 2
	MealTimeTypeLunch          MealTimeType = 3
	MealTimeTypeDinner         MealTimeType = 4
	MealTimeTypeAfternoonSnack MealTimeType = 5
)

var (
	MealTimeTypeMap = map[MealTimeType]string{
		MealTimeTypeBreakfast:      "Breakfast",
		MealTimeTypeMorningSnack:   "Morning Snack",
		MealTimeTypeLunch:          "Lunch",
		MealTimeTypeDinner:         "Dinner",
		MealTimeTypeAfternoonSnack: "Afternoon Snack",
	}
)

func (c MealTimeType) String() string {
	return MealTimeTypeMap[c]
}

func ValueOfMealTimeType(value string) MealTimeType {
	for k, v := range MealTimeTypeMap {
		if v == value {
			return k
		}
	}
	return MealTimeTypeUnknown
}

func (c MealTimeType) IsValid() bool {
	_, ok := MealTimeTypeMap[c]
	return ok
}
