{{define "content"}}
<div class="modal fade student-modal" id="update-modal" data-bs-backdrop="static">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <h5>{{index .translation "EditStudent"}}</h5>
            </div>
            <div class="modal-body">
                <input type="hidden" class="waiting-id">
                <div class="col">
                    <div class="alert alert-danger col-10 m-auto text-center"></div>
                </div>
                <div class="row">
                    <div class="col col-md-4">{{index .translation "Name"}}</div><div class="col col-md-7"><input type="text" class="waiting-name"></div>
                </div>
                <div class="row">
                    <div class="col col-md-4">{{index .translation "Surname"}}</div><div class="col col-md-7"><input type="text" class="waiting-surname"></div>
                </div>
                <div class="row">
                    <div class="col col-md-4">{{index .translation "DOB"}}</div><div class="col col-md-7"><input type="date" class="waiting-dob"></div>
                </div>
                <div class="row">
                    <div class="col col-md-4">{{index .translation "Department"}}</div><div class="col col-md-7">
                        <select class="waiting-departament">
                            <option value="undefined" selected>{{index .translation "Select"}}</option>
                        </select>
                    </div>
                </div>
                <div class="row">
                    <div class="col col-md-4">{{index .translation "Sex"}}</div><div class="col col-md-7">
                        <select class="waiting-sex">
                            <option value="undefined" selected>{{index .translation "Select"}}</option>
                            <option value="0">{{index .translation "Male"}}</option>
                            <option value="1">{{index .translation "Female"}}</option>
                        </select>
                    </div>
                </div>
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-secondary" data-bs-dismiss="modal" >{{index .translation "Cancel"}}</button>
                <button type="button" class="btn btn-primary" >{{index .translation "MenuEdit"}}</button>
            </div>
        </div>
    </div>
</div>
<div class="table-layout col-sm-10 m-auto mt-5">
    <table class="content-table student-table table" id="modify-table">
        <thead>
            <tr>
                <th scope="col">{{index .translation "NO"}}</th>
                <th scope="col">{{index .translation "Name"}}</th>
                <th scope="col">{{index .translation "Surname"}}</th>
                <th scope="col">{{index .translation "DOB"}}</th>
                <th scope="col">{{index .translation "Department"}}</th>
                <th scope="col">{{index .translation "Sex"}}</th>
                <th scope="col">{{index .translation "MenuEdit"}}</th>
                <th scope="col">{{index .translation "MenuDelete"}}</th>
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