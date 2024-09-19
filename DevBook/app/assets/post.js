function likePost(id) {
    fetch(`/posts/${id}/like`, {
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

function unlikePost(id) {
    fetch(`/posts/${id}/like`, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json'
        }
    })
    .then(res => {
        location.reload();
    })
}

function deletePost(id) {
    console.log(`delete post with id: ${id}`);
    fetch(`/posts/${id}`, {
        method: 'DELETE'
    })
    .then(res => {
        location.reload();
    })
}