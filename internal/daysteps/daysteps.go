package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(data string) error {
	// TODO: реализовать функцию
	fields := strings.Split(data, ",")
	if len(fields) != 2 {
		return errors.New("ожидаются шаги и длительность прогулки")
	}

	steps, err := strconv.Atoi(fields[0])
	if err != nil {
		return err
	}
	duration, err := time.ParseDuration(fields[1])
	if err != nil {
		return err
	}
	if steps <= 0 || duration <= 0 {
		return errors.New("шаги и длительность должны быть больше нуля")
	}

	ds.Steps = steps
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories), nil
}
