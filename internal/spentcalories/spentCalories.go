package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// ваш код ниже
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, errors.New("Ошибка")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, errors.New("Ошибка")
	}
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, errors.New("Ошибка")
	}
	trainingType := parts[1]

	return steps, trainingType, duration, nil
}

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
func distance(steps int) float64 {
	// ваш код ниже
	return float64(steps) * lenStep / float64(mInKm)
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий(число шагов при ходьбе и беге).
// duration time.Duration — длительность тренировки.
func meanSpeed(steps int, duration time.Duration) float64 {
	// ваш код ниже
	if duration <= 0 {
		return 0
	}
	realDistance := distance(steps)
	durationInHours := duration.Hours()

	return realDistance / durationInHours
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// data string - строка с данными.
// weight, height float64 — вес и рост пользователя.
func TrainingInfo(data string, weight, height float64) string {
	// ваш код ниже
	steps, trainingType, duration, err := parseTraining(data)
	if err != nil {
		return ""
	}
	var trainingInfo string
	switch trainingType {

	case "Ходьба":
		realDistance := distance(steps)
		realMeanSpeed := meanSpeed(steps, duration)
		walkingCalories := WalkingSpentCalories(steps, weight, height, duration)

		trainingInfo = fmt.Sprintf("Тип тренировки: %v\nДлительность: %v ч.\nДистанция: %.2f км.\nСкорость: %.2f км.ч\nСожгли калорий: %.2f", trainingType, duration, realDistance, realMeanSpeed, walkingCalories)
	case "Бег":
		realDistance := distance(steps)
		realMeanSpeed := meanSpeed(steps, duration)
		runningCalories := RunningSpentCalories(steps, weight, duration)

		trainingInfo = fmt.Sprintf("Тип тренировки: %v\nДлительность: %v ч.\nДистанция: %.2f км.\nСкорость: %.2f км.ч\nСожгли калорий: %.2f", trainingType, duration, realDistance, realMeanSpeed, runningCalories)

	default:
		fmt.Println("Неизвестный тип тренировки")
	}
	return trainingInfo
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// duration time.Duration — длительность тренировки.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
	// ваш код здесь
	realMeanSpeed := meanSpeed(steps, duration)

	return ((runningCaloriesMeanSpeedMultiplier * realMeanSpeed) - runningCaloriesMeanSpeedShift*weight)
}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// steps int - количество шагов.
// duration time.Duration — длительность тренировки.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
	// ваш код здесь
	realMeanSpeed := meanSpeed(steps, duration)
	durationInHours := duration.Hours()

	return ((walkingCaloriesWeightMultiplier * weight) + (realMeanSpeed*realMeanSpeed/height)*walkingSpeedHeightMultiplier*durationInHours*minInH)

}
