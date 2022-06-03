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
	Query string
	Args  interface{}
}
type BaseQuery struct {
	OrStatus    bool
	Expressions []Option
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
type _{{$obj.StructName}}QueryOptions  struct {
	*domain.BaseQuery
}

func New{{$obj.StructName}}QueryOptions() *_{{$obj.StructName}}QueryOptions {
	return &_{{$obj.StructName}}QueryOptions{
		BaseQuery: new(domain.BaseQuery),
	}
}

func (options *_{{$obj.StructName}}QueryOptions ) GetBaseQuery() *domain.BaseQuery {
	return options.BaseQuery
}

// GetTableName get sql table name.获取数据库名字
func (options *_{{$obj.StructName}}QueryOptions ) GetTableName() string {
	return "{{GetTablePrefixName $obj.TableName}}"
}


{{range $oem := $obj.Em}}
func (options *_{{$obj.StructName}}QueryOptions ) With{{$oem.ColStructName}}In({{CapLowercase $oem.ColStructName}}s []{{$oem.Type}}) *_{{$obj.StructName}}QueryOptions  {
	options.Expressions = append(options.Expressions, domain.Option{
		Or:    options.OrStatus,
		Query: "{{$oem.ColName}} in (?)",
		Args:  {{CapLowercase $oem.ColStructName}}s ,
	})
	options.OrStatus = false
	return options
}

func (options *_{{$obj.StructName}}QueryOptions ) With{{$oem.ColStructName}}Eq({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) *_{{$obj.StructName}}QueryOptions  {
	options.Expressions = append(options.Expressions, domain.Option{
		Or:    options.OrStatus,
		Query: "{{$oem.ColName}} = ?",
		Args:  {{CapLowercase $oem.ColStructName}},
	})
	options.OrStatus = false
	return options
}


func (options *_{{$obj.StructName}}QueryOptions ) With{{$oem.ColStructName}}Gt({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) *_{{$obj.StructName}}QueryOptions  {
	options.Expressions = append(options.Expressions, domain.Option{
		Or:    options.OrStatus,
		Query: "{{$oem.ColName}} > ?",
		Args:  {{CapLowercase $oem.ColStructName}},
	})
	options.OrStatus = false
	return options
}

func (options *_{{$obj.StructName}}QueryOptions ) With{{$oem.ColStructName}}Gte({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) *_{{$obj.StructName}}QueryOptions  {
	options.Expressions = append(options.Expressions, domain.Option{
		Or:    options.OrStatus,
		Query: "{{$oem.ColName}} >= ?",
		Args:  {{CapLowercase $oem.ColStructName}},
	})
	options.OrStatus = false
	return options
}

func (options *_{{$obj.StructName}}QueryOptions ) With{{$oem.ColStructName}}Lt({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) *_{{$obj.StructName}}QueryOptions  {
	options.Expressions = append(options.Expressions, domain.Option{
		Or:    options.OrStatus,
		Query: "{{$oem.ColName}} < ?",
		Args:  {{CapLowercase $oem.ColStructName}},
	})
	options.OrStatus = false
	return options
}

func (options *_{{$obj.StructName}}QueryOptions ) With{{$oem.ColStructName}}Lte({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) *_{{$obj.StructName}}QueryOptions  {
	options.Expressions = append(options.Expressions, domain.Option{
		Or:    options.OrStatus,
		Query: "{{$oem.ColName}} <= ?",
		Args:  {{CapLowercase $oem.ColStructName}},
	})
	options.OrStatus = false
	return options
}

func (options *_{{$obj.StructName}}QueryOptions ) With{{$oem.ColStructName}}Ne({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) *_{{$obj.StructName}}QueryOptions  {
	options.Expressions = append(options.Expressions, domain.Option{
		Or:    options.OrStatus,
		Query: "{{$oem.ColName}} != ?",
		Args:  {{CapLowercase $oem.ColStructName}},
	})
	options.OrStatus = false
	return options
}
{{end}}
func  (options *_{{$obj.StructName}}QueryOptions )  Or() *_{{$obj.StructName}}QueryOptions {
	options.OrStatus = true
	return options
}
`
)
