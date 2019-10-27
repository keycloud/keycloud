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
            ...
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-warning" data-dismiss="modal">Close</button>
                <button type="button" class="btn btn-success" data-dismiss="modal" aria-label="Close">Use</button>
            </div>
        </div>
    </div>
</div>`);
$("#addPwPopup").modal("show");