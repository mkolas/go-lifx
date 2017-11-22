function makeFloaters() {
    // IE 11 doesn't render this correctly, but maybe I'll fix it later.
    let emojis = ['🚨','🚨','🚨','🚨','🚨','🚨','🚨','🚨','🚨',
                  '🚨','🚨','🚨','🚨','🚨','🚨','🚨','🚨','🚨'];

    // Fewer emojis on smaller screens.
    if (screen.width < 1000) {
        emojis.length = emojis.length/2;
    }

    let floaters = [];
    emojis.forEach(function(e) {
        let child = document.createElement('span');
        child.classList.add('floater');
        child.innerHTML = e;
        floaters.push(child);
        document.getElementById('emoji-background').appendChild(child);
    });

    let i=1;
    let animCss = '';
    floaters.forEach(function(floater) {
        let topStart = Math.round(Math.random() * 100);
        let topEnd = Math.round(Math.random() * 100);
        let duration = Math.round(Math.random() * 20 + 8);
        let delay = Math.round(Math.random() * 20);
        let rotationStart = Math.round(Math.random() * 360);
        let rotationEnd = Math.round(Math.random() * 360 + 360);
        let leftStrings = ['left: -120px;\n', 'left: calc(100vw + 120px);\n'];
        if (Math.random() > 0.5) {
            leftStrings.unshift(leftStrings.pop());
        }

        floater.classList.add('floater-' + i);
        animCss += '@keyframes emoji-float-' + i + ' ' +
            '{ \n' +
            '  from {\n' +
            '    ' + leftStrings.pop() +
            '    top: ' + topStart + 'vh;\n' +
            '    transform: rotateZ(' + rotationStart + 'deg);\n' +
            '  }\n' +
            '  to {\n' +
            '    ' + leftStrings.pop() +
            '    top: ' + topEnd + 'vh;\n' +
            '    transform: rotateZ(' + rotationEnd + 'deg);\n' +
            '  }\n' +
            '}\n';

        animCss += '#emoji-background .floater-' + i + ' ' +
            '{ \n' +
            '  animation: emoji-float-' + i + ' ' +
            duration + 's ' +
            delay + 's linear infinite alternate; \n' +
            '}\n';
        i++;
    });
    let styleElement = document.createElement('style');
    styleElement.innerHTML = animCss;
    document.body.appendChild(styleElement);
}

let name = "cheater";
let contentHolder = document.getElementById('content-holder');
let nameContent = document.getElementById('name-content');
let formContent = document.getElementById('form-content');
let formLog = document.getElementById('form-log');
let formLogBody = document.getElementById('form-log-body');
let colorPickerObject = document.getElementById('color-picker');
let colorPicker;


function saveName() {
    name = document.getElementById("name").value;
    let request = new XMLHttpRequest();
    request.open('GET', '/recent');
    request.send();
    request.onload = function() {
        colorPicker = new iro.ColorPicker(colorPickerObject);
        formLogBody.innerHTML = this.response;
        setForm();
    }
}

function setName() {
    nameContent.style.display = 'block';
    formContent.style.display = 'none';
}

function setForm() {

    nameContent.style.display = 'none';
    formContent.style.display = 'block';
}

function sendColor() {
    let request = new XMLHttpRequest();
    request.open('POST', '/change');
    let formData = new FormData();
    formData.append('color', colorPicker.color.hexString);
    formData.append('name', name);
    request.send(formData);
    request.onload = function() {
        formLog.style.display = 'block';
        formLogBody.innerHTML = this.response;
    }
}

contentHolder.appendChild(nameContent);
contentHolder.appendChild(formContent);

makeFloaters();
setName();
