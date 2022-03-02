const rawResponse = await fetch("https://ymp3cli-api.herokuapp.com/user", {
  method: "POST",
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
  },
  body: JSON.stringify({
    id:"",
    Name: "pai",
  }),
});
const content = await rawResponse.text();
console.log(content);
