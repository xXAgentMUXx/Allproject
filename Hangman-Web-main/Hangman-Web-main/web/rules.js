function myFunction() {
    document.getElementById("Whangman").classList.toggle("show");
}

function myFunction2() {
    document.getElementById("Rhangman").classList.toggle("show");
}

function myFunctionDifficult(){
    document.getElementById("Dhangman").classList.toggle("show");
}

window.onclick = function(event) {
    // Fermer le menu associé à `.dropW`
    if (!event.target.matches('.dropW')) {
        var dropdownsW = document.getElementsByClassName("dropdown-content-W");
        for (var i = 0; i < dropdownsW.length; i++) {
            var openDropdownW = dropdownsW[i];
            if (openDropdownW.classList.contains('show')) {
                openDropdownW.classList.remove('show');
            }
        }
    }

    // Fermer le menu associé à `.dropR`
    if (!event.target.matches('.dropR')) {
        var dropdownsR = document.getElementsByClassName("dropdown-content-R");
        for (var i = 0; i < dropdownsR.length; i++) {
            var openDropdownR = dropdownsR[i];
            if (openDropdownR.classList.contains('show')) {
                openDropdownR.classList.remove('show');
            }
        }
    }

    // Fermer le menu associé à `.dropD``
    if (!event.target.matches('.dropD')) {
        var dropdownsD = document.getElementsByClassName("dropdown-content-D");
        for (var i = 0; i < dropdownsD.length; i++) {
            var openDropdownD = dropdownsD[i];
            if (openDropdownD.classList.contains('show')) {
                openDropdownD.classList.remove('show');
            }
        }
    }
};
