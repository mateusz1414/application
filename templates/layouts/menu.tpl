<header>
    <h1 class="logo">{{index .translation "H1Banner"}}</h1>
    <nav>
        <ul class="page-nav">
            <li><a href="/{{.language}}/showstudents/">{{index .translation "MenuDisplay"}}</a></li>
            {{if (eq .permission "dean")}}<li><a href="/{{.language}}/addstudents/">{{index .translation "MenuAdd"}}</a></li>{{end}}
            {{if (eq .permission "dean")}}<li><a href="/{{.language}}/deletestudents/">{{index .translation "MenuDelete"}}</a></li>{{end}}
            {{if (eq .permission "dean")}}<li><a href="/{{.language}}/editstudents/">{{index .translation "MenuEdit"}}</a></li>{{end}}
            {{if (eq .permission "dean")}}<li><a href="/{{.language}}/dean/">{{index .translation "MenuPermission"}}</a></li>{{end}}
            {{if (eq .permission "student")}}<li><a href="/{{.language}}/grade/">{{index .translation "MenuGrades"}}</a></li>{{end}}
            {{if (eq .permission "teacher")}}<li><a href="/{{.language}}/grades/">{{index .translation "MenuGrades"}}</a></li>{{end}}
            <li><a href="/{{.language}}/teacher/">{{index .translation "MenuTeachers"}}</a></li>
            <li><a href="/{{.language}}/departament/">{{index .translation "MenuDepartaments"}}</a></li>
            <li><a href="/{{.language}}/register/">{{index .translation "MenuRegister"}}</a></li>
        </ul>
        {{if .isLogined}}
{{index .translation "DisplayLogged"}}
<button class="logoutButton">{{index .translation "LogoutButton"}}</button>
{{end}}
    </nav>
</header>