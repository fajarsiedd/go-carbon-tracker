package utils

import (
	"context"
	"fmt"
	"go-carbon-tracker/entities"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GetTipsFromGemini(trip entities.Trip) (result string) {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	journey := fmt.Sprintf(
		"Dalam perjalanan saya dari %s menuju %s sejauh %d km menggunakan %s berbahan bakar %s menghasilkan emisi karbon sebanyak %f gram.",
		trip.StartLocation.Address,
		trip.EndLocation.Address,
		int(trip.DistanceKM),
		trip.Vehicle.VehicleType,
		trip.Vehicle.FuelType,
		trip.CarbonEmission,
	)

	prompt := "Berdasarkan emisi karbon yang dihasilkan, apa pendapatmu? berikan tips. Sajikan jawaban pendek dalam 1 paragraf saja dan sertakan perjalanan saya!"

	model := client.GenerativeModel("gemini-1.5-flash")

	resp, err := model.GenerateContent(ctx,
		genai.Text(journey),
		genai.Text(prompt),
	)

	if err != nil {
		log.Fatal(err)
	}

	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				result += fmt.Sprintf("%s", part)
			}
		}
	}

	result = strings.ReplaceAll(result, "\n", "")

	return result
}
