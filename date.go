package date

import (
	"log"
	"strings"
	"time"
)

type Date interface {
	timeNowDB() (string, string)
	timeNowBR() (string, string)
	formatDate() (string, string)
	addDaysDateBR(string, int) (string, string)
	removeDaysDateBR(string, int) (string, string)
	week(int) string
	month(int) string
	weekBR(int) string
	monthBR(int) string
}

func formatDate(t time.Time, format string) (string, string) {
	mask := "02/01/2006"
	if strings.ToUpper(format) == "DB" {
		mask = "2006-01-02"
	}
	date := t.Format(mask)
	time := t.Format("03:04:05")
	return date, time
}

func timeNowDB() (string, string) {
	return formatDate(time.Now(), "DB")
}

func timeNowBR() (string, string) {
	return formatDate(time.Now(), "BR")
}

func addDaysDateBR(dateInit string, numberDays int) (string, string) {
	mask := "02/01/2006"
	t, err := time.Parse(mask, dateInit)
	if err != nil {
		log.Println("Error addDaysDateBR", err)
		return "", ""
	}
	return formatDate(t, "BR")
}

func removeDaysDateBR(dateBR string, numberDays int) (string, string) {
	mask := "02/01/2006"
	t, err := time.Parse(mask, dateBR)
	if err != nil {
		log.Println("Error removeDaysDateBR", err)
	}
	return formatDate(t, "BR")
}

func week(week int) (dayWeek string) {
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

func month(month int) (dayMonth string) {
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

func weekBR(week int) (dayWeek string) {
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

func monthBR(month int) (dayMonth string) {
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
