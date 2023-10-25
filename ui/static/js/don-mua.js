// const navLinks = document.querySelectorAll('.nav-item a');

// navLinks.forEach(link => {
//     link.addEventListener('click', function () {
//         navLinks.forEach(item => item.classList.remove('active'));
//         this.classList.add('active');
//     });
// });
document.addEventListener("DOMContentLoaded", function () {
    const navLinks = document.querySelectorAll('.nav-item a');
    navLinks[0].classList.add('active');

    navLinks.forEach(link => {
        link.addEventListener('click', function () {
            navLinks.forEach(item => item.classList.remove('active'));
            this.classList.add('active');
        });
    });
});
