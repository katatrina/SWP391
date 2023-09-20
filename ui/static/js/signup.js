function showHideForms() {
    var roleSelect = document.getElementById("role");
    var businessInfo = document.getElementById("businessInfo");

    if (roleSelect.value === "provider") {
        businessInfo.style.display = "block";
    } else {
        businessInfo.style.display = "none";
    }
}