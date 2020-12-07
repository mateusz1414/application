{{define "page-aside"}}
{{if .isLogined}}
{{index .translation "DisplayLogged"}}
{{else}}
{{index .translation "DisplayLoginIn"}}</br>
<p class="errors loginMessageFirst">
{{index .translation .loginErrorFirst}}
</p>
<p class="errors loginMessageSecond">
{{index .translation .loginErrorSecond}}
</p>
<table class="login-table">
    <thead>
        <tr>
            <th>{{index .translation "DisplayLogin"}}:</th>
            <td><input type="text" class="loginUser"></td>
        </tr>
        <tr>
            <th>{{index .translation "DisplayPassword"}}:</th>
            <td><input type="password" class="loginPassword"></td>
        </tr>
        <tr>
            <td colspan="2" style="text-align: center;"><button class="loginButton">{{index .translation "DisplayLoginIn"}}</button></td>
        </tr>
    </thead>
</table>
{{end}}
{{end}}