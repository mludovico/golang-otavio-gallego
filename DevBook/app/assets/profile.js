document.querySelector('#edit-form').addEventListener('submit', function(e) {
    e.preventDefault();
    const formData = new FormData(e.target);
    const data = Object.fromEntries(formData);
    const resource = e.target.getAttribute('data-resource');
    console.log(data, resource);
    let endpoint = "";
    if (resource === 'profile') {
        endpoint = '/profile/update';
    } else if (resource === 'password') {
        if (data.new !== data.confirmation) {
            return alert('New password and confirmation do not match');
        }
        endpoint = '/profile/password';
    }
    fetch(endpoint, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json'
        },
        body: JSON.stringify(data),
    })
    .then(res => {
        if (!res.ok) {
            throw new Error('Error:', res.status);
        }
        window.location.replace('/profile');
    })
    .catch(error => console.error('Error:', error));
});

function deleteAccount() {
    if (confirm('Are you sure you want to delete your account?')) {
        fetch('/profile/delete', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json'
            },
        })
        .then(res => {
            if (!res.ok) {
                throw new Error('Error:', res.status);
            }
            window.location.replace('/');
        })
        .catch(error => console.error('Error:', error));
    }
}