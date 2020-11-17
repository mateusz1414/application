
{{define "content"}}
<form method="post" action="/action/edit/">
    <table class="content-table">
        <thead>
            <tr>
                <th>Imie:</th>
                <td><input type="text" value="{{.student.StudentFirstName}}" name="studentFirstName"></td>
            </tr>
            <tr>
                <th>Nazwisko:</th>
                <td><input type="text" value="{{.student.StudentLastName}}" name="studentLastName"></td>
            </tr>
            <tr>
                <th>Data urodzenia:</th>
                <td><input type="date" value="{{.student.DateOfBrith}}" name="studentDateOfBrith"></td>
            </tr>
            <tr>
                <th>Wydział:</th>
                <td><input type="text" value="{{.student.StudentFaciulty}}" name="studentFaciulty"></td>
            </tr>
            <tr>
                <th>Płeć:</th>
                <td><select name="studentGender">
                    <option value="0">Mężczyzna</option>
                    <option value="1" {{if eq .student.StudentFaciulty "1"}}selected{{end}}>Kobieta</option>
                </select></td>
            </tr>
            <tr>
                <td colspan="2" style="text-align: center;"><button>ZMIEŃ</button></td>
            </tr>
        </thead>
    </table>
</form>
{{end}}