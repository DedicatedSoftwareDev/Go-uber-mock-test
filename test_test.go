package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

//go:generate mockgen -destination=mock_fetcher.go -package=main github.com/yourusername/yourproject DataFetcher

func TestProcessData(t *testing.T) {
	// Create a new mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// Create a new mock fetcher
	mockFetcher := NewMockDataFetcher(mockCtrl)

	// Set expectations
	mockFetcher.EXPECT().
		FetchData().
		Return("mock data")

	// Test the ProcessData function
	result := ProcessData(mockFetcher)
	assert.Equal(t, "Processed: mock data", result)
}
