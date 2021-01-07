<nav class="navbar navbar-dark navbar-expand-md">
    <div class="collapse navbar-collapse" id="nav-menu">
        
        <ul class="nav-main navbar-nav me-auto">
            <li class="nav-item">
                <a href="/{{.language}}/students/">{{index .translation "H1Banner"}}</a>
            </li>
            <li class="nav-item">
                <a href="/{{.language}}/teachers/">{{index .translation "Teachers"}}</a>
            </li>
            {{if .user.IsLogined}}
                <li class="nav-item">
                    {{if (eq .user.Permissions "dean")}}
                        <a href="/{{.language}}/modify/">{{index .translation "Modify"}}</a>
                    {{else if (eq .user.Permissions "teacher")}}
                        <a href="/{{.language}}/addgrades/">{{index .translation "AddGrades"}}</a>
                    {{else if (eq .user.Permissions "student")}}
                        <a href="/{{.language}}/getgrades/">{{index .translation "ViewGrades"}}</a>
                    {{else if (eq .user.Permissions "user")}}
                        <a href="#" id="joinToWaitingList">{{index .translation "Join"}}</a>
                    {{end}}
                </li>
            {{end}}
        </ul>
        <ul class="nav-user navbar-nav">
            {{if .user.IsLogined}}
                <li class="nav-item">
                    <a href="/{{.language}}/user/">{{.user.Email}}</a>
                </li>
                <li class="nav-item">
                    <a href="/{{.language}}/logout/">{{index .translation "LogoutButton"}}</a>
                </li>
            {{else}}
                <li class="nav-item">
                    <a href="/{{.language}}/login/">{{index .translation "LoginNoun"}}</a>
                </li>
                <li class="nav-item">
                    <a href="/{{.language}}/register/">{{index .translation "RegisterNoun"}}</a>
                </li>
            {{end}}
        </ul>
    </div>
    <!--tłumaczenie aria-label-->
    <!--js naprawić toolbar-->
    <button class="navbar-toggler order-first" type="button" data-bs-toggle="collapse" data-bs-target="#nav-menu" aria-controls="nav-menu" aria-expanded="false" aria-label="navigation-toggle">
        <span class="navbar-toggler-icon"></span>
    </button>
</nav>