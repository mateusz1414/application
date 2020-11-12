{{define "content"}}
<table class="content-table">
    <thead>
        <tr>
            <th>Lp.</th>
            <th>Imie</th>
            <th>Nazwisko</th>
            <th>Data urodzenia</th>
            <th>Wydział</th>
            <th>Plec</th>
            <th>Edytuj</th>
        </tr>
    </thead>
    <tbody>
        {{ range $index,$value :=.studentsList}}
        <tr>
            <td>{{$index}}</td>
            <td>{{.StudentFirstName}}</td>
            <td>{{.StudentLastName}}</td>
            <td>{{.DateOfBrith}}</td>
            <td>{{.StudentFaciulty}}</td>
            <td>{{if eq .StudentGender "0"}}Mężczyzna{{else}}Kobieta{{end}}</td>
            <td><a href="editstudentform/{{.StudentID}}"><button>EDYTUJ</button></a></td>
        </tr>
        {{end}}
    </tbody>
</table>
{{end}}