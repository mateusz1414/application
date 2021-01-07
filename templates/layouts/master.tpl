<!DOCTYPE html>
<html lang="pl">
    <head>
        {{include "layouts/head"}}
    </head>
    <body>
        <header class="page-header col-sm-9">
            <div class="banner">
                <h1>Studenci</h1>
            </div>
            {{include "layouts/menu"}}
        </header>
        <div class="modal fade" id="modal-alert">
            <div class="modal-dialog">
                <div class="modal-content">
                    <div class="modal-body">
                        
                    </div>
                    <div class="modal-footer">
                        <button type="button" class="btn-secondary" data-bs-dismiss="modal">Zamknij</button>
                    </div>
                </div>
            </div>
        </div>
        {{if (eq .user.Permissions "user")}}
            <div class="modal fade student-modal" id="waiting-modal" data-bs-backdrop="static">
                <div class="modal-dialog">
                    <div class="modal-content">
                        <div class="modal-header">
                            <h5>{{index .translation "JoinToWaiting"}}</h5>
                        </div>
                        <div class="modal-body">
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
                            <button type="button" class="btn btn-primary" >{{index .translation "SendRequest"}}</button>
                        </div>
                    </div>
                </div>
            </div>
        {{end}}
        <main class="page-main-section col-sm-9">
            <section>
                {{template "content" .}}
            </section>
        </main>
        {{include "layouts/footer"}}
    </body>
    <script>
        var userID ={{.user.UserID}};
        var translation={
            addOnWaiting: '{{index .translation "AddOnWaiting"}}',
            sessionExpired: '{{index .translation "SessionExpired"}}',
            male: '{{index .translation "Male"}}',
            female: '{{index .translation "Female"}}',
            edit: '{{index .translation "MenuEdit"}}',
            delete: '{{index .translation "MenuDelete"}}',
            dob: '{{index .translation "DOB"}}',
            sex: '{{index .translation "Sex"}}',
            departament: '{{index .translation "Department"}}',
            waitingList: '{{index .translation "WaitingList"}}',
            average: '{{index .translation "Average"}}',
            addGrade: '{{index .translation "AddGrade"}}',
            incorrectGrade: '{{index .translation "IncorrectGrade"}}',
            incorrectEmailOrPassword: '{{index .translation "IncorrectEmailOrPassword"}}',
            incorrectEmail: '{{index .translation "IncorrectEmail"}}',
            passwordIsTooShort: '{{index .translation "PasswordIsTooShort"}}',
            passwordsDoNotMatch: '{{index .translation "PasswordsDoNotMatch"}}',
            busyEmail: '{{index .translation "BusyEmail"}}',
            allFields: '{{index .translation "AllFields"}}',
            onList: '{{index .translation "OnList"}}',
            serverError: '{{index .translation "ServerError"}}',
        }
    </script>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta1/dist/js/bootstrap.bundle.min.js" integrity="sha384-ygbV9kiqUc6oa4msXn9868pTtWMgiQaeYH7/t7LECLbyPA2x65Kgf80OJFdroafW" crossorigin="anonymous"></script>
    <script type="text/javascript" src="/js/config.js"></script>
    <script src = "http://ajax.googleapis.com/ajax/libs/jquery/1.10.2/jquery.min.js"></script>
    <script type="text/javascript" src="/js/main.js"></script>
</html>