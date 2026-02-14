const themes = ['plain', 'mint', 's.berry', 'banana', 'peanut', 'grape', 'melon'];

let stored_theme = window.localStorage.getItem('theme')

if (stored_theme == null) {
    stored_theme = 0
}

window.localStorage.setItem('theme', stored_theme)

document.documentElement.setAttribute('data-theme', themes[stored_theme]);

update_theme_text()


function change_theme() {
    stored_theme = (stored_theme + 1) % themes.length;
    document.documentElement.setAttribute('data-theme', themes[stored_theme]);
    window.localStorage.setItem('theme', stored_theme)
    update_theme_text()
}

function update_theme_text() {
    let themetext = document.getElementById("theme")
    themetext.textContent = themes[stored_theme]
}

window.onload = function () {
    update_theme_text()
};