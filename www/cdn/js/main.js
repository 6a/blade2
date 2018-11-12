const MAX_IMG_SIZE = 1024000;

function setOpacityZero(id) {
    var element = document.getElementById(id);
    element.classList.add("opacity-0");
}

function switchWindow(index) {
    if (!this.firstRun && localStorage.windowIndex == index) { return; }
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

    if (index == this.profileWindowIndex) {
        this.filter.classList.add("display-none");
    } else {
        this.filter.classList.remove("display-none");
        if (this.previousWindowIndex == this.profileWindowIndex) {
            if (index == this.searchWindowIndex) {
                this.filter.placeholder = "Search by name or player ID";
                this.filter.value = this.storedSearchValue;
            } else {
                this.filter.placeholder = "Filter results";
                this.filter.value = this.storedFilterValue;
            }
        } else if (this.previousWindowIndex == this.searchWindowIndex) {
            this.filter.placeholder = "Filter results";
            this.storedSearchValue = this.filter.value;
            this.filter.value = this.storedFilterValue;
        } else {
            this.storedFilterValue = this.filter.value;
            if (index == this.searchWindowIndex) {
                this.filter.placeholder = "Search by name or player ID";
                this.filter.value = this.storedSearchValue;
            }
        }
    }
    this.previousWindowIndex = index;
    this.firstRun = false;
}

function init() {
    this.contentWindows = document.getElementsByClassName("body-content");
    this.windowSelectButtons = document.getElementsByClassName("body-banner-opt");
    this.filter = document.getElementById("filter");
    this.profilePicture = document.getElementById("profile-pic");

    this.profileWindowIndex = this.contentWindows.length - 1;
    this.searchWindowIndex = 3;
    this.storedSearchValue = "";
    this.storedFilterValue = "";
    this.previousWindowIndex = 0;
    this.firstRun = true;

    for (let i = 0; i < this.windowSelectButtons.length; i++) {
        this.windowSelectButtons[i].addEventListener("click", function () { switchWindow(i) });
    }

    if (localStorage.windowIndex) {
        switchWindow(localStorage.windowIndex);
    }

    this.profilePicture.addEventListener("click", function () {
        var fileSelector = document.createElement('input');
        fileSelector.setAttribute('type', 'file');
        fileSelector.setAttribute('accept', 'image/*');
        fileSelector.click();
        // fileSelector.addEventListener("submit", function(e) {console.log(e);});
        fileSelector.addEventListener("change", function (e) {
            var fReader = new FileReader();
            fReader.readAsDataURL(e.path[0].files[0]);
            fReader.onloadend = function(event){
                if (event.total > MAX_IMG_SIZE) {
                    window.alert("Invalid image: Must be an image file, no larger than 1024Kb");
                } else {
                    profilePicture.style.backgroundImage = "url('" + event.target.result +"')";
                }
            }
        });
    });
}

init();