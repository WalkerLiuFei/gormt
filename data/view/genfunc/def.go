package genfunc

const (
	genTnf = `
// TableName get sql table name.获取数据库表名
func (m *{{.StructName}}) TableName() string {
	return "{{.TableName}}"
}
`
	genBase = `
    package {{.PackageName}}
type Option struct {
	Or    bool
	query string
	args  interface{}
}
type _BaseQuery struct {
	OrStatus    bool
	expressions []Option
	PageInfo    *PageInfo
}
type PageInfo struct {
	Offset int32
	Limit  int32
}

`
	genColumn = `
// {{.StructName}}Columns get sql column name.获取数据库列名
var {{.StructName}}Columns = struct { {{range $em := .Em}}
	{{$em.StructName}} string{{end}}    
	}{ {{range $em := .Em}}
		{{$em.StructName}}:"{{$em.ColumnName}}",  {{end}}           
	}
  
`

	genlogic = `{{$obj := .}}{{$list := $obj.Em}}
type _{{$obj.StructName}}Options struct {
	*_BaseQuery
}

// GetTableName get sql table name.获取数据库名字
func (options *_{{$obj.StructName}}Options) GetTableName() string {
	return "{{GetTablePrefixName $obj.TableName}}"
}


{{range $oem := $obj.Em}}
func (options *_{{$obj.StructName}}Options) With{{$oem.ColStructName}}In({{CapLowercase $oem.ColStructName}}s []{{$oem.Type}}) *_{{$obj.StructName}}Options {
	options.expressions = append(options.expressions, Option{
		Or:    options.OrStatus,
		query: "{{$oem.ColName}} in (?)",
		args:  {{CapLowercase $oem.ColStructName}}s ,
	})
	options.OrStatus = false
	return options
}

func (options *_{{$obj.StructName}}Options) With{{$oem.ColStructName}}Eq({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) *_{{$obj.StructName}}Options {
	options.expressions = append(options.expressions, Option{
		Or:    options.OrStatus,
		query: "{{$oem.ColName}} = ?",
		args:  {{CapLowercase $oem.ColStructName}},
	})
	options.OrStatus = false
	return options
}


func (options *_{{$obj.StructName}}Options) With{{$oem.ColStructName}}Gt({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) *_{{$obj.StructName}}Options {
	options.expressions = append(options.expressions, Option{
		Or:    options.OrStatus,
		query: "{{$oem.ColName}} > ?",
		args:  {{CapLowercase $oem.ColStructName}},
	})
	options.OrStatus = false
	return options
}

func (options *_{{$obj.StructName}}Options) With{{$oem.ColStructName}}Gte({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) *_{{$obj.StructName}}Options {
	options.expressions = append(options.expressions, Option{
		Or:    options.OrStatus,
		query: "{{$oem.ColName}} >= ?",
		args:  {{CapLowercase $oem.ColStructName}},
	})
	options.OrStatus = false
	return options
}

func (options *_{{$obj.StructName}}Options) With{{$oem.ColStructName}}Lt({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) *_{{$obj.StructName}}Options {
	options.expressions = append(options.expressions, Option{
		Or:    options.OrStatus,
		query: "{{$oem.ColName}} < ?",
		args:  {{CapLowercase $oem.ColStructName}},
	})
	options.OrStatus = false
	return options
}

func (options *_{{$obj.StructName}}Options) With{{$oem.ColStructName}}Lte({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) *_{{$obj.StructName}}Options {
	options.expressions = append(options.expressions, Option{
		Or:    options.OrStatus,
		query: "{{$oem.ColName}} <= ?",
		args:  {{CapLowercase $oem.ColStructName}},
	})
	options.OrStatus = false
	return options
}

func (options *_{{$obj.StructName}}Options) With{{$oem.ColStructName}}Ne({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) *_{{$obj.StructName}}Options {
	options.expressions = append(options.expressions, Option{
		Or:    options.OrStatus,
		query: "{{$oem.ColName}} != ?",
		args:  {{CapLowercase $oem.ColStructName}},
	})
	options.OrStatus = false
	return options
}
{{end}}
func  (options *_{{$obj.StructName}}Options)  Or() *_{{$obj.StructName}}Options{
	options.OrStatus = true
	return options
}
`
)
