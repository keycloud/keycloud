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

function updateModal(){
    $(".custom-field-row-added").remove();
}

function removeEntry(event) {
    console.log(event);
}