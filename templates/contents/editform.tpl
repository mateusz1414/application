
{{define "content"}}
<table class="content-table">
    <thead>
        <tr>
            <th>Imie:</th>
            <td><input type="text" value="{{.student.StudentFirstName}}"></td>
        </tr>
        <tr>
            <th>Nazwisko:</th>
            <td><input type="text" value="{{.student.StudentLastName}}"></td>
        </tr>
        <tr>
            <th>Data urodzenia:</th>
            <td><input type="date" value="{{.student.DateOfBrith}}"></td>
        </tr>
        <tr>
            <th>Wydział:</th>
            <td><input type="text" value="{{.student.StudentFaciulty}}"></td>
        </tr>
        <tr>
            <th>Płeć:</th>
            <td><select>
                <option value="0">Mężczyzna</option>
                <option value="1" {{if eq .student.StudentFaciulty "1"}}selected{{end}}>Kobieta</option>
            </select></td>
        </tr>
        <tr>
            <td colspan="2" style="text-align: center;"><button>ZMIEŃ</button></td>
        </tr>
    </thead>
</table>
{{end}}