{{define "content"}}
<div class="login_register login-panel row col-lg-10 col-xl-8">
    <div class="content col-md-7 form-group mb-5">
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
        <div class="auth-group">
            <div class="row">
                <div class="col-md-10 m-auto">
                  <a class="btn col-8 col-md-10 btn-outline-light" href="/a" role="button">
                    <img width="20px" style="margin-bottom:3px; margin-right:5px" alt="Google sign-in" src="https://upload.wikimedia.org/wikipedia/commons/thumb/5/53/Google_%22G%22_Logo.svg/512px-Google_%22G%22_Logo.svg.png" />
                    Google
                  </a>
                </div>
              </div>
        </div>
    </div>
</div>
{{end}}
{{define "script"}}
    <script type="text/javascript" src="/js/loginregister.js"></script>
{{end}}