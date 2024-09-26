package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// Интерфейс для всех животных
type Animal interface {
	Speak() string
	Move() string
	Eat() string
	Sleep() string
}

// Интерфейс для животных, которые умеют плавать
type Swimmer interface {
	CanSwim() bool
}

type Monkey struct{}

func (t Monkey) Speak() string {
	return "Кричит"
}
func (t Monkey) Move() string {
	return "Прыгает по деревьям"
}
func (t Monkey) Eat() string {
	return "Ест бананы"
}
func (t Monkey) Sleep() string {
	return "Спит в вольере"
}
func (t Monkey) CanSwim() bool {
	return false
}
func (m Monkey) Climb() string {
	return "Лазает по деревьям"
}

type Shark struct{}

func (t Shark) Speak() string {
	return "не издаёт звук"
}
func (t Shark) Move() string {
	return "Плавает"
}
func (t Shark) Sleep() string {
	return "Спит на дне"
}
func (t Shark) Eat() string {
	return "Съела мясо"
}
func (t Shark) CanSwim() bool {
	return true
}
func (s Shark) Hunt() string {
	return "Охотится на рыбу"
}

type Eagle struct{}

func (t Eagle) Speak() string {
	return "Орёт"
}
func (t Eagle) Move() string {
	return "Летит"
}
func (t Eagle) Eat() string {
	return "Ест мышь"
}
func (t Eagle) Sleep() string {
	return "Спит в гнезде"
}
func (t Eagle) CanSwim() bool {
	return false
}
func (e Eagle) Fly() string {
	return "Летает высоко"
}

type Bear struct{}

func (t Bear) Speak() string {
	return "Рычит"
}
func (t Bear) Move() string {
	return "Бежит"
}
func (t Bear) Eat() string {
	return "Ест малину"
}
func (t Bear) Sleep() string {
	return "Спит в берлоге"
}
func (t Bear) CanSwim() bool {
	return true
}
func (b Bear) Hibernate() string {
	return "Спит зимой"
}

type Whale struct{}

func (t Whale) Speak() string {
	return "Издаёт ултразвук"
}
func (t Whale) Move() string {
	return "Медленно плывёт у поверхности"
}
func (t Whale) Eat() string {
	return "Ест планктон"
}
func (t Whale) Sleep() string {
	return "Спит вертикально"
}
func (t Whale) CanSwim() bool {
	return true
}
func (w Whale) Dive() string {
	return "Ныряет глубоко"
}

// UnknownAnimal структура, которая реализует интерфейс Animal, но генерирует ошибку при вызове метода Speak
type UnknownAnimal struct{}

func (t UnknownAnimal) Speak() string {
	panic("UnknownAnimal cannot speak")
}
func (t UnknownAnimal) Move() string {
	return "Неизвестно"
}
func (t UnknownAnimal) Eat() string {
	return "Неизвестно"
}
func (t UnknownAnimal) Sleep() string {
	return "Неизвестно"
}

// Функция для получения животного по его типу
func getAnimal(animalType string) (Animal, error) {
	switch strings.ToLower(animalType) {
	case "monkey":
		return Monkey{}, nil
	case "shark":
		return Shark{}, nil
	case "eagle":
		return Eagle{}, nil
	case "bear":
		return Bear{}, nil
	case "whale":
		return Whale{}, nil
	default:
		return UnknownAnimal{}, errors.New("unknown animal type")
	}
}

// Функция для логирования ошибок
func logError(err error) {
	logFile, _ := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer logFile.Close()

	logger := log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(err)
}

func main() {
	// Создаем лог-файл для ошибок
	logFile, err := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Не удалось создать файл лога: %v", err)
	}
	defer logFile.Close()

	// Устанавливаем логгер для стандартного вывода ошибок
	log.SetOutput(logFile)

	// Запрашиваем у пользователя тип животного
	fmt.Print("Введите тип животного (monkey, shark, eagle, bear, whale): ")
	var animalType string
	fmt.Scanln(&animalType)

	// Получаем животное по его типу
	animal, err := getAnimal(animalType)
	if err != nil {
		logError(err)
		fmt.Println("Неизвестный тип животного. Попробуйте снова.")
		return
	}

	// Выводим информацию о животном
	fmt.Printf("Животное: %T\n", animal)
	fmt.Printf("Звук: %v\n", animal.Speak())
	fmt.Printf("Движение: %v\n", animal.Move())
	fmt.Printf("Еда: %v\n", animal.Eat())
	fmt.Printf("Сон: %v\n", animal.Sleep())

	// Проверка на умение плавать
	if swimmer, ok := animal.(Swimmer); ok {
		fmt.Printf("Умеет плавать: %v\n", swimmer.CanSwim())
	}

	// Выводим дополнительную информацию в зависимости от типа животного
	switch a := animal.(type) {
	case Monkey:
		fmt.Printf("Лазание: %v\n", a.Climb())
	case Shark:
		fmt.Printf("Охота: %v\n", a.Hunt())
	case Eagle:
		fmt.Printf("Полет: %v\n", a.Fly())
	case Bear:
		fmt.Printf("Зимовка: %v\n", a.Hibernate())
	case Whale:
		fmt.Printf("Ныряние: %v\n", a.Dive())
	case UnknownAnimal:
		fmt.Println("Неизвестное животное")
	}

	fmt.Println("-----------------------------")
}
