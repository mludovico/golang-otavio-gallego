let form = document.querySelector('form');

form.addEventListener('submit', function (e) {
    e.preventDefault();
    let formData = new FormData(form);
    switch (form.id) {
        case 'login':
            login(formData);
            break;
        case 'signup':
            register(formData);
            break;
    }
});

function login(formData) {
    let username = formData.get('username');
    let password = formData.get('password');
    let post = { email: username, password: password };
    fetch('/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(post),
    })
        .then((response) => {
            console.log('Response: ', response);
            if (response.status === 200) {
                window.location.href = '/home';
            } else {
                alert('Login failed');
            }
        })
        .catch((error) => {
            console.error('Error:', error);
        });
}

function register(formData) {
    let username = formData.get('name');
    let email = formData.get('email');
    let nick = formData.get('nick');
    let password = formData.get('password');
    let password_confirmation = formData.get('password_confirmation');
    if (password !== password_confirmation) {
        alert('Passwords do not match');
        return;
    }
    let post = { 
        name: username,
        email: email, 
        nick: nick, 
        password: password 
    };
    fetch('/signup', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(post),
        mode: 'no-cors',
    })
        .then((response) => response.json())
        .then((data) => {
            if (data.status === 'success') {
                window.location.href = '/';
            } else {
                alert('Registration failed');
            }
        })
        .catch((error) => {
            console.error('Error:', error);
        });
}
