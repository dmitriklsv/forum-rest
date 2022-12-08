const header = `
          <div class="header">
              <h2 class="header-title">forum</h2>
              <a class="header-profile" href="javascript:;" onclick="onClickProfile()">Profile</a>
          </div>
      `;

document.querySelector("#root").innerHTML = header;

const onClickProfile = () => {
  console.log(getCookie("{81851175-761f-4d46-80dc-c7bdf5c8387f}"));
};

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
