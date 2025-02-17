let title = '';
let description = '';
let author = '';
let date = '';
let avatar = '';
let avatar_name = '';
let hero = '';
let hero_name = '';
let content = '';

const form = document.getElementById('published');
form.addEventListener('click', publish);

async function publish(event) {
    event.preventDefault();
    let form = document.querySelector('.page__form');
    let contentTextArea = form.querySelector('textarea[name="content"]')

    if (contentTextArea !== null) {
        content = contentTextArea.value;
    }    

    if (!validateForm()) {

        if ((hero && avatar) != '') {
            hero = hero.split(',')[1];
            avatar = avatar.split(',')[1];
        }

        let post = {
            title,
            description,
            author,
            date,
            avatar,
            avatar_name,
            hero,
            hero_name,
            content,
        }
        console.log(post);

        let XHR = new XMLHttpRequest();
        XHR.open('POST', '/api/post');
        XHR.send(JSON.stringify(post));
    }
}

function validateForm() {
    let errorTitle = false;
    let errorDiscp = false;
    let errorAuthor = false;
    let errorDate = false;
    let errorContent = false;

    const pErrors = document.querySelectorAll('.error__p');
    if (pErrors.length > 0) {
        for (let pE of pErrors) {
            pE.remove();
        }
    }

    const inputs = document.querySelectorAll('.input__text');
    if (inputs.length > 0) {
        for (let inp of inputs) {
            inp.style.borderColor = '#EAEAEA';
            inp.style.background = '#FFFFFF';
        }
    }

    const textA = document.querySelector('.content__text_value');
    if (textA !== null) {
        textA.style.borderColor = '#EAEAEA';
        textA.style.background = '#FFFFFF';
    }

    const message = document.getElementById('message');
    while (message.firstChild) {
        message.firstChild.remove();
    }
    message.classList.remove('error__block');
    message.classList.remove('ok__block');

    const pMessage = document.createElement('p');
    const img = document.createElement('img');
    if (title === '') {
        const pTitle = document.createElement('p');
        const labelTitel = document.getElementById('label-title');
        const inputTitle = document.getElementById('input-title');
        inputTitle.style.borderColor = '#E86961';
        inputTitle.style.background = '#FFFFFF';
        pTitle.innerText = 'Title is required.';
        pTitle.classList.add('error__p');
        labelTitel.appendChild(pTitle);
        errorTitle = true;
    }
    if (description === '') {
        const pDiscp = document.createElement('p');
        const labelDiscp = document.getElementById('label-discp');
        const inputDiscp = document.getElementById('input-discription');
        inputDiscp.style.borderColor = '#E86961';
        inputDiscp.style.background = '#FFFFFF';
        pDiscp.innerText = 'Discription is required.';
        pDiscp.classList.add('error__p');
        labelDiscp.appendChild(pDiscp);
        errorDiscp = true;
    }
    if (author === '') {
        const pName = document.createElement('p');
        const labelName = document.getElementById('label-name');
        const inputName = document.getElementById('input-name');
        inputName.style.borderColor = '#E86961';
        inputName.style.background = '#FFFFFF';
        pName.innerText = 'Author name is required.';
        pName.classList.add('error__p');
        labelName.appendChild(pName);
        errorAuthor = true;
    }
    if (date === '') {
        const pDate = document.createElement('p');
        const labelDate = document.getElementById('label-date');
        const inputDate = document.getElementById('input-date');
        inputDate.style.borderColor = '#E86961';
        inputDate.style.background = '#FFFFFF';
        pDate.innerText = 'Date is required.';
        pDate.classList.add('error__p');
        labelDate.appendChild(pDate);
        errorDate = true;
    }
    if (content === '') {
        const pContent = document.createElement('p');
        const labelContent = document.getElementById('label-content');
        const inputContent= document.getElementById('textarea-content');
        inputContent.style.borderColor = '#E86961';
        inputContent.style.background = '#FFFFFF';
        pContent.innerText = 'Content is required.';
        pContent.classList.add('error__p');
        labelContent.appendChild(pContent);
        errorContent = true;
    }

    let error = false;
    if (errorTitle || errorDiscp || errorAuthor || errorDate || errorContent) {
        img.src = '/static/image/alert-circle.svg';
        pMessage.innerText = 'Whoops! Some fields need your attention :o';
        pMessage.classList.add('error__message');
        message.appendChild(img);
        message.appendChild(pMessage);
        message.classList.add('error__block');
        error = true;
    } else {
        img.src = '/static/image/check-circle.svg';
        pMessage.innerText = 'Publish Complete!';
        pMessage.classList.add('error__message');
        message.appendChild(img);
        message.appendChild(pMessage);
        message.classList.add('ok__block');
    }
    return error;
}

