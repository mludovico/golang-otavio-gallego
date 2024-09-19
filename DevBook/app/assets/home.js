document.getElementById('current-year').innerText = new Date().getFullYear();
document.getElementById('hamburger-menu').addEventListener('click', () => {
    document.querySelector('.menu').classList.toggle('active');
});
