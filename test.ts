const rawResponse = await fetch("https://ymp3cli-api.herokuapp.com/user", {
  method: "POST",
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
  },
  body: JSON.stringify({
    Name: "pai",
    Password:"d24250200f82994615e9f4c7336c4aa4457e28d39491cce206f60f89b73594ea98a047edfd8e68afe7705a7b2098cd81cc14043cb52e87c391b8e68c9f65c354"
  }),
});
const content = await rawResponse.text();
console.log(content);