async function loadImage(event) {
    let input = event.target;
    let reader = new FileReader();
    reader.onload = () => {
        let dataURL = reader.result;
        let image = document.createElement('img');
        image.src = dataURL;
        hero = dataURL;

        let elementsBox = [];
        let labelsImg = document.querySelectorAll('.form__hero-img');
        boxShowImg = document.querySelector('.box-show__img');
        boxShowCardImg = document.querySelector('.box-show__card_img');
        elementsBox.push(boxShowImg);
        elementsBox.push(boxShowCardImg);
        elementsBox.push(...labelsImg);
        for (let element of elementsBox) {
            element.innerHTML = '';
            element.style.backgroundImage = `url('${dataURL}')`;
            element.classList.add('addImg');
        }  //прогрузка фото автора на блоки превью 

        for (let labelImg of labelsImg) {
            let divImg = document.createElement('div');
            labelImg.parentNode.replaceChild(divImg, labelImg);
            divImg.classList.add('form__hero-img');
            if (labelImg.classList.contains('form__hero-img-big_size')) {
                divImg.classList.add('form__hero-img-big_size');
            } else {
                divImg.classList.add('form__hero-img-little_size');
            }
            divImg.innerHTML = '';
            divImg.style.backgroundImage = `url('${dataURL}')`;
            divImg.classList.add('addImg');
        } //прогрузка фото автора на блок form__hero

        let pSubnames = document.querySelectorAll('.form__subname_font');
        if (pSubnames.length > 0) {
            for (let pSubname of pSubnames) {
                pSubname.remove();
            }
        }  //удаление subname

        let buttonsDisplay = document.querySelectorAll('.buttons__display');
        if (buttonsDisplay.length > 0) {
            for (let button__display of buttonsDisplay) {
                button__display.remove();
            }
        }

        let divHeroImgs = document.querySelectorAll('.form__hero-image');
        for (let divHeroImg of divHeroImgs) {
            let labelChangeImg = document.createElement('label');
            labelChangeImg.classList.add('addImg__button_display');
            let divRemoveImg = document.createElement('div');
            divRemoveImg.classList.add('addImg__button_display');
            divRemoveImg.setAttribute('onclick', 'removeImage(event)');

            let divChangeInput = document.createElement('div');
            divChangeInput.classList.add('addImg__button');
            divChangeInput.classList.add('addImg__size');

            let inputChangeAuthorPhoto = document.createElement('input');
            inputChangeAuthorPhoto.name = 'avatar';
            inputChangeAuthorPhoto.type = 'file';
            inputChangeAuthorPhoto.setAttribute('onchange', 'loadImage(event)');
            divChangeInput.appendChild(inputChangeAuthorPhoto);

            let pChangeImg = document.createElement('p');
            pChangeImg.innerText = 'Upload New';
            pChangeImg.classList.add('form__author-photo_upload');

            let divRemoveButton = document.createElement('div');
            divRemoveButton.classList.add('removeImg__button');
            divRemoveButton.classList.add('addImg__size');
            
            let pRemoveImg = document.createElement('p');
            pRemoveImg.innerText = 'Remove';
            pRemoveImg.style.color = '#E86961';
            pRemoveImg.classList.add('form__author-photo_upload');

            let divBottomBotun = document.createElement('div');
            divBottomBotun.classList.add('buttons__display');

            labelChangeImg.appendChild(divChangeInput);
            labelChangeImg.appendChild(pChangeImg);
        
            divBottomBotun.appendChild(labelChangeImg);
            divRemoveImg.appendChild(divRemoveButton);
            divRemoveImg.appendChild(pRemoveImg);
            divBottomBotun.appendChild(divRemoveImg);

            divHeroImg.appendChild(divBottomBotun);
        }  //создание кнопок новой загрузки и удаления
    };
    reader.readAsDataURL(input.files[0]);
    hero_name = input.files[0].name;
}

