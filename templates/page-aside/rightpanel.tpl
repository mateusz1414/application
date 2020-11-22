{{define "page-aside"}}
{{if .isLogined}}
{{index .translation "DisplayLogged"}}
{{else}}
{{index .translation "DisplayLoginIn"}}</br>
{{if .loginErrorFirst}}
<p class="errors">
{{index .translation .loginErrorFirst}}
</p>
{{end}}
{{if .loginErrorSecond}}
<p class="errors">
{{index .translation .loginErrorSecond}}
</p>
{{end}}
<table class="login-table">
<form action="/{{.language}}/user/login/" method="POST">
    <thead>
        <tr>
            <th>{{index .translation "DisplayLogin"}}:</th>
            <td><input type="text" name="user"></td>
        </tr>
        <tr>
            <th>{{index .translation "DisplayPassword"}}:</th>
            <td><input type="password" name="password"></td>
        </tr>
        <tr>
            <td colspan="2" style="text-align: center;"><button>{{index .translation "DisplayLoginIn"}}</button></td>
        </tr>
    </thead>
    </form>
</table>
{{end}}
{{end}}