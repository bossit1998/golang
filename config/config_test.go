package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoad(t *testing.T) {
	c := Load()

	ast := assert.New(t)

	ast.NotNil(c)
	ast.Equal("127.0.0.1", c.PostgresHost)
}