function removeImage(event) {
    let divImg = document.querySelector('.form__hero-img');
    let labelImgBig = document.createElement('label');
    divImg.parentNode.replaceChild(labelImgBig, divImg);
    labelImgBig.classList.add('form__hero-img');
    labelImgBig.classList.add('form__hero-img-big_size');

    let inputChangeImg = document.createElement('input');
    inputChangeImg.name = 'image';
    inputChangeImg.type = 'file';
    inputChangeImg.setAttribute('onchange', 'loadImage(event)');
    labelImgBig.appendChild(inputChangeImg);

    let imgIcon = document.createElement('img');
    imgIcon.src = '/static/image/camera.svg';
    labelImgBig.appendChild(imgIcon);

    let pUpload = document.createElement('p');
    pUpload.classList.add('form__author-photo_upload');
    pUpload.innerText = 'Upload'
    labelImgBig.appendChild(pUpload);

    let boxShowCardImg = document.querySelector('.box-show__card_img');
    boxShowCardImg.style = '';
    boxShowCardImg.classList.remove('addImg');
    boxShowCardImg.classList.add('box-show__card_img');

    let boxShowTitleImg = document.querySelector('.box-show__img');
    boxShowTitleImg.style = '';
    boxShowTitleImg.classList.remove('addImg');
    boxShowTitleImg.classList.add('box-show__img');

    let divBoxs = document.querySelectorAll('.buttons__display');
    for (let divBox of divBoxs) {
        divBox.remove();
    }

    let divImgBig = document.querySelector('.form__hero-image');

    let pHint = document.createElement('p');
    pHint.innerText = 'Size up to 10mb. Format: png, jpeg, gif.';
    pHint.classList.add('form__subname_font');
    divImgBig.appendChild(pHint);
}

async function loadAvatar(event) {
    let input = event.target;
    let reader = new FileReader();
    reader.onload = () => {
        let dataURL = reader.result;
        let image = document.createElement('img');
        image.src = dataURL;
        avatar = dataURL;
        

        boxShowCardImg = document.querySelector('.box-show__card_author-img');
        boxShowCardImg.innerHTML = '';
        boxShowCardImg.style.backgroundImage = `url('${dataURL}')`;
        boxShowCardImg.classList.add('addImg');

        let divsChange = document.getElementById('load-avatar');
        if (divsChange !== null) {
            divsChange.remove();
        }

        let labelAuthorPhoto = document.querySelector('.form__author-photo');

        let divButtons = document.createElement('div');
        divButtons.id = 'load-avatar';
        divButtons.classList.add('buttons__display');

        let labelChangeAuthorPhoto = document.createElement('label');
        labelChangeAuthorPhoto.classList.add('addImg__button_display');
        let divRemoveAuthorPhoto = document.createElement('div');
        divRemoveAuthorPhoto.classList.add('addImg__button_display');
        divRemoveAuthorPhoto.setAttribute('onclick', 'removeAvatar(event)');

        let divChangeInput = document.createElement('div');
        divChangeInput.classList.add('addImg__button');
        divChangeInput.classList.add('addImg__size');

        let inputChangeAuthorPhoto = document.createElement('input');
        inputChangeAuthorPhoto.name = 'avatar';
        inputChangeAuthorPhoto.type = 'file';
        inputChangeAuthorPhoto.setAttribute('onchange', 'loadAvatar(event)');
        divChangeInput.appendChild(inputChangeAuthorPhoto);

        let pChangeAuthorPhoto = document.createElement('p');
        pChangeAuthorPhoto.innerText = 'Upload New';
        pChangeAuthorPhoto.classList.add('form__author-photo_upload');

        let divRemoveButton = document.createElement('div');
        divRemoveButton.classList.add('removeImg__button');
        divRemoveButton.classList.add('addImg__size');
        
        let pRemoveAuthorPhoto = document.createElement('p');
        pRemoveAuthorPhoto.innerText = 'Remove';
        pRemoveAuthorPhoto.style.color = '#E86961';
        pRemoveAuthorPhoto.classList.add('form__author-photo_upload');
    
        let pForm = document.querySelector('.form__author-photo_upload');
        if (pForm !== null) {
            pForm.remove();
        }

        let divAuthorPhoto = document.createElement('div');
        labelAuthorPhoto.parentNode.replaceChild(divAuthorPhoto, labelAuthorPhoto);

        let imgAuthor = document.createElement('div');
        imgAuthor.style.backgroundImage = `url('${dataURL}')`;
        imgAuthor.classList.add('addImg__size');
        imgAuthor.classList.add('addImg');
        divAuthorPhoto.appendChild(imgAuthor);

        labelChangeAuthorPhoto.appendChild(divChangeInput);
        labelChangeAuthorPhoto.appendChild(pChangeAuthorPhoto);
        divAuthorPhoto.classList.add('form__author-photo');
        divAuthorPhoto.appendChild(labelChangeAuthorPhoto);

        divRemoveAuthorPhoto.appendChild(divRemoveButton);
        divRemoveAuthorPhoto.appendChild(pRemoveAuthorPhoto);

        divAuthorPhoto.appendChild(divRemoveAuthorPhoto);
    };
    reader.readAsDataURL(input.files[0]);
    avatar_name = input.files[0].name;
}

