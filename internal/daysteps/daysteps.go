package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-4-sprint-final/internal/spentcalories"
)

var (
	StepLength = 0.65 // длина шага в метрах
)

func parsePackage(data string) (int, time.Duration, error) {
	// ваш код ниже
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("Elements must be 2")
	}
	numOfSteps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("Conversion error: %w", err)
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("Conversion error: %w", err)
	}

	return numOfSteps, duration, nil

}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {
	// ваш код ниже
	numOfSteps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if numOfSteps <= 0 {
		return ""
	}

	distanceInMetres := StepLength * float64(numOfSteps)
	distanceInKilometres := distanceInMetres / 1000.00
	walkingCalories := spentcalories.WalkingSpentCalories(numOfSteps, weight, height, duration)

	dayActionInfo := fmt.Sprintf("Количество шагов: %d.\nДистанция составила: %.2f км.\nВы сожгли  %.2f ккал.", numOfSteps, distanceInKilometres, walkingCalories)
	return dayActionInfo
}
