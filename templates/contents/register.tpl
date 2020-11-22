{{define "content"}}
{{if .registerErrorFirst}}
<p class="errors">
{{index .translation .registerErrorFirst}}
</p>
{{end}}
{{if .registerErrorSecond}}
<p class="errors">
{{index .translation .registerErrorSecond}}
</p>
{{end}}
<table class="content-table">
<form action="/{{.language}}/user/register/" method="POST">
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
            <th>{{index .translation "DisplayConfirmPassword"}}:</th>
            <td><input type="password" name="confirmpassword"></td>
        </tr>
        <tr>
            <td colspan="2" style="text-align: center;"><button>{{index .translation "DisplayRegister"}}</button></td>
        </tr>
    </thead>
    </form>
</table>
{{end}}