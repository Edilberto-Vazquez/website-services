db.createUser({
  user: "root",
  pwd: "root",
  roles: [
    {
      role: "readWrite",
      db: "db-dev",
    },
  ],
});

db.createCollection("coll-dev");