function removeAvatar(event) {
    let divAuthorPhoto = document.querySelector('.form__author-photo');
    let labelAuthorPhoto = document.createElement('label');
    divAuthorPhoto.parentNode.replaceChild(labelAuthorPhoto, divAuthorPhoto);
    labelAuthorPhoto.classList.add('form__author-photo');
    
    let divAuthorImgPhoto = document.createElement('div');
    divAuthorImgPhoto.classList.add('form__author-photo_img');

    let inputChangeAuthorPhoto = document.createElement('input');
    inputChangeAuthorPhoto.name = 'avatar';
    inputChangeAuthorPhoto.type = 'file';
    inputChangeAuthorPhoto.setAttribute('onchange', 'loadAvatar(event)');
    divAuthorImgPhoto.appendChild(inputChangeAuthorPhoto);

    let pForm = document.createElement('p');
    pForm.classList.add('form__author-photo_upload');
    pForm.innerText = 'Upload';

    labelAuthorPhoto.appendChild(divAuthorImgPhoto);
    labelAuthorPhoto.appendChild(pForm);

    let boxShowCardImg = document.querySelector('.box-show__card_author-img');
    boxShowCardImg.style = '';
    boxShowCardImg.classList.remove('addImg');
    boxShowCardImg.classList.add('box-show__card_author-img');
}

async function setTitle() {
    let element = document.getElementById('input-title');
    let inputValue = element.value;
    let elementInsert = document.querySelector('.box-show__title');
    let elementInsert2 = document.querySelector('.box-show__card_title');
    elementInsert.innerText = inputValue;
    elementInsert2.innerText = inputValue;
    title = inputValue;
}

async function setDiscription() {
    let element = document.getElementById('input-discription');
    let inputValue = element.value;
    let elementInsert = document.querySelector('.box-show__subtitle');
    let elementInsert2 = document.querySelector('.box-show__card_subtitle');
    elementInsert.innerText = inputValue;
    elementInsert2.innerText = inputValue;
    description = inputValue;
}

async function setAuthorName() {
    let element = document.getElementById('input-name');
    let inputValue = element.value;
    let elementInsert = document.getElementById('avtar-name');
    elementInsert.innerText = inputValue;
    author = inputValue;
}

async function setDate() {
    let element = document.getElementById('input-date');
    let inputValue = element.value;
    let elementInsert = document.getElementById('date');
    elementInsert.innerText = inputValue;
    date = inputValue;
}
