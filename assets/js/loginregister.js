$(()=>{
    if($('#login-form').length){
        var url = $(location).attr('pathname');
        var array = url.split('/');
        var last = array[array.length-1];
        if(last == 'activated'){
            $('#modal-alert .modal-body').text(translation.activated);
            $('#modal-alert').modal('show');
        }else
        if(last == 'invalidToken' || last == 'notFound'){
            $('#modal-alert .modal-body').text(translation.invalidToken);
            $('#modal-alert').modal('show');
        }
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
});

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