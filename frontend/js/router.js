const route = (event) => {
  event = event || window.event;
  event.preventDefault();
  window.history.pushState({}, "", event.target.href);
  handleLocation();
};

const routes = {
  404: "/frontend/pages/404.html",
  "/frontend/": "/frontend/pages/index.html",
  "/registration": "/frontend/pages/registration.html",
  "/post/:id": "/frontend/pages/post.html",
};
//todo - match post route
const handleLocation = async () => {
  const path = window.location.pathname;
  const route = routes[path] || routes[404];
  console.log(path);
  console.log(route);
  const html = await fetch(route).then((data) => data.text());

  document.getElementById("root").innerHTML = html;
};

window.onpopstate = handleLocation;
window.route = route;

handleLocation();
