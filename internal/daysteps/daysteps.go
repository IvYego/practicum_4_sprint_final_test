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
	if len(parts) != 2 {
		log.Println("неверный формат строки")
		return 0, 0, errors.New("неверный формат данных шагов и времени")
	} else if data == "" {
		log.Println("неверный формат строки")
		return 0, 0, errors.New("пустая строка")
	}

	steps, err := strconv.Atoi(parts[0])
	if steps <= 0 || err != nil {
		log.Println(err)
		return 0, 0, errors.New("недостаточно шагов")
	}

	duration, err := time.ParseDuration(parts[1])
	if duration <= 0 || err != nil {
		log.Println(err)
		return 0, 0, errors.New("неверный формат времени")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {

	steps, duration, err := parsePackage(data)
	if err != nil {
		return ""
	}

	distKm := (float64(steps) * stepLength) / mInKm
	amountCalories, _ := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps, distKm, amountCalories)

	return result

}
