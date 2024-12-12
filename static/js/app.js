function copyToClipboard() {
  const url = window.location.href; 
  navigator.clipboard.writeText(url)
      .then(() => {
          alert('URL copied to clipboard!');
      })
      .catch(err => {
          alert('Failed to copy URL: ' + err);
      });
}


function toggleContent() {
  const serviceCard = document.getElementById("serviceCard");
  serviceCard.classList.toggle("expanded");

  // Cambia el texto y el icono del botón dependiendo del estado
  const seeMoreBtn = document.querySelector(".see-more");
  if (serviceCard.classList.contains("expanded")) {
    seeMoreBtn.textContent = "Ver menos ";
    seeMoreBtn.appendChild(document.createElement("span")).textContent = "▲";
  } else {
    seeMoreBtn.textContent = "Ver más ";
    seeMoreBtn.appendChild(document.createElement("span")).textContent = "▼";
  }
}

/*MENU MOBILE*/
document.addEventListener("DOMContentLoaded", function () {
  var menuToggle = document.querySelector(".menu-toggle");
  var navMenu = document.querySelector(".nav-menu_mob");
  var closeButton = document.querySelector(".close-menu");

  menuToggle.addEventListener("click", function () {
    navMenu.classList.toggle("open");
  });
  closeButton.addEventListener("click", function () {
    navMenu.classList.remove("open");
  });

  document.querySelectorAll(".nav-menu_mob .nav-link").forEach((link) => {
    link.addEventListener("click", (event) => {
      if (link.id === "casos-de-uso-link" || link.id === "idioma-link" || link.id === "blog-link") {
        event.preventDefault();
        return;
      }
      navMenu.classList.remove("open");
    });
  });
});
/*FIN MENU MOBILE*/

/*SUBMENU*/
const casosDeUsoLink = document.getElementById("casos-de-uso-link");
const casosDeUsoSubmenu = document.getElementById("casos-de-uso-submenu");
const backCasosDeUso = document.getElementById("back-casos-de-uso");

const idiomaLink = document.getElementById("idioma-link");
const idiomaSubmenu = document.getElementById("idioma-submenu");
const backIdioma = document.getElementById("back-idioma");

casosDeUsoLink.addEventListener("click", function (event) {
  event.preventDefault(); // Evita el comportamiento predeterminado del enlace
  casosDeUsoSubmenu.classList.add("open");
});

backCasosDeUso.addEventListener("click", function () {
  casosDeUsoSubmenu.classList.remove("open");
});
backIdioma.addEventListener("click", function () {
  idiomaSubmenu.classList.remove("open");
});
idiomaLink.addEventListener("click", function (event) {
  event.preventDefault(); // Evita el comportamiento predeterminado del enlace
  idiomaSubmenu.classList.add("open");
});

/* Handle Blog - New Dropdown */
const blogLink = document.getElementById("blog-link");
const blogSubmenu = document.getElementById("blog-submenu");
const backBlog = document.getElementById("back-blog");

blogLink.addEventListener("click", function (event) {
  event.preventDefault(); // Prevents default link behavior
  blogSubmenu.classList.add("open");
});

backBlog.addEventListener("click", function () {
  blogSubmenu.classList.remove("open");
});

/**/
var currentPath = window.location.pathname;
var navLinks = document.querySelectorAll(".nav-menu .nav-link");
navLinks.forEach(function (link) {
  if (link.getAttribute("href") === currentPath) {
    link.classList.add("active");
  }
});

var modal = document.getElementById("myModal");
var span = document.getElementsByClassName("close")[0];
var cta = document.getElementsByClassName("bg_cta")[0];
var buttons = document.querySelectorAll(".modal-opener");

buttons.forEach(function (button) {
  button.onclick = function () {
    modal.style.display = "flex";
    cta.style.zIndex = "0";
  };
});

span.onclick = function () {
  modal.style.display = "none";
  cta.style.zIndex = 9999;
};

window.onclick = function (event) {
  if (event.target == modal) {
    modal.style.display = "none";
  }
};
