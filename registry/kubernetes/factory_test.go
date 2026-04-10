package kubernetes

import (
	"testing"

	conf "github.com/52dev/kratos-bootstrap/api/gen/go/conf/v1"
	"github.com/stretchr/testify/assert"
)

func TestNewKubernetesRegistry(t *testing.T) {
	var cfg conf.Registry
	reg, err := NewRegistry(&cfg)
	assert.Nil(t, err)
	assert.NotNil(t, reg)
}
