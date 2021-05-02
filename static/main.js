$("document").ready(() => {

    var localization = Cookies.get('loc');
    if (localization == undefined) {
        localization = 'UA';
    }
    Cookies.set('loc', localization, { expires: 365});

    $('map').imageMapResize();
    
    let loginBtn = $("#buttonLogin"); // Signing in
    let regBtn = $("#buttonReg"); // Signing up

    // var mapHeight
    // var mapWidth
    // $("#mainMap").on("click", function () {
    //     mapHeight = this.height;
    //     mapWidth = this.width;

    //     var elem = document.getElementById('mainMap');
    //     const rect = elem.getBoundingClientRect();
    //     const x = event.clientX - rect.left;
    //     const y = event.clientY - rect.top;
    //     console.log(x + "," + y + ",");
    //     //window.location.href = "/"
    // });

    $("#localization").on("click", function () {
        if (localization == "UA") {
            localization = "EN"
        } else {
            localization = "UA"
        }
        Cookies.remove('loc');
        Cookies.set('loc', localization, { expires: 365});
        window.location.reload()
    });

    // Signing in
    loginBtn.click(() => {
        $("#modalClose").click();
        var name = $("#username").val();
        var password = $("#password").val();
        var check
        
        if ($("#check").is(':checked')) {
            check = "true";
        } else {
            check = "false";
        }

        $.post("/login", {
            username: name,
            password: password,
            check: check,
        }, function(result){
            if (result.error != null) {
                errorAlert(result.message);
            } else {
                successAlert(result.message);
                $( "#full-menu" ).load(window.location.href + " #full-menu" );
            }
        });
    });

    regBtn.click(() => {
        $("#modalCloseReg").click();
        var email = $("#emailReg").val();
        var name = $("#usernameReg").val();
        var password = $("#passwordReg").val();

        $.post("/register", {
            email: email,
            username: name,
            password: password,
        }, function(result){
            if (result.error != null) {
                errorAlert(result.message);
            } else {
                successAlert(result.message);
                $( "#full-menu" ).load(window.location.href + " #full-menu" );
            }
        });
    });

    function successAlert(message) {
        document.getElementById("errorMenu").innerHTML = "<p class=\"bg-success dropdown-item text-white font-weight-bold\">"+message+"</p>";
        setTimeout(() => $( "#errorMenu" ).load(window.location.href + " #errorMenu" ), 2500);
    };

    function errorAlert(message) {
        document.getElementById("errorMenu").innerHTML = "<p class=\"bg-danger dropdown-item text-white font-weight-bold\">"+message+"</p>";
        setTimeout(() => $( "#errorMenu" ).load(window.location.href + " #errorMenu" ), 3500);
    };

});
