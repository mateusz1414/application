{{define "content"}}  
 <form action="/action/add/" method="post">
    <table class="content-table">
        <thead>
            <tr>
                <th>Imie:</th>
                <td><input type="text" name="studentFirstName"></td>
            </tr>
            <tr>
                <th>Nazwisko:</th>
                <td><input type="text" name="studentLastName"></td>
            </tr>
            <tr>
                <th>Data urodzenia:</th>
                <td><input type="date" name="studentDateOfBrith"></td>
            </tr>
            <tr>
                <th>Wydział:</th>
                <td><input type="text" name="studentFaciulty"></td>
            </tr>
            <tr>
                <th>Płeć:</th>
                <td>
                    <select name="studentGender">
                        <option value="0">Mężczyzna</option>
                        <option value="1">Kobieta</option>
                    </select>
                </td>
            </tr>
            <tr>
                <td colspan="2" style="text-align: center;"><button>DODAJ</button></td>
            </tr>
        </thead>
    </table>
</form>
{{end}}