// let paymentButtons = document.querySelectorAll('.checkout-payments button');

// paymentButtons.forEach(function (button) {
//     button.addEventListener('click', function () {
//         // Remove 'active' class from all buttons
//         paymentButtons.forEach(function (btn) {
//             btn.classList.remove('active');
//         });

//         // Add 'active' class to the clicked button
//         button.classList.add('active');

//         // Send the selected payment method to the backend
//         let selectedPaymentMethod = button.getAttribute('data-payment-method');
//         // Perform the necessary action to send the 'selectedPaymentMethod' to the backend
//         console.log('Selected Payment Method:', selectedPaymentMethod);
//     });
// });

let paymentButtons = document.querySelectorAll('.checkout-payments button');

// Thêm lớp 'active' cho phương thức thanh toán mặc định
paymentButtons[0].classList.add('active');

paymentButtons.forEach(function (button) {
    button.addEventListener('click', function () {
        // Loại bỏ lớp 'active' từ tất cả các nút khác
        paymentButtons.forEach(function (btn) {
            btn.classList.remove('active');
        });

        // Thêm lớp 'active' vào nút được nhấp
        button.classList.add('active');

        // Gửi phương thức thanh toán được chọn đến backend
        let selectedPaymentMethod = button.getAttribute('data-payment-method');
        // Thực hiện hành động cần thiết để gửi 'selectedPaymentMethod' đến backend
        console.log('Selected Payment Method:', selectedPaymentMethod);
    });
});
