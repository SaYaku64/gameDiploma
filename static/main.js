$("document").ready(() => {

    var localization = Cookies.get('loc');
    if (localization == undefined) {
        localization = 'UA';
    }
    Cookies.set('loc', localization, { expires: 180});

    var mapHeight
    var mapWidth

    // $('img[usemap]').mapster('resize');

    // $('img[usemap]').mapster({
    //     fillColor: 'ff0000',
    //     stroke: true,
    //     singleSelect: true
    // });

	$('map').imageMapResize();
    
    $("#mainMap").on("click", function () {
        mapHeight = this.height;
        mapWidth = this.width;

        var elem = document.getElementById('mainMap');
        const rect = elem.getBoundingClientRect();
        const x = event.clientX - rect.left;
        const y = event.clientY - rect.top;
        console.log(x + "," + y + ",");
        //window.location.href = "/"
    });

    $("#localization").on("click", function () {
        if (localization == "UA") {
            localization = "EN"
        } else {
            localization = "UA"
        }
        Cookies.remove('loc');
        Cookies.set('loc', localization, { expires: 180});
        window.location.reload()
    });

    let loginBtn = $("#buttonLogin"); // Signing in
    
    // let check = "false";
    // // Handle the checkbox
    // $("#checkLogin").click(function (e) {
    //     if ($(this).is(':checked')) {
    //         check = "true";
    //     } else {
    //         check = "false";
    //     }
    // });

    // Signing in
    loginBtn.click(() => {
        var name = $("#usernameLogin").val();
        var password = $("#passwordLogin").val();
        var check
        
        if ($("#checkLogin").is(':checked')) {
            check = "true";
        } else {
            check = "false";
        }

        $.post("/login", {
            usernameLogin: name,
            passwordLogin: password,
            checkLogin: check,
        }, function(result){
            if (result.err != null) {
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
