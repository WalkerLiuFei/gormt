package genfunc

const (
	genTnf = `
// TableName get sql table name.获取数据库表名
func (m *XXX_{{.StructName}}) TableName() string {
	return "{{.TableName}}"
}
func (m *XXX_{{.StructName}}Updates) TableName() string {
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

	genImplementation = `

func (repo *_repository) Create{{.StructName}}(ctx context.Context, obj domain.{{.StructName}}) error {
	daoObject := &XXX_{{.StructName}}{}
	if err := copier.Copy(daoObject, obj); err != nil {
		return err
	}
	return repo.db.WithContext(ctx).Model(&XXX_{{.StructName}}{}).Create(daoObject).Error
}

func (repo *_repository) Query{{.StructName}}ByOptions(ctx context.Context, queryOptions *options.{{.StructName}}QueryOptions) ([]domain.{{.StructName}}, error) {
	//TODO implement me
	daoResult := make([]XXX_{{.StructName}}, 0)
	if err := options.ApplyOptions(ctx, repo.db.Model(&XXX_{{.StructName}}{}), queryOptions.GetBaseQuery()).Find(&daoResult).Error; isUnexpectError(err) {
		return nil, err
	}

	result := make([]domain.{{.StructName}}, len(daoResult))
	for i, v := range daoResult {
		if err := copier.Copy(&result[i], &v); err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (repo *_repository) Update{{.StructName}}ByOptions(ctx context.Context, update domain.{{.StructName}}Updates, queryOptions *options.{{.StructName}}QueryOptions) error {
	daoObject := &XXX_{{.StructName}}Updates{}
	if err := copier.Copy(daoObject, update); err != nil {
		return err
	}
	if err := options.ApplyOptions(ctx, repo.db.Model(daoObject), queryOptions.GetBaseQuery()).UpdateColumns(daoObject).Error; err != nil {
		return err
	}
	return nil
}

func (repo *_repository) Count{{.StructName}}ByOptions(ctx context.Context, queryOptions *options.{{.StructName}}QueryOptions) (int64, error) {
	//TODO implement me
	result := new(int64)
	if err := options.ApplyOptions(ctx, repo.db.Model(&XXX_{{.StructName}}{}), queryOptions.GetBaseQuery()).Count(result).Error; err != nil || result == nil {
		return -1, err
	}

	return *result, nil
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

	// general functions.
	genInterface = `
    type {{.StructName}}Repository interface {
       Create{{.StructName}}(ctx context.Context, botProfit {{.StructName}}) error

	   Query{{.StructName}}ByOptions(ctx context.Context, queryOptions *options.{{.StructName}}QueryOptions) ([]{{.StructName}}, error)

	   Update{{.StructName}}ByOptions(ctx context.Context, update {{.StructName}}Updates, queryOptions *options.{{.StructName}}QueryOptions) error

	   Count{{.StructName}}ByOptions(ctx context.Context, queryOptions *options.{{.StructName}}QueryOptions) (int64, error)
	}
    `

	genlogic = `{{$obj := .}}{{$list := $obj.Em}}
type {{$obj.StructName}}QueryOptions  struct {
	*BaseQuery
}

func New{{$obj.StructName}}QueryOptions() *{{$obj.StructName}}QueryOptions {
	return &{{$obj.StructName}}QueryOptions{
		BaseQuery: new(BaseQuery),
	}
}

func (options *{{$obj.StructName}}QueryOptions ) GetBaseQuery() *BaseQuery {
	return options.BaseQuery
}

// GetTableName get sql table name.获取数据库名字
func (options *{{$obj.StructName}}QueryOptions ) GetTableName() string {
	return "{{GetTablePrefixName $obj.TableName}}"
}


{{range $oem := $obj.Em}}
func (options *{{$obj.StructName}}QueryOptions ) With{{$oem.ColStructName}}In({{CapLowercase $oem.ColStructName}}s []{{$oem.Type}}) *{{$obj.StructName}}QueryOptions  {
	options.Expressions = append(options.Expressions, Option{
		Or:    options.OrStatus,
		Query: "{{$oem.ColName}} in (?)",
		Args:  {{CapLowercase $oem.ColStructName}}s ,
	})
	options.OrStatus = false
	return options
}

func (options *{{$obj.StructName}}QueryOptions ) With{{$oem.ColStructName}}Eq({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) *{{$obj.StructName}}QueryOptions  {
	options.Expressions = append(options.Expressions, Option{
		Or:    options.OrStatus,
		Query: "{{$oem.ColName}} = ?",
		Args:  {{CapLowercase $oem.ColStructName}},
	})
	options.OrStatus = false
	return options
}


func (options *{{$obj.StructName}}QueryOptions ) With{{$oem.ColStructName}}Gt({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) *{{$obj.StructName}}QueryOptions  {
	options.Expressions = append(options.Expressions, Option{
		Or:    options.OrStatus,
		Query: "{{$oem.ColName}} > ?",
		Args:  {{CapLowercase $oem.ColStructName}},
	})
	options.OrStatus = false
	return options
}

func (options *{{$obj.StructName}}QueryOptions ) With{{$oem.ColStructName}}Gte({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) *{{$obj.StructName}}QueryOptions  {
	options.Expressions = append(options.Expressions, Option{
		Or:    options.OrStatus,
		Query: "{{$oem.ColName}} >= ?",
		Args:  {{CapLowercase $oem.ColStructName}},
	})
	options.OrStatus = false
	return options
}

func (options *{{$obj.StructName}}QueryOptions ) With{{$oem.ColStructName}}Lt({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) *{{$obj.StructName}}QueryOptions  {
	options.Expressions = append(options.Expressions, Option{
		Or:    options.OrStatus,
		Query: "{{$oem.ColName}} < ?",
		Args:  {{CapLowercase $oem.ColStructName}},
	})
	options.OrStatus = false
	return options
}

func (options *{{$obj.StructName}}QueryOptions ) With{{$oem.ColStructName}}Lte({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) *{{$obj.StructName}}QueryOptions  {
	options.Expressions = append(options.Expressions, Option{
		Or:    options.OrStatus,
		Query: "{{$oem.ColName}} <= ?",
		Args:  {{CapLowercase $oem.ColStructName}},
	})
	options.OrStatus = false
	return options
}

func (options *{{$obj.StructName}}QueryOptions ) With{{$oem.ColStructName}}Ne({{CapLowercase $oem.ColStructName}} {{$oem.Type}}) *{{$obj.StructName}}QueryOptions  {
	options.Expressions = append(options.Expressions, Option{
		Or:    options.OrStatus,
		Query: "{{$oem.ColName}} != ?",
		Args:  {{CapLowercase $oem.ColStructName}},
	})
	options.OrStatus = false
	return options
}
func (options *{{$obj.StructName}}QueryOptions ) OrderBy{{$oem.ColStructName}}(desc bool) *{{$obj.StructName}}QueryOptions  {
	options.OrderByInfo = append(options.OrderByInfo, OrderByInfo{
		Column:  "{{$oem.ColName}}",
		Desc: desc,
	})
	return options
}
{{end}}
func  (options *{{$obj.StructName}}QueryOptions )  Or() *{{$obj.StructName}}QueryOptions {
	options.OrStatus = true
	return options
}
`
)
