{{define "content"}}
    <div class="user-panel col-11 col-lg-9 row">
        <div class="col
        {{if (eq .user.Permissions "dean")}}
            col-12 col-lg-6
        {{else}}
            col-12
        {{end}}">
            <h5>{{.user.Email}}</h5>
            <div class="col-12 user-functions">
                {{index .translation "ComingSoon"}}
            </div>
        </div>
        {{if (eq .user.Permissions "dean")}}
            <div class="col dean-functions col-12 col-lg-6">
                <div class="col-12 m-auto m-lg-0">
                    <div class="waiting-list ">
                        <div class="col waiting-element">{{index .translation "WaitingList"}}</div>
                      <!--  <div class="waiting-element col-12 row">
                            
                            <div class="col col-8 list-email" data-bs-toggle="collapse" href="#collapse-user-1">KtostamKtostam</div>
                            <div class="col col-2"><button type="button" class="btn btn-success">✔</button></div>
                            <div class="col col-2"><button type="button" class="btn btn-danger">X</button></div>
                            <div class="collapse" id="collapse-user-1">
                                <div class="row">
                                    <div class="col">Imie</div><div class="col">Adam</div>
                                </div>
                                <div class="row">
                                    <div class="col">Nazwkisko</div><div class="col">Nowak</div>
                                </div>
                            </div>
                        </div>-->
                    </div>
                </div>
            </div>
        {{end}}
    </div>
{{end}}