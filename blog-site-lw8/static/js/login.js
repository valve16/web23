function showPassword() {
  let img = document.querySelector(".form__input_icon");
  img.src = "/static/img/eye.png";
  img.setAttribute("onclick", "hidePassword()");
  let input = document.getElementById("password");
  input.type = "text";
}

function hidePassword() {
  let img = document.querySelector(".form__input_icon");
  img.src = "/static/img/eye-off.png";
  img.setAttribute("onclick", "showPassword()");
  let input = document.getElementById("password");
  input.type = "password";
}

document.querySelectorAll("input").forEach(function (input) {
  input.addEventListener("input", function () {
    if (input.value) {
      input.classList.add("form__input_filled");
    } else {
      input.classList.remove("form__input_filled");
    }
  });
});
