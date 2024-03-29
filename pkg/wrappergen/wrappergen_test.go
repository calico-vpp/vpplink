package wrappergen_test

import (
	"embed"
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"go.fd.io/govpp/binapigen"
	"go.fd.io/govpp/binapigen/vppapi"

	"github.com/calico-vpp/vpplink/pkg/wrappergen"
)

//go:embed testdata
var templates embed.FS

func TestRequirementSatisfied(t *testing.T) {
	data := wrappergen.NewDataFromFiles(
		"go.fd.io/govpp/binapi",
		"somepackage",
		[]*binapigen.File{{
			Desc:    vppapi.File{Name: "ipip"},
			Version: "2.1.0",
		}})
	//data.RequirementSatisfied("ip_types", ">= 3.0.0", "interface_types", ">= 1.0.0", "ipip", ">= 2.0.2")
	require.True(t, data.RequirementSatisfied("ipip", ">= 2.0.2"))
	require.False(t, data.RequirementSatisfied("ipip", "< 2.0.2"))
}

func TestReqirementSatisfiedInTemplate(t *testing.T) {
	outputDir := "./testdata/output"
	data := wrappergen.NewDataFromFiles(
		"go.fd.io/govpp/binapi",
		"somepackage",
		[]*binapigen.File{},
	)
	templates, err := fs.Sub(templates, "testdata")
	require.NoError(t, err)
	tmpl, err := wrappergen.ParseFS(templates, "*.tmpl")
	require.NoError(t, err)
	err = tmpl.ExecuteAll(outputDir, data, nil)
	require.NoError(t, err)
	require.NoError(t, os.RemoveAll(outputDir))
}
