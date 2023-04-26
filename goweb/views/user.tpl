<!DOCTYPE html>
<html lang="zh-CN">
<head>
	<title>学习go模板</title>
	<meta charset="utf-8">
</head>
<body>
<h1>学习模板标签</h1><br/>
<ol>
<li>使用 . 来访问当前位置的上下文 </li>
<li>使用 $ 来引用当前模板根级的上下文</li>
<li>使用 $var 来访问创建的变量</li>
</ol>

<h3>网站: {{.Website}}</h3>
{{ if .user }}
用户名: {{.user.Username}}
{{else}}
查找不到用户
{{ end }}

{{if .Website}}
<br/>
存在网站
{{else}}
    {{if .Website}}
           网站是个博客
    {{end}}
{{end}}
<hr/>
<h2>
range ... end
</h2><br/>
循环打印数组<br/>
{{range .Pages}}
    <!-- 使用 .Num 输出子元素的 Num 属性，使用 $. 引用模板中的根级上下文 -->
    {{.Num}} of {{$.Total}}
<!-- 打印的结果：10 of 100 20 of 100 30 of 100 -->
{{end}}

<br/>使用创建的变量<br/>
{{range $index,$elem:=.Pages}}
<!-- 注意,这里的$index,$elem 相当于 key,value-->
{{$index}}-{{$elem.Num}} -{{.Num}} of {{$.Total}}
{{end}}
<hr/>
<h1>with ...end</h1>
{{with .Field.NestField.SubField}}
    {{.Var}}
{{end}}
</br>对变量进行赋值</br>
{{with $value:="My name is %s"}}
    {{printf . "slence"}}
{{end}}
</br>
{{with pipeline}}
{{else}}
    {{/* 当pipeline为空时会执行这里 */}}
{{end}}

</body>
</html>