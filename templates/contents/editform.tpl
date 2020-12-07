{{define "content"}}
<table class="content-table edit-form">
    <thead>
        <tr>
            <th>{{index .translation "DisplayName"}}:</th>
            <td><input type="text" value="" id="studentFirstName"></td>
        </tr>
        <tr>
            <th>{{index .translation "DisplayLastName"}}:</th>
            <td><input type="text" value="" id="studentLastName"></td>
        </tr>
        <tr>
            <th>{{index .translation "DisplayDateOfBirth"}}:</th>
            <td><input type="date" value="" id="studentDateOfBrith"></td>
        </tr>
        <tr>
            <th>{{index .translation "DisplayDepartment"}}:</th>
            <td><input type="text" value="" id="studentFaciulty"></td>
        </tr>
        <tr>
            <th>{{index .translation "DisplayGender"}}:</th>
            <td><select id="studentGender">
                <option value="0" id="studentMen">{{index .translation "DisplaySelectMen"}}</option>
                <option value="1" id="studentWomen">{{index .translation "DisplaySelectWomen"}}</option>
            </select></td>
        </tr>
        <tr>
            <td colspan="2" style="text-align: center;"><button id="editButton">{{index .translation "DisplayChange"}}</button></td>
        </tr>
    </thead>
</table>
{{end}}