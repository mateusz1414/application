{{define "page-aside"}}
Logowanie</br>
{{if .isLogined}}
zalogowano
{{else}}
<p class="errors">{{if .loginError}}
{{range .loginError}}
{{.}}</br>
{{end}}</p>
{{end}}
<table class="login-table">
<form action="/{{.language}}/user/login/" method="POST">
    <thead>
        <tr>
            <th>Login:</th>
            <td><input type="text" name="user"></td>
        </tr>
        <tr>
            <th>Hasło:</th>
            <td><input type="password" name="password"></td>
        </tr>
        <tr>
            <td colspan="2" style="text-align: center;"><button>ZALOGUJ</button></td>
        </tr>
    </thead>
    </form>
</table>
{{end}}
{{end}}