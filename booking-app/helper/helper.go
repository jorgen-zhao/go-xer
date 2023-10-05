package helper

import "strings"

func ValidUserInput(firstName string, lastName string, emailAdress string, userTickets uint, remainTickets uint) (bool, bool, bool){
	isvalidName := len(firstName) > 2 && len(lastName) > 2
	isvalidEmailAddress := strings.Contains(emailAdress, "@")
	isvalidTickets := userTickets > 0 && userTickets < remainTickets
	return isvalidName, isvalidEmailAddress, isvalidTickets
}