
function toastInfo(msg) {
    bootoast({
        message: msg,
        type: 'info',
        position: 'bottom-center',
        icon: undefined,
        timeout: 2,
        animationDuration: 300,
        dismissable: true
    });
}

function toastWarning(msg) {
    bootoast({
        message: msg,
        type: 'warning',
        position: 'bottom-center',
        icon: undefined,
        timeout: 2,
        animationDuration: 300,
        dismissable: true
    });
}

function toastSuccess(msg) {
    bootoast({
        message: msg,
        type: 'success',
        position: 'bottom-center',
        icon: undefined,
        timeout: 2,
        animationDuration: 300,
        dismissable: true
    });
}

function toastDanger(msg) {
    bootoast({
        message: msg,
        type: 'danger',
        position: 'bottom-center',
        icon: undefined,
        timeout: 2,
        animationDuration: 300,
        dismissable: true
    });
}

function post2Server(url, data, success_func) {
    request(url, 'post', data, 'json', success_func)
}

function request(url, type, data, dataType, success_func) {
    if (!success_func || success_func === 'undefined') {
        success_func = function (data) {
            if (data.code === "success") {
                successAlert(data.message, function () {
                    window.location.reload();
                });
            } else {
                errorAlert(data.message);
            }
        }
    }
    $.ajax({
        url       : url,
        type      : type,
        dataType  : dataType,
        data      : data,
        beforeSend: function () {
            $('.loading-modal').modal('show');
        },
        success   : function (data) {
            $('.loading-modal').modal('hide');
            success_func(data);
        }
    }).fail(function (XMLHttpRequest, textStatus, errorThrown) {
        $('.loading-modal').modal('hide');
        errorAlert('服务器出错' + errorThrown);
    });
}

function postFile2Server(url, data, success_func) {
    if (!success_func || success_func === 'undefined') {
        success_func = function (data) {
            if (data.code === 0) {
                successAlert(data.message, function () {
                    window.location.reload();
                });
            } else {
                errorAlert(data.message);
            }
        }
    }
    $.ajax({
        url        : url,
        type       : 'post',
        data       : data,
        cache      : false,
        processData: false,
        contentType: false,
        dataType   : 'json',
        beforeSend : function () {
            $('.loading-modal').modal('show');
        },
        success    : function (data) {
            $('.loading-modal').modal('hide');
            success_func(data);
        }
    }).fail(function (XMLHttpRequest, textStatus, errorThrown) {
        $('.loading-modal').modal('hide');
        errorAlert('服务器出错' + errorThrown);
    });
}