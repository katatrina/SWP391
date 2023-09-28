function showHideForms() {
    var roleSelect = document.getElementById("role");
    var userInfo = document.getElementById("userInfo");
    var providerInfo = document.getElementById("providerInfo");

    if (roleSelect.value === "user") {
        userInfo.style.display = "block";
    } else {
        providerInfo.style.display = "none";
    }

    if (roleSelect.value === "provider") {
        providerInfo.style.display = "block";
    } else {
        userInfo.style.display = "none";
    }

}