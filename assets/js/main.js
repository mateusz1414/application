$(()=> {
    if($('.student-table').length){
        makeTable('student');
    }
    if($('.teacher-table').length){
        makeTable('teacher');
    }
    if($('#login-form').length){
        $('#login-form').submit((event)=>{
            event.preventDefault();
            loginRegister("login");
        })
    }
    if($('#register-form').length){
        $('#register-form').submit((event)=>{
            event.preventDefault();
            loginRegister("register");
        })
    }
    if($('.student-grades').length){
        url = config.apiAddress+'student/'+userID;
        createSpinner($('.spinner-field'));
        sendHttpRequest('GET',url,null,null,generateGrades);
        //loadStudents(createRow,"teacher","getAll",$('.teacher-table'));
    }
    if($('.teacher-grades').length){
        url = config.apiAddress+'grade/getAll';
        createSpinner($('.spinner-field'));
        getJWT((response)=>{
            sendHttpRequest("GET",url,null,response.jwt,addGradeList)
        });
    }
    if($('#waiting-modal').length){
        generateDepartamentsSelect($('.waiting-departament'));
        studentModal("sendRequest");
    }
    if($('#update-modal').length){
        generateDepartamentsSelect($('.waiting-departament'));
        studentModal('');
    }
    if($('.dean-functions').length){
        applicationListRequest();
    }
});

function studentModal(what){
    $('.student-modal .btn-primary').click(()=>{
        var body = {
            name: $('.waiting-name').val(),
            surname: $('.waiting-surname').val(),
            dob: $('.waiting-dob').val(),
            departamentID: parseInt($('.waiting-departament').val()),
            sex: parseInt($('.waiting-sex').val()),
        }
        if(body.name == '' || body.surname == '' || body.dob == '' || body.departament == 'undefined' || body.sex == 'undefined')
        {
            alertlog({errorCode: "Not all"});
            return;
        }
        var url = config.apiAddress+'student/'+what;
        var method = "POST";
        if(what==''){
            method='PUT';
            url+=$('.waiting-id').val();
        }
        getJWT(response=>{
            sendHttpRequest(method,url,body,response.jwt,()=>{
                $('.student-modal').modal('hide');
                if(what!=''){
                    $('#modal-alert .modal-body').text(translation.addOnWaiting);
                    $('#modal-alert').modal('show');
                }else
                {
                    refreshTable();
                }
            },alertlog);
        });
    });
}

function makeTable(who){
    url = config.apiAddress+who+'/getAll';
    createSpinner($('.spinner-col'));
    sendHttpRequest('GET',url,null,null,generateTable)
}

function applicationListRequest(){
    getJWT(response=>{
        var url = config.apiAddress+'management/applicationList' ;
        sendHttpRequest('GET',url,null,response.jwt,createApplicationList)
    });
}

$('#joinToWaitingList').click(()=>{
    $('#waiting-modal').modal('show');
});

function generateDepartamentsSelect(select){
    var url = config.apiAddress+'departament/getAll';
    sendHttpRequest('GET',url,null,null,response=>{
        console.log(response);
        response.departaments.forEach(departament=>{
            var option = $('<option value="'+departament.departamentID+'">'+departament.name+'</option>');
            select.append(option);
        });
    });
}

function sendHttpRequest(method,url,body,token,ifTrue,ifFalse){
    var XHR = new XMLHttpRequest();
    XHR.onload = ()=>{
        var response = JSON.parse(XHR.response);
        if(XHR.status==200){
            ifTrue(response);
        }else
        if(XHR.status==401){
            $('#waiting-modal').modal('hide');
            $('#modal-alert .modal-body').text(translation.sessionExpired);
            $('#modal-alert').modal('show');
            $('#modal-alert').on("hidden.bs.modal",()=>{
                $(location).prop('href',config.serverAddress+getLanguage()+'/logout/');
            });
        }else
        {
            ifFalse(response);
        }
    }
    XHR.open(method,url);
    XHR.setRequestHeader("Content-Type", "application/json");
    XHR.setRequestHeader("Authorization", "Bearer " + token);
    XHR.send(JSON.stringify(body));
}

function createSpinner(where){
    var spinner = $('<div class="spinner-border spinner" role="status"><span class="visually-hidden">Loading...</span></div>');
    where.append(spinner);
}

function generateTable(response){
    var array;
    var table = $('.content-table');
    var spinner = $('.spinner-row');
    var index = 1;
    if(response.students!=undefined){
        array =  response.students;
    }else
    {
        array = response.teachers;
    }
    array.forEach(person => {
        createRow(table,person,index,spinner);
        index++;
    });
    spinner.remove();

}

function refreshTable(){
    $('.content-table').find('tbody').text('');
    $('.content-table').find('tbody').append('<tr class="spinner-row"><td colspan="100%" class="justify-content-center spinner-col" ></td></tr> ');
    makeTable('student');
}

