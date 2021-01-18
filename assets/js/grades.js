$(()=>{
    if($('.student-grades').length){
        url = config.apiAddress+'student/'+userID;
        createSpinner($('.spinner-field'));
        sendHttpRequest('GET',url,null,null,generateGrades);
    }
    if($('.teacher-grades').length){
        url = config.apiAddress+'grade/getAll';
        createSpinner($('.spinner-field'));
        getJWT((response)=>{
            sendHttpRequest("GET",url,null,response.jwt,addGradeList)
        });
    }
});


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