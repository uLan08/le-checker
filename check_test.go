package lechecker

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCAId(t *testing.T) {
	ids := getCAIds()
	fmt.Println(ids)
	assert.NotEmpty(t, ids)
}
