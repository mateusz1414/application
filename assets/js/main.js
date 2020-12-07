$(()=> {
    if($('.student-table').length){
        loadStudents(createRow,"getAll",$('.student-table'));
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
        var loadWheel = $('<tr><td class="loadWheel" colspan="'+count+'"><img src="/css/img/wheel.png"></td></tr>');
        contentTable.append(loadWheel);
    }
    var XHR = new XMLHttpRequest();
    XHR.onload = function(){
        if(XHR.status==200){
            var response = JSON.parse(XHR.response);
            if(id=="getAll"){
                var index = 1;
                response.Students.forEach(student => {
                    fun(contentTable,student,index,loadWheel);
                    index++;
                });
                loadWheel.remove();
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

function createRow(contentTable,student,index,loadWheel){
    var tableRow = $('<tr>');
    loadWheel.before(tableRow);
    var td;
    for(var key in student){
        var text;
        if(key == "studentID"){
            text = index;
        }else
        if(key == "sex"){
            if(student[key]==0){
                text = "Mężczyzna";
            }else
            {
                text = "Kobieta";
            }
        }else
        {
            text = student[key];
        }
        td = $('<td>').text(text);
        tableRow.append(td);
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

function setSession(key,value){
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
    XHR.send(JSON.stringify({
        key:key,
        value:value
    }));
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
                ifTrue('jwt',response.AuthToken);
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