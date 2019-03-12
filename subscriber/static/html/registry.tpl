<!doctype html>
<html lang="ru">
	<head><meta charset="utf-8"></head>
	<body style="font-size: 18px; background: #fafbfc; color: #6c8093; font-family: 'PT sans', 'Trebuchet MS', 'Nimbus Sans L', 'Helvetica CY', sans-serif;">
		<table>
		{{range .}}
			<tr>
				<td><b>Topic: </b>{{.Topic}}</td>
				<td><b>         </b></td>
				<td><b>Value: </b>{{.Payload}}</td>
			</tr>
		{{end}}
		</table>
	</body>
</html>
