// JavaScript code to add 'active' class to the current link
// var currentLocation = window.location.href;
// var menuItems = document.querySelectorAll('.nav-main a');

// menuItems.forEach(function (item) {
//     if (item.href === currentLocation) {
//         item.classList.add('active');
//     }
// });

// JavaScript code to add 'active' class to the current link
var currentPath = window.location.pathname;
var menuItems = document.querySelectorAll('a.nav-link.mx-2');

menuItems.forEach(function (item) {
    if (item.getAttribute('href') === currentPath) {
        item.classList.add('active');
    }
});
