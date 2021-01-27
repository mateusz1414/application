{{define "content"}}
<p class="errors registerMessageFirst"></p>
<p class="errors registerMessageSecond"></p>
<table class="content-table short-table">
    <thead>
        <tr>
            <th>{{index .translation "DisplayLogin"}}:</th>
            <td><input type="text" class="registerUser"></td>
        </tr>
        <tr>
            <th>{{index .translation "DisplayPassword"}}:</th>
            <td><input type="password" class="registerPassword"></td>
        </tr>
        <tr>
            <th>{{index .translation "DisplayConfirmPassword"}}:</th>
            <td><input type="password" class="registerConfirmPassword"></td>
        </tr>
        <tr>
            <td colspan="2" style="text-align: center;"><button class="registerButton">{{index .translation "DisplayRegister"}}</button></td>
        </tr>
    </thead>
</table>
{{if .isLogined}}
{{index .translation "DisplayLogged"}}
<button class="logoutButton">{{index .translation "LogoutButton"}}</button>
{{else}}
<p class="errors loginMessageFirst">

</p>
<p class="errors loginMessageSecond">

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