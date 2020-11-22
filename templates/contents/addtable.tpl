{{define "content"}}  
 <form action="/{{.language}}/action/add/" method="post">
    <table class="content-table">
        <thead>
            <tr>
                <th>{{index .translation "DisplayName"}}:</th>
                <td><input type="text" name="studentFirstName"></td>
            </tr>
            <tr>
                <th>{{index .translation "DisplayLastName"}}:</th>
                <td><input type="text" name="studentLastName"></td>
            </tr>
            <tr>
                <th>{{index .translation "DisplayDateOfBirth"}}:</th>
                <td><input type="date" name="studentDateOfBrith"></td>
            </tr>
            <tr>
                <th>{{index .translation "DisplayDepartment"}}:</th>
                <td><input type="text" name="studentFaciulty"></td>
            </tr>
            <tr>
                <th>{{index .translation "DisplayGender"}}:</th>
                <td>
                    <select name="studentGender">
                        <option value="0">{{index .translation "DisplaySelectMen"}}</option>
                        <option value="1">{{index .translation "DisplaySelectWomen"}}</option>
                    </select>
                </td>
            </tr>
            <tr>
                <td colspan="2" style="text-align: center;"><button>{{index .translation "DisplayAdd"}}</button></td>
            </tr>
        </thead>
    </table>
</form>
{{end}}