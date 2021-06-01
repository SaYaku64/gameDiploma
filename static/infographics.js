$("document").ready(() => {

    
    $("#clearInfo").click(() => {
        $.get("/clearInfo", function(result) {
            successAlert(result.message);
          });
    });

    $(".radios").on("change", function () {
        $("#changeCheck").prop( "disabled", false );
        $("#placedCheck").prop( "disabled", false );
        $("#check1").prop( "disabled", false );
        $("#check2").prop( "disabled", false );
        $("#check3").prop( "disabled", false );
        $("#check4").prop( "disabled", false );
        $("#check5").prop( "disabled", false );
        $("#check6").prop( "disabled", false );

        if ( $("#pieRadio").is(':checked') ) {
            $("#changeCheck").prop( "disabled", true );
            $("#placedCheck").prop( "disabled", true );
            $("#check3").prop( "disabled", true );
            $("#check4").prop( "disabled", true );
            $("#check5").prop( "disabled", true );
            $("#check6").prop( "disabled", true );
        }
        
        // if ( $("#radarRadio").is(':checked') ) {
        //     $("#changeCheck").prop( "disabled", true );
        //     $("#placedCheck").prop( "disabled", true );
        //     $("#check34").prop( "disabled", true );
        // }
        
        // if ( $("#wordRadio").is(':checked') ) {
        //     $("#check3").prop( "disabled", true );
        //     $("#check4").prop( "disabled", true );
        //     $("#check5").prop( "disabled", true );
        //     $("#check6").prop( "disabled", true );
        // }
    })

    let infoFilters = $("#infoFilters");

    infoFilters.click(() => {
        var radio = "barRadio"

        var changeCheck = false
        var placedCheck = false
        var check1 = false
        var check2 = false
        var check3 = false
        var check4 = false
        var check5 = false
        var check6 = false
        var check7 = false

        if ($("#barRadio").is(':checked')) {radio = "barRadio";}
        if ($("#pieRadio").is(':checked')) {radio = "pieRadio";}
        // if ($("#radarRadio").is(':checked')) {radio = "radarRadio";}
        // if ($("#wordRadio").is(':checked')) {radio = "wordRadio";}

        if ($("#changeCheck").is(':checked')) {changeCheck = true;}
        if ($("#placedCheck").is(':checked')) {placedCheck = true;}
        if ($("#check1").is(':checked')) {check1 = true;}
        if ($("#check2").is(':checked')) {check2 = true;}
        if ($("#check3").is(':checked')) {check3 = true;}
        if ($("#check4").is(':checked')) {check4 = true;}
        if ($("#check5").is(':checked')) {check5 = true;}
        if ($("#check6").is(':checked')) {check6 = true;}
        if ($("#check7").is(':checked')) {check7 = true;}

        // var checks = [changeCheck, placedCheck, check1, check2, check34, check5, check6]

        // if ( $("#pieRadio").is(':checked') ) {
        //     checks = [check1, check2]
        // }
        
        // if ( $("#radarRadio").is(':checked') ) {
        //     checks = [check1, check2, check5, check6]
        // }

        // if ( $("#wordRadio").is(':checked') ) {
        //     checks = [changeCheck, placedCheck, check1, check2]
        // }

        //$.get("/graph", { radio: radio, "checks[]": checks } );
        
        // $('<a id="fred99" />').attr('href', '/bar').text('LINK').appendTo('body').get(0).click();
        // $('#fred99').remove();

        $.post("/infographics", {
            radio:       radio,
            changeCheck: changeCheck,
            placedCheck: placedCheck,
            check1:      check1,
            check2:      check2,
            check3:     check3,
            check4:     check4,
            check5:      check5,
            check6:      check6,
            check7:      check7,
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
