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
    <td><button type="button" class="btn btn-info" id="${value.i}" onclick="copyToClipboard(this.id)"><i class="fa fa-clipboard" ></i> Copy to Clipboard</button></td>
    <td><button type="button" class="btn btn-danger" id="${value.i}" onclick="removeEntry(this.id)"><i class="fa fa-remove"></i></button></td>
    </tr>`);
}

function updateModal(){
    $(".custom-field-row-added").remove();
}

function removeEntry(id) {
    exampleEntries.splice(id, 1);
    renderTable(); // works so far, needs some work done on the indices
}

function generatePassword() {
    document.getElementById("pwInput").value = genPW();
}

function copyToClipboard(id) {
    let pw = exampleEntries[id].Password;
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
    })
});
