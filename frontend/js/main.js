const getCookie = (name) => {
  return document.cookie
    .split(";")
    .map((pair) => pair.trim().split("="))
    .filter((key) => key === name)?.[1];
};

function onClickPost(id) {
  fetch(`https://dummyjson.com/posts/${id}`)
    .then((res) => res.json())
    .then((data) => {
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

      getComments(id).then((response) => {
        document.querySelector(".layout-container").innerHTML += response;
      });
    });
}

const getPostMarkup = (post) => {
  const postDiv = document.createElement("div");
  postDiv.className = "post";
  postDiv.innerHTML = `
    <h4 class="post-title">${post.title}</h4>
    <p class="post-content">${post.body}</p>
    <a href="javascript:;" id="post-${post.id}">Discuss</a>
    `;

  postDiv
    .querySelector(`a#post-${post.id}`)
    .addEventListener("click", function (event) {
      onClickPost(post.id);
    });

  return postDiv;
};

const check = (event) => {
  alert(event.target.value);
};

const getPosts = () =>
  fetch("https://dummyjson.com/posts")
    .then((res) => res.json())
    .then((json) => {
      const container = document.createElement("div");
      container.className = "layout-container";
      json.posts.forEach((post) => {
        container.append(getPostMarkup(post));
      });

      document.querySelector("#root").append(container);
    });

getPosts();

const history = [];

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
