{{define "content"}}
<table class="content-table">
    <thead>
        <tr>
            <th>{{index .translation "Display#"}}</th>
            <th>{{index .translation "DisplayName"}}</th>
            <th>{{index .translation "DisplayLastName"}}</th>
            <th>{{index .translation "DisplayDateOfBirth"}}</th>
            <th>{{index .translation "DisplayDepartment"}}</th>
            <th>{{index .translation "DisplayGender"}}</th>
            <th>{{index .translation "DisplayEdit"}}</th>
        </tr>
    </thead>
    <tbody>
        {{$displayedit:=index .translation "DisplayEdit"}}
        {{$displayselectmen:=index .translation "DisplaySelectMen"}}
        {{$displayselectwomen:=index .translation "DisplaySelectWomen"}}
        {{$language:=.language}}
        {{ range $index,$value :=.studentsList}}
        <tr>
            <td>{{$index}}</td>
            <td>{{.StudentFirstName}}</td>
            <td>{{.StudentLastName}}</td>
            <td>{{.DateOfBrith}}</td>
            <td>{{.StudentFaciulty}}</td>
            <td>{{if eq .StudentGender "0"}}{{$displayselectmen}}{{else}}{{$displayselectwomen}}{{end}}</td>
            <td><a href="/{{$language}}/editstudentform/{{.StudentID}}/"><button>{{$displayedit}}</button></a></td>
        </tr>
        {{end}}
    </tbody>
</table>
{{end}}