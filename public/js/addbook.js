let submitButton = document.getElementById('submit');
submitButton.addEventListener("click", submittingForm);
let completeButton = document.getElementById('complete');
completeButton.addEventListener("click", closingForm);

function submittingForm(e) {
    var isbn = document.getElementById('InputISBN');

    axios.post('/addbook', {
        isbn: isbn.value
    })
        .then((response) => {
            var myModal = new bootstrap.Modal(document.querySelector("#myModal"), {});
            console.log(response.data)
            document.getElementById("key").value = response.data["BookResponse"]["Key"];
            document.getElementById("title").value = response.data["BookResponse"]["Title"];
            document.getElementById("author").value = response.data["BookResponse"]["Authors"][0];
            //TODO make it loop over multiple authors
            document.getElementById("pages").value = response.data["BookResponse"]["PageNumber"];
            let genreDOM = document.getElementById("genre");
            for(let i = 0; i < response.data["Genres"].length; i++){
                let option = document.createElement("option")
                option.value = i;
                option.text = response.data["Genres"][i];
                genreDOM.appendChild(option);
            }
            myModal.show();

            console.log(response)
        })
        .catch((error) => {
            console.log(error)
        });
}

function closingForm(e) {
    var review = document.getElementById("review");
    var tags = document.getElementById("tags");
    var title = document.getElementById("title");
    var author = document.getElementById("author");
    var pages = document.getElementById("pages");
    var genre = document.getElementById("genre");
    console.log(review, tags, title, author, pages, genre);
    //separate the tags out at the commas into an array.
    axios.post('/confirmbook', {
        review: review.value,
        tags: tags.value,
        title: title.value,
        author: author.value,
        pages: pages.value,
        genre: genre.value,
        isbn: document.getElementById("isbn").value,
        key: document.getElementById("key").value
    })
}
    //TODO close the modela once the information is sent back with a popup message that says book added.
    //book should then pop up on screen with the option to add book with information, edit information, or manually add book

