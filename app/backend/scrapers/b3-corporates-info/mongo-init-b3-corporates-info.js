db.createUser({
  user: "mongouser",
  pwd: "mongopass",
  roles: [
    {
      role: "readWrite",
      db: "b3-corporates-info",
    },
  ],
});
