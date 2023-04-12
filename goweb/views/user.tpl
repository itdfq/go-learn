<!DOCTYPE html>
<html lang="zh-CN">
<head>
	<title>Demo</title>
	<meta charset="utf-8">
</head>
<body>
<h1>网站: {{.Website}}</h1>
{{ if .user }}
用户名: {{.user.Username}}
{{else}}
查找不到用户
{{ end }}
</body>
</html>