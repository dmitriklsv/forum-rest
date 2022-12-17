const getFormData = (formId) => {
  if (!formId) {
    console.error("No form id");
  }

  const formRef = document.getElementById(formId);

  if (!formRef) {
    console.error("Unable to find form by id");
  }

  const formData = new FormData(formRef);
  const data = {};
  for (const [key, value] of formData.entries()) {
    data[key] = value;
  }

  return data;
};

function serveRegisterMarkup() {
  const registerFormMarkup = `
    <form id="registrationForm">
      <div class="inputGroup">
        <label for="email">Your e-mail</label>
        <input id="email" type="email" name="email"></input>
      </div>
      <div class="inputGroup">
        <label for="password">Your password</label>
        <input id="password" type="password" name="password"></input>
      </div>
      <div class="inputGroup">
        <label for="repeatedPassword">Repeat your password</label>
        <input id="repeatedPassword" type="password"></input>
      </div>
      <div class="inputGroup">
        <button class="button primary" id="registerButton">Register</button>
      </div>
    </form>
  `;

  document.querySelector("#root").innerHTML = registerFormMarkup;

  const registerButton = document.getElementById("registerButton");
  console.log(registerButton);
  registerButton.addEventListener("click", onRegister);
}

const onRegister = async () => {
  const formId = "registrationForm";

  const formData = getFormData(formId);

  const response = await fetch("http://localhost:8080/signup", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(formData),
  });

  return await response.json();
};

const header = `
<div class="header">
    <h2 class="header-title">forum</h2>
    <div class="header-profile">
      <a class="" href="javascript:;">Login</a>
      or
      <a class="" href="javascript:;" id="header-register">Register</a>
    </div>
</div>
`;

document.querySelector("#root").innerHTML = header;

const headerRegister = document.getElementById("header-register");
headerRegister.onclick = serveRegisterMarkup;
console.log(headerRegister);

const getCookie = (name) => {
  return document.cookie
    .split(";")
    .map((pair) => pair.trim().split("="))
    .filter((key) => key === name)?.[1];
};

const getPostMarkup = (post) => `
    <div class="post">
        <h4 class="post-title">${post.title}</h4>
        <p class="post-content">${post.body}</p>
        <a href="javascript:;" onclick="onClickPost(${post.id})">Discuss</a>
    </div>
`;

const getPosts = () =>
  fetch("https://dummyjson.com/posts")
    .then((res) => res.json())
    .then((json) => {
      const posts = `
            <div class="layout-container">
                ${json.posts.map((post) => getPostMarkup(post)).join("")}
            </div>
        `;
      document.querySelector("#root").innerHTML += posts;
    });

getPosts();

const history = [];

const onClickPost = (id) => {
  fetch(`https://dummyjson.com/posts/${id}`)
    .then((res) => res.json())
    .then(async (data) => {
      const post = `
                <div class="post">
                    <h2 class="post-title">${data.title}</h2>
                    <p class="post-content">${data.body}</p>
                    <ul class="tags">
                        ${data?.tags
                          ?.map(
                            (tag) => `
                            <li class="tag">${tag}</li>
                        `
                          )
                          .join("")}
                    </ul>
                </div>
            `;

      history.push(document.querySelector(".layout-container").innerHTML);

      document.querySelector(".layout-container").innerHTML = post;

      const comments = await getComments(id).then((response) => {
        document.querySelector(".layout-container").innerHTML += response;
      });
    });
};

const onGoBack = () => {
  if (!history.length) {
    return;
  }

  const prevPage = history.splice(history.length - 1, 1);

  document.querySelector(".layout-container").innerHTML = prevPage;

  history.pop();
};

const getComments = (postId) =>
  fetch(`https://dummyjson.com/posts/${postId}/comments`)
    .then((res) => res.json())
    .then(
      (json) => `
            <div class="comments">
                ${json.comments
                  .map((comment) => getCommentMarkup(comment))
                  .join("")}
            </div>
        `
    );

const getCommentMarkup = (comment) => `
    <div class="comment">
        <p class="comment-username"><strong>${comment.user.username}</strong> says...</p>
        <p class="comment-content">${comment.body}</p>
    </div>
`;
