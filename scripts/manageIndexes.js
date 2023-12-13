print("operating on killmail collection");
// killmail collections
print("-- creating killmail id index");
db.killmails.createIndex( { "killmailID": -1 }, { unique: true } );
print("-- creating systemID and killmailTime index");
db.killmails.createIndex( { "solarSystemID": -1, "killmailTime": -1 } );


// esi status collection
print("operating on esiStatus collection");
print("-- creating esiStatus index");
db.server_statuses.createIndex( { "created_at": -1 }, { unique: true });

// esi universe collection
print("operating on esiStatus collection");
print("-- creating esiStatus index");
db.universe_types.createIndex( { "typeID": -1 }, { unique: true } );