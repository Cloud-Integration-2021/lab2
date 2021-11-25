package models

import "github.com/brianvoe/gofakeit"

type Actor struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	BirthDate string `json:"birthDate"`
}

func generateRandomActor()(a Actor){

	a.FirstName = gofakeit.FirstName()
	a.LastName = gofakeit.LastName()
	a.BirthDate = gofakeit.Date().Format("2006-01-02")

	return
}

func GenerateRandomActors()(actors []Actor){

	randomNumber := gofakeit.Number(1,10)

	for i := 0; i < randomNumber; i++ {
		actors = append(actors, generateRandomActor())
	}

	return
}