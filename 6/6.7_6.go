package main

import (
	"fmt"
	"time"
)

// Проверяю, есть ли люди на этаже.
// Если да, то сколько человек.
func checkPeople(peopleOnTheFloor [24]int, currentFloor int) (int, bool) {
	if peopleOnTheFloor[currentFloor-1] > 0 {
		return peopleOnTheFloor[currentFloor-1], true
	} else {
		return 0, false
	}
}

// Проверяю, могу ли взять пассажиров в лифт.
// Если да, то покажу, сколько человек уже есть в лифте.
func checkPickUp(elevatorFullness int, maxElevatorFulness int) (int, bool) {
	if elevatorFullness < maxElevatorFulness {
		return elevatorFullness, true
	} else {
		return 0, false
	}
}

// Люди заходят в лифт
func pickUp(
	peopleOnTheFloor *[24]int,
	elevatorFullness *int,
	countPeople,
	currentFloor,
	maxElevatorFulness int) {

	if countPeople > maxElevatorFulness-*elevatorFullness {
		countPeople = maxElevatorFulness - *elevatorFullness
	}
	*elevatorFullness += countPeople
	peopleOnTheFloor[currentFloor-1] -= countPeople
}

func main() {
	const (
		maxElevatorFulness int           = 5  // Сколько человек помещается в лифт
		elevatorSpeed      time.Duration = 50 // За сколько мс проезжает один этаж
	)
	var (
		peopleOnTheFloor        [24]int
		elevatorCurrentPosition int
		elevatorFullness        int
		movement                string = "up"
		isReturn                bool
	)

	peopleOnTheFloor[3] = 4  // 4 этаж
	peopleOnTheFloor[6] = 3  // 7 этаж
	peopleOnTheFloor[9] = 3  // 10 этаж
	peopleOnTheFloor[16] = 5 // 17 этаж

	fmt.Println("Лифт поехал...")

	for {
		if movement == "up" {
			elevatorCurrentPosition++
			time.Sleep(time.Millisecond * elevatorSpeed)
			fmt.Print(elevatorCurrentPosition, " ")

			if elevatorCurrentPosition == 24 {
				movement = "down"
				fmt.Println("\n\nДостигнув последнего этажа, лифт спускается и по дороге останавливается, чтобы взять пассажира.\n")
			}
		} else {

			if elevatorCurrentPosition == 0 {

				if isReturn {
					movement = "up"
					isReturn = false
					elevatorFullness = 0

					fmt.Println("\n\nПоднимемся ещё раз, чтобы при спуске забрать желающих.\n")
				} else {
					break
				}

			} else {
				time.Sleep(time.Millisecond * elevatorSpeed)
				fmt.Print(elevatorCurrentPosition, " ")

				// Проверяем, есть ли люди на этаже
				if countPeopleOnTheFloor, ok := checkPeople(peopleOnTheFloor, elevatorCurrentPosition); ok {

					// Проверяем, есть ли место в лифте.
					if _, ok := checkPickUp(elevatorFullness, maxElevatorFulness); ok {

						pickUp(
							&peopleOnTheFloor,
							&elevatorFullness,
							countPeopleOnTheFloor,
							elevatorCurrentPosition,
							maxElevatorFulness)

						fmt.Printf("\n-- На этаже №%d есть %d человек.\n-- -- Сейчас в лифте: %d\n-- -- На этаже остались: %d\n",
							elevatorCurrentPosition,
							countPeopleOnTheFloor,
							elevatorFullness,
							peopleOnTheFloor[elevatorCurrentPosition-1])

						// Надо вернуться за пассажирами
						// Если на этаже ещё остались люди, т.е. не все зашли в лифт (не все поместились)
						if peopleOnTheFloor[elevatorCurrentPosition-1] > 0 {
							isReturn = true
						}

					} else {
						fmt.Print("\n-- На этаже есть люди, но в лифте нет места. Заберём в следующий заход.\n")

						// Надо вернуться за пассажирами
						// Попадаю сюда, если на этаже есть люди, а места в лифте вообще нет.
						isReturn = true
					}
				}

				elevatorCurrentPosition--
			}
		}

	}
	fmt.Println()
}
