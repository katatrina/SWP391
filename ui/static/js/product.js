// Lấy danh sách các phần tử <li>
const filterItems = document.querySelectorAll('.item-list-group');

// Lấy danh sách các sản phẩm
const products = document.querySelectorAll('.product');

// Bắt đầu bằng việc hiển thị tất cả sản phẩm
function showAllProducts() {
    products.forEach(product => {
        product.style.display = 'inline-block';
    });
}

// Xử lý sự kiện click trên từng phần tử <li>
filterItems.forEach(filterItem => {
    filterItem.addEventListener('click', () => {
        // Lấy giá trị data-filter của phần tử <li> được click
        const filterValue = filterItem.getAttribute('data-filter');

        // Ẩn tất cả sản phẩm
        products.forEach(product => {
            product.style.display = 'none';
        });

        if (filterValue === 'all') {
            // Hiển thị tất cả sản phẩm nếu chọn "Tất Cả"
            showAllProducts();
        } else {
            // Hiển thị sản phẩm tương ứng với giá trị data-filter
            const filteredProducts = document.querySelectorAll('.' + filterValue);
            filteredProducts.forEach(product => {
                product.style.display = 'inline-block';
            });
        }
    });
});

// Mặc định hiển thị tất cả sản phẩm khi trang web được tải
showAllProducts();
