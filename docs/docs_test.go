package docs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocs_init(t *testing.T) {
	assert.NotNil(t, len(SwaggerInfo.InstanceName()))
}
