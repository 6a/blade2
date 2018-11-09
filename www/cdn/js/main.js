function setOpacityZero(id) {
    var element = document.getElementById(id);
    element.classList.add("opacity-0");
}

function switchWindow(index) {
    localStorage.windowIndex = index;
    for (let i = 0; i < this.contentWindows.length; i++) {
        if (i == index) {
            this.contentWindows[i].classList.add("content-active");
            this.windowSelectButtons[i].classList.add("select-active");
        } else {
            this.contentWindows[i].classList.remove("content-active");
            this.windowSelectButtons[i].classList.remove("select-active");
        }
    }

    if (index == 3) {
        this.filter.classList.add("display-none");
    } else {
        this.filter.classList.remove("display-none");
    }
}

function init() {
    this.contentWindows = document.getElementsByClassName("body-content");
    this.windowSelectButtons = document.getElementsByClassName("body-banner-opt");
    this.filter = document.getElementById("filter");
    for (let i = 0; i < this.windowSelectButtons.length; i++) {
        this.windowSelectButtons[i].addEventListener("click", function () { switchWindow(i) });
    }

    if (localStorage.windowIndex) {
        switchWindow(localStorage.windowIndex);
    }
}

init();