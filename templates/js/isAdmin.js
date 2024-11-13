  document.addEventListener("DOMContentLoaded", function() {
    const jwt = localStorage.getItem("jwt");
    const navContainer = document.querySelector(".navbar-nav");

    if (jwt) {
      navContainer.innerHTML = `
        <li class="nav-item">
             <a class="nav-link active" aria-current="page" href="/">Search</a>
         </li>
        <li class="nav-item">
          <a class="nav-link" href="/admin/upload">Upload</a>
        </li>
      `;
    } else {
      navContainer.innerHTML = `
        <li class="nav-item">
          <a class="nav-link active" aria-current="page" href="/">Search</a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="login_page.html">Login</a>
        </li>
      `;
    }
  });
