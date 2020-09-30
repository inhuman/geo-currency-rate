package rate_client

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRate(t *testing.T) {

	client, err := NewRateClient()
	assert.NoError(t, err)

	resp, err := client.GetRate()
	assert.NoError(t, err)
	fmt.Printf("%+v\n", resp)
}
