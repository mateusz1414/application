{{define "content"}}
<table class="content-table">
    {{index .translation "h1Banner"}}
    <thead>
        <tr>
            <th>Lp.</th>
            <th>Imie</th>
            <th>Nazwisko</th>
            <th>Data urodzenia</th>
            <th>Wydział</th>
            <th>Plec</th>
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
        </tr>
        {{end}}
    </tbody>
</table>
{{end}}