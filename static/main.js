$("document").ready(() => {

    let btn1 = $("#buttonLogin"); // Signing in
    let btn2 = $("#buttonComment"); // Creates comment
    let btn3 = $("#buttonComDel"); // Deletes comments
    let btn4 = $("#buttonArtDel"); // Deletes article
    let btn5 = $("#buttonArticleCreation"); // Creates article
    let btn6 = $("#buttonArticleDelete"); // Deletes article
    let btn7 = $("#buttonArticleEdit"); // Edit article
    let btn8 = $("#clearBusket"); // clear busket
    
    
    let check = "false";
    // Handle the checkbox
    $("#checkLogin").click(function (e) {
        if ($(this).is(':checked')) {
            check = "true";
        } else {
            check = "false";
        }
    });

    btn8.click(() => {
        $.post("/clear", function(result){
            if (result.message == "Cleared") {
                $("#busketShowBackdrop").modal('hide');
            }
        });
    });

    // Signing in
    btn1.click(() => {
        var name = $("#usernameLogin").val();
        var password = $("#passwordLogin").val();
        
        $.post("/u/login", {
            usernameLogin: name,
            passwordLogin: password,
            checkLogin: check,
        }, function(result){
            if (result.message == "Successful login") {
                document.getElementById("errorMenu").innerHTML = "<p class=\"bg-success dropdown-item text-white font-weight-bold\">"+result.message+"</p>";
                $( "#navbarTogglerDemo01" ).load(window.location.href + " #navbarTogglerDemo01" );
                setTimeout(() => $( "#errorMenu" ).load(window.location.href + " #errorMenu" ), 2500);
            } else {
                document.getElementById("errorMenu").innerHTML = "<p class=\"bg-danger dropdown-item text-white font-weight-bold\">"+result.message+"</p>";
            }
        });
    });

    // Creates comment
    

    btn6.click(() => {
        var title = $("#title").val();
        $('.close').trigger('click');

        $.post("/article/delete", {
            title: title,
        }, function(result){
            if (result.message == "Successfully deleted") {
                document.getElementById("errorMenu").innerHTML = "<p class=\"bg-success dropdown-item text-white font-weight-bold\">"+result.message+"</p>";
                // document.getElementById("commentPlace").innerHTML = result.message;
                $( "#commentPlace" ).load(window.location.href + " #commentPlace" );
                //$( "#errorMenu" ).load(window.location.href + " #errorMenu" );
                setTimeout(() => $( "#errorMenu" ).load(window.location.href + " #errorMenu" ), 2500);
            } else {
                document.getElementById("errorMenu").innerHTML = "<p class=\"bg-danger dropdown-item text-white font-weight-bold\">"+result.message+"</p>";
                //$( "#commentPlace" ).load(window.location.href + " #commentPlace" );
                setTimeout(() => $( "#errorMenu" ).load(window.location.href + " #errorMenu" ), 2500);
            }
        });
    });


    btn2.click(() => {
        var amount = $("#amount").val();
        var title = document.getElementById('articleTitle').innerHTML
        var price = document.getElementById('price').innerHTML
        $('.close').trigger('click');

        $.post("/article/comment", {
            amount: amount,
            title: title,
            price: price,
        }, function(result){
            if (result.message == "Успішно додано") {
                document.getElementById("errorMenu").innerHTML = "<p class=\"bg-success dropdown-item text-white font-weight-bold\">"+result.message+"</p>";
                // document.getElementById("commentPlace").innerHTML = result.message;
                $( "#commentPlace" ).load(window.location.href + " #commentPlace" );
                //$( "#errorMenu" ).load(window.location.href + " #errorMenu" );
                setTimeout(() => $( "#errorMenu" ).load(window.location.href + " #errorMenu" ), 2500);
            } else {
                document.getElementById("errorMenu").innerHTML = "<p class=\"bg-danger dropdown-item text-white font-weight-bold\">"+result.message+"</p>";
                //$( "#commentPlace" ).load(window.location.href + " #commentPlace" );
                setTimeout(() => $( "#errorMenu" ).load(window.location.href + " #errorMenu" ), 2500);
            }
        });
    });

    // Deletes comment
    btn3.click(() => {
        var title = document.getElementById('articleTitle').innerHTML

        $.post("/article/comment/delete", {
            title: title,
        }, function(result){
            if (result.message == "Successfully removed") {
                document.getElementById("errorMenu").innerHTML = "<p class=\"bg-success dropdown-item text-white font-weight-bold\">"+result.message+"</p>";
                // document.getElementById("commentPlace").innerHTML = result.message;
                $( "#commentPlace" ).load(window.location.href + " #commentPlace" );
                setTimeout(() => $( "#errorMenu" ).load(window.location.href + " #errorMenu" ), 2500);
                //setTimeout(() => location.reload(), 1500);
            } else {
                document.getElementById("errorMenu").innerHTML = "<p class=\"bg-danger dropdown-item text-white font-weight-bold\">"+result.message+"</p>";
                setTimeout(() => $( "#errorMenu" ).load(window.location.href + " #errorMenu" ), 2500);
            }
        });

    });

    // Deletes article
    btn4.click(() => {
        var r = confirm("Are you sure?");
        if (r == true) {
            var title = document.getElementById('articleTitle').innerHTML
            var authorName = document.getElementById('authorName').innerHTML

            $.post("/article/delete", {
                title: title,
                authorName: authorName,
            }, function(result){
                if (result.message == "Successfully deleted") {
                    
                    setTimeout(() => window.location.href = '/article', 2500);
                    document.getElementById("errorMenu").innerHTML = "<p class=\"bg-success dropdown-item text-white font-weight-bold\">"+result.message+"</p>";
                    setTimeout(() => $( "#errorMenu" ).load(window.location.href + " #errorMenu" ), 2500);
                } else {
                    document.getElementById("errorMenu").innerHTML = "<p class=\"bg-danger dropdown-item text-white font-weight-bold\">"+result.message+"</p>";
                    setTimeout(() => $( "#errorMenu" ).load(window.location.href + " #errorMenu" ), 2500);
                }
            });
        }
    });

    // Creates article
    btn5.click(() => {
        var title = $("#title").val();
        var content = $("#content").val();
        var picture = $("#picture").val();
        var price = $("#price").val();
        
        $.post("/article/create", {
            title: title,
            content: content,
            picture: picture,
            price: price,
        }, function(result){
            if (result.message == "Successful creation") {
                setTimeout(() => window.location.href = '/article', 2500);
                document.getElementById("errorMenu").innerHTML = "<p class=\"bg-success dropdown-item text-white font-weight-bold\">"+result.message+"</p>";
                setTimeout(() => $( "#errorMenu" ).load(window.location.href + " #errorMenu" ), 2500);
            } else {
                document.getElementById("errorMenu").innerHTML = "<p class=\"bg-danger dropdown-item text-white font-weight-bold\">"+result.message+"</p>";
                setTimeout(() => $( "#errorMenu" ).load(window.location.href + " #errorMenu" ), 2500);
            }
        });

    });

    // Edits article
    btn7.click(() => {
        var titleEdit = $("#titleEdit").val();
        var title = $("#title").val();
        var content = $("#content").val();
        var picture = $("#picture").val();
        var price = $("#price").val();
        
        $.post("/article/edit", {
            titleEdit: titleEdit,
            title: title,
            content: content,
            picture: picture,
            price: price,
        }, function(result){
            if (result.message == "Successful edition") {
                setTimeout(() => window.location.href = '/article', 2500);
                document.getElementById("errorMenu").innerHTML = "<p class=\"bg-success dropdown-item text-white font-weight-bold\">"+result.message+"</p>";
                setTimeout(() => $( "#errorMenu" ).load(window.location.href + " #errorMenu" ), 2500);
            } else {
                document.getElementById("errorMenu").innerHTML = "<p class=\"bg-danger dropdown-item text-white font-weight-bold\">"+result.message+"</p>";
                setTimeout(() => $( "#errorMenu" ).load(window.location.href + " #errorMenu" ), 2500);
            }
        });

    });
    // Попытка сделать удаление одного комментария, но всё упирается в Айди, если будет желание - доделаю
     
    // btn3.click(() => {
    //     //alert("AAAAAAAAA")
    //     var comContent =  $("[id^=\"comContent\"][id^=\"buttonComDel\"]").innerHTML
    //     var comName = $("[id^=\"comName\"]").innerHTML
    //     var comTime = $("[id^=\"comTime\"]").innerHTML

    //     // var comContent = document.getElementById('comContent').innerHTML
    //     // var comName = document.getElementById('comName').innerHTML
    //     // var comTime = document.getElementById('comTime').innerHTML
    //     alert(comContent)
    //     alert(comName)
    //     alert(comTime)

    //     $.post("/article/comment/delete", {
    //         comContent: comContent,
    //         comName: comName,
    //         comTime: comTime,
    //     }, function(result){
    //         if (result.message == "Successfully added") {
    //             document.getElementById("errorMenu").innerHTML = "<p class=\"bg-success dropdown-item text-white font-weight-bold\">"+result.message+"</p>";
    //             // document.getElementById("commentPlace").innerHTML = result.message;
    //             $( "#commentPlace" ).load(window.location.href + " #commentPlace" );
    //             //setTimeout(() => location.reload(), 1500);
    //         } else {
    //             document.getElementById("errorMenu").innerHTML = "<p class=\"bg-danger dropdown-item text-white font-weight-bold\">"+result.message+"</p>";
    //             setTimeout(() => location.reload(), 1500);
    //         }
    //     });
    // });
});
