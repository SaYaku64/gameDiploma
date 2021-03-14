$("document").ready(() => {

    $(".tilePicture").on("click", function(){
        var dataId = $(this).attr("data-id");
        alert("The data-id of clicked item is: " + dataId);
    });

});
