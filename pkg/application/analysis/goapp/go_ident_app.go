package goapp

import (
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/phodal/coca/pkg/infrastructure/ast/cocago"
	"io/ioutil"
	"strings"
)

type GoIdentApp struct {
	Extensions     interface{}
	PackageManager core_domain.CodePackageManagerInfo
}

// todo: support multiple project
func (g *GoIdentApp) AnalysisPackageManager(path string) core_domain.CodePackageManagerInfo {
	content, _ := ioutil.ReadFile(path + "/" + "go.mod")
	pmInfo := &core_domain.CodePackageManagerInfo{}

	if content != nil {
		mod := string(content)
		splitContent := strings.Split(mod, "\n")
		modLine := splitContent[0]

		moduleName := modLine[len("module "):]
		pmInfo.ProjectName = strings.TrimSpace(moduleName)
	}
	g.PackageManager = *pmInfo
	return *pmInfo
}

func (g *GoIdentApp) Analysis(code string, fileName string) core_domain.CodeFile {
	parser := cocago.NewCocagoParser()
	var codeMembers []core_domain.CodeMember
	if g.Extensions != nil {
		codeMembers = g.Extensions.([]core_domain.CodeMember)
	}
	if g.PackageManager.ProjectName != "" {
		parser.SetPackageManager(g.PackageManager)
	}
	return *parser.ProcessString(code, fileName, codeMembers)
}

func (g *GoIdentApp) IdentAnalysis(code string, fileName string) []core_domain.CodeMember {
	parser := cocago.NewCocagoParser()
	codeFile := parser.IdentAnalysis(code, fileName)
	return codeFile.Members
}

func (g *GoIdentApp) SetExtensions(extension interface{}) {
	g.Extensions = extension
}
