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

function getYourGames() {
    function getHTML (data) {
        const item = `
        <div class="body-content-entry">
            <div class="body-content-item text-active-highlight-desc">${data.gameId}</div>
            <div class="body-content-item">${data.winnerName}</div>
            <div class="body-content-item">${data.winnerElo} (+${data.winnerEloDelta})</div>
            <div class="body-content-item">${data.loserName}</div>
            <div class="body-content-item">${data.loserElo} (${data.loserEloDelta})</div>
            <div class="body-content-item">${data.dateTime}</div>
        </div>
        `;

        return item;
    }

    // TODO replace with actual data getting
    const data = [
        {gameId: "7FFFFFFFFFFFFFFF", winnerName: "james", winnerElo: 2400, winnerEloDelta: 24, loserName: "ryan", loserElo: 2400, loserEloDelta: -34, dateTime: "11:45 04/12/1991"},
        {gameId: "7FFFFFFFFFAFFFFF", winnerName: "james", winnerElo: 2376, winnerEloDelta: 23, loserName: "dave", loserElo: 2398, loserEloDelta: -23, dateTime: "11:22 04/12/1991"},
        {gameId: "7FFFFFF3DFFFFFFF", winnerName: "james", winnerElo: 2353, winnerEloDelta: 25, loserName: "mark", loserElo: 2322, loserEloDelta: -23, dateTime: "11:01 04/12/1991"},
        {gameId: "7FFAAFFFFFFFFFFF", winnerName: "james", winnerElo: 2332, winnerEloDelta: 21, loserName: "pork", loserElo: 2300, loserEloDelta: -41, dateTime: "08:22 03/12/1991"},
        {gameId: "7FFAAFFFFFFFFFFF", winnerName: "matthew", winnerElo: 976, winnerEloDelta: 21, loserName: "james", loserElo: 2300, loserEloDelta: -11, dateTime: "00:22 01/11/1991"}
    ]; 

    data.forEach(element => {
        $('#dynamic-content-yourgames').append(getHTML(element));
    });
}

function getTopRatings() {
    function getHTML (data) {
        const item = `
        <div class="body-content-entry decorate-on-hover">
            <div class="body-content-item body-content-id-field text-active-highlight-asc">${data.rank}</div>
            <div class="body-content-item">${data.name}</div>
            <div class="body-content-item">${data.elo}</div>
            <div class="body-content-item">${data.wins}</div>
            <div class="body-content-item">${data.losses}</div>
            <div class="body-content-item">${data.draws}</div>
            <div class="body-content-item">${data.ratio}</div>
            <a class="search-entry-link" href="./accounts/${data.id}/profile"></a>
        </div>
        `;

        return item;
    }

    // TODO replace with actual data getting
    const data = [
        {id: "7FFFFFFFFFFFFFFF", rank: 1, name: "james", elo: 2400, wins: 65, losses: 44, draws: 5, ratio: 1.48},
        {id: "7FFFFFFFFFAFFFFF", rank: 2, name: "dave", elo: 2398, wins: 64, losses: 45, draws: 1, ratio: 1.42},
        {id: "7FFFFFF3DFFFFFFF", rank: 3, name: "mark", elo: 2322, wins: 48, losses: 40, draws: 3, ratio: 1.2},
        {id: "7FFAAFFFFFFFFFFF", rank: 4, name: "pork", elo: 2300, wins: 153, losses: 148, draws: 3, ratio: 1.03}
    ]; 

    data.forEach(element => {
        $('#dynamic-content-elo').append(getHTML(element));
    });
}

function sort(event) {
    var source = event.srcElement;
    var data = source.dataset;
    var window = $("#" + data.window);
    var header = window[0].children[0];
    var children = ([].slice.call(window[0].children, 0)).splice(1, window[0].children.length - 1);
    var sortType = data.sort;

    function sortArray(array, column, ascending) {
        array.sort(function(a, b) {
            var secondarySortIndex = 0;
            function compare(a, b, column) {
                var valueA = (a.children[column].textContent);
                var valueB = (b.children[column].textContent);
                if (isNaN(valueA)) {
                    if (valueA.length < valueB.length) {
                        valueA = valueA.padStart(valueB.length, "0");
                    } else if (valueA.length > valueB.length) {
                        valueB = valueB.padStart(valueA.length, "0");
                    }

                    if (ascending) this.comparison = valueA.localeCompare(valueB);
                    else this.comparison = -valueA.localeCompare(valueB);
                    if (this.comparison == 0) {
                        column = 0 + secondarySortIndex;
                        secondarySortIndex++;
                        return compare(a, b, column);
                    }

                    return this.comparison;
                } else {
                    if (ascending) this.comparison = valueA - valueB;
                    else this.comparison = valueB - valueA;

                    if (this.comparison == 0) {
                        column = 0 + secondarySortIndex;
                        secondarySortIndex++;
                        return compare(a, b, column);
                    }

                    return this.comparison;
                }
            }
            
            return compare(a, b, column);
        });
    }
    
    window.html(header);

    switch (sortType) {
        case "n":
            if (data.default == "a") {
                sortArray(children, parseInt(data.column), true);
            } else {
                sortArray(children, parseInt(data.column), false);
            }

            source.setAttribute("data-sort", data.default);
            break;
        case "a":
            sortArray(children, parseInt(data.column), false);
            
            source.setAttribute("data-sort", "d");
            break;  
        case "d":
            sortArray(children, parseInt(data.column), true);


            source.setAttribute("data-sort", "a");
            break;  
    }

    [].slice.call(header.children, 0).forEach(element => {
        if (element != source) {
            element.setAttribute("data-sort", "n");
        }
    });

    children.unshift(header);

    children.forEach(element => {
        for (let index = 0; index < element.children.length; index++) {
            if (index == parseInt(data.column)) {
                if (data.sort == "a") {
                    element.children[index].classList.add("text-active-highlight-asc");
                    element.children[index].classList.remove("text-active-highlight-desc");
                } else {
                    element.children[index].classList.add("text-active-highlight-desc");
                    element.children[index].classList.remove("text-active-highlight-asc");
                }
            } else {
                element.children[index].classList.remove("text-active-highlight-desc");
                element.children[index].classList.remove("text-active-highlight-asc");
            }
        }
    });

    window.html(children);
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

    [].slice.call(document.getElementsByClassName("body-content-item-header"), 0).forEach(element => {
        element.addEventListener("click", e => this.sort(e));
    }); 

    getTopRatings();
    getYourGames();
}

init();