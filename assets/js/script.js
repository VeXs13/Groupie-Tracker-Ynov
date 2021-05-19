//getting all required elements
const searchWrapper = document.querySelector(".searchauto");
const inputBox = searchWrapper.querySelector("input");
const suggest = searchWrapper.querySelector(".autocomplete");
const suggestions = ['Queen', 'SOJA', 'Pink Floyd', 'Scorpions', 'XXXTentacion', 'Mac Miller', 'Joyner Lucas', 'Kendrick Lamar', 'ACDC', 'Pearl Jam', 'Katy Perry', 'Rihanna', 'Genesis', 'Phil Collins', 'Led Zeppelin', 'The Jimi Hendrix Experience', 'Bee Gees', 'Deep Purple', 'Aerosmith', 'Dire Straits', 'Mamonas Assassinas', 'Thirty Seconds to Mars', 'Imagine Dragons', 'Juice Wrld', 'Logic', 'Alec Benjamin', 'Bobby McFerrins', 'R3HAB', 'Post Malone', 'Travis Scott', 'J. Cole', 'Nickelback', 'Mobb Deep', "'Guns N' Roses'", 'NWA', 'U2', 'Arctic Monkeys', 'Fall Out Boy', 'Gorillaz', 'Eagles', 'Linkin Park', 'Red Hot Chili Peppers', 'Eminem', 'Green Day', 'Metallica', 'Coldplay', 'Maroon 5', 'Twenty One Pilots', 'The Rolling Stones', 'Muse', 'Foo Fighters', 'The Chainsmokers'];
inputBox.onkeyup = (e) => {
    let userData = e.target.value; //user entered data

    let emptyArray = [];
    if (userData) {
        emptyArray = suggestions.filter((data) => {
            //filtering array value and user char to lowercase and return only those word/sentence which starts with user entered word
            return data.toLocaleLowerCase().startsWith(userData.toLocaleLowerCase());
        });
        emptyArray = emptyArray.map((data) => {
            return (data = "<li>" + data + "</li>");
        });
        console.log(emptyArray);
        searchWrapper.classList.add("active") //show autocomplete box
        showsuggestions(emptyArray)
        let alllist = document.querySelectorAll("li")
        for (let i = 0; i < alllist.length; i++) {
            //adding onclick attribute in all li tags
            alllist[i].setAttribute("onclick", "select(this)")
        }
        const cardContent = document.querySelectorAll('.card');

        for (i = 0; i < cardContent.length; i++) {

            const cardHeader = cardContent[i].querySelector('.card-header');

            txtValue = cardHeader.textContent || cardHeader.innerText;

            if (txtValue.toUpperCase().indexOf(userData.toUpperCase()) > -1) {

                cardContent[i].style.display = "";

            } else {

                cardContent[i].style.display = "none";

            }

        }
    } else {
        searchWrapper.classList.remove("active") //hide autocomplete box
    }
};

function select(element) {
    let selectedData = element.textContent
    inputBox.value = selectedData //passing the selected Data in the search box
}

function showsuggestions(list) {
    let Datalist;
    if (!list.length) {
        Datalist = '<li>' + inputBox.value + '</li>'

    } else {
        Datalist = list.join("");
    }
    suggest.innerHTML = Datalist;
}