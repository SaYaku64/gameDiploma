$("document").ready(() => {
    var mapHeight
    var mapWidth

    $("#mainMap").ready(function () {

        $("#mainMap").on("click", function () {
            mapHeight = this.height;
            mapWidth = this.width;

            var elem = document.getElementById('mainMap');
            const rect = elem.getBoundingClientRect();
            const x = event.clientX - rect.left;
            const y = event.clientY - rect.top;
            console.log("x: " + x + " y: " + y);
            console.log("mapWidth: " + mapWidth + " mapHeight: " + mapHeight);

            $.ajax({
                type: "POST",
                url: "/coord",
                data: {
                    x:         x,
                    y:         y,
                    mapWidth:  mapWidth,
                    mapHeight: mapHeight,
                },
                success: function(result){
                    //alert(result)
                },
                error: function(result){
                    document.getElementById("errorMenu").innerHTML = "<p class=\"bg-danger dropdown-item text-white font-weight-bold\">"+result.responseText+"</p>";
                    setTimeout(() => $( "#errorMenu" ).load(window.location.href + " #errorMenu" ), 2500);
                }
            });

            // $.post("/coord", {
            //     x:         x,
            //     y:         y,
            //     mapWidth:  mapWidth,
            //     mapHeight: mapHeight,
            // }, function(result){
                
            //     // if (result.message == "Successfully deleted") {
            //     //     document.getElementById("errorMenu").innerHTML = "<p class=\"bg-success dropdown-item text-white font-weight-bold\">"+result.message+"</p>";
            //     //     setTimeout(() => $( "#errorMenu" ).load(window.location.href + " #errorMenu" ), 2500);
            //     // } else {
            //     //     document.getElementById("errorMenu").innerHTML = "<p class=\"bg-danger dropdown-item text-white font-weight-bold\">"+result.message+"</p>";
            //     // }
            // });
        });
    });

});
