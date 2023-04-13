package def

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/scaf-fold/db-gknife/utils"
)

// DDoc define document
type DDoc struct {
	Comments []string
	Fields   []DField
}

// DField define field
type DField struct {
	Name string
	Typ  DToken
	Tag  reflect.StructTag
}

func DParser(dir, defName string) *DDoc {
	doc := new(DDoc)
	DTraversalDir(dir, defName, func(name string, typ DToken, tag reflect.StructTag) {
		doc.Fields = append(doc.Fields, DField{
			name,
			typ,
			tag,
		})
	}, func(text string) {
		doc.Comments = append(doc.Comments, text)
	})
	return doc
}

func DTraversalDir(dir string, defName string, fields func(name string, typ DToken, tag reflect.StructTag), comments func(text string)) {
	pk, err := OpenDir(dir)
	if err != nil {
		panic(err)
	}
	for _, v := range pk {
		for _, f := range v.Files {
			if no := f.Scope.Lookup(defName); no != nil {
				importMap := make(map[string]string)
				// 文件中包含定义结构体
				ast.Inspect(f, func(node ast.Node) bool {
					switch n := node.(type) {
					case *ast.ImportSpec:
						s := strings.ReplaceAll(n.Path.Value, "\"", "")
						importMap[filepath.Base(s)] = s
					case *ast.Comment:
						// table basic define
						comments(n.Text)
					case *ast.Field:
						// 符合要求相同
						switch t := n.Type.(type) {
						case *ast.Ident:
							// 基础定义类型
							if LookUp(t.Name).IsTypeDef() {
								if n.Tag != nil {
									// 目标结果
									tag := reflect.StructTag(strings.ReplaceAll(n.Tag.Value, "`", ""))
									fields(n.Names[0].Name, LookUp(t.Name), tag)
								}
							} else {
								// 同包不同文件
								if t.Obj == nil || (t.Obj != nil && t.Obj.Kind != ast.Typ) {
									DTraversalDir(dir, t.Name, fields, comments)
								}
							}
						case *ast.SelectorExpr:
							// 其它包中
							if x, ok := t.X.(*ast.Ident); ok {
								f, exist := importMap[x.Name]
								if exist {
									path, err := utils.GetImportPath(f)
									if err != nil {
										panic(err)
									}
									DTraversalDir(path, t.Sel.Name, fields, comments)
								}
							}
						}
					}
					return true
				})
			}
		}
	}
}

func OpenDir(path string, targetFile ...string) (map[string]*ast.Package, error) {
	return parser.ParseDir(token.NewFileSet(), path, func(info fs.FileInfo) bool {
		if info.IsDir() || (len(targetFile) == 1 && info.Name() != targetFile[0]) {
			return false
		}
		return true
	}, parser.ParseComments)
}
