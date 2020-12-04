{{define "content"}}
<table class="content-table student-table">
    {{index .translation "h1Banner"}}
    <thead>
        <tr>
            <th>{{index .translation "Display#"}}</th>
            <th>{{index .translation "DisplayName"}}</th>
            <th>{{index .translation "DisplayLastName"}}</th>
            <th>{{index .translation "DisplayDateOfBirth"}}</th>
            <th>{{index .translation "DisplayDepartment"}}</th>
            <th>{{index .translation "DisplayGender"}}</th>
        </tr>
    </thead>
    <tbody>
       <!-- {{$displayselectmen:=index .translation "DisplaySelectMen"}}
        {{$displayselectwomen:=index .translation "DisplaySelectWomen"}}
        {{ range $index,$value :=.studentsList}}
        <tr>
            <td>{{$index}}</td>
            <td>{{.StudentFirstName}}</td>
            <td>{{.StudentLastName}}</td>
            <td>{{.DateOfBrith}}</td>
            <td>{{.StudentFaciulty}}</td>
            <td>{{if eq .StudentGender "0"}}{{$displayselectmen}}{{else}}{{$displayselectwomen}}{{end}}</td>
        </tr>
        {{end}}-->
    </tbody>
</table>
{{end}}