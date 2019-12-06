var exampleEntries = [
    {
        "i" : 0,
        "Username" : "Mark",
        "Url" : "https://www.google.com",
        "Password" : "example",
    },{
        "i" : 1,
        "Username" : "Jacob",
        "Url" : "https://www.google.com",
        "Password" : "example",
    },{
        "i" : 2,
        "Username" : "Larry",
        "Url" : "https://www.google.com",
        "Password" : "example",
    }
];

function clickTab(elem){
    for(let i = 0; i < $("a.nav-link").length; i++){
        $("a.nav-link")[i].classList.remove("active");
    }
    elem.classList.add("active");
}

function updateUserData(){
    fetch('/user', {
        method: 'GET',
        credentials: 'same-origin',
        headers: {
            'Accept': 'application/json'
        }
    }).then(async function (response) {
        if(response.status !== 200){
            throw new Error(response.statusText);
        }else {
            let resp = await response.json();
            $("#master-pw").val(resp.MasterPassword);
        }
    })
}

function addCustomField(){
    $(`<div class="form-row custom-field-row-added" style="margin-bottom: 15px">
                                    <div class="col">
                                        <input type="text" class="form-control" placeholder="Custom Field Name">
                                    </div>
                                    <div class="col">
                                        <input type="text" class="form-control" placeholder="Custom Field content">
                                    </div>
                                    <div class="col">
                                        <input class="form-check-input big-checkbox" type="checkbox">
                                        <label class="form-check-label" style="font-size: x-large;margin-left: 15px;">
                                            Encrypt
                                        </label>
                                    </div>
                                </div>`).insertBefore("#btn-add-field-group");
}

function addTableRow(value) {
    $("#pwTable").prepend(`<tr class="entry">
    <th scope="row">${value.i}</th>
    <td>${value.Username}</td>
    <td><a target="_blank" rel="noopener noreferrer" href="${value.Url}">${value.Url}</a></td>
    <td><button type="button" class="btn btn-info" id="cp${value.i}" onclick="copyToClipboard(this.id)"><i class="fa fa-clipboard" ></i> Copy to Clipboard</button></td>
    <td><button type="button" class="btn btn-danger" id="rm${value.i}" onclick="removeEntry(this.id)"><i class="fa fa-remove"></i></button></td>
    </tr>`);
}

function updateModal(){
    $(".custom-field-row-added").remove();
}

function removeEntry(id) {
    exampleEntries.splice(id.slice(2,), 1);
    renderTable(); // works so far, needs some work done on the indices
    $('.toast').toast('show');
}

function generatePassword() {
    document.getElementById("pwInput").value = genPW();
}

function copyToClipboard(id) {
    let pw = exampleEntries[id.slice(2,)].Password;
    window.prompt("Copy to clipboard: Ctrl+C", pw); // Workaround for now, could use other, prettier techniques
}

function saveNewEntry() {
    let newEntry = {
        "i" : exampleEntries.length,
        "Username" : "",
        "Url" : "",
        "Password" : "",
    };
    newEntry["Username"] = document.getElementById("usernameInput").value;
    newEntry["Url"] = document.getElementById("urlInput").value;
    newEntry["Password"] = document.getElementById("pwInput").value;
    exampleEntries.push(newEntry);
    renderTable();
}

function renderTable() {
    $(".entry").remove(); // clear table
    exampleEntries.reverse();  // bc of callback
    exampleEntries.forEach(addTableRow);
    exampleEntries.reverse();
}

$("document").ready(function() {
    renderTable();
    $(function () {
        $('[data-toggle="tooltip"]').tooltip()
    });
    $('.toast').toast()
});

class WebAuthn {
    // Decode a base64 string into a Uint8Array.
    static _decodeBuffer(value) {
        return Uint8Array.from(atob(value), c => c.charCodeAt(0));
    }

    // Encode an ArrayBuffer into a base64 string.
    static _encodeBuffer(value) {
        return btoa(new Uint8Array(value).reduce((s, byte) => s + String.fromCharCode(byte), ''));
    }

    // Checks whether the status returned matches the status given.
    static _checkStatus(status) {
        return res => {
            if (res.status === status) {
                return res;
            }
            throw new Error(res.statusText);
        };
    }

    add2FA() {
        return fetch('/webauthn/registration/start', {
            method: 'POST',
            body: JSON.stringify({'username': webAuthnConfig.username}),
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            }
        })
            .then(WebAuthn._checkStatus(200))
            .then(res => res.json())
            .then(res => {
                res.publicKey.challenge = WebAuthn._decodeBuffer(res.publicKey.challenge);
                res.publicKey.user.id = WebAuthn._decodeBuffer(res.publicKey.user.id);
                res.publicKey.authenticatorSelection.userVerification = "required";
                if (res.publicKey.excludeCredentials) {
                    for (var i = 0; i < res.publicKey.excludeCredentials.length; i++) {
                        res.publicKey.excludeCredentials[i].id = WebAuthn._decodeBuffer(res.publicKey.excludeCredentials[i].id);
                    }
                }
                return res;
            })
            .then(res => navigator.credentials.create(res))
            .then(credential => {
                return fetch('/webauthn/registration/finish', {
                    method: 'POST',
                    headers: {
                        'Accept': 'application/json',
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        id: credential.id,
                        rawId: WebAuthn._encodeBuffer(credential.rawId),
                        response: {
                            attestationObject: WebAuthn._encodeBuffer(credential.response.attestationObject),
                            clientDataJSON: WebAuthn._encodeBuffer(credential.response.clientDataJSON)
                        },
                        type: credential.type,
                        username: webAuthnConfig.username
                    }),
                })
                .then(function (res) {
                    console.log(res);
                    // TODO: Show any kind of success/fail - message
                })
            })
    }
}
