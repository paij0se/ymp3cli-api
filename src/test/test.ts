const rawResponse = await fetch("http://localhost:5000/user", {
  method: "POST",
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
  },
  body: JSON.stringify({
    User: "pai",
  }),
});
const content = await rawResponse.json();
console.log(content);
