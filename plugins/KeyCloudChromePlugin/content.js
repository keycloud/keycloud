$("body").append(`<div class="modal" id="addPwPopup" tabindex="-1" role="dialog" aria-labelledby="exampleModalLabel" aria-hidden="true">
    <div class="modal-dialog modal-dialog-centered" role="document">
        <div class="modal-content" style="background-color: #333333">
            <div class="modal-header">
                <h5 class="modal-title" style="color: white">Login with KeyCloud</h5>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                    <span aria-hidden="true" style="color: red">&times;</span>
                </button>
            </div>
            <div class="modal-body" id="customFieldList">
                <form id="fLogin">
                  <div class="form-group">
                    <label style="color: white" for="fUsername">Username</label>
                    <input type="text" class="form-control" id="fUsername" placeholder="Enter username">
                  </div>
                  <div class="form-group">
                    <label style="color: white" for="fMasterpassword">Masterpassword</label>
                    <input type="password" class="form-control" id="fMasterpassword" placeholder="Masterpassword">
                  </div>
                  <div class="modal-footer">
                      <button type="button" class="btn btn-warning" data-dismiss="modal">Close</button>
                      <button type="submit" class="btn btn-success" aria-label="Close">Use</button>
                      <button type="button" class="btn btn-info" id="bPasswords">GET /passwords</button>
                  </div>
                </form>
            </div>
        </div>
    </div>
</div>`);
$("#addPwPopup").modal("show");

$( "#fLogin" ).on("submit", function(e) {
    e.preventDefault(); // avoid to execute the actual submit of the form.
    $.ajax({
        type: "POST",
        url: "https://keycloud-dev.zeekay.dev/standard/login",
        contentType: "application/json",
        data: JSON.stringify({ "username": e.target[0].value, "masterpassword": e.target[1].value }),
        success: function(data)
        {
            console.log(window.location.href)
            $.ajax({
                type: "GET",
                url: "https://keycloud-dev.zeekay.dev/password-by-url",
                contentType: "application/json",
                processData: false,
                data: JSON.stringify({"url": "zeekay.dev"}),
                success: function(data)
                {
                    alert(data);
                },
                onFailure: {}
            });
        },
        onFailure: {}
    });
});
