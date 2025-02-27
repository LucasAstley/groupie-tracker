const clientId = "";
const clientSecret = "";

async function getSpotifyToken() {
    /*
		This function retrieves a Spotify token using the client ID and client secret
	 */
    const response = await fetch("https://accounts.spotify.com/api/token", {
        method: "POST",
        headers: {
            "Content-Type": "application/x-www-form-urlencoded",
            "Authorization": "Basic " + btoa(clientId + ":" + clientSecret)
        },
        body: "grant_type=client_credentials"
    });

    const data = await response.json();
    return data.access_token;
}

async function loadArtist(artistName) {
    /*
		This function loads the Spotify player with the artist's top tracks
	 */
    const token = await getSpotifyToken();

    const searchResponse = await fetch(`https://api.spotify.com/v1/search?q=${encodeURIComponent(artistName)}&type=artist`, {
        headers: {"Authorization": "Bearer " + token}
    });

    const searchData = await searchResponse.json();
    if (searchData.artists.items.length === 0) {
        alert("Artiste introuvable !");
        return;
    }

    const artistId = searchData.artists.items[0].id;

    document.getElementById("player").innerHTML = `
        <iframe style="border-radius:12px" src="https://open.spotify.com/embed/artist/${artistId}?utm_source=generator" 
        width="100%" height="351vh" frameBorder="0" allowfullscreen="" 
        allow="autoplay; clipboard-write; encrypted-media; fullscreen; picture-in-picture" loading="lazy">
        </iframe>
    `;
}

window.onload = function () {
    const artistName = document.querySelector(".header-title").textContent;
    loadArtist(artistName);
};
