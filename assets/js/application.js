require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
require("./sb-admin-2.min.js");
$(() => {
    jQuery(document).ready(function(){
        $("div.p-switch input[type=checkbox]").change(function(){
            window.location.replace("/auth/scm/"+$(this).attr("name"));
        });
    })
});
