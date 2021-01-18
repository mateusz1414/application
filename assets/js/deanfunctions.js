$(()=>{
    if($('.dean-functions').length){
        applicationListRequest();
    }
});

function applicationListRequest(){
    getJWT(response=>{
        var url = config.apiAddress+'management/applicationList' ;
        sendHttpRequest('GET',url,null,response.jwt,createApplicationList)
    });
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