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
        </main>
        {{include "layouts/footer"}}
    </body>
    <script type="text/javascript" src="/js/config.js"></script>
    <script src = "http://ajax.googleapis.com/ajax/libs/jquery/1.10.2/jquery.min.js"></script>
    <script type="text/javascript" src="/js/main.js"></script>
</html>