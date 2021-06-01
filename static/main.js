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

    $(".changeTerr").on("click", function () {
        $.post("/change", {
            territory: $(this).attr('terr'),
            build: $(this).attr('id'),
        }, function(result){
            if (result.error != null) {
                errorAlert(result.message);
            } else {
                successAlert(result.message);
            }
        });
    });

    let surveyBtn = $("#buttonSurvey");
    surveyBtn.click(() => {

        var selY = $("#selectYear").val()
        var selU = $("#selectUniversity").val()
        var selD = $("#selectDorm").val()
        

        if (selY == "ch") {
            $("#selectYear").addClass("is-invalid")
            if (selU == "ch") {
                $("#selectUniversity").addClass("is-invalid")
            }
            return
        }
        if (selU == "ch") {
            $("#selectUniversity").addClass("is-invalid")
            return
        }
        var chk1 = false
        var chk2 = false
        var chk3 = false
        var chk4 = false
        var chk5 = false
        var chk6 = false
        var chk7 = false
        var chk8 = false

        if ($("#chk1").is(':checked')) {chk1 = true;}
        if ($("#chk2").is(':checked')) {chk2 = true;}
        if ($("#chk3").is(':checked')) {chk3 = true;}
        if ($("#chk4").is(':checked')) {chk4 = true;}
        if ($("#chk5").is(':checked')) {chk5 = true;}
        if ($("#chk6").is(':checked')) {chk6 = true;}
        if ($("#chk7").is(':checked')) {chk7 = true;}
        if ($("#chk8").is(':checked')) {chk8 = true;}

        $.post("/survey", {
            selY: selY,
            selU: selU,
            rangeCum: $('#rangeMarkCum').val(),
            rangeCorp: $('#rangeMarkCorp').val(),
            chk1: chk1,
            chk2: chk2,
            chk3: chk3,
            chk4: chk4,
            chk5: chk5,
            chk6: chk6,
            chk7: chk7,
            chk8: chk8,
            selD: selD,
        }, function(result){
            if (result.error != null) {
                errorAlert(result.message);
            } else {
                successAlert(result.message);
            }
        });
    });

    // Signing in
    loginBtn.click(() => {
        $("#modalClose").click();
        var name = $("#username").val();
        var password = $("#password").val();
        var check = false
        
        if ($("#check").is(':checked')) {
            check = true;
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
            }
        });
    });

    function successAlert(message) {
        document.getElementById("errorMenu").innerHTML = "<p class=\"bg-success dropdown-item text-white font-weight-bold\">"+message+"</p>";
        setTimeout(() => window.location.reload(), 1000);
    };

    function errorAlert(message) {
        document.getElementById("errorMenu").innerHTML = "<p class=\"bg-danger dropdown-item text-white font-weight-bold\">"+message+"</p>";
        setTimeout(() => $( "#errorMenu" ).load(window.location.href + " #errorMenu" ), 3500);
    };

});
