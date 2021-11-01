package date

import (
	"log"
	"strings"
	"time"
)

func FormatDate(t time.Time, format string) (string, string) {
	mask := "02/01/2006"
	if strings.ToUpper(format) == "DB" {
		mask = "2006-01-02"
	}
	date := t.Format(mask)
	time := t.Format("03:04:05")
	return date, time
}

func TimeNowDB() (string, string) {
	return FormatDate(time.Now(), "DB")
}

func TimeNowBR() (string, string) {
	return FormatDate(time.Now(), "BR")
}

func AddDaysDateBR(dateInit string, numberDays int) (string, string) {
	mask := "02/01/2006"
	t, err := time.Parse(mask, dateInit)
	if err != nil {
		log.Println("Error addDaysDateBR", err)
		return "", ""
	}
	return FormatDate(t, "BR")
}

func RemoveDaysDateBR(dateBR string, numberDays int) (string, string) {
	mask := "02/01/2006"
	t, err := time.Parse(mask, dateBR)
	if err != nil {
		log.Println("Error removeDaysDateBR", err)
	}
	return FormatDate(t, "BR")
}

func IsTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func ThisTimeIsAfter(date time.Time) bool {
	return date.After(date)
}

func ThisTimeIsBefore(date time.Time) bool {
	return date.Before(date)
}

func ParseDate(date string) (time.Time, bool) {
	mask := "2006-01-02"
	if len(date) > 10 {
		mask = "2006-01-02 15:04:05"
	}

	if strings.Index(date, "-") < 1 {
		mask = "02/01/2006"
		if len(date) > 10 {
			mask = "02/01/2006 15:04:05"
		}
	}

	time, err := time.Parse(mask, date)
	if err != nil {
		return time, false
	}
	return time, true
}

func IsDateEqual(dateInitial, dateFinal string) bool {
	dateInitialConverted, right := ParseDate(dateInitial)
	if right == false {
		return false
	}

	dateFinalConverted, right := ParseDate(dateFinal)
	if right == false {
		return false
	}

	if dateInitialConverted.Equal(dateFinalConverted) {
		return true
	}
	return false
}

func IsDateLonger(dateInitial, dateFinal string) bool {
	dateInitialConverted, right := ParseDate(dateInitial)
	if right == false {
		return false
	}

	dateFinalConverted, right := ParseDate(dateFinal)
	if right == false {
		return false
	}

	if dateInitialConverted.After(dateFinalConverted) {
		return true
	}
	return false
}

func IsDateShorter(dateInitial, dateFinal string) bool {
	dateInitialConverted, right := ParseDate(dateFinal)
	if right == false {
		return false
	}

	dateFinalConverted, right := ParseDate(dateFinal)
	if right == false {
		return false
	}

	if dateInitialConverted.Before(dateFinalConverted) {
		return true
	}
	return false
}

func Week(week int) (dayWeek string) {
	switch week {
	case 0:
		dayWeek = "Sunday"
	case 1:
		dayWeek = "Monday"
	case 2:
		dayWeek = "Tuesday"
	case 3:
		dayWeek = "Wednesday"
	case 4:
		dayWeek = "Thursday"
	case 5:
		dayWeek = "Friday"
	case 6:
		dayWeek = "Saturday"
	default:
		dayWeek = "Non-existent day week."
	}
	return dayWeek
}

func Month(month int) (dayMonth string) {
	switch month {
	case 1:
		dayMonth = "January"
	case 2:
		dayMonth = "February"
	case 3:
		dayMonth = "March"
	case 4:
		dayMonth = "April"
	case 5:
		dayMonth = "May"
	case 6:
		dayMonth = "June"
	case 7:
		dayMonth = "July"
	case 8:
		dayMonth = "August"
	case 9:
		dayMonth = "Setember"
	case 10:
		dayMonth = "October"
	case 11:
		dayMonth = "November"
	case 12:
		dayMonth = "December"
	default:
		dayMonth = "does not exist this month"
	}
	return dayMonth
}

func WeekBR(week int) (dayWeek string) {
	switch week {
	case 0:
		dayWeek = "Domingo"
	case 1:
		dayWeek = "Segunda-Feira"
	case 2:
		dayWeek = "Terça-Feira"
	case 3:
		dayWeek = "Quarta-Feira"
	case 4:
		dayWeek = "Quinta-Feira"
	case 5:
		dayWeek = "Sexta-Feira"
	case 6:
		dayWeek = "Sábado"
	default:
		dayWeek = "Non-existent day week."
	}
	return dayWeek
}

func MonthBR(month int) (dayMonth string) {
	switch month {
	case 1:
		dayMonth = "Janeiro"
	case 2:
		dayMonth = "Fevereiro"
	case 3:
		dayMonth = "Março"
	case 4:
		dayMonth = "Abril"
	case 5:
		dayMonth = "Maio"
	case 6:
		dayMonth = "Junho"
	case 7:
		dayMonth = "Julho"
	case 8:
		dayMonth = "Agosto"
	case 9:
		dayMonth = "Setembro"
	case 10:
		dayMonth = "Outubro"
	case 11:
		dayMonth = "Novembro"
	case 12:
		dayMonth = "Dezembro"
	default:
		dayMonth = "does not exist this month."
	}
	return dayMonth
}
