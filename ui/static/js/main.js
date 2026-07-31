
const handleNav = () => {
	const navLinks = document.querySelectorAll("nav a");
	
	if(navLinks){
		navLinks.forEach(( link ) => {
			if (link.getAttribute('href') == window.location.pathname) {
				link.classList.add("active");
			}
		})
	}
}

document.addEventListener("DOMContentLoaded", handleNav);