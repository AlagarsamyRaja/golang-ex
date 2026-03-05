package main

import (
	"fmt"
	"time"
)

func checkDate(DateStr string) (time.Time, error) {
	birthDate, err := time.Parse("2006-01-02", DateStr)

	if err != nil {
		fmt.Println("Date must be in given format")
		return time.Time{}, err
	}

	fmt.Println(birthDate.Format("2006-01-02"))
	return birthDate, nil
}

func main() {
	var dateOfBirthStr, motherBirthDateStr string

	fmt.Print("Enter your Date of birth(YYYY-MM-DD):")
	fmt.Scanln(&dateOfBirthStr)

	fmt.Print("Enter your mother's Date of birth(YYYY-MM-DD):")
	fmt.Scanln(&motherBirthDateStr)

	dateOfBirth, _ := checkDate(dateOfBirthStr)

	motherBirthDate, _ := checkDate(motherBirthDateStr)

	dobYear := dateOfBirth.Year()
	motherYear := motherBirthDate.Year()

	year := dobYear - motherYear
	if year < 18 {
		fmt.Println("Enter valid date of birth for mother")
		return
	}

	now := time.Now()
	age := now.Year() - dateOfBirth.Year()

	if now.YearDay() < dateOfBirth.YearDay() {
		age--
	}

	fmt.Println("Your current age:", age)

	motherAge := dateOfBirth.Year() - motherBirthDate.Year()

	if dateOfBirth.YearDay() < motherBirthDate.YearDay() {
		motherAge--
	}

	fmt.Println("Your mother's age when you born:", motherAge)

	nextBirthDay := time.Date(now.Year(), dateOfBirth.Month(), dateOfBirth.Day(), 0, 0, 0, 0, time.UTC)

	if now.After(nextBirthDay){
		nextBirthDay = nextBirthDay.AddDate(1,0,0)
	}

	daysLeft := nextBirthDay.Sub(now)
	day := int(daysLeft.Hours() / 24)
	fmt.Println("Days left for your birthday:", day)
}