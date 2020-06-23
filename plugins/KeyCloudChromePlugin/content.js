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
                  </div>
                </form>
            </div>
        </div>
    </div>
</div>`);

let form = findLoginForm();
let inputs = getFormInputs(form);
let usernameInput = findUsernameInput(inputs);
let passwordInput = findPasswordInput(inputs);

// only show popup if fillable login form is found
if (usernameInput && passwordInput) { $("#addPwPopup").modal("show"); }

$( "#fLogin" ).on("submit", function(e) {
    e.preventDefault(); // avoid to execute the actual submit of the form.
    $.ajax({
        type: "POST",
        url: "https://keycloud-dev.zeekay.dev/standard/login",
        contentType: "application/json",
        crossDomain: true,
        crossOrigin: true,
        data: JSON.stringify({ "username": e.target[0].value, "masterpassword": e.target[1].value }),
        success: function(data)
        {
            $.ajax({
                type: "GET",
                url: "https://keycloud-dev.zeekay.dev/password-by-url",
                crossDomain: true,
                crossOrigin: true,
                data: {
                    "url": window.location.href
                },
                success: function(data)
                {
                    data = JSON.parse(data);
                    fillInputs(data, usernameInput, passwordInput);
                    $("#addPwPopup").modal("hide");
                },
                onFailure: {}
            });
        },
        onFailure: {}
    });
});

function findLoginForm() {
    return $('form[id*="login"]')
}

function getFormInputs(form) {
    return form.find(':input');
}

function findUsernameInput(inputs) {
    let usernameInput = null;
    // https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/forEach#Description
    // there is no way to stop or break a .each expect return false
    inputs.each(function() {
        if ($(this)[0].type === "email") {
            usernameInput = $(this)[0];
            return false;
        } else if ($(this)[0].name.match(/username/i)) {
            usernameInput = $(this)[0];
            return false
        }
    });
    return usernameInput;
}

function findPasswordInput(inputs) {
    let passwordInput = null;
    inputs.each(function() {
       if ($(this)[0].type === "password") {
           passwordInput = $(this)[0];
           return false
       }
    });
    return passwordInput;
}

function fillInputs(data, usernameInput, passwordInput) {
    if (usernameInput) { usernameInput.value = data[0]["username"] }
    if (passwordInput) { passwordInput.value = data[0]["password"] }
}
