<header>
    <h1 class="logo">{{index .translation "H1Banner"}}</h1>
    <nav>
        <ul class="page-nav">
            <li><a href="/{{.language}}/showstudents/">{{index .translation "MenuDisplay"}}</a></li>
            <li><a href="/{{.language}}/addstudents/">{{index .translation "MenuAdd"}}</a></li>
            <li><a href="/{{.language}}/deletestudents/">{{index .translation "MenuDelete"}}</a></li>
            <li><a href="/{{.language}}/editstudents/">{{index .translation "MenuEdit"}}</a></li>
            <li><a href="/{{.language}}/register/">{{index .translation "MenuRegister"}}</a></li>
        </ul>
    </nav>
</header>