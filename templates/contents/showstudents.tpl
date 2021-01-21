{{define "content"}}
<div class="table-layout col-sm-10 m-auto mt-5">
    <table class="content-table student-table table ">
        <thead>
            <tr>
                <th scope="col">{{index .translation "NO"}}ei</th>
                <th scope="col">{{index .translation "Name"}}</th>
                <th scope="col">{{index .translation "Surname"}}</th>
                <th scope="col">{{index .translation "DOB"}}</th>
                <th scope="col">{{index .translation "Department"}}</th>
                <th scope="col">{{index .translation "Sex"}}</th>
            </tr>
        </thead>
        <tbody>
            <tr class="spinner-row">
                <td colspan="100%" class="justify-content-center spinner-col" ></td>
            </tr> 
        </tbody>
    </table>
</div>
{{end}}
{{define "script"}}
{{end}}