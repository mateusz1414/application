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
    </tbody>
</table>
{{end}}