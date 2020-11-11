<!DOCTYPE html>
<html lang="pl">
    <head>
        {{include "layouts/head"}}
    </head>
    <body>
        {{include "layouts/menu"}}
        <main>
            <section>
                <div class="page-section">
                    {{template "content" .}}
                </div>
            </section>
            <aside>
                <div class="page-aside">
                    {{template "page-aside" .}}
                </div>
            </aside>
        </main>
        {{include "layouts/footer"}}
    </body>
</html>