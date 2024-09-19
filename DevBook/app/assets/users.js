function followUser(user_id) {
    fetch(`/users/${user_id}/follow`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json'
        }
    })
    .then(res => {
        location.reload();
    })
}

function unfollowUser(user_id) {
    fetch(`/users/${user_id}/unfollow`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json'
        }
    })
    .then(res => {
        location.reload();
    })
}
