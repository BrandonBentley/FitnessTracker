let userData;

$(function () {
    result = $.getJSON("/json", function (data) {
        userData = data;
        displayContent();
        generateQuickItems();
    });
    if (result.fail()) {
        checkForWorkouts();
        console.log("failed");
    }
});

function displayContent() {
    clearElements();
    generateCards();
    applyHandlers();
    console.log("display refreshed");
}

function clearElements() {
    $('.noContent').remove();
    $('.workoutCard').remove();
    //$('.rightNavItem').remove();
}

function checkForWorkouts() {
    if (userData == null) {
        $(".cardCenter").append(
            `<div class="noContent">
                <h2 class="text-secondary">No Workouts Yet :(</h2>
                    ${circleSpinner}
             </div>`
        );
    }
}

function applyHandlers() {
    $('.workoutCard').on("click", displayContent);
}

function generateCards() {
    for (let item of userData.workouts) {
        $(".cardCenter").append(
            `<div class="card workoutCard">
                    <div class ="card-body workoutCardContent">
                        <p>${item.title}</p>
                    </div>
                </div>`
        );
    }
}

function generateQuickItems() {
    for (item of userData.quickItems) {
        $(".RightOptionMenuList").append(
            `<li class="nav-item rightNavItem">
            <a class="nav-link text-secondary" href="#">${item.title}</a>
         </li>`
        );
    }
}