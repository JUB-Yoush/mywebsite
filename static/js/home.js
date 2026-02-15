const gif_tuple = [["https://media1.tenor.com/m/MZRbDClAvOgAAAAC/capybara-bath.gif", "capybara"], ["https://img.itch.zone/aW1nLzEzNjI4MjU0LnBuZw==/original/11RCpN.png", "There is no gif. Play Void Stranger."], ["https://64.media.tumblr.com/f2d0490d0300ce46a1e7a7f15dc74165/tumblr_nbbb83YqbU1swqiquo1_500.gif", "Get 8 hours of sleep"], ["https://media.tenor.com/WkI4kRS7eqMAAAAM/broly-dbz-fanmade.gif", "mfw 8 hours of sleep"], ["https://media3.giphy.com/media/v1.Y2lkPTc5MGI3NjExMHh5NGlvdWhrZDZqY3Q4dHhmdGF0MGEzYXNyM3M2NDNlaXc4bWZ1dyZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/ZIIyYOR6vXTGeEXTZu/giphy.gif", "Check out the site's colour themes."]]
let gif_number = Math.floor(Math.random() * gif_tuple.length)


document.getElementById("home_gif_id").src = gif_tuple[gif_number][0]
document.getElementById("gif_caption").textContent = gif_tuple[gif_number][1]

function set_gif() {

    document.getElementById("home_gif_id").src = gif_tuple[gif_number][0]
}

function change_gif() {
    gif_number = (gif_number + 1) % gif_tuple.length;

    document.getElementById("home_gif_id").src = gif_tuple[gif_number][0]
    document.getElementById("gif_caption").textContent = gif_tuple[gif_number][1]

}
