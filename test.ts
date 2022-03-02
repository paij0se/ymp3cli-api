const rawResponse = await fetch("http://localhost:8080/user", {
  method: "POST",
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
  },
  body: JSON.stringify({
    id: "pai",
    password:""
  }),
});
const content = await rawResponse.json();
console.log(content);
