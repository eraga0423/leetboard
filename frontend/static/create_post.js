document.addEventListener("DOMContentLoaded", function() {
    const form = document.getElementById('createPost');
    form.addEventListener('submit', function(event){
        event.preventDefault();
        const formData = new FormData(form);
        fetch('/new_post', {
            method: 'POST',
            body: formData
        })
        .then(response =>response.json())
        .then(data =>{
            alert('форма отправлена')
        })
        .catch((error) => {
            console.error('ошибка', error)
        }); 
    });
});
