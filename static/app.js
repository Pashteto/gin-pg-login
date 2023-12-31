
function qs(selector) {
	return document.querySelector(selector)
}

function qsa(selector) {
	return document.querySelectorAll(selector)
}

function showBigLoader() {
	let bigLoader = qs("#big-loader")
	bigLoader.classList.toggle('hidden')
	let loadingOverlay = qs('#big-loader-overlay')
	loadingOverlay.classList.toggle('hidden')
}

function hideBigLoader() {
	let bigLoader = qs("#big-loader")
	let loadingOverlay = qs('#big-loader-overlay')
	if (bigLoader != null) {
		bigLoader.classList.add('hidden')
	}
	if (loadingOverlay != null) {
		loadingOverlay.classList.add('hidden')
	}
}

function openNav() {
	let hamburger = qs('#nav-hamburger')
	let menu = qs('#nav-menu')
	let x = qs('#nav-x')
	let overlay = qs('#nav-overlay')
	links = menu.querySelectorAll('li')
	links.forEach(link => {
		let href = link.children[0].getAttribute('href')
		if (href == window.location.pathname) {
			link.classList.add('bg-gray-200')
			link.classList.add('border-gray-400')
			link.children[0].classList.add('bg-gray-200')
		}
	});
	if (hamburger != null) {
		hamburger.classList.add("hidden")
	}
	if (menu != null) {
		menu.classList.remove("hidden")
	}
	if (x != null) {
		x.classList.remove("hidden")
	}
	if (overlay != null) {
		overlay.classList.remove('hidden')
	}
}

function closeNav() {
	let hamburger = qs('#nav-hamburger')
	let menu = qs('#nav-menu')
	let x = qs('#nav-x')
	let overlay = qs('#nav-overlay')
	if (hamburger != null) {
		hamburger.classList.remove("hidden")
	}
	if (menu != null) {
		menu.classList.add('hidden')
	}
	if (x != null) {
		x.classList.add("hidden")
	}
	if (overlay != null) {
		overlay.classList.add('hidden')
	}
}

// Function to parse URL data
function parseUrlData() {
	const fragment = window.location.hash.substring(1);
	const params = new URLSearchParams(fragment);
	const userParam = params.get('tgWebAppData');
	const decodedUserParam = decodeURIComponent(userParam);
	const userJson = decodedUserParam.substring(decodedUserParam.indexOf('{'), decodedUserParam.lastIndexOf('}') + 1);
	const user = JSON.parse(decodeURIComponent(userJson));
	return user;
}

function handleFormSubmit(event) {
	event.preventDefault();  // prevent form from submitting immediately
	showBigLoader();

	// You can then submit the form programmatically after your logic
	document.getElementById('signup-form').submit();
}


document.getElementById('login-form').addEventListener('submit', function() {
	// Get current URL fragment
	const fragment = window.location.hash;

	// Append it to the form's action URL
	this.action += fragment;

	const fragment1 = sessionStorage.getItem('urlFragment');
	if (fragment1) {
		window.location.hash = fragment1;
	}
});

// Save fragment to session storage when page loads
if (window.location.hash) {
	sessionStorage.setItem('urlFragment', window.location.hash);
}