function createRow(contentTable,person,index,loadWheel){
    var tableRow = $('<tr>');
    var td;
    var text;
    loadWheel.before(tableRow);
    for(var key in person){
        switch(key){
            case "studentID":
                text=index
            break;
            case "sex":
                if(person[key]==0){
                    text = translation.male;
                }else
                {
                    text = translation.female;
                }
            break;
            case "departament":
                text = person[key].name;
            break;
            case "subject":
                text = person[key].name;
            break;
            default:
                text = person[key];
        }
        td = $('<td>').text(text);
        tableRow.append(td);
    }
    if(contentTable.attr('id') == 'modify-table'){
        var editBtn = $('<td><button class="btn btn-secondary">'+translation.edit+'</button></td>')
        tableRow.append(editBtn);
        var deleteBtn = $('<td><button class="btn btn-secondary">'+translation.delete+'</button></td>')
        tableRow.append(deleteBtn);
        editBtn.find('.btn').click(()=>{
            console.log($('.waiting-departament option[value='+person.departament.departamentID+']'));
            $('#update-modal').modal('show');
            $('.waiting-name').val(person.name);
            $('.waiting-surname').val(person.surname);
            $('.waiting-dob').val(person.dob);
            $('.waiting-id').val(person.studentID);
            $('.waiting-departament option[value='+person.departament.departamentID+']').attr('selected','selected');
            $('.waiting-sex option[value='+person.sex+']').attr('selected','selected');
        });
        deleteBtn.find('.btn').click(()=>{
            getJWT(response=>{
                var url = config.apiAddress+'student/'+person.studentID;
                sendHttpRequest('DELETE',url,null,response.jwt,()=>{
                    refreshTable();
                });
            });
        });
    }
}

function createApplicationList(response){
    response.students.forEach(person=>{
        var element = $('<div class="waiting-element col-12 row"></div>');
        var name ='';
        if(person.name =='' && person.surname ==''){
            name = 'Brak nazwy'
        }else
        {
            name = person.name+' '+person.surname;
        }
        var col = $('<div class="col col-8 list-name" data-bs-toggle="collapse" href="#collapse-user-'+person.studentID+'">'+name+'</div>');
        var btns = $('<div class="col col-2"><button type="button" class="btn btn-success">✔</button></div><div class="col col-2"><button type="button" class="btn btn-danger">X</button></div>');
        var collapse =$('<div class="collapse" id="collapse-user-'+person.studentID+'"></div>'); 
        for(var key in person){
            var index=false;
            if(key=='dob' || key=='sex' || key=='departament'){
                var value =person[key];
                switch(key){
                    case 'sex':
                        if(value==0){
                            value=translation.male;
                        }else
                        {
                            value=translation.female;
                        }
                    break;
                    case 'departament':
                        value = value.name;
                    break;
                }
                one = $('<div class="row"><div class="col">'+translation[key]+'</div><div class="col">'+value+'</div></div>');
                collapse.append(one)
            }
        }
        element.append(col);
        element.append(btns);
        element.append(collapse);
        $('.waiting-list').append(element);
        btns.find('.btn-success').click(()=>{
            getJWT(response=>{
                var url = config.apiAddress+'management/'+person.studentID;
                sendHttpRequest('PUT',url,null,response.jwt,()=>{
                    $('.waiting-list').text('');
                    $('.waiting-list').append('<div class="col waiting-element">'+translation.waitingList+'</div>');
                    applicationListRequest();
                });
            });
        });
        btns.find('.btn-danger').click(()=>{
            getJWT(response=>{
                var url = config.apiAddress+'management/'+person.studentID;
                sendHttpRequest('DELETE',url,null,response.jwt,()=>{
                    $('.waiting-list').text('');
                    $('.waiting-list').append('<div class="col waiting-element">'+translation.waitingList+'</div>');
                    applicationListRequest();
                });
            });
        });
    });
}

function getLanguage(){
    var url = $(location).attr('pathname');
    var language = url.split("/")[1];
    return language;
}

function getJWT(next){
    var url = config.serverAddress + 'session/jwt';
    sendHttpRequest("GET",url,null,null,next);
}

function generateGrades(response){
    var student = response.students[0];
    $('.student-name').text(student.name + ' ' + student.surname);
    $('.student-name').removeClass('spinner-field');
    var url = config.apiAddress + 'grade/myGrades';
    getJWT((response)=>{
        sendHttpRequest("GET",url,null,response.jwt,createGradesList);
    });  
}

