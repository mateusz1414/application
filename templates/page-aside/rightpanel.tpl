{{define "page-aside"}}
Logowanie</br>
{{if .isLogined}}
zalogowano
{{else}}
<table class="login-table">
    <thead>
        <tr>
            <th>Login:</th>
            <td><input type="text"></td>
        </tr>
        <tr>
            <th>Hasło:</th>
            <td><input type="password"></td>
        </tr>
        <tr>
            <td colspan="2" style="text-align: center;"><button>ZALOGUJ</button></td>
        </tr>
    </thead>
</table>
{{end}}
{{end}}