{{define "content"}}
{{if .registerError}}
<p class="errors">{{range .registerError}}
{{.}}</br>
{{end}}
{{end}}</p>
<table class="content-table">
<form action="/user/register/" method="POST">
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
            <th>Powtórz hasło:</th>
            <td><input type="password" name="confirmpassword"></td>
        </tr>
        <tr>
            <td colspan="2" style="text-align: center;"><button>ZAREJESTRUJ</button></td>
        </tr>
    </thead>
    </form>
</table>
{{end}}