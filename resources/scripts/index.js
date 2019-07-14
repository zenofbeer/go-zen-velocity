var resources = resources || {};
resources.scripts = resources.scripts || {};
resources.scripts.index = function() {
    var init = function() {
        handleAdminClicks('');
        registerButtonClicks();
    },
    handleAdminClicks = function(adminTool) {
        var adminTools = [
            'selectWorkstreamDiv',
            'createWorkstreamDiv',
            'editWorkstreamDiv'
        ];
        for (var i = 0; i < adminTools.length; i++) {
            if(adminTools[i] == adminTool) {
                $('#' + adminTools[i]).show();
            } else {
                $('#' + adminTools[i]).hide();
            }
        }
    },
    registerButtonClicks = function() {
        $('#selectWorkstreamButton').click(function() {
            getWorkstreamSelectionList();
            handleAdminClicks('selectWorkstreamDiv');
        });
        $('#createWorkstreamButton').click(function() {
            handleAdminClicks('createWorkstreamDiv');
        });
        $('#editWorkstreamButton').click(function() {
            handleAdminClicks('editWorkstreamDiv');
        });
        $('#navigateToWorkstreamHome').click(function() {
            navigateToWorkstreamHome();
        });
    },
    navigateToWorkstreamHome = function() {
        var displayNameVal = $('#WorkstreamAdminSelect option:selected').text();
        var idVal = $('#WorkstreamAdminSelect option:selected').attr('id');
        if (idVal == -1) {
            alert("Please select a workstream from the list");
        } else {
            $.get('velocity/workstreamHome',{
                displayName : displayNameVal,
                id : idVal
            });
            window.location.replace("velocity/workstreamHome");
        }
    },
    getWorkstreamSelectionList = function() {
        $.ajax({
            url:'velocity/workstreamNames',
            type:'POST',
            dataType:'JSON',
            data:{ajaxpostdata:'hello'},
            success:function(response){
                // console.log(JSON.stringify(response))
                $('#WorkstreamAdminSelect').empty();
                $.each(response.WorkstreamNames, function(i, workstreamName){
                    $('#WorkstreamAdminSelect').append($('<option></option').attr('id', workstreamName.ID).text(workstreamName.Name));
                });
            }
        });
    };
    return {
        Init : init
    }
}();

$('document').ready(function() {
    var $page = $('#index');
    if ($page && $page.length > 0) {
        var view = resources.scripts.index;
        view.Init();
    }
});