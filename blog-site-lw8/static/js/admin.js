document.querySelectorAll("input").forEach(function (input) {
  if (input.type === "text") {
    input.addEventListener("input", function () {
      if (input.value) {
        input.classList.add("form__input_filled");
      } else {
        input.classList.remove("form__input_filled");
      }
    });
  }
  input.addEventListener("change", () => {
    fillCardPreview();
  });
});

function previewAuthorPhoto() {
  const preview = document.getElementById("form__author-photo");
  const cardPreview = document.getElementById("card__author-avatar");
  const file = document.getElementById("author-photo").files[0];
  const reader = new FileReader();

  reader.addEventListener(
    "load",
    () => {
      preview.src = reader.result;
      cardPreview.src = reader.result;
    },
    false
  );

  if (file) {
    reader.readAsDataURL(file);
  }

  document.getElementById("form__author-photo-upload").classList.add("hidden");
  document.getElementById("form__author-photo-upload-new").classList.remove("hidden");
  document.getElementById("form__author-photo-remove").classList.remove("hidden");
}

function removeAuthorPhoto(event) {
  event.preventDefault();
  const preview = document.getElementById("form__author-photo");
  const cardPreview = document.getElementById("card__author-avatar");
  preview.src = "/static/img/author-photo.png";
  cardPreview.src = "/static/img/card-preview-author.png";

  document.getElementById("form__author-photo-upload").classList.remove("hidden");
  document.getElementById("form__author-photo-upload-new").classList.add("hidden");
  document.getElementById("form__author-photo-remove").classList.add("hidden");
}

function previewArticleImage() {
  const preview = document.getElementById("form__article-image");
  const articlePreview = document.getElementById("article__image");
  const file = document.getElementById("article-image").files[0];
  const reader = new FileReader();

  reader.addEventListener(
    "load",
    () => {
      preview.src = reader.result;
      articlePreview.src = reader.result;
    },
    false
  );

  if (file) {
    reader.readAsDataURL(file);
  }

  document.getElementById("form__file-description_article").classList.add("hidden");
  document.getElementById("form__article-image-upload-new").classList.remove("hidden");
  document.getElementById("form__article-image-remove").classList.remove("hidden");
}

function removeArticleImage(event) {
  event.preventDefault();
  const preview = document.getElementById("form__article-image");
  const articlePreview = document.getElementById("article__image");
  preview.src = "/static/img/upload-article.png";
  articlePreview.src = "/static/img/article-preview-image.png";

  document.getElementById("form__file-description_article").classList.remove("hidden");
  document.getElementById("form__article-image-upload-new").classList.add("hidden");
  document.getElementById("form__article-image-remove").classList.add("hidden");
}

function previewCardImage() {
  const preview = document.getElementById("form__card-image");
  const cardPreview = document.getElementById("card__image");
  const file = document.getElementById("card-image").files[0];
  const reader = new FileReader();

  reader.addEventListener(
    "load",
    () => {
      preview.src = reader.result;
      cardPreview.src = reader.result;
    },
    false
  );

  if (file) {
    reader.readAsDataURL(file);
  }

  document.getElementById("form__file-description_card").classList.add("hidden");
  document.getElementById("form__card-image-upload-new").classList.remove("hidden");
  document.getElementById("form__card-image-remove").classList.remove("hidden");
}

function removeCardImage(event) {
  event.preventDefault();
  const preview = document.getElementById("form__card-image");
  const cardPreview = document.getElementById("card__image");
  preview.src = "/static/img/upload-card.png";
  cardPreview.src = "/static/img/card-preview-image.png";

  document.getElementById("form__file-description_card").classList.remove("hidden");
  document.getElementById("form__card-image-upload-new").classList.add("hidden");
  document.getElementById("form__card-image-remove").classList.add("hidden");
}

function fillCardPreview() {
  document.getElementById("card__title").textContent = document.getElementById("title").value;
  document.getElementById("card__subtitle").textContent = document.getElementById("description").value;
  document.getElementById("card__author-name").textContent = document.getElementById("author-name").value;
  document.getElementById("card__date").textContent = document.getElementById("date").value;
  document.getElementById("article__title").textContent = document.getElementById("title").value;
  document.getElementById("article__description").textContent = document.getElementById("description").value;
}

function publish(event) {
  event.preventDefault();
  const form = document.getElementsByClassName("form")[0];
  const formData = new FormData(form);

  const data = {};
  for (let [key, value] of formData.entries()) {
    data[key] = value;
  }
  console.log(data);
}
