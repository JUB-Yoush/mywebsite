const hobbies = ["remaking this website.", "still modding Nintendo stuff.", "learning another Vim shortcut.", "breaking my linux installation.", "encouraging people to try Raylib.", "finishing my drawabox homework.", "learning about the Godot engine.", "playing the piano.", "making jazz playlists.", "thinking about JRPGs.", "making up new serves in ping pong.", "mashing in Guilty Gear.", "reading the new One Piece chapter.", "emulating retro games.", "beating my cybergrind high-score.", "being glad they ported Final Fantasy Tactics.", "trying to put on 5 lbs of muscle.", "designing Mercury Man lore.", "gifting you Void Stranger on Steam.", "organizing my life in plaintext.", "making apps for my Pebble.", "losing the lane in Deadlock."]

let order = []
let current_hobby = ''
function shuffle_order() {
    while (order.length < hobbies.length) {
        let rng = Math.floor(Math.random() * hobbies.length)
        if (!order.includes(rng)) {
            order.push(rng)
        }
    }
    next_hobby()
}

function next_hobby() {
    current_hobby = hobbies[order.pop()]
    if (order.length == 0) {
        shuffle_order()
    }
    let anchor_text = document.getElementById("hobby")
    anchor_text.textContent = current_hobby

}
document.addEventListener("DOMContentLoaded", function () {
    shuffle_order()
});