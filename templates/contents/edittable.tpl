{{define "content"}}
<table class="content-table student-table edit-table">
    <thead>
        <tr>
            <th>{{index .translation "Display#"}}</th>
            <th>{{index .translation "DisplayName"}}</th>
            <th>{{index .translation "DisplayLastName"}}</th>
            <th>{{index .translation "DisplayDateOfBirth"}}</th>
            <th>{{index .translation "DisplayDepartment"}}</th>
            <th>{{index .translation "DisplayGender"}}</th>
            <th class="thEditButton">{{index .translation "DisplayEdit"}}</th>
        </tr>
    </thead>
    <tbody>
    </tbody>
</table>
{{end}}