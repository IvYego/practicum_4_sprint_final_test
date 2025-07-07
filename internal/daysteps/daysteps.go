package daysteps

import (
	"errors"
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {

	parts := strings.SplitN(data, ",", 2)

	if data == "" {
		log.Println("пустая строка")
		return 0, 0, errors.New("invalid format")
	}

	if len(parts) != 2 {
		log.Println("неверный формат строки")
		return 0, 0, errors.New("invalid format")
	}

	steps, err := strconv.Atoi(parts[0])

	if steps <= 0 {
		log.Printf("нулевое или отрицательное значение шагов:%v", err)
		return 0, 0, errors.New("conversion error steps: negative or zero value ")
	}

	if err != nil {
		log.Printf("ошибка при расчёте шагов: %v", err)
		return 0, 0, errors.New("conversion error steps")
	}

	duration, err := time.ParseDuration(parts[1])

	if duration <= 0 {
		log.Printf("нулевое или отрицательное значение времени: %v", err)
		return 0, 0, errors.New("conversion error duration: negative or zero value")
	}

	if err != nil {
		log.Printf("ошибка при расчёте времени: %v", err)
		return 0, 0, errors.New("conversion error duration")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {

	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println("error parsing package:", err)
		return ""
	}

	distKm := (float64(steps) * stepLength) / mInKm
	amountCalories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Printf("ошибка при расчёте калорий: %v", err)
		return ""
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps, distKm, amountCalories)

	return result

}
