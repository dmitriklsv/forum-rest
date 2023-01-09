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
