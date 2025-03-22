package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

type ExportedItem struct {
	Name      string
	Kind      string
	File      string
	Package   string
	Directory string
	RecvType  string // For methods, the receiver type
}

const pkgTemplate = `// Package {{.Package}} provides access to github.com/sourcegraph/zoekt/internal/{{.Directory}}
//
// This package is a mirror of the internal package with the same name, providing
// access to its exported identifiers.
package {{.Package}}

import original "github.com/sourcegraph/zoekt/internal/{{.Directory}}"

{{range .Vars}}
// {{.Name}} is re-exported from the internal package.
var {{.Name}} = original.{{.Name}}
{{end}}

{{range .Consts}}
// {{.Name}} is re-exported from the internal package.
const {{.Name}} = original.{{.Name}}
{{end}}

{{range .Types}}
// {{.Name}} is re-exported from the internal package.
type {{.Name}} = original.{{.Name}}
{{end}}

{{range .Funcs}}
// {{.Name}} is re-exported from the internal package.
var {{.Name}} = original.{{.Name}}
{{end}}
`

func main() {
	if len(os.Args) > 1 && os.Args[1] == "-h" {
		fmt.Println("Usage: go run main.go")
		fmt.Println("Walks through all Go files in the internal/ directory and creates a symmetric package structure under pkg/")
		return
	}

	// Get the root directory of the project
	rootDir, err := filepath.Abs(".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting current directory: %v\n", err)
		os.Exit(1)
	}

	// Make sure we're in the project root
	for !isProjectRoot(rootDir) {
		parentDir := filepath.Dir(rootDir)
		if parentDir == rootDir {
			fmt.Fprintf(os.Stderr, "Could not find project root (directory containing internal/)\n")
			os.Exit(1)
		}
		rootDir = parentDir
	}

	internalDir := filepath.Join(rootDir, "internal")
	if _, err := os.Stat(internalDir); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "internal/ directory not found at %s\n", internalDir)
		os.Exit(1)
	}

	var exportedItems []ExportedItem

	// Walk through all files in the internal directory
	err = filepath.Walk(internalDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			// Skip internal directories under internal/
			if info.Name() == "internal" && path != internalDir {
				fmt.Printf("Skipping internal directory: %s\n", path)
				return filepath.SkipDir
			}
			return nil
		}

		// Only process Go files
		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		// Skip test files
		if strings.HasSuffix(path, "_test.go") {
			return nil
		}

		// Parse the Go file
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing %s: %v\n", path, err)
			return nil
		}

		// Get relative path from internal/
		relPath, err := filepath.Rel(internalDir, path)
		if err != nil {
			relPath = path
		}

		// Get directory relative to internal/
		dir := filepath.Dir(relPath)

		// Find all exported declarations
		for _, decl := range file.Decls {
			switch d := decl.(type) {
			case *ast.GenDecl:
				// Handle vars, consts, and types
				switch d.Tok {
				case token.VAR, token.CONST, token.TYPE:
					for _, spec := range d.Specs {
						switch s := spec.(type) {
						case *ast.ValueSpec:
							// Handle vars and consts
							for _, name := range s.Names {
								if ast.IsExported(name.Name) {
									kind := "var"
									if d.Tok == token.CONST {
										kind = "const"
									}
									exportedItems = append(exportedItems, ExportedItem{
										Name:      name.Name,
										Kind:      kind,
										File:      filepath.Base(path),
										Package:   file.Name.Name,
										Directory: dir,
									})
								}
							}
						case *ast.TypeSpec:
							// Handle types
							if ast.IsExported(s.Name.Name) {
								exportedItems = append(exportedItems, ExportedItem{
									Name:      s.Name.Name,
									Kind:      "type",
									File:      filepath.Base(path),
									Package:   file.Name.Name,
									Directory: dir,
								})
							}
						}
					}
				}
			case *ast.FuncDecl:
				// Handle functions and methods
				if ast.IsExported(d.Name.Name) {
					kind := "func"
					recvType := ""
					if d.Recv != nil {
						kind = "method"
						// Get the receiver type
						if len(d.Recv.List) > 0 {
							recvField := d.Recv.List[0]
							switch t := recvField.Type.(type) {
							case *ast.StarExpr:
								if ident, ok := t.X.(*ast.Ident); ok {
									recvType = ident.Name
								}
							case *ast.Ident:
								recvType = t.Name
							}
						}
					}
					exportedItems = append(exportedItems, ExportedItem{
						Name:      d.Name.Name,
						Kind:      kind,
						File:      filepath.Base(path),
						Package:   file.Name.Name,
						Directory: dir,
						RecvType:  recvType,
					})
				}
			}
		}

		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error walking directory: %v\n", err)
		os.Exit(1)
	}

	// Sort the exported items by directory, package, kind, and name
	sort.Slice(exportedItems, func(i, j int) bool {
		if exportedItems[i].Directory != exportedItems[j].Directory {
			return exportedItems[i].Directory < exportedItems[j].Directory
		}
		if exportedItems[i].Package != exportedItems[j].Package {
			return exportedItems[i].Package < exportedItems[j].Package
		}
		if exportedItems[i].Kind != exportedItems[j].Kind {
			return exportedItems[i].Kind < exportedItems[j].Kind
		}
		return exportedItems[i].Name < exportedItems[j].Name
	})

	// Group exported items by directory and package
	packageMap := make(map[string]map[string][]ExportedItem)
	for _, item := range exportedItems {
		if _, ok := packageMap[item.Directory]; !ok {
			packageMap[item.Directory] = make(map[string][]ExportedItem)
		}
		packageMap[item.Directory][item.Package] = append(packageMap[item.Directory][item.Package], item)
	}

	// Create pkg directory if it doesn't exist
	pkgDir := filepath.Join(rootDir, "pkg")
	if err := os.MkdirAll(pkgDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating pkg directory: %v\n", err)
		os.Exit(1)
	}

	// Create mirror packages
	for dir, packages := range packageMap {
		for pkg, items := range packages {
			// Skip empty packages
			if len(items) == 0 {
				continue
			}

			// Create package directory
			pkgPath := filepath.Join(pkgDir, dir)
			if err := os.MkdirAll(pkgPath, 0755); err != nil {
				fmt.Fprintf(os.Stderr, "Error creating directory %s: %v\n", pkgPath, err)
				continue
			}

			// Group items by kind
			vars := []ExportedItem{}
			consts := []ExportedItem{}
			types := []ExportedItem{}
			funcs := []ExportedItem{}

			for _, item := range items {
				switch item.Kind {
				case "var":
					vars = append(vars, item)
				case "const":
					consts = append(consts, item)
				case "type":
					types = append(types, item)
				case "func":
					// Only include functions, not methods
					funcs = append(funcs, item)
					// Skip methods - they'll be accessible through the type aliases
				}
			}

			// Create package file
			packageFile := filepath.Join(pkgPath, pkg+".go")

			// Prepare template data
			data := struct {
				Package   string
				Directory string
				Vars      []ExportedItem
				Consts    []ExportedItem
				Types     []ExportedItem
				Funcs     []ExportedItem
			}{
				Package:   pkg,
				Directory: dir,
				Vars:      vars,
				Consts:    consts,
				Types:     types,
				Funcs:     funcs,
			}

			// Generate package file
			tmpl, err := template.New("pkg").Parse(pkgTemplate)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing template: %v\n", err)
				continue
			}

			var buf bytes.Buffer
			if err := tmpl.Execute(&buf, data); err != nil {
				fmt.Fprintf(os.Stderr, "Error executing template: %v\n", err)
				continue
			}

			// Write package file
			if err := ioutil.WriteFile(packageFile, buf.Bytes(), 0644); err != nil {
				fmt.Fprintf(os.Stderr, "Error writing file %s: %v\n", packageFile, err)
				continue
			}

			fmt.Printf("Created package file: %s\n", packageFile)
		}
	}

	fmt.Println("\nDone! The pkg/ directory now contains mirror packages for all exported identifiers in internal/.")
}

// isProjectRoot checks if the given directory is the project root
// by checking if it contains the internal/ directory
func isProjectRoot(dir string) bool {
	internalDir := filepath.Join(dir, "internal")
	info, err := os.Stat(internalDir)
	if err != nil {
		return false
	}
	return info.IsDir()
}
