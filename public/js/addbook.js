let submitButton = document.getElementById('submit');
submitButton.addEventListener("click", submittingForm);

function submittingForm(e) {
    var isbn = document.getElementById('InputISBN')

    axios.post('/addbook', {
        isbn: isbn.value
    })
        .then((response) => {
            console.log(response)
        })
        .catch( (error) => {
            console.log(error)
        });

    //post the form data as barcode to the backend, which will run return a book with information
    //book should then pop up on screen with the option to add book with information, edit information, or manually add book

}