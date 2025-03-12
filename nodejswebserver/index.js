const express = require("express");
const app = express();
const port = 3000;

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.get("/", (req, res) => {
  res.status(200).send("welcome to the node js server");
});

app.get("/get", (req, res) => {
  res.status(200).send({ message: "hello from the node js server" });
});

app.post("/post", (req, res) => {
  let myjson = req.body;
  res.status(200).send(myjson);
});

app.post("/postform", (req, res) => {
  let myjson = req.body;

  res.status(200).send(JSON.stringify(myjson));
});

app.listen(port, () => {
  console.log("Example app listing at port: ", port);
});
