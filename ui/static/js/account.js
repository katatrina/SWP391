function toggleForm(formId) {
    var forms = document.getElementsByClassName('form');
    for (var i = 0; i < forms.length; i++) {
        forms[i].style.display = 'none';
    }
    document.getElementById(formId).style.display = 'block';
}