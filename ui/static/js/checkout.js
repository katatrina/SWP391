let paymentButtons = document.querySelectorAll('.checkout-payments button');

paymentButtons.forEach(function (button) {
    button.addEventListener('click', function () {
        // Remove 'active' class from all buttons
        paymentButtons.forEach(function (btn) {
            btn.classList.remove('active');
        });

        // Add 'active' class to the clicked button
        button.classList.add('active');

        // Send the selected payment method to the backend
        let selectedPaymentMethod = button.getAttribute('data-payment-method');
        // Perform the necessary action to send the 'selectedPaymentMethod' to the backend
        console.log('Selected Payment Method:', selectedPaymentMethod);
    });
});