function createGradesList(response){
    var subjects = response.subjects;
    subjects.forEach(subject=>{
        var spinner = $('.spinner-field')
        var li = $('<li class="list-group-item list-group-item-dark"></li>')
        var subjectName = $('<div class="list-group-item list-group-item-secondary subject">'+subject.name+'</div>');
        var grades = $('<ul class="list-group list-group-horizontal"></ul>');
        var sum = 0;
        spinner.before(li);
        li.append(subjectName);
        li.append(grades);
        subject.grades.forEach(grade=>{
            var gradeLi = $('<li class="list-group-item list-group-item-dark">'+grade.value+'</li>');
            grades.append(gradeLi);
            sum+=grade.value;
        });
        var average = $('<div class="list-group-item list-group-item-secondary average">'+translation.average+' '+sum/subject.grades.length+'</div>');
        li.append(average);
    });
    $('.spinner-field').remove();
}


function addGradeList(response){
    $('.subject-name-add').text(response.subjectName);
    $('.subject-name-add').removeClass('spinner-field');
    var spinner = $('.spinner-field');
    response.students.forEach(student=>{
        var  li = $('<li class="list-group-item list-group-item-dark"></li>');
        var div = $('<div class="list-group-item list-group-item-dark student"></div>');
        var row = $('<div class="row"></div>');
        var person = $('<div class="col col-md-6">'+student.name+' '+student.surname+'</div>');
        var addGrade = $('<div class="col col-md-6 text-end"><input id="student-'+student.studentID+'-input" type="number" min="1" max="6"><button id="student-'+student.studentID+'" class="add-grade-button">'+translation.addGrade+'</button></div>');
        var ul = $('<ul class="list-group list-group-horizontal" id="student-'+student.studentID+'-ul"></ul>');
        spinner.before(li);
        li.append(div);
        div.append(row);
        row.append(person);
        row.append(addGrade);
        var sum = ulWriteGrades(student.grades,ul);
        li.append(ul);
        var avg = sum/student.grades.length;
        if(isNaN(avg)){
            avg=0;
        }
        var average = $('<div class="list-group-item list-group-item-dark average" id="student-'+student.studentID+'-avg">'+translation.average+' '+avg+'</div>');
        li.append(average);
        $('#student-'+student.studentID).click(event=>{
            var studentID = event.target.id.split('-')[1];
            var input = $('#'+event.target.id+'-input');
            var grade = input.val();
            if(grade<1 || grade>5){
                $('#modal-alert .modal-body').text(translation.incorrectGrade);
                $('#modal-alert').modal('show');
                return;
            }
            var url = config.apiAddress+"grade/";
            var body = {
                value: parseInt(grade),
                studentID: parseInt(studentID),
            }
            getJWT((response)=>{
                sendHttpRequest("POST",url,body,response.jwt,generateStudentGrade,showModalGrade)
            })
        });
    });
    spinner.remove();
}

function ulWriteGrades(grades,ul){
    var sum=0;
    grades.forEach(grade=>{
        var point = $('<li class="list-group-item list-group-item-dark">'+grade.value+'</li>');
        sum += grade.value;
        ul.append(point);
    });
    return sum;
}

function generateStudentGrade(response){
    var ul = $('#student-'+response.studentID+'-ul');
    ul.text('');
    var avg = ulWriteGrades(response.studentGrades,ul)/response.studentGrades.length;
    $('#student-'+response.studentID+'-avg').text(translation.average+' '+avg);
    
}

function showModalGrade(response){
    $('#modal-alert .modal-body').text(translation.incorrectGrade);
    $('#modal-alert').modal(response.message);
}

function loginRegister(what){
    user = {
        email:  $('.email').val(),
        password: $('.password').val(),
        confirmpassword: $('.confirm-password').val(),
    };
    var url = config.apiAddress+'user/'+what;
    sendHttpRequest('POST',url,user,null,login,alertlog);
}

function login(response){
    if(response.message=="Logged"){
        var body = [{
            'key': 'jwt',
            'value': response.authToken
        },
        {
            'key': 'email',
            'value': response.email
        },
        {
            'key': 'userID',
            'value': String(response.userID)
        },
        {
            'key': 'permissions',
            'value': response.permissions
        }
        ];
        sendHttpRequest('POST',config.serverAddress+'session/',body,null,()=>{
            $(location).prop('href',config.serverAddress+getLanguage()+'/');
        });
    }else
    {
        $(location).prop('href',config.serverAddress+getLanguage()+'/login/');
    }
}

function alertlog(response){
    switch(response.errorCode){
        case "Invalid":
            $('.alert-danger').text(translation.incorrectEmailOrPassword);
        break;
        case "Invalid email":
            $('.alert-danger').text(translation.incorrectEmail);
        break;
        case "Password is to short":
            $('.alert-danger').text(translation.iasswordIsTooShort);
        break;
        case "Password do not match":
            $('.alert-danger').text(translation.passwordsDoNotMatch);
        break;
        case "Email taken":
            $('.alert-danger').text(translation.busyEmail);
        break;
        case "Not all":
            $('.alert-danger').text(translation.allFields);
        break;
        case "On list":
            $('.alert-danger').text(translation.onList);
        break;
        default:
            $('.alert-danger').text(translation.serverError);
    }
    $('.alert-danger').show();
}