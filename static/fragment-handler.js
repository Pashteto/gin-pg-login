document.addEventListener('DOMContentLoaded', function() {
    let fragment = window.location.hash;

    let links = document.querySelectorAll('a');
    links.forEach(link => {
        link.href += fragment;
    });
    const user = parseUrlData();
    if (document.getElementById('firstName')) {
        document.getElementById('firstName').value = user.first_name;
        document.getElementById('lastName').value = user.last_name;
        document.getElementById('username').value = user.username;
    }
});

document.addEventListener('click', function(e) {
    if (e.target.tagName === 'A') {
        e.preventDefault();
        window.location.href = e.target.href;
    }
});