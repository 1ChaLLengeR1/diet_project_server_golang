package helpers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"internal/internal/data"
// 	"os"
// )

// func CreateFile(data data.CollectionWeek) error{

// 	file, err := os.Create("dane.txt")
// 	if err != nil{
// 		return fmt.Errorf("file not create %v", err)
// 	}
// 	defer file.Close()

// 	for _, v := range data.Data{

// 		var down string = fmt.Sprintf("Waga w doł: %.2f kg \n", v.WeightDown)
// 		if v.WeightDown > 0{
// 			down = fmt.Sprintf("Waga w doł: +%.2f kg \n", v.WeightDown)
// 		}

// 		line := fmt.Sprintf("Tydzień: %d \n", v.Week) +
// 		fmt.Sprintf("Startowa Waga: %.2f \n", v.StartWeight) +
// 		fmt.Sprintf("Waga końcowa: %.2f \n", v.EndWeight) +
// 		down +
// 		fmt.Sprintf("Suma kilogramów: %.2f \n", v.SumKg) +
// 		fmt.Sprintf("Średnia waga: %.2f \n", v.AvgKg) +
// 		fmt.Sprintf("Suma fitatu: %.2f \n", v.SumFitatu)

// 		for _, t := range v.Training{
// 			for _, k := range t.Data{
// 				line += fmt.Sprintf("Typ treningu: %v, liczba treningów: %d, suma spalonych kalorii: %d \n", k.Type, k.Currecnt, k.SumKcal)
// 			}

// 		}

// 		line += fmt.Sprintf("Suma czasu treningów: %v \n\n", v.TrainingTime)

// 		_, err := file.WriteString(line)
// 		if err != nil{
// 			return fmt.Errorf("not wirte to file %v", err)
// 		}
// 	}

// 	return nil
// }
