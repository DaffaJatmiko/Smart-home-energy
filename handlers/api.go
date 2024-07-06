package handlers

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/DaffaJatmiko/smart-home-energy/model"
	"github.com/DaffaJatmiko/smart-home-energy/service"
	"github.com/DaffaJatmiko/smart-home-energy/util"
	"github.com/gin-gonic/gin"
)

type AIHandlerInterface interface {
	GetAnswer(c *gin.Context)
}

type AIhandler struct {
	aiService *service.AIService
}

func NewAIHandler(aiService *service.AIService) *AIhandler {
	return &AIhandler{
		aiService: aiService,
	}
}

func (h *AIhandler) GetAnswer(c *gin.Context) {
	token := os.Getenv("HUGGINGFACE_TOKEN")
	if token == "" {
		log.Fatal("HUGGINGFACE_TOKEN is not set")
	}
	
	dataMap, err := util.ReadCSV("data-series.csv")
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error reading CSV file",
		})
		return
	}

	var query model.InputQuery
	if err := c.ShouldBindJSON(&query); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request payload",
		})
		return
	}

	input := model.Inputs{
		Table: dataMap,
		Query:  query.Query,
	}

	log.Printf("Input: %+v\n", input)
	log.Printf("Token: %s\n", token)

	response, err := h.aiService.GetAnswer(input, token)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error getting answer from AI model",
		})
		return
	}

	switch response.Aggregator {
	case "SUM":
		totalEnergyConsumption := 0.0
		for _, cell := range response.Cells {
			value, err := strconv.ParseFloat(cell, 64)
			if err != nil {
				log.Printf("Error parsing energy consumption value: %v\n", err)
				continue
			}
			totalEnergyConsumption += value
		}
		c.JSON(200, gin.H{
			"result": fmt.Sprintf("Total Energy Consumption: %.2f", totalEnergyConsumption),
			"data":   response,
		})
	case "AVERAGE":
		sum := 0.0
		for _, cell := range response.Cells {
			value, err := strconv.ParseFloat(cell, 64)
			if err != nil {
				log.Printf("Error parsing energy consumption value: %v\n", err)
				continue
			}
			sum += value
		}
		average := sum / float64(len(response.Cells))
		c.JSON(200, gin.H{
			"result": fmt.Sprintf("Average Energy Consumption: %.2f", average),
			"data":   response,
		})
	case "MAX":
		max := -1.0
		for _, cell := range response.Cells {
			value, err := strconv.ParseFloat(cell, 64)
			if err != nil {
				log.Printf("Error parsing energy consumption value: %v\n", err)
				continue
			}
			if value > max {
				max = value
			}
		}
		c.JSON(200, gin.H{
			"result": fmt.Sprintf("Maximum Energy Consumption: %.2f", max),
			"data":   response,
		})
	case "MIN":
		min := 1000000.0
		for _, cell := range response.Cells {
			value, err := strconv.ParseFloat(cell, 64)
			if err != nil {
				log.Printf("Error parsing energy consumption value: %v\n", err)
				continue
			}
			if value < min {
				min = value
			}
		}
		c.JSON(200, gin.H{
			"result": fmt.Sprintf("Minimum Energy Consumption: %.2f", min),
			"data":   response,
		})
	default:
		c.JSON(500, gin.H{
			"message": fmt.Sprintf("Unexpected aggregator: %s", response.Aggregator),
		})
	}
}
