{{define "content"}}
<table class="content-table">
    <thead>
        <tr>
            <th>Imie:</th>
            <td><input type="text"></td>
        </tr>
        <tr>
            <th>Nazwisko:</th>
            <td><input type="text"></td>
        </tr>
        <tr>
            <th>Data urodzenia:</th>
            <td><input type="datetime-local"></td>
        </tr>
        <tr>
            <th>Wydział:</th>
            <td><input type="text"></td>
        </tr>
        <tr>
            <th>Płeć:</th>
            <td><input type="text"></td>
        </tr>
        <tr>
            <td colspan="2" style="text-align: center;"><button>DODAJ</button></td>
        </tr>
    </thead>
</table>
{{end}}