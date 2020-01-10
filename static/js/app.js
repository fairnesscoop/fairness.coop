
var app = {
    init: function() {
       console.log('init');

       var buttonMenu = document.querySelector('#buttonMenu');

       buttonMenu.addEventListener('click', app.handleClick);
    },

    handleClick: function(){
        console.log('click');
    }
};

document.addEventListener('DOMContentLoaded', app.init);
