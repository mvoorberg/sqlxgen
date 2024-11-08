package store

import (
	"bytes"
	_ "embed"
	"go/format"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"text/template"

	"github.com/joomcode/errorx"
	"github.com/mvoorberg/sqlxgen/internal/utils"
	"github.com/mvoorberg/sqlxgen/internal/utils/casing"
	"github.com/mvoorberg/sqlxgen/internal/utils/writer"
)

type Package struct {
	WriterCreator writer.Creator `json:"-"`
	PackageName   string         `json:"package_name"`
	PackageDir    string         `json:"package_dir"`
	GenDir        string         `json:"gen_dir"`
	Options       map[string]string
}

func (p Package) Generate() error {
	slog.Debug("generating store package")

	helpers := template.FuncMap{
		"GetOption": func(opt string) string {
			optVal, ok := p.Options[opt]
			if !ok {
				return ""
			}
			return optVal
		},
	}

	tmpl, err := template.New(p.PackageName).Funcs(helpers).Parse(storeTemplate)

	if err != nil {
		return errorx.IllegalFormat.Wrap(err, "unable to parse store template")
	}

	var storeFileBuffer bytes.Buffer

	err = tmpl.Execute(
		&storeFileBuffer,
		map[string]interface{}{
			"PackageName": p.PackageName,
		},
	)

	if err != nil {
		return errorx.InternalError.Wrap(err, "unable to execute store template")
	}

	formatted, err := format.Source(storeFileBuffer.Bytes())

	if err != nil {
		return errorx.InternalError.Wrap(err, "unable to format store template")
	}

	err = os.MkdirAll(p.GenDir, 0755)

	if err != nil {
		return errorx.InitializationFailed.Wrap(err, "unable to create directory for store")
	}

	genFileName := utils.FilenameWithGen("store.go")

	storeFilePath := path.Join(p.GenDir, genFileName)

	pen := p.WriterCreator(storeFilePath, string(formatted))

	err = pen.Write()

	if err != nil {
		return errorx.InternalError.Wrap(err, "unable to write store file")
	}

	slog.Debug("generated store package")

	return nil
}

func NewPackage(
	writerCreator writer.Creator,
	packageDir string,
	genDir string,
	opts map[string]string,
) (Package, error) {
	parentDir := filepath.Base(packageDir)

	packageName, err := casing.SnakeCase(parentDir)

	if err != nil {
		return Package{}, errorx.IllegalState.Wrap(err, "unable to generate package name")
	}

	p := Package{
		WriterCreator: writerCreator,
		PackageName:   packageName,
		PackageDir:    packageDir,
		GenDir:        genDir,
		Options:       opts,
	}

	return p, nil
}

//go:embed store.go.tmpl
var storeTemplate string
