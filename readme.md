~~~
/***
// table name
// @def [index|unique_index(sechema i_name filed_name)]
// @def [primay (sechema filed_name)]
// @def others
type TableDef struct{
  FiledDef (var|type) `db:"[filedname],[default=|size=|autoincreate]"`
}
*/
~~~

通过struct 定义结表结构，方便OOP化处理数据库后续DDL，DML以及DCL，先简单讲定义内容 做如下阐述。

在struct的 comment group中定义索引、唯一索引、联合索引、主键等，其中关键字包含：<br/>
<b>@def，index/unique_index,primary等</b>，<u>使用规则如下</u>:

### @def >

comment group中所有关键字使用必须使用的前置字段，且后续定义采用空格作为分词; <br/>

~~~
// @def index i_xxx f_id
~~~

### index/unique_index

字段中添加索引，唯一索引，联合索引

~~~
// @def index i_name f_name
// @def index i_name_age f_name f_age
// @def unique_index i_name f_name
~~~

建议索引字段名称按照i_xxx,关联数据表字段名称f_xxx

### primary

数据表主键定义(基本规则如上)

~~~
@def primary f_id
~~~

### 结构补充

~~~
\pkg
    \def 主定义包 
        —— struct_def.go 表定义
        —— struct_def_another.go 共同表抽象定义
    \def_another 其它引用包
        —— struct_def_common.go
~~~



