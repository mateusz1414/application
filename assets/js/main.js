$(()=> {
    if($('.student-table').length){
        loadStudents(createRow,"getAll",$('.student-table'));
    }
    if($('.grades-table').length){
        loadGrades(createRowOfGrades,"getAllOfStudent",$('.grades-table'));
    }
    if($('.teachers-table').length){
        loadTeachers(createRowOfTeachers,"GetOnce",$('.teachers-table'));
    }
    if($('.grades-table-teacher').length){
        loadGradesForTeacher(createRowOfGradesForTeacher,"getAllOfStudents",$('.grades-table-teacher'));
    }
    if($('.permission-table').length){
        loadPermissionForDean(createRowOfNoPermission,"/",$('.permission-table'));
    }
    if($('.departament-table').length){
        loadDepartaments(createRowOfDepartaments,"/",$('.departament-table'));
    }
    if($('.edit-form').length){
        var url = $(location).attr('pathname').split('/');
        var id = url[url.length-2];
        loadStudents(editForm,id)
    }
    if($('.add-form').length){
        addStudent();
    }
    if($('.registerButton').length){
        loginRegister('register');
    }
    if($('.loginButton')){
        loginRegister('login');
    }
    if($('.logoutButton')){
        logoutPanel();
    }
});

function loadStudents(fun,id,contentTable){
    if(contentTable!=undefined){
        $('.content-table tbody tr').remove();
        var count = 6;
        if($('.edit-table').length || $('.delete-table').length){
            count=7;
        }
        //var loadWheel = $('<tr><td class="loadWheel" colspan="'+count+'"><img src="/css/img/wheel.png"></td></tr>');
        //contentTable.append(loadWheel);
    }
    var XHR = new XMLHttpRequest();
    XHR.onload = function(){
        if(XHR.status==200){
            var response = JSON.parse(XHR.response);
            if(id=="getAll"){
                var index = 1;
                response.Student.forEach(student => {
                    fun(contentTable,student,index);
                    index++;
                });
                //loadWheel.remove();
            }else
            {
                fun(response.Students[0]);
            }
        }else
        {

        }
    }
    XHR.open('GET',config.apiAddress+'student/'+id);
    //XHR.setRequestHeader("Access-Control-Allow-Origin","https://studenci.herokuapp.com/");
    XHR.send();
}

function createRow(contentTable,student,index){
    var tableRow = $('<tr>');
    //loadWheel.before(tableRow);
    var td;
    for(var key in student){
        var text;
        if(key == "idstudent"){
            text = index;
        }else
        if(key == "sex"){
            if(student[key]==0){
                text = "Mężczyzna";
            }else
            {
                text = "Kobieta";
            }
        }
        else 
        if(key == "IDDepartament"){
            text =student.Departament["departament"];
            
        }else if(key!= "Departament")
        {
            text = student[key];
        }
        if(key != "Departament"){
        td = $('<td>').text(text);
        tableRow.append(td);
        }
        contentTable.append(tableRow);
    }
    if(contentTable.find(".thEditButton").length){
        td = $('<td>');
        tableRow.append(td);
        var a = $('<a>');
        var url = $(location).attr('pathname')
        var language = url.split("/")[1];
        td.append(a);
        var button = $('<button>').text('Edytuj');
        button.on('click',()=>{
            getSession('jwt',(jwt,url)=>{
                $(location).prop('href', url);
            },'/'+language+'/editstudentform/'+student['studentID']+'/');
        })
        a.append(button);
    }
    if(contentTable.find(".thDelButton").length){
        td = $('<td>');
        tableRow.append(td);
        button = $('<button>').text('Usuń');
        td.append(button);
        button.on('click',()=>{
            getSession("jwt",actionStudent,student,'DELETE')
            //napraw api jak nie znajdzie nie ma wywalac error
        });
    }
        
}

function getSession(key,callback,student,method){
    var XHR = new XMLHttpRequest;
    XHR.onload = () => {
        if(XHR.status==200){
            var response = JSON.parse(XHR.response);
            if(response[key]==null){
                showError("Nie jesteś zalogowany","Kliknij poza okno aby się zalogować",()=>{
                    var url = $(location).attr('pathname');
                    var language = url.split("/")[1];
                    $(location).prop('href', config.serverAddress+language+'/register/');
                });
                return
            }
            callback(response[key],student,method);
        }else
        {
            showError("Upps","Coś poszło nie tak");
        }
    }
    XHR.open("GET",config.serverAddress+'session/'+key);
    XHR.send();
}

