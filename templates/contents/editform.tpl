
{{define "content"}}
<form method="post" action="/{{.language}}/action/edit/">
    <table class="content-table">
        <thead>
            <tr>
                <th>{{index .translation "DisplayName"}}:</th>
                <td><input type="text" value="{{.student.StudentFirstName}}" name="studentFirstName"></td>
            </tr>
            <tr>
                <th>{{index .translation "DisplayLastName"}}:</th>
                <td><input type="text" value="{{.student.StudentLastName}}" name="studentLastName"></td>
            </tr>
            <tr>
                <th>{{index .translation "DisplayDateOfBirth"}}:</th>
                <td><input type="date" value="{{.student.DateOfBrith}}" name="studentDateOfBrith"></td>
            </tr>
            <tr>
                <th>{{index .translation "DisplayDepartment"}}:</th>
                <td><input type="text" value="{{.student.StudentFaciulty}}" name="studentFaciulty"></td>
            </tr>
            <tr>
                <th>{{index .translation "DisplayGender"}}:</th>
                <td><select name="studentGender">
                    <option value="0">{{index .translation "DisplaySelectMen"}}</option>
                    <option value="1" {{if eq .student.StudentFaciulty "1"}}selected{{end}}>{{index .translation "DisplaySelectWomen"}}</option>
                </select></td>
            </tr>
            <tr>
                <td colspan="2" style="text-align: center;"><button>{{index .translation "DisplayChange"}}</button></td>
            </tr>
        </thead>
    </table>
</form>
{{end}}