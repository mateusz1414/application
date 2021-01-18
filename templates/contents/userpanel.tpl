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
                    </div>
                </div>
            </div>
        {{end}}
    </div>
{{end}}
{{define "script"}}
    <script type="text/javascript" src="/js/deanfunctions.js"></script>
{{end}}