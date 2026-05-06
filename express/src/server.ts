import express from "express";

const app = express();

app.get("/health", (_req, res) => {
  res.status(200).json({ ok: true });
});

app.get("/api/v1/users/:id", (req, res) => {
  res.status(200).json({
    id: req.params.id,
    message: "user: " + req.params.id
  });
});

app.get("/api/v1/users/", (req, res) => {
  res.status(200).json({
    message: "users"
  });
});

const port = Number(process.env.PORT) || 3000;
app.listen(port, () => {
  console.log(`mini express listening on :${port}`);
});
