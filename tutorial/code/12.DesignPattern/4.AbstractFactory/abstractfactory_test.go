package abstractfactory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestFactoryCreateProgrammer
func TestFactoryCreateProgrammer(t *testing.T) {
	ffe := &FrontendFactory{}
	fa := ffe.CreateArchitect()
	fp := ffe.CreateProgrammer()
	assert.Equal(t, "FrontendArchitect designs html and css with Adobe XD", fa.Design())
	assert.Equal(t, "FrontendProgrammer writes html and css with DreamWeaver", fp.WriteCode())
	fbe := &BackendFactory{}
	ba := fbe.CreateArchitect()
	bp := fbe.CreateProgrammer()
	assert.Equal(t, "BackendArchitect designs microservices and try to deploy them with kubernetes", ba.Design())
	assert.Equal(t, "BackendProgrammer writes golang with vscode and vim-plugin", bp.WriteCode())
}
