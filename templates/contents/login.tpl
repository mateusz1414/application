{{define "content"}}
<div class="login_register login-panel row col-lg-10 col-xl-8">
    <div class="content col-md-7 form-group">
        <form method="POST" id="login-form">
            <div class="col">
                <div class="alert alert-danger py-1 mx-auto col-10" role="alert"></div>
            </div>
            <div class="row">
                <div class="label col col-3 col-sm-4">
                    {{index .translation "Email"}}:
                </div>
                <div class="col col-9 col-sm-5 col-md-8">
                    <input type="email" class="email form-control"></br>
                </div>
            </div>
            <div class="row">
                <div class="label col col-3 col-sm-4">
                    {{index .translation "Password"}}:
                </div>
                <div class="col col-9 col-sm-5 col-md-8">
                    <input type="password" class="password form-control"></br>
                </div>
            </div>
            <div class="col">
                <button class="button col-5 col-sm-4">{{index .translation "LoginVerb"}}</button>
            </div>
        </form>
    </div>
    <div class="login-options col-md-5">
        {{index .translation "ComingSoon"}}
    </div>
</div>
{{end}}
{{define "script"}}
    <script type="text/javascript" src="/js/loginregister.js"></script>
{{end}}