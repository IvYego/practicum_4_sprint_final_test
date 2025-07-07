package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
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
	// TODO: реализовать функцию
	parts := strings.SplitN(data, ",", 3)
	if len(parts) != 3 {
		return 0, "", 0, errors.New("неверный формат данных шагов, типа тренировки и времени")
	}

	steps, err := strconv.Atoi(parts[0])
	if steps <= 0 || err != nil {
		log.Println("ошибка")
		return 0, "", 0, errors.New("недостаточно шагов")
	}

	duration, err := time.ParseDuration(parts[2])
	if duration <= 0 || err != nil {
		log.Println("ошибка")
		return 0, "", 0, errors.New("неверный формат времени")
	}

	typeTrain := parts[1]

	return steps, typeTrain, duration, nil

}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	distKm := ((height * stepLengthCoefficient) * float64(steps)) / mInKm
	return distKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}

	Dist := distance(steps, height)
	averageSpeed := Dist / duration.Hours()

	return averageSpeed

}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию

	steps, typeTrain, timeTrain, err := parseTraining(data)
	if err != nil {
		log.Println("Ошибка при разборе данных:", err)
		return "", err
	}

	distKm := distance(steps, height)
	averageSpeed := meanSpeed(steps, height, timeTrain)

	var calories float64

	switch typeTrain {
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, timeTrain)
		if err != nil {
			log.Println("Ошибка при подсчёте калорий (Бег):", err)
			return "", err
		}
	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, timeTrain)
		if err != nil {
			log.Println("Ошибка при подсчёте калорий (Ходьба):", err)
			return "", err
		}
	default:
		err := errors.New("неизвестный тип тренировки")
		log.Println(err)
		return "", err
	}

	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		typeTrain, timeTrain.Hours(), distKm, averageSpeed, calories)

	return result, nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		log.Println("ошибка")
		return 0, errors.New("некорректные данные шагов")
	}

	if weight <= 0 {
		log.Println("ошибка")
		return 0, errors.New("некорректные данные веса")
	}

	if height <= 0 {
		log.Println("ошибка")
		return 0, errors.New("некорректные данные роста")
	}

	if duration <= 0 {
		log.Println("ошибка")
		return 0, errors.New("некорректные данные времени")
	}

	averageSpeed := meanSpeed(steps, height, duration)
	amountCalories := (weight * averageSpeed * duration.Minutes()) / minInH

	return amountCalories, nil

}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, errors.New("некорректные данные шагов")
	}

	if weight <= 0 {
		return 0, errors.New("некорректные данные веса")
	}

	if height <= 0 {
		return 0, errors.New("некорректные данные роста")
	}

	if duration <= 0 {
		return 0, errors.New("некорректные данные времени")
	}

	averageSpeed := meanSpeed(steps, height, duration)
	amountCalories := (weight * averageSpeed * duration.Minutes()) / minInH
	caloriesCoefficient := amountCalories * walkingCaloriesCoefficient

	return caloriesCoefficient, nil

}
