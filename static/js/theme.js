const themes = ['mint', 'plain', 's.berry', 'banana', 'peanut', 'grape', 'melon'];

let stored_theme = window.localStorage.getItem('theme')

if (stored_theme == null) {
    stored_theme = 0
}

window.localStorage.setItem('theme', stored_theme)

document.documentElement.setAttribute('data-theme', themes[stored_theme]);
console.log(document.documentElement.getAttribute("data-theme"))


function change_theme() {
    if (browser) {
        let next_theme = (stored_theme + 1) % themes.length;
        document.documentElement.setAttribute('data-theme', themes[stored_theme]);
        stored_theme = next_theme
        window.localStorage.setItem('theme', stored_theme)
    }
}