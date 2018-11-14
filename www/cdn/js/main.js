const MAX_IMG_SIZE = 1024000;
const VALID_AVATAR_TYPES = ['image/png', 'image/jpeg', 'image/gif'];

function b2Kb (bytes) {
    return bytes / 1000;
}

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
        fileSelector.addEventListener("change", function (e) {
            if (e.path.length && e.path[0].files && e.path[0].files.length) {
                var image = new Image();
                image.onload = function (e) {
                    if ('naturalHeight' in this) {
                        if (this.naturalHeight + this.naturalWidth === 0) {
                            image.onerror();
                            return;
                        }
                    } 

                    if (this.width + this.height == 0) {
                        image.onerror();
                        return;
                    }

                    profilePicture.style.backgroundImage = "url('" + this.src + "')";                  
                }

                image.onerror = function () {
                    window.alert("Invalid file: Must be a jpeg, png or gif, no larger than " + b2Kb(MAX_IMG_SIZE) + "Kb.");
                    URL.revokeObjectURL(this.src);
                }

                if (e.path[0].files[0].size > MAX_IMG_SIZE) {
                    image.onerror();
                    return;
                }

                if (!VALID_AVATAR_TYPES.includes(e.path[0].files[0].type)) {
                    image.onerror();
                    return;
                }

                var url = window.URL || window.webkitURL;
                image.src = url.createObjectURL(e.path[0].files[0]);
            }
        });
    });
}

init();