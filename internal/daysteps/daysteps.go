package daysteps

import (
	"time"
	"fmt"
	"strings"
	"errors"
	"strconv"
	"github.com/Yandex-Practicum/tracker/internal/spentcalories"

	
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	dataSlice := strings.Split(data, ",")
	if len(dataSlice) != 2 {
		return  0, 0, errors.New ("не верный формат - разделить их нельзя")
	} 
	numberSteps, err := strconv.Atoi (dataSlice[0])
	if err!= nil {
		return 0, 0, err
	}
	if numberSteps == 0 {
		return 0, 0, errors.New ("количество шагов равно 0")
	} 
	timeDistance, err := time.ParseDuration (dataSlice[1])
	if err !=nil {
		return 0, 0, err
	}
	return numberSteps, timeDistance, nil
}

func DayActionInfo(data string, weight, height float64) string {
	const mInKm = 1000
	numberSteps, timeDistance, err :=  parsePackage(data)
	if err != nil {
		return ""	 
	}
	if numberSteps == 0 {
		return ""
	}
	distance:=  ((float64 (numberSteps)) * stepLength)/float64(mInKm)
	numberCall, err := spentcalories.WalkingSpentCalories(numberSteps, weight, height, timeDistance)
	if err != nil {
		return "" 
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", numberSteps, distance, numberCall )
}
