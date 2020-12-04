{{define "content"}}  
<table class="content-table add-form">
    <thead>
        <tr>
            <th>{{index .translation "DisplayName"}}:</th>
            <td><input type="text" id="studentFirstName"></td>
        </tr>
        <tr>
            <th>{{index .translation "DisplayLastName"}}:</th>
            <td><input type="text" id="studentLastName"></td>
        </tr>
        <tr>
            <th>{{index .translation "DisplayDateOfBirth"}}:</th>
            <td><input type="date" id="studentDateOfBrith"></td>
        </tr>
        <tr>
            <th>{{index .translation "DisplayDepartment"}}:</th>
            <td><input type="text" id="studentFaciulty"></td>
        </tr>
        <tr>
            <th>{{index .translation "DisplayGender"}}:</th>
            <td>
                <select id="studentGender">
                    <option value="0">{{index .translation "DisplaySelectMen"}}</option>
                    <option value="1">{{index .translation "DisplaySelectWomen"}}</option>
                </select>
            </td>
        </tr>
        <tr>
            <td colspan="2" style="text-align: center;"><button id="addButton">{{index .translation "DisplayAdd"}}</button></td>
        </tr>
    </thead>
</table>
{{end}}