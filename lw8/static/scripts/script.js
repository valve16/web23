const fileInput = document.getElementById('file-input');
const uploadLabel = document.getElementById('upload-label');
const uploadImage = document.getElementById('upload-image');



fileInput.addEventListener('change', function() {
  const file = this.files[0];
  if (file) {
    const reader = new FileReader();
    reader.addEventListener('load', function() {
      uploadImage.setAttribute('src', this.result);
      uploadLabel.setAttribute('for', '');
    });
    reader.readAsDataURL(file);
  }
});