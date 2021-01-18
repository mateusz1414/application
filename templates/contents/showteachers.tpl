{{define "content"}}
<div class="table-layout col-sm-10 m-auto mt-5">
    <table class="content-table teacher-table table ">
        {{index .translation "h1Banner"}}
        <thead>
            <tr>
                <th scope="col">{{index .translation "NO"}}</th>
                <th scope="col">{{index .translation "Name"}}</th>
                <th scope="col">{{index .translation "Surname"}}</th>
                <th scope="col">{{index .translation "Subject"}}</th>
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