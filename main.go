package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

func main() {
	// Abre el archivo de entrada
	inputFile, err := os.Open("GOOG.csv")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	// Lee el contenido del archivo CSV
	reader := csv.NewReader(inputFile)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// Realiza el Bubble Sort en función de la columna "Volume" (posición 6 en los registros)
	for i := 1; i < len(records)-1; i++ {
		for j := 1; j < len(records)-i; j++ {
			volumeA, _ := strconv.Atoi(records[j][6])
			volumeB, _ := strconv.Atoi(records[j+1][6])

			//TODO completar
			if _________ {
				//swap
			}
		}
	}

	// Crea el archivo de salida
	outputFile, err := os.Create("output.csv")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// Escribe los registros ordenados en el archivo de salida
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Escribe la cabecera
	if err := writer.Write(records[0]); err != nil {
		panic(err)
	}

	// Escribe los registros ordenados
	for _, record := range records[1:] {
		if err := writer.Write(record); err != nil {
			panic(err)
		}
	}

	println("Archivo de salida 'output.csv' creado exitosamente.")
}
