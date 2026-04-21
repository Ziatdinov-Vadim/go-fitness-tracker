package spentcalories

import (
	"time"
	"strconv"
	"strings"
	"errors"
	"log"
	"fmt"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	dataSlice := strings.Split(data, ",")
	if len(dataSlice) !=3 {
		return 0, "", 0, errors.New("неверные данные ")
	}
	numberSteps, err := strconv.Atoi(dataSlice[0])
	if err!= nil {
		return 0, "",  0, err
	}
	if numberSteps ==0 {
		 return 0, "", 0, errors.New("")
	}
		timeDistance, err := time.ParseDuration (dataSlice[2]) 
	if err != nil{
		return 0, "", 0, errors.New("")
	}
	return numberSteps, dataSlice[1], timeDistance, nil
}

func distance(steps int, height float64) float64 {
	longStep := height * stepLengthCoefficient 
	numberSteps := float64(steps) * longStep
	dist := numberSteps/float64(mInKm)
	return dist
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	dist := distance(steps, height)
	spead := dist/duration.Hours()
	return spead
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	numberSteps, typeWork, timeDistance, err  := parseTraining(data)
	if err !=nil {
		log.Println(err)
		return "", err
	}
	switch typeWork {
	case "Ходьба":
		dist:= distance (numberSteps, height )
		speed:= meanSpeed (numberSteps, height, timeDistance)
		call, err:= WalkingSpentCalories(numberSteps, weight, height, timeDistance)
		if err !=nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f", 
		typeWork, timeDistance.Hours(), dist, speed, call), nil

	case "Бег":
		dist:= distance (numberSteps, height )
		speed:= meanSpeed (numberSteps, height, timeDistance)
		call, err:= RunningSpentCalories(numberSteps, weight, height, timeDistance)
		if err !=nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция:%.2f км.\n Скорость:%.2f км/ч\nСожгли калорий:%.2f",
		 typeWork, timeDistance.Hours(), dist, speed, call), nil
	default:
		return "", errors.New("неизвестный тип тренировки")
	}


}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps == 0{
		return 0, errors.New("количество шагов 0")
	}
	if weight <= 0 {
		return 0, errors.New("введите корректный вес")
	} 
	if height <= 0 {
		return 0, errors.New ("введите корректный рост")
	}
	if duration <=0 {
		return 0, errors.New("время равно 0 или отицательное")
	}
	everageSpeed := meanSpeed(steps,  height, duration)
	call:= (everageSpeed*weight*duration.Minutes())/float64(minInH) 
	return call, nil

}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps == 0 {
		return 0, errors.New("количество шагов равно 0")
	} 
	if weight <= 0 {
		return 0, errors.New(" введите корректный вес")
	}
	if height <=0 {
		return 0, errors.New("введите корректный рост")
	}
	if duration <=0 {
		return 0, errors.New("время отрицательное или равно 0")
	}
	everageSpeed := meanSpeed(steps,  height, duration)
	call := (everageSpeed*weight*duration.Minutes())/float64(minInH)
  kcall:=call*walkingCaloriesCoefficient
  return kcall, nil
}
