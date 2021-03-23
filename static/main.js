$("document").ready(() => {
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

    // function sendMapSize(param) {
    //     mapHeight = param.height;
    //     mapWidth = param.width;
    //     $.ajax({
    //         type: "POST",
    //         url: "/mapSize",
    //         data: {
    //             mapWidth:  mapWidth,
    //             mapHeight: mapHeight,
    //         },
    //         success: function(result){
    //             setShapes(result)
    //         },
    //         error: function(){
    //             clearInterval(setInt)
    //         }
    //     });
    // }

    // function setShapes(params) {
    //     console.log(JSON.stringify(params))
    // }

    // var setInt = setInterval(function () {
    //     $("#mainMap").click();
    // }, 2000);

    // $("#mainMap").on("click", function () {
    //     sendMapSize(this)

    //     // var elem = document.getElementById('mainMap');
    //     // const rect = elem.getBoundingClientRect();
    //     // const x = event.clientX - rect.left;
    //     // const y = event.clientY - rect.top;
    //     // console.log("x: " + x + " y: " + y);
    //     // console.log("mapWidth: " + mapWidth + " mapHeight: " + mapHeight);

    //     // $.ajax({
    //     //     type: "POST",
    //     //     url: "/coord",
    //     //     data: {
    //     //         x:         x,
    //     //         y:         y,
    //     //     },
    //     //     success: function(result){
    //     //         //alert(result)
    //     //     },
    //     //     error: function(result){
    //     //         document.getElementById("errorMenu").innerHTML = "<p class=\"bg-danger dropdown-item text-white font-weight-bold\">"+result.responseText+"</p>";
    //     //         setTimeout(() => $( "#errorMenu" ).load(window.location.href + " #errorMenu" ), 2500);
    //     //     }
    //     // });

    //     // $.post("/coord", {
    //     //     x:         x,
    //     //     y:         y,
    //     //     mapWidth:  mapWidth,
    //     //     mapHeight: mapHeight,
    //     // }, function(result){
            
    //     //     // if (result.message == "Successfully deleted") {
    //     //     //     document.getElementById("errorMenu").innerHTML = "<p class=\"bg-success dropdown-item text-white font-weight-bold\">"+result.message+"</p>";
    //     //     //     setTimeout(() => $( "#errorMenu" ).load(window.location.href + " #errorMenu" ), 2500);
    //     //     // } else {
    //     //     //     document.getElementById("errorMenu").innerHTML = "<p class=\"bg-danger dropdown-item text-white font-weight-bold\">"+result.message+"</p>";
    //     //     // }
    //     // });
    // });

});