function setSession(jwt,permission){
    var XHR = new XMLHttpRequest; 
    XHR.onload = () => {
        if(XHR.status==200){
            location.reload();
        }else
        {
            showError("Upps","Coś poszło nie tak");
        }
    }
    XHR.open("POST",config.serverAddress+'session/');
    XHR.setRequestHeader("Content-Type", "application/json");
    XHR.send(JSON.stringify([{
        key:"jwt",
        value:jwt
    },
    {
        key:"permission",
        value:permission
    }
]));
}

function actionStudent(jwt,student,method){
    var XHR = new XMLHttpRequest;
    XHR.onload=() =>{
        if(XHR.status == 200){
            if(method=="DELETE"){
                loadStudents(createRow,"getAll",$('.student-table'));
            }else
            {
                var url = $(location).attr('pathname');
                var language = url.split("/")[1];
                $(location).prop('href', config.serverAddress+language+'/');
            }
        }else
        if(XHR.status == 401){
            showError("Twoja sesja wygasła","Kliknij poza okno aby się zalogować",()=>{
                logout();
            });
        }
        else
        {
            showError("Wystąpił problem serwera","Prosimy spróbować ponownie");
        }
    }
    if(method == 'POST'){
        XHR.open(method,config.apiAddress+"student/");
    }else
    {
        XHR.open(method,config.apiAddress+"student/"+student['studentID']);
    }
    XHR.setRequestHeader('Authorization','Bearer '+jwt);
    XHR.setRequestHeader("Content-Type", "application/json");
    XHR.send(JSON.stringify(student));

}

function editForm(student){
    $('#studentFirstName').val(student.name);
    $('#studentLastName').val(student.surname);
    $('#studentDateOfBrith').val(student.dateofbrith);
    $('#studentFaciulty').val(student.departament);
    if(student.sex==1){
        $('#studentWomen').prop("selected",true);
    }
    $('#editButton').on('click',() =>{
        student.name =  $('#studentFirstName').val();
        student.surname =  $('#studentLastName').val();
        student.dateofbrith =  $('#studentDateOfBrith').val();
        student.departament = $('#studentFaciulty').val();
        student.sex = $('#studentGender').val();
        getSession('jwt',actionStudent,student,'PUT');
    });
}

function addStudent(){
    $('#addButton').on('click',() =>{
        student = {
            name: $('#studentFirstName').val(),
            surname: $('#studentLastName').val(),
            dateofbrith: $('#studentDateOfBrith').val(),
            departament: $('#studentFaciulty').val(),
            sex: $('#studentGender').val(),
        }
        getSession("jwt",actionStudent,student,"POST");
    })
}

function showError(message1,message2,callback){
    $('body').append('<div id="widok"></div>');
    $('body').append('<div id="err"></div>');
    var paragraph = $('<p>',{
        text: message1,
        class: "errorFirst"
    });
    $('#err').append(paragraph);
    paragraph = $('<p>',{
        text: message2,
        class: "errorSecond"
    });
    $('#err').append(paragraph);
    $("#err").animate({
        width:"25%",
        height:"25%",
        fontSize:"24px",
    }, 500);
    $('#widok').click(()=>{
        $('#widok').remove();
        $('#err').remove();
        callback();
    })
}

function send(user,action,ifFalse,ifTrue){
    var XHR = new XMLHttpRequest;
    XHR.onload = ()=>{
        var response = JSON.parse(XHR.response); 
        if(XHR.status==200){
            if(ifTrue==undefined){
                sendMessage(response,"register");
            }else
                ifTrue(response.AuthToken,response.Permission);
        }else
        {
            ifFalse(response,action);
        }
    };
    XHR.open('POST',config.apiAddress+'user/'+action);
    XHR.setRequestHeader("Content-Type", "application/json");
    XHR.send(JSON.stringify(user))

}

function loginRegister(what){
    $('.'+what+'Button').on('click',() =>{
        user = {
            login:  $('.'+what+'User').val(),
            password: $('.'+what+'Password').val(),
            confirmpassword: $('.'+what+'ConfirmPassword').val(),
        };
        if(what=='login'){
            send(user,what,sendMessage,setSession);
        }else
            send(user,what,sendMessage);
    });
}


function sendMessage(response,action){
    $('.'+action+'MessageFirst').text(response.Message);
    $('.'+action+'MessageSecond').text(response.ErrorCode);
}

function  logoutPanel(){
    $('.logoutButton').on('click',logout);
}

function logout(){
    XHR = new XMLHttpRequest;
    XHR.onload = () =>{
        if(XHR.status!=200){
            showError("Wystąpił problem serwera","Prosimy spróbować ponownie");
        }else
        {
            var url = $(location).attr('pathname');
            var language = url.split("/")[1];
            $(location).prop('href', config.serverAddress+language+'/register/');
        }
    }
    XHR.open('DELETE',config.serverAddress+'session/jwt');
    XHR.send();
}
function loadGrades(){
    getSession('jwt',createGrades,null,'GET');
}

