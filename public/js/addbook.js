let submitButton = document.getElementById('submit');
submitButton.addEventListener("click", submittingForm);
let completeButton = document.getElementById('complete');
completeButton.addEventListener("click", closingForm);
var myModal = new bootstrap.Modal(document.querySelector("#myModal"), {});
var isbn = document.getElementById('InputISBN');



function submittingForm(e) {
    e.preventDefault();
    axios.post('/addbook', {
        isbn: isbn.value
    })
        .then((response) => {
            console.log(response.data);
            document.getElementById("bookKey").value = response.data["BookResponse"]["Key"];
            document.getElementById("title").value = response.data["BookResponse"]["Title"];
           let authors = response.data["BookResponse"]["Authors"];
           let authorName = "";
           let authorKey = "";
            for (let x = 0; x < authors.length; x++) {
               if (authorName !== "") {
                  authorName = authorName.concat(", ", authors[x]["Name"]);
               } else {
                   authorName = authors[x]["Name"];
               }
               if (authorKey !== "") {
                   authorKey.concat(",", authors[x]["Key"]);
               } else {
                   authorKey = authors[x]["Key"];
               }
           }
            document.getElementById("author").value = authorName;
            document.getElementById("authorsKeys").value = authorKey;
            document.getElementById("pages").value = response.data["BookResponse"]["PageNumber"];
            let genreDOM = document.getElementById("genre");
            for(let i = 0; i < response.data["Genres"].length; i++){
                let option = document.createElement("option");
                option.value = i;
                option.text = response.data["Genres"][i];
                genreDOM.appendChild(option);
            }
            myModal.show();
        })
        .catch((error) => {
            console.log(error);
        });
}

function closingForm(e) {
    e.preventDefault();
    var review = document.getElementById("review");
    var tags = document.getElementById("tags");
    var title = document.getElementById("title");
    var author = document.getElementById("authorsKeys");
    var pages = document.getElementById("pages");
    var genre = document.getElementById("genre");
    //separate the tags out at the commas into an array.
    axios.post('/confirmbook', {
        review: review.value,
        tags: tags.value,
        title: title.value,
        authors: author.value,
        pages: pages.value,
        genre: genre.value,
        isbn: isbn.value,
        key: document.getElementById("bookKey").value
    })
        .then((response) => {
            myModal.hide();
            isbn.value = "";
        })
}


