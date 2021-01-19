
$('#menu').click(function(e){
    e.preventDefault();
    $('.sidenav').show('500')
});
$('#closeMenu').click(function(e){
    e.preventDefault();
    $('.sidenav').hide('500')
});