const parallaxImages = document.querySelectorAll('img');

/*
    Parallax effect on images
*/

function handleMouseParallax(event) {
    const image = event.currentTarget;
    const rect = image.getBoundingClientRect();
    const x = event.clientX - rect.left;
    const y = event.clientY - rect.top;

    const rotateX = (rect.height / 2 - y) / rect.height * 30;
    const rotateY = (rect.width / 2 - x) / rect.width * -30;

    image.style.transform = `perspective(1000px) rotateX(${rotateX}deg) rotateY(${rotateY}deg) scale(1.1)`;
}

function resetParallax(event) {
    event.currentTarget.style.transform = 'perspective(1000px) rotateX(0) rotateY(0) scale(1)';
}

parallaxImages.forEach(image => {
    image.addEventListener('mousemove', handleMouseParallax);
    image.addEventListener('mouseleave', resetParallax);
});

const artists = document.querySelectorAll('.artist');

/*
    Scroll fade-in reveal effect on main page
*/

function handleScroll() {
    const scrollPosition = window.scrollY;
    const windowHeight = window.innerHeight;

    artists.forEach(artist => {
        const rect = artist.getBoundingClientRect();
        const offset = rect.top + scrollPosition;

        if (scrollPosition + windowHeight > offset) {
            artist.style.opacity = 1;
            artist.style.transform = 'translateY(0)';
        } else {
            artist.style.opacity = 0;
            artist.style.transform = 'translateY(50px)';
        }
    });
}

window.addEventListener('scroll', handleScroll);
document.addEventListener('DOMContentLoaded', handleScroll);

/*
    Search bar and filter animation handlers
*/

document.addEventListener('DOMContentLoaded', function () {
    const searchButton = document.querySelector('.search-button');
    const searchInput = document.querySelector('.js-search');

    searchButton.addEventListener('click', function () {
        searchInput.focus();
    });

    const filterButton = document.querySelector('.filter-button');
    const filterSection = document.querySelector('.filter-section');

    filterButton.addEventListener('click', function () {
        filterSection.style.display = filterSection.style.display === 'flex' ? 'none' : 'flex';
    });

});

const creationDateStart = document.getElementById('creation-date-start');
const creationDateEnd = document.getElementById('creation-date-end');
const creationDateStartLabel = document.getElementById('creation-date-start-label');
const creationDateEndLabel = document.getElementById('creation-date-end-label');

const firstAlbumDateStart = document.getElementById('first-album-date-start');
const firstAlbumDateEnd = document.getElementById('first-album-date-end');
const firstAlbumDateStartLabel = document.getElementById('first-album-date-start-label');
const firstAlbumDateEndLabel = document.getElementById('first-album-date-end-label');

function updateCreationDateRange() {
    let start = parseInt(creationDateStart.value);
    let end = parseInt(creationDateEnd.value);

    if (start > end) {
        [start, end] = [end, start];
        creationDateStart.value = start;
        creationDateEnd.value = end;
    }

    creationDateStartLabel.textContent = start;
    creationDateEndLabel.textContent = end;
}

function updateFirstAlbumDateRange() {
    let start = parseInt(firstAlbumDateStart.value);
    let end = parseInt(firstAlbumDateEnd.value);

    if (start > end) {
        [start, end] = [end, start];
        firstAlbumDateStart.value = start;
        firstAlbumDateEnd.value = end;
    }

    firstAlbumDateStartLabel.textContent = start;
    firstAlbumDateEndLabel.textContent = end;
}

creationDateStart.addEventListener('input', updateCreationDateRange);
creationDateEnd.addEventListener('input', updateCreationDateRange);
firstAlbumDateStart.addEventListener('input', updateFirstAlbumDateRange);
firstAlbumDateEnd.addEventListener('input', updateFirstAlbumDateRange);

updateCreationDateRange();
updateFirstAlbumDateRange();