function createRowOfGrades(contentTable,response){
        response.Value.forEach(grade => {
            var tr=$('<tr>');
            var td=$('<td>').text(grade.Subjects[0].name);
            tr.append(td);
            var td=$('<td>').text(grade.Teachers[0].Name+' '+grade.Teachers[0].Surname);
            tr.append(td);
            var td=$('<td>').text(grade.grade);
            tr.append(td);
            contentTable.append(tr);
        });
        

        
}

function createGrades(jwt,body,method){
    var XHR = new XMLHttpRequest;
    XHR.onload = () =>{
        if(XHR.status==200){
            var response = JSON.parse(XHR.response);
            createRowOfGrades($(".grades-table"),response)
        }
    }
    XHR.open(method,config.apiAddress+'grade/getAllOfStudent');
    XHR.setRequestHeader("Authorization", "Bearer " +jwt);
    XHR.send(body);
}

function loadTeachers(contentTable){
    if(contentTable!=undefined){
        createTeachers('GET');
    }
    
}

function createRowOfTeachers(contentTable,response){
    response.Teacher.forEach(teacher => {
        var tr=$('<tr>');
        console.log(teacher)
        var td=$('<td>').text(teacher.name);
        tr.append(td);
        var td=$('<td>').text(teacher.surname);
        tr.append(td);
        var td=$('<td>').text(teacher.Subject.subject);
        tr.append(td);
        contentTable.append(tr);
    });
    

    
}

function createTeachers(method){
var XHR = new XMLHttpRequest;
XHR.onload = () =>{
    if(XHR.status==200){
        var response = JSON.parse(XHR.response);
        createRowOfTeachers($(".teachers-table"),response)
    }
}
XHR.open(method,config.apiAddress+'teacher/getAll');
XHR.send();
}

function loadGradesForTeacher(){
    getSession('jwt',createRowGradesForTeacher,null,'GET');
}

function createRowOfGradesForTeacher(contentTable,response){
console.log(response);
        response.Value.forEach(grades => {
            var tr=$('<tr>');
            var td=$('<td>').text(grades.name+' '+grades.surname);
            tr.append(td);
            var listgrades="";
            grades.Grades.forEach(grade => {
                listgrades += grade.Grade+",";
            });
            
            var td=$('<td>').text(listgrades.substr(0,listgrades.length-1));
            tr.append(td);
            contentTable.append(tr);
        });
        

        
}

function createRowGradesForTeacher(jwt,body,method){
    var XHR = new XMLHttpRequest;
    XHR.onload = () =>{
        if(XHR.status==200){
            var response = JSON.parse(XHR.response);
            createRowOfGradesForTeacher($(".grades-table-teacher"),response)
        }
    }
    XHR.open(method,config.apiAddress+'grade/getAllOfStudents');
    XHR.setRequestHeader("Authorization", "Bearer " +jwt);
    XHR.send(body);
}
function loadPermissionForDean(){
    getSession('jwt',createRowPermission,null,'GET');
}

function createRowOfNoPermission(contentTable,response){
console.log(response);
        response.User.forEach(permission => {
            var tr=$('<tr>');
            var td=$('<td>').text(permission.login);
            tr.append(td);
            contentTable.append(tr);
        });
        

        
}

function createRowPermission(jwt,body,method){
    var XHR = new XMLHttpRequest;
    XHR.onload = () =>{
        if(XHR.status==200){
            var response = JSON.parse(XHR.response);
            createRowOfNoPermission($(".permission-table"),response)
        }
    }
    XHR.open(method,config.apiAddress+'dean/');
    XHR.setRequestHeader("Authorization", "Bearer " +jwt);
    XHR.send(body);
}
//////////////////////////////////
function loadDepartaments(contentTable){
    if(contentTable!=undefined){
        createDepartaments('GET');
    }
    
}

function createRowOfDepartaments(contentTable,response){
    console.log(response.Departament)
    response.Departament.forEach(departament => {
        var tr=$('<tr>');
        var td=$('<td>').text(departament.name);
        tr.append(td);
        //var td=$('<td>').text(departament.Subjects[0].name);
        //tr.append(td);
        var listsubjects="";
        departament.Subjects.forEach(depart => {
                listsubjects += depart.name+",";
            });
            
            var td=$('<td>').text(listsubjects.substr(0,listsubjects.length-1));
            tr.append(td);
        contentTable.append(tr);
    });
    

    
}

function createDepartaments(method){
var XHR = new XMLHttpRequest;
XHR.onload = () =>{
    if(XHR.status==200){
        var response = JSON.parse(XHR.response);
        createRowOfDepartaments($(".departament-table"),response)
    }
}
XHR.open(method,config.apiAddress+'departament/');
XHR.send();
}