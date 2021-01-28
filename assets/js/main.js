$(()=> {
    if($('.student-table').length){
        makeTable('student');
    }
    if($('.teacher-table').length){
        makeTable('teacher');
    }
    if($('#waiting-modal').length){
        generateDepartamentsSelect($('.waiting-departament'));
        studentModal("sendRequest");
    }
    if($('#update-modal').length){
        generateDepartamentsSelect($('.waiting-departament'));
        studentModal('');
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
        console.log(body);
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
            $('.waiting-departament option').removeAttr("selected");
            $('.waiting-sex option').removeAttr("selected");
            $('#update-modal form')[0].reset();
            $('.waiting-name').val(person.name);
            $('.waiting-surname').val(person.surname);
            $('.waiting-dob').val(person.dob);
            $('.waiting-id').val(person.studentID);
            $('.waiting-departament option[value='+person.departament.departamentID+']').attr('selected','selected');
            $('.waiting-sex option[value='+person.sex+']').attr('selected','selected');
            $('#update-modal').modal('show');
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

function getLanguage(){
    var url = $(location).attr('pathname');
    var language = url.split("/")[1];
    return language;
}

function getJWT(next){
    var url = config.serverAddress + 'session/jwt';
    sendHttpRequest("GET",url,null,null,next);
}

function alertlog(response){
    console.log(response)
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
        case "Not active":
            $('#modal-alert .modal-body').html(translation.shouldActive+' <a href="'+response.activationURL+'">'+translation.click+'</a> '+translation.toActive);
            $('#modal-alert').modal('show');
        break;
        default:
            $('.alert-danger').text(translation.serverError);
    }
    $('.alert-danger').show();
